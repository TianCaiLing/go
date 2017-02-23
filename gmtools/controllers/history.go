package controllers

import (
	"encoding/json"
	_ "strconv"
	"time"

	"github.com/astaxie/beego"
)

type SystemNoticeHistoryResult struct {
	Title   string
	Content string
	Color   string
	Time    string
}

type ServerNoticeHistoryResult struct {
	Title        string
	Content      string
	ContentColor string
	Time         string
}

type RollNoticeHistoryResult struct {
	SendType string
	Content  string
	Timestr  string
	Time     string
}
type InsertMailHistoryResult struct {
	Title      string
	Sender     string
	Content    string
	Recvers    string
	Stritemids string
	Stritemsks string
	Lowlevel   int
	Highlevel  int
	Time0      string
	Time1      string
	SendType   string
	Time       string
}
type AccumulateRechargeHistoryResult struct {
	Name      string
	OpenTime  string
	CloseTime string
	Status    string
	Time      string
}
type DiscountStoreHistoryResult struct {
	Name      string
	OpenTime  string
	CloseTime string
	Status    string
	Time      string
}

type SingleRechargeHistoryResult struct {
	Name      string
	OpenTime  string
	CloseTime string
	Status    string
	Time      string
}
type HotspotPartnerHistoryResult struct {
	Name      string
	OpenTime  string
	CloseTime string
	Status    string
	Time      string
}
type ExtractPartnerHistoryResult struct {
	Name      string
	OpenTime  string
	CloseTime string
	Status    string
	Time      string
}
type LoginActivityHistoryResult struct {
	Name      string
	OpenTime  string
	CloseTime string
	Status    string
	Time      string
}
type MoneyPartnerHistoryResult struct {
	Name      string
	OpenTime  string
	CloseTime string
	Status    string
	Time      string
}
type History struct {
	beego.Controller
}

//系统公告
func (c *History) query_system_notice() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM notice WHERE types='system_notice'")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []SystemNoticeHistoryResult{}
	for records.Next() {

		var types string

		tmp := SystemNoticeHistoryResult{}

		err := records.Scan(&types, &tmp.Title, &tmp.Content, &tmp.Color, &tmp.Time)

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

//区服公告
func (c *History) query_server_notice() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM Service WHERE types='2'")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []ServerNoticeHistoryResult{}

	for records.Next() {

		var types int

		tmp := ServerNoticeHistoryResult{}

		err := records.Scan(&types, &tmp.Title, &tmp.Content, &tmp.ContentColor, &tmp.Time)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}

		results = append(results, tmp)
	}
	//	c.TplName = "server_notice_history.html"
	//	c.Data["Results"] = results

	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))
}

//滚动公告
func (c *History) query_roll_notice() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM Roll WHERE types='3'")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []RollNoticeHistoryResult{}
	beego.Debug(results)
	for records.Next() {

		var types int

		tmp := RollNoticeHistoryResult{}

		err := records.Scan(&types, &tmp.SendType, &tmp.Content, &tmp.Timestr, &tmp.Time)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}

		results = append(results, tmp)

	}
	//	c.TplName = "roll_notice_history.html"
	//	c.Data["Results"] = results

	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))
}

//历史邮件
func (c *History) query_insert_mail() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM mail")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []InsertMailHistoryResult{}

	for records.Next() {

		tmp := InsertMailHistoryResult{}

		err := records.Scan(&tmp.Title, &tmp.Sender, &tmp.Content, &tmp.Recvers, &tmp.Stritemids, &tmp.Stritemsks, &tmp.Lowlevel, &tmp.Highlevel, &tmp.Time0, &tmp.Time1, &tmp.SendType, &tmp.Time)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}

		results = append(results, tmp)

	}

	//	c.TplName = "insert_mail_history.html"
	//	c.Data["Results"] = results

	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))
}

