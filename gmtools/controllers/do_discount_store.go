package controllers

import (
	"encoding/json"
	"strconv"
	//	"strings"
	"time"

	"github.com/astaxie/beego"
)

type DiscountReward struct {
	ItemId uint32

	Discount float64
	Price    uint32
	LimitNum uint32
}
type DiscountStoreCmd struct {
	Type      string
	OpenTime  uint64
	CloseTime uint64
	Rewards   []DiscountReward
}

type DoDiscountStore struct {
	beego.Controller
}

func (c *DoDiscountStore) Post() {
	c.Ctx.Request.ParseForm()

	servs := ToServIds(c.GetStrings("servs"))

	if len(servs) == 0 {
		c.Redirect("/error?param=没有选择服务器", 302)
		return
	}

	cmd := DiscountStoreCmd{}

	cmd.Type = "GMT_DiscountStore"

	thetime, _ := time.ParseInLocation("2006-1-02 15:04", c.Input().Get("open_time1"), time.Local)
	cmd.OpenTime = uint64(thetime.Unix())

	thetime1, _ := time.ParseInLocation("2006-1-02 15:04", c.Input().Get("close_time1"), time.Local)
	cmd.CloseTime = uint64(thetime1.Unix())

	cmd.Rewards = []DiscountReward{}

	strs0 := c.GetStrings("ckb")
	strs1 := c.GetStrings("ckeb1")
	strs2 := c.GetStrings("ckeb2")
	strs3 := c.GetStrings("ckeb3")

	for i := 0; i < len(strs0); i++ {

		rwd := DiscountReward{}

		LimitNum, errNum := strconv.Atoi(strs3[i])
		if errNum != nil {
			c.Ctx.WriteString(errNum.Error())
			return
		}
		rwd.LimitNum = uint32(LimitNum)

		ItemId, errId := strconv.Atoi(strs0[i])
		if errId != nil {
			c.Ctx.WriteString(errId.Error())
			return
		}
		rwd.ItemId = uint32(ItemId)

		Price, errPrice := strconv.Atoi(strs1[i])
		if errPrice != nil {
			c.Ctx.WriteString(errPrice.Error())
			return
		}
		rwd.Price = uint32(Price)

		Discount, _ := strconv.ParseFloat(strs2[i], 64)
		beego.Debug(Discount)
		rwd.Discount = Discount

		cmd.Rewards = append(cmd.Rewards, rwd)
	}
	if thetime.Unix() > thetime1.Unix() {
		c.Redirect("/error?param=活动关闭时间不能小于开始时间", 302)
		return
	}
	jparam, _ := json.Marshal(cmd)
	beego.Debug(string(jparam))
	results := PostGMTServs(servs, jparam)
	jresults, _ := json.Marshal(results)
	DiscountStoreInsert(c.Input().Get("open_time1"), c.Input().Get("close_time1"), time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
	c.Redirect("/alert?param="+string(jresults), 302)
}
