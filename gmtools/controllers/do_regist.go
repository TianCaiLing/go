package controllers

import (
	"github.com/astaxie/beego"
)

type Regist struct {
	beego.Controller
}

func (c *Regist) checkUser(username string, password string) bool {

	db, err := Database()

	if err != nil {
		beego.Error("DB fail", err)
		return false
	}
	defer db.Close()

	//向后台取数据
	records, err := db.Query("SELECT * FROM `gmtoos` WHERE name=?", username)

	if err != nil {
		beego.Error("DB fail", err)
		return false
	}

	db_username := ""
	db_password := ""
	var db_level int
	db_Admin := ""
	db_tel := ""
	db_time := ""
	var db_id int

	for records.Next() {
		err := records.Scan(&db_id, &db_username, &db_password, &db_level, &db_Admin, &db_tel, &db_time)

		if err != nil {
			beego.Error("DB fail", err)
			return false
		}
	}

	beego.Debug(db_username, db_password, db_level)

	if db_level == 0 {
		return true
	}

	return false

}
func (c *Regist) Post() {
	c.Ctx.Request.ParseForm()
	username := c.GetString("username")
	password := c.GetString("password")
	beego.Debug(username, password)
	if c.checkUser(username, password) {
		//验证成功  跳转
		c.Ctx.WriteString(string("ok"))
	}
	///验证失败

}
