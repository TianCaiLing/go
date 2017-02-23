package controllers

import (
	_ "database/sql"
	//"encoding/json"
	_ "errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	_ "gmtools/models"
	_ "strconv"
)

type DoChannelCmd struct {
	Type string
}

type DoChannel struct {
	beego.Controller
}

func (c *DoChannel) Post() {
	c.Ctx.Request.ParseForm()
	beego.Debug(c.Input())
	severid := c.GetString("id")
	severname := c.GetString("name")
	version := c.GetString("version")
	CDN := c.GetString("CDN")
	qd := c.GetString("qd")

	req := httplib.Post(beego.AppConfig.String("carriershost") + "servs/update")
	req.Param("servid", severid)
	req.Param("channels", qd)
	req.Param("version", version)
	req.Param("resVersion", CDN)
	str, err := req.String()
	if err != nil {
		beego.Debug(err)
	}

	beego.Debug(severid, severname, version, CDN, qd)

	c.Ctx.WriteString(str)

}
