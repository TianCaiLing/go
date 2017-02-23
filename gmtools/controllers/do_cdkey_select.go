package controllers

import (
	 "encoding/json"
	
	"github.com/astaxie/beego"
)

type CdkeySelectCmd struct {
	Type            string
	SelectCondition []string
}

type DoCdkeySelect struct {
	beego.Controller
}

func (c *DoCdkeySelect) Post() {
	c.Ctx.Request.ParseForm()
	beego.Debug(c.Input())
	
//	servs := ToServIds(c.GetStrings("servs"))
//	if len(servs) == 0 {
//		c.Redirect("/error?param=没有选择服务器", 302)
//		return
//	}
	
	var cmd CdkeySelectCmd
	
	cmd.Type = "GMT_CdkeySelect"
	
	cmd.SelectCondition = c.GetStrings("cdkey_select")
	
	beego.Debug(cmd)
	
    jparam, _ := json.Marshal(cmd)
	
//	results := PostGMTServs(servs, jparam)
//	jresults, _ := json.Marshal(results)
//	c.Ctx.WriteString(string(jparam))

	c.Redirect("/alert?param="+string(jparam), 302)
}