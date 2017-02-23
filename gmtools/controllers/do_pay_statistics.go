package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	_ "strconv"
	_ "strings"
	_ "time"
)

type PayStatisticsResult struct {
	Sortid    int
	Game      string
	Pfid      string
	Pfname    string
	Orderid   string
	Accountid string
	Payment   float32
	Paytime   string
}
type PayStatisticsCmd struct {
	Type string
}

type DoPayStatistics struct {
	beego.Controller
}

func (c *DoPayStatistics) Post() {
	c.Ctx.Request.ParseForm()

	fb, err := Database1(3)
	beego.Debug(fb)
	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer fb.Close()

	records, err := fb.Query("SELECT * FROM `order`")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []PayStatisticsResult{}

	for records.Next() {

		tmp := PayStatisticsResult{}
		beego.Debug(tmp)
		err := records.Scan(&tmp.Sortid, &tmp.Game, &tmp.Pfid, &tmp.Pfname, &tmp.Orderid, &tmp.Accountid, &tmp.Payment, &tmp.Paytime)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}

		results = append(results, tmp)

	}
	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))

}
