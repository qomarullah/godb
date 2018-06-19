package controllers

import (
	"godb/models"
	"strings"

	"github.com/astaxie/beego"
)

// Operations about query
type QueryController struct {
	beego.Controller
}
type Resp struct {
	status string
	count  int
}

// @Title Query Select
// @Description query from config
// @Param	ds		query 	string	true		"Datasource from config"
// @Param	sqlid		query 	string	true		"SQL id predefine from config"
// @Success 200 {string} success
// @Failure 403 data not found
// @router /select [get]
func (q *QueryController) Select() {
	_ds := q.GetString("ds")
	_sqlid := q.GetString("sqlid")
	_id := q.GetString("id")

	//ds := "apps:aplikasi@tcp(localhost:3306)/test"
	ds := beego.AppConfig.String("ds." + _ds)
	sqlid := beego.AppConfig.String("sqlid." + _sqlid)
	query := strings.Replace(sqlid, "[id]", _id, -1)

	var resp map[string]interface{}
	if resp == nil {
		resp = make(map[string]interface{})
	}
	resp["count"] = 0
	resp["desc"] = "-"
	resp["success"] = false
	var xdata []map[string]string
	resp["data"] = xdata

	if ds == "" || sqlid == "" {
		resp["desc"] = "config not found"
		q.Data["json"] = resp
		q.ServeJSON()
		return
	}

	//fmt.Println(ds, sqlid, query)

	//connect
	db, err := models.Connect(ds)
	if err != nil {
		resp["desc"] = err
		q.Data["json"] = resp
		q.ServeJSON()
		return

	}
	//query
	data, err := models.Select(db, query)
	if err != nil {
		resp["desc"] = err
		q.Data["json"] = resp
		q.ServeJSON()
		return

	}
	resp["data"] = data
	resp["success"] = true
	q.Data["json"] = resp
	beego.Info(resp)
	q.ServeJSON()

}
