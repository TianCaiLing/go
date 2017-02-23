package controllers

import (
	 _"encoding/json"
	_ "strconv"
	_ "strings"
	_ "time"

	"github.com/astaxie/beego"
)

type CdkeyGenerateCmd struct {
	Type string
}

type DoCdkeyGenerate struct {
	beego.Controller
}

func (c *DoCdkeyGenerate) Post() {
	c.Ctx.Request.ParseForm()
	
	
//	jparam, _ := json.Marshal(cmd)
//	beego.Debug(string(jparam))
	
//	jresults, _ := json.Marshal(jparam)
//	c.Redirect("/alert?param="+string(jresults), 302)
	

}
