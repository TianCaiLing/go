package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	_ "strings"
	_ "time"
)

type OrderInquiryResult struct {
	Game      string
	Paytime   string
	Orderid   string
	Pfid      string
	Pfname    string
	Accountid string
	Roleid    int
	Payment   float32
}
type OrderInquiryCmd struct {
	Type string
}

type DoOrderInquiry struct {
	beego.Controller
}

func (c *DoOrderInquiry) Post() {
	c.Ctx.Request.ParseForm()
	beego.Debug(c.Input())
	sever, _ := strconv.Atoi(c.GetString("id"))
	val := c.GetString("val")
	beego.Debug(sever, val)
	fb, err := Database1(sever)
	beego.Debug(fb)
	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer fb.Close()
	var records *sql.Rows
	if val == "" {

		records, err = fb.Query("SELECT `game`,`paytime`,`orderid`,`pfid`,`pfname`,`accountid`,`roleid`,`payment` FROM `order`")

		if err != nil {
			beego.Error("DB fail", err)
			return
		}
	} else {
		var err error
		records, err = fb.Query("SELECT `game`,`paytime`,`orderid`,`pfid`,`pfname`,`accountid`,`roleid`,`payment` FROM `order` WHERE `pfid` = ?", val)

		if err != nil {
			beego.Error("DB fail", err)
			return
		}
	}

	results := []OrderInquiryResult{}

	for records.Next() {
		tmp := OrderInquiryResult{}
		err := records.Scan(&tmp.Game, &tmp.Paytime, &tmp.Orderid, &tmp.Pfid, &tmp.Pfname, &tmp.Accountid, &tmp.Roleid, &tmp.Payment)
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
