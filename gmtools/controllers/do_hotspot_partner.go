package controllers

import (
	"encoding/json"
	"strconv"

	"time"

	"github.com/astaxie/beego"
)

type HotspotReward struct {
	LimitNum []string
	RoleId   uint32
	Price    uint32
	BuyLimit uint32
}
type HotspotPartnerCmd struct {
	Type string
	/*PropQuantity    int*/
	OpenTime  uint64
	CloseTime uint64
	Items     []HotspotReward
}

type DoHotspotPartner struct {
	beego.Controller
}

func (c *DoHotspotPartner) Post() {
	c.Ctx.Request.ParseForm()
	beego.Debug(c.Input())
	servs := ToServIds(c.GetStrings("servs"))
	if len(servs) == 0 {
		c.Redirect("/error?param=没有选择服务器", 302)
		return
	}

	cmd := HotspotPartnerCmd{}
	cmd.Type = "GMT_HotRole"

	thetime, _ := time.ParseInLocation("2006-1-02 15:04", c.Input().Get("opentime1"), time.Local)
	cmd.OpenTime = uint64(thetime.Unix())

	thetime1, _ := time.ParseInLocation("2006-1-02 15:04", c.Input().Get("closetime1"), time.Local)
	cmd.CloseTime = uint64(thetime1.Unix())

	beego.Debug(c.Input().Get("opentime1"), thetime1)
	cmd.Items = []HotspotReward{}

	Price := c.GetStrings("partner_price")

	PropId := c.GetStrings("prop_id")

	BuyLimit := c.GetStrings("purchased_number")

	EntityType := c.GetStrings("partner_type")

	for i := 0; i < len(Price); i++ {

		rwd := HotspotReward{}

		Price, errNum := strconv.Atoi(Price[i])
		if errNum != nil {
			c.Ctx.WriteString(errNum.Error())
			return
		}
		rwd.Price = uint32(Price)

		PropId, errId := strconv.Atoi(PropId[i])
		if errId != nil {
			c.Ctx.WriteString(errId.Error())
			return
		}
		rwd.RoleId = uint32(PropId)

		BuyLimit, errPrice := strconv.Atoi(BuyLimit[i])
		if errPrice != nil {
			c.Ctx.WriteString(errPrice.Error())
			return
		}
		rwd.BuyLimit = uint32(BuyLimit)

		rwd.LimitNum = EntityType

		cmd.Items = append(cmd.Items, rwd)
	}

	if thetime.Unix() > thetime1.Unix() {
		c.Redirect("/error?param=活动关闭时间不能小于开始时间", 302)
		return
	}
	beego.Debug(cmd.Items)
	jparam, _ := json.Marshal(cmd)
	beego.Debug(string(jparam))
	results := PostGMTServs(servs, jparam)
	jresults, _ := json.Marshal(results)
	time := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	HotspotPartnerInsert(c.Input().Get("opentime1"), c.Input().Get("closetime1"), time)
	//c.Ctx.WriteString(string(jparam))
	c.Redirect("/alert?param="+string(jresults), 302)
}
