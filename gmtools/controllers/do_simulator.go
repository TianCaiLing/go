package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type SimulatorCmd struct {
	Type     string
	PlayerId int
	ShopId   int
	Payment  int
	OrderId  string
}

type DoSimulator struct {
	beego.Controller
}

func (c *DoSimulator) Post() {
	c.Ctx.Request.ParseForm()
	beego.Debug(c.Input())
	serv := ToServIds(c.GetStrings("servs"))[0]
	beego.Debug(serv)
	playerid := c.Input().Get("PlayerId")
	PlayerId, errid := strconv.Atoi(playerid)
	if errid != nil {
		c.Ctx.WriteString(errid.Error())
		return
	}
	shopid := c.Input().Get("ShopId")
	ShopId, errshop := strconv.Atoi(shopid)
	if errshop != nil {
		c.Ctx.WriteString(errshop.Error())
		return
	}
	payment := c.Input().Get("Payment")
	Payment, errpayment := strconv.Atoi(payment)
	if errpayment != nil {
		c.Ctx.WriteString(errpayment.Error())
		return
	}
	orderid := c.Input().Get("Orderid")
	cmd := SimulatorCmd{}
	cmd.Type = "GNT_MakeOrder"
	cmd.PlayerId = PlayerId
	cmd.ShopId = ShopId
	cmd.Payment = Payment
	cmd.OrderId = orderid

	jparam, _ := json.Marshal(cmd)
	beego.Debug(string(jparam))
	results := PostGMTServ(serv, jparam)
	jresults, _ := json.Marshal(results)
	beego.Debug(string(jparam))

	c.Ctx.WriteString(string(jresults))
}
