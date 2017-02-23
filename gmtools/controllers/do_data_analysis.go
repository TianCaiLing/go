package controllers

import (
	"encoding/json"
	_ "strconv"
	_ "strings"
	_ "time"

	"github.com/astaxie/beego"
)

type DataAnalysisCmd struct {
	Type string
}

type DoDataAnalysis struct {
	beego.Controller
}

func (c *DoDataAnalysis) Post() {
	c.Ctx.Request.ParseForm()
	beego.Debug(c.GetStrings("ckb"))
	servs := MakeServIds(c.Input(), "data_function_serv_")

	if len(servs) == 0 {
		c.Redirect("/error?param=没有选择服务器", 302)
		return
	}
	cmd := DataAnalysisCmd{}
	cmd.Type = "GMT_DataAnalysis"

	jparam, _ := json.Marshal(cmd)
	beego.Debug(string(jparam))
	results := PostGMTServs(servs, jparam)
	jresults, _ := json.Marshal(results)
	c.Redirect("/alert?param="+string(jresults), 302)

}
