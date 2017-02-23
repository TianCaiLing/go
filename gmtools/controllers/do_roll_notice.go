package controllers

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type NoticeCmd struct {
	Type     string `json:"Type"`
	SendType string `json:"NoticeSendType"`
	Content  string `json:"Content"`
	TheTime  int64  `json:"TheTime"`
	ItvTime  int64  `json:"ItvTime"`
}

type DoRollNotice struct {
	beego.Controller
}

func (c *DoRollNotice) Post() {
	c.Ctx.Request.ParseForm()
	beego.Debug(c.Input())
	servs := ToServIds(c.GetStrings("servs"))

	if len(servs) == 0 {
		c.Redirect("/error?param=没有选择服务器", 302)
		return
	}

	var cmd NoticeCmd
	cmd.Type = "GMT_Notice"
	cmd.SendType = c.GetString("types")
	cmd.Content = c.GetString("content")
	if len(cmd.Content) == 0 {
		c.Redirect("/error?param=没有公告内容", 302)
		return
	}
	time0 := c.GetString("param_name0")
	var sendType = cmd.SendType
	var content = cmd.Content
	var timestr = time0
	beego.Debug(sendType)
	beego.Debug(timestr)

	if cmd.SendType != "NST_Immediately" {
		param := strings.Replace(time0, "/", "-", -1) + " " + ":00"

		beego.Debug(param)

		if cmd.SendType == "NST_Timming" {
			thetime, _ := time.Parse("2006-01-02 15:04", param)
			beego.Debug(param, thetime)

			if thetime.Unix() < time.Now().Unix() {
				c.Redirect("/error?param=定时发送时间已经过去(请检查时间格式 2006-01-02 15:04:05", 302)
				return
			}
			cmd.TheTime = thetime.Unix()
		} else {

			strtime := c.Input().Get("param_name3")
			param1 := strings.Replace(c.Input().Get("param_name4"), "/", "-", -1) + " " + ":00"

			thetime, _ := time.ParseInLocation("2006-1-02 15:04", param1, time.Local)

			itvtime, _ := strconv.Atoi(strtime)
			beego.Debug(param, param1, strtime)

			if thetime.Unix() < time.Now().Unix() {
				c.Ctx.WriteString("/failed?param=循环发送结束时间已经过去(请检查时间格式 2006-01-02 15:04:05,30")
				return
			}
			if itvtime <= 0 {
				c.Ctx.WriteString("/error?param=循环发送间隔时间错误(请检查时间格式 2006-01-02 15:04:05,30")
				return
			}
			cmd.TheTime = thetime.Unix()
			cmd.ItvTime = int64(itvtime)

		}

	}
	time := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	jparam, _ := json.Marshal(cmd) //把数据转成json数据
	results := PostGMTServs(servs, jparam)
	jresults, _ := json.Marshal(results)
	beego.Debug(string(jparam))
	RollNoticeInsert(sendType, content, timestr, time)
	c.Ctx.WriteString(string(jresults))
}