//登录活动
func (c *History) query_login_activity() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM Login WHERE name='登录活动' ")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []LoginActivityHistoryResult{}

	for records.Next() {

		tmp := LoginActivityHistoryResult{}

		err := records.Scan(&tmp.Name, &tmp.OpenTime, &tmp.CloseTime, &tmp.Time)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}

		thetime, _ := time.ParseInLocation("2006-1-02 15:04", tmp.OpenTime, time.Local)
		opentime := thetime.Unix()

		thetime1, _ := time.ParseInLocation("2006-1-02 15:04", tmp.CloseTime, time.Local)
		closetime := thetime1.Unix()

		nowtime := time.Now().Unix()

		if nowtime <= closetime && nowtime >= opentime {
			tmp.Status = "正在进行"
		} else if nowtime < opentime {
			tmp.Status = "未开始"
		} else if nowtime > closetime {
			tmp.Status = "已结束"
		}

		results = append(results, tmp)
	}

	//	c.TplName = "login_activity_history.html"
	//	c.Data["Results"] = results

	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))
}

//累积充值
func (c *History) query_accumulate_recharge() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM Accumulate WHERE name='累积充值' ")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []AccumulateRechargeHistoryResult{}

	for records.Next() {

		tmp := AccumulateRechargeHistoryResult{}

		err := records.Scan(&tmp.Name, &tmp.OpenTime, &tmp.CloseTime, &tmp.Time)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}

		thetime, _ := time.ParseInLocation("2006-1-02 15:04", tmp.OpenTime, time.Local)
		opentime := thetime.Unix()

		thetime, _ = time.ParseInLocation("2006-1-02 15:04", tmp.CloseTime, time.Local)
		closetime := thetime.Unix()
		beego.Debug(thetime, thetime)
		nowtime := time.Now().Unix()
		if nowtime <= closetime && nowtime >= opentime {
			tmp.Status = "正在进行"
		} else if nowtime < opentime {
			tmp.Status = "未开始"
		} else if nowtime > closetime {
			tmp.Status = "已结束"
		}

		results = append(results, tmp)
	}

	//	c.TplName = "accumulate_recharge_history.html"
	//	c.Data["Results"] = results
	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))
}

//打折商店
func (c *History) query_discount_store() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM Discount WHERE name='打折商店'")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []DiscountStoreHistoryResult{}

	for records.Next() {

		tmp := DiscountStoreHistoryResult{}

		err := records.Scan(&tmp.Name, &tmp.OpenTime, &tmp.CloseTime, &tmp.Time)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}
		thetime, _ := time.ParseInLocation("2006-1-02 15:04", tmp.OpenTime, time.Local)
		opentime := thetime.Unix()
		thetime, _ = time.ParseInLocation("2006-1-02 15:04", tmp.CloseTime, time.Local)
		closetime := thetime.Unix()
		nowtime := time.Now().Unix()
		if nowtime <= closetime && nowtime >= opentime {
			tmp.Status = "正在进行"
		} else if nowtime < opentime {
			tmp.Status = "未开始"
		} else if nowtime > closetime {
			tmp.Status = "已结束"
		}
		results = append(results, tmp)
	}

	//	c.TplName = "discount_store_history.html"
	//	c.Data["Results"] = results

	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))
}

//单笔充值
func (c *History) query_single_recharge() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM Single WHERE name='单笔充值'")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []SingleRechargeHistoryResult{}

	for records.Next() {

		tmp := SingleRechargeHistoryResult{}

		err := records.Scan(&tmp.Name, &tmp.OpenTime, &tmp.CloseTime, &tmp.Time)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}
		thetime, _ := time.ParseInLocation("2006-1-02 15:04", tmp.OpenTime, time.Local)
		opentime := thetime.Unix()
		thetime, _ = time.ParseInLocation("2006-1-02 15:04", tmp.CloseTime, time.Local)
		closetime := thetime.Unix()
		nowtime := time.Now().Unix()
		if nowtime <= closetime && nowtime >= opentime {
			tmp.Status = "正在进行"
		} else if nowtime < opentime {
			tmp.Status = "未开始"
		} else if nowtime > closetime {
			tmp.Status = "已结束"
		}
		results = append(results, tmp)
	}

	//	c.TplName = "single_recharge_history.html"
	//	c.Data["Results"] = results

	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))
}

