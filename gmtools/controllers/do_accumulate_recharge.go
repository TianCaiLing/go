package controllers

import (
	"encoding/json"
	"strconv"
	//	"strings"
	"time"

	"github.com/astaxie/beego"
)

type Reward struct {
	ItemId    uint32
	ItemStack uint32
}
type Content struct {
	LimitNum uint32
	Reward   []Reward
}

type AccmulateRechargeCmd struct {
	Type      string
	OpenTime  uint64
	CloseTime uint64
	Content   []Content
}

type DoAccmulateRecharge struct {
	beego.Controller
}

func (c *DoAccmulateRecharge) Post() {
	c.Ctx.Request.ParseForm()
	beego.Debug(c.Input())
	servs := ToServIds(c.GetStrings("servs"))
	if len(servs) == 0 {
		c.Redirect("/error?param=没有选择服务器", 302)
		return
	}

	cmd := AccmulateRechargeCmd{}
	cmd.Type = "GMT_ChargeTotal"

	thetime, _ := time.ParseInLocation("2006-1-02 15:04", c.Input().Get("appointed_day1"), time.Local)
	cmd.OpenTime = uint64(thetime.Unix())

	thetime1, _ := time.ParseInLocation("2006-1-02 15:04", c.Input().Get("close_day1"), time.Local)
	cmd.CloseTime = uint64(thetime1.Unix())

	cmd.Content = []Content{}

	var strLimit string
	for i := 0; i < 8; i++ {
		content := Content{}

		strLimit = c.GetString("ckb" + strconv.FormatInt(int64(i), 10))
		itemLimit, errLimit := strconv.Atoi(strLimit)
		if errLimit != nil {
			c.Ctx.WriteString(errLimit.Error())
			return
		}
		content.LimitNum = uint32(itemLimit)

		strLimit += strLimit + ","

		var stritemId string
		var stritemStack string
		for j := 1; j < 4; j++ {

			stritemId = c.GetString("ckab" + strconv.FormatInt(int64(i*3+j), 10))
			stritemStack = c.GetString("ckeb" + strconv.FormatInt(int64(i*3+j), 10))

			itemid, errid := strconv.Atoi(stritemId)
			itemsk, errsk := strconv.Atoi(stritemStack)

			if errid != nil {
				c.Ctx.WriteString(errid.Error())
				return
			}
			if errsk != nil {
				c.Ctx.WriteString(errsk.Error())
				return
			}
			if itemid != 0 && itemsk != 0 {
				if content.Reward == nil {
					content.Reward = []Reward{}
				}
				rwd := Reward{}
				rwd.ItemId = uint32(itemid)
				rwd.ItemStack = uint32(itemsk)

				content.Reward = append(content.Reward, rwd)
			}

			stritemId += stritemId + ","
			stritemStack += stritemStack + ","

		}
		cmd.Content = append(cmd.Content, content)
	}
	if thetime.Unix() > thetime1.Unix() {
		c.Redirect("/error?param=活动关闭时间不能小于开始时间", 302)
		return
	}

	jparam, _ := json.Marshal(cmd)
	beego.Debug(string(jparam))
	results := PostGMTServs(servs, jparam)
	jresults, _ := json.Marshal(results)
	time := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	AccumulateRechargeInsert(c.Input().Get("appointed_day1"), c.Input().Get("close_day1"), time)
	c.Redirect("/alert?param="+string(jresults), 302)
}
