package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, _ := errgroup.WithContext(ctx)
	// 启动http1
	http1 := http.Server{
		Addr:              ":8111",
	}
	g.Go(func() error {
		if err := http1.ListenAndServe(); err != nil {
			cancel()
			return err
		}
		return nil
	})
	// 启动http2
	http2 := http.Server{
		Addr:              ":8112",
	}
	g.Go(func() error {
		if err := http2.ListenAndServe(); err != nil {
			cancel()
			return err
		}
		return nil	})


	// 监听sig信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for  {
			select {
			case s := <- c:
				switch s {
				case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
					cancel()
				default:
				}
			}
		}
	}()

	// context取消后，关闭http server
	go func() {
		select {
		case <- ctx.Done():
			log.Println(ctx.Err())
			http1.Shutdown(context.Background())
			http1.Close()
			http2.Shutdown(context.Background())
			return
		}

	}()

	if err := g.Wait(); err != nil {
		log.Println(err)
		log.Println("all exit")
	}
}
