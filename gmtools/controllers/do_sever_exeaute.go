package controllers

import (
	_ "database/sql"
	//"encoding/json"
	_ "errors"
	"github.com/astaxie/beego"
	"gmtools/models"
	"strconv"
)

type DoSeverCmd struct {
	Type string
}

type DoSever struct {
	beego.Controller
}

func (c *DoSever) open(severid int) {
	host := models.GetGMTServHost(severid)
	beego.Debug(host)
	c.Ctx.WriteString(string(host))
}

func (c *DoSever) close(severid int) {
	host := models.GetGMTServHost(severid)
	beego.Debug(host)
	c.Ctx.WriteString(string(host))
}

func (c *DoSever) backup(severid int) {
	host := models.GetGMTServHost(severid)
	beego.Debug(host)
	c.Ctx.WriteString(string(host))
}

func (c *DoSever) Post() {
	c.Ctx.Request.ParseForm()
	beego.Debug(c.Input())
	severid, _ := strconv.Atoi(c.GetString("sever_id"))
	button := c.GetString("button")
	beego.Debug(severid, button)
	if button == "open" {
		c.open(severid)
		c.Ctx.WriteString(string("ok"))
	} else if button == "close" {
		c.close(severid)
	} else if button == "backup" {
		c.backup(severid)
	}

}
