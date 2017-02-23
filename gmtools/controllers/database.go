package controllers

import (
	"database/sql"
	_ "encoding/json"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "strconv"
	_ "strings"
	_ "time"
	_ "unsafe"
)

func Database() (*sql.DB, error) {
	dsn := beego.AppConfig.String("dbuser") + ":" + beego.AppConfig.String("dbpass") + "@tcp(" + beego.AppConfig.String("dbhost") + ":" + beego.AppConfig.String("dbport") + ")/" + beego.AppConfig.String("dbname")
	beego.Debug(dsn)
	return sql.Open("mysql", dsn)
}

///公告相关记录

//系统公告添加
func SystemNoticeInsert(title string, content string, color string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `notice`(types,Title,Content,Color,time) VALUES(?,?,?,?,?)", "system_notice", title, content, color, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}

//区服公告添加
func ServiceNoticeInsert(title string, content string, color string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `Service`(types,title,content,color,time) VALUES(?,?,?,?,?)", "2", title, content, color, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}

//滚动公告
func RollNoticeInsert(sendType string, content string, timestr string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `Roll`(types,sendType,content,timestr,time) VALUES(?,?,?,?,?)", "3", sendType, content, timestr, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}

//历史邮件
func InsertMailInsert(Title string, Sender string, Content string, recvers string, stritemids string, stritemsks string, LowLevel int, HighLevel int, time0 string, time1 string, SendType string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `Mail`(Title,Sender,Content,recvers,stritemids,stritemsks,LowLevel,HighLevel,time0,time1,SendType,time) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)", Title, Sender, Content, recvers, stritemids, stritemsks, LowLevel, HighLevel, time0, time1, SendType, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}

//登录活动
func LoginActivityInsert(OpenTime string, CloseTime string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `Login`(name,OpenTime,CloseTime,time) VALUES(?,?,?,?)", "登录活动", OpenTime, CloseTime, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}

//累积充值
func AccumulateRechargeInsert(OpenTime string, CloseTime string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `Accumulate`(name,OpenTime,CloseTime,time) VALUES(?,?,?,?)", "累积充值", OpenTime, CloseTime, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}

//打折商店
func DiscountStoreInsert(OpenTime string, CloseTime string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `Discount`(name,OpenTime,CloseTime,time) VALUES(?,?,?,?)", "打折商店", OpenTime, CloseTime, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}

//单笔充值
func SingleRechargeInsert(OpenTime string, CloseTime string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `Single`(name,OpenTime,CloseTime,time) VALUES(?,?,?,?)", "单笔充值", OpenTime, CloseTime, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}

//热点伙伴
func HotspotPartnerInsert(OpenTime string, CloseTime string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `hotspot`(name,OpenTime,CloseTime,time) VALUES(?,?,?,?)", "热点伙伴", OpenTime, CloseTime, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}

//抽伙伴活动
func ExtractPartnerInsert(OpenTime string, CloseTime string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `extract`(name,OpenTime,CloseTime,time) VALUES(?,?,?,?)", "顶级招募", OpenTime, CloseTime, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}

//小额礼包
func MoneyRechargeInsert(OpenTime string, CloseTime string, time string) {
	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `Money`(name,OpenTime,CloseTime,time) VALUES(?,?,?,?)", "小额礼包", OpenTime, CloseTime, time)

	if err != nil {
		beego.Error("DB fail", err)
		return
	}
}
