package controllers

import (
	"github.com/astaxie/beego"
)

type Error struct {
	beego.Controller
}

func (c *Error) Get() {
	c.Ctx.Request.ParseForm()
	c.TplName = "error.html"

	c.Data["ErrorDesc"] = c.Input().Get("param")
}
