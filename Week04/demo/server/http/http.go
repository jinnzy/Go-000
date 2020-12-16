package http

import (
	"demo/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	svc *service.Service
)

func Init(s *service.Service) {
	svc = s
	svr := &http.Server{
		Addr:           fmt.Sprintf(":8080"),
		Handler:        initRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}
	fmt.Println("http init")
}


func initRouter() *gin.Engine {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK,"../front/dist/index.html",gin.H{})
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "test",
		})
	})

	return r
}
