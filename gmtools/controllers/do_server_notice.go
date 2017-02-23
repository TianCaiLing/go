package controllers

import (
	"encoding/json"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type DoServerNotice struct {
	beego.Controller
}

func (c *DoServerNotice) Post() {
	c.Ctx.Request.ParseForm()
	beego.Debug(c.Input())
	c.Ctx.WriteString("OK")

	servs := ToServIds(c.GetStrings("servs"))

	if len(servs) == 0 {
		c.Redirect("/error?param=没有选择服务器", 302)
		return
	}
	content := c.GetString("content")
	title := c.GetString("title")
	color := c.GetString("color")

	if len(color) == 0 {
		color = "00FF00"
	}
	if len(content) != 0 {
		content = "[" + color + "]" + content + "[-]"
		content = title + ";" + content
	}
	strservs, _ := json.Marshal(servs)

	req := httplib.Post(beego.AppConfig.String("carriershost") + "gmtools/notice/server")
	req.Header("content-type", "application/x-www-form-urlencoded")
	req.Param("servid", (string)(strservs))
	req.Param("content", content)
	str, err := req.String()

	time := time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05")
	ServiceNoticeInsert(title,content,color,time)
	if err == nil {
		beego.Info(content)
		c.Redirect("/alert?param="+str, 302)
		
	} else {
		c.Redirect("/error?param="+err.Error(), 302)
	}

}