//热点伙伴
func (c *History) query_hotspot_partner() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM Hotspot WHERE name='热点伙伴'")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []HotspotPartnerHistoryResult{}

	for records.Next() {

		tmp := HotspotPartnerHistoryResult{}

		err := records.Scan(&tmp.Name, &tmp.OpenTime, &tmp.CloseTime, &tmp.Time)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}
		thetime, _ := time.ParseInLocation("2006-1-02 15:04", tmp.OpenTime, time.Local)
		opentime := thetime.Unix()
		thetime, _ = time.ParseInLocation("2006-1-02 15:04", tmp.CloseTime, time.Local)
		closetime := thetime.Unix()
		nowtime := time.Now().Unix()
		if nowtime <= closetime && nowtime >= opentime {
			tmp.Status = "正在进行"
		} else if nowtime < opentime {
			tmp.Status = "未开始"
		} else if nowtime > closetime {
			tmp.Status = "已结束"
		}
		results = append(results, tmp)
	}

	//	c.TplName = "hotspot_partner_history.html"
	//	c.Data["Results"] = results

	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))
}

//抽取伙伴
func (c *History) query_extract_partner() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM Extract WHERE name='顶级招募'")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []ExtractPartnerHistoryResult{}

	for records.Next() {

		tmp := ExtractPartnerHistoryResult{}

		err := records.Scan(&tmp.Name, &tmp.OpenTime, &tmp.CloseTime, &tmp.Time)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}
		thetime, _ := time.ParseInLocation("2006-1-02 15:04", tmp.OpenTime, time.Local)
		opentime := thetime.Unix()
		thetime, _ = time.ParseInLocation("2006-1-02 15:04", tmp.CloseTime, time.Local)
		closetime := thetime.Unix()
		nowtime := time.Now().Unix()
		if nowtime <= closetime && nowtime >= opentime {
			tmp.Status = "正在进行"
		} else if nowtime < opentime {
			tmp.Status = "未开始"
		} else if nowtime > closetime {
			tmp.Status = "已结束"
		}

		results = append(results, tmp)
	}

	//	c.TplName = "extract_partner_history.html"
	//	c.Data["Results"] = results

	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))
}

//小额礼包
func (c *History) query_money_partner() {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()

	records, err := db.Query("SELECT * FROM Money WHERE name='小额礼包'")

	if err != nil {
		beego.Error("DB fail", err)
		return
	}

	results := []MoneyPartnerHistoryResult{}

	for records.Next() {

		tmp := MoneyPartnerHistoryResult{}

		err := records.Scan(&tmp.Name, &tmp.OpenTime, &tmp.CloseTime, &tmp.Time)
		if err != nil {
			beego.Error("DB fail", err)
			return
		}
		thetime, _ := time.ParseInLocation("2006-1-02 15:04", tmp.OpenTime, time.Local)
		opentime := thetime.Unix()
		thetime, _ = time.ParseInLocation("2006-1-02 15:04", tmp.CloseTime, time.Local)
		closetime := thetime.Unix()
		nowtime := time.Now().Unix()
		if nowtime <= closetime && nowtime >= opentime {
			tmp.Status = "正在进行"
		} else if nowtime < opentime {
			tmp.Status = "未开始"
		} else if nowtime > closetime {
			tmp.Status = "已结束"
		}

		results = append(results, tmp)
	}

	//	c.TplName = "extract_partner_history.html"
	//	c.Data["Results"] = results

	beego.Debug(results)

	jparam, _ := json.Marshal(results)

	c.Ctx.WriteString(string(jparam))
}

func (c *History) Post() {
	c.Ctx.Request.ParseForm()

	typename := c.GetString("type")
	beego.Debug(typename)

	if typename == "system_notice" {
		c.query_system_notice()
	} else if typename == "server_notice" {
		c.query_server_notice()
	} else if typename == "roll_notice" {
		c.query_roll_notice()
	} else if typename == "insert_mail" {
		c.query_insert_mail()
	} else if typename == "accumulate_recharge" {
		c.query_accumulate_recharge()
	} else if typename == "discount_store" {
		c.query_discount_store()
	} else if typename == "single_recharge" {
		c.query_single_recharge()
	} else if typename == "hotspot_partner" {
		c.query_hotspot_partner()
	} else if typename == "extract_partner" {
		c.query_extract_partner()
	} else if typename == "login_activity" {
		c.query_login_activity()
	} else if typename == "Money_partner" {
		c.query_money_partner()
	}

}
