package dao

import "demo/conf"

type Dao struct {
}

// 后续可以替换为传递 c 变量到各个init中
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
	}

	return
}

