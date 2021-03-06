package controller

import (
	"fmt"
	"github.com/cleango/gallop"
	"github.com/cleango/gallop/example/app/config"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
	Demo *config.Demo  `inject:""`
	Demo1 *config.Demo `inject:"demo"`
	Cfg *config.Configuration `inject:""`
}

func (ctr *HelloController) Hello(c *gallop.Context) string {
	return fmt.Sprint(ctr.Demo,ctr.Demo1,ctr.Cfg.B.C)
}

func (ctr *HelloController) Json(c *gallop.Context) gallop.Json{
	return gin.H{
		"hello":"world",
	}
}