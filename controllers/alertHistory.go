package controllers

import (
	"PrometheusAlert/logic"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)

type AlertHistoryController struct {
	beego.Controller
}


//alert history接口
func (c *AlertHistoryController) PrometheusAlertList() {
	alert,err := logic.GetAlertList()
	log.Println("err:",err)
	res,err := json.Marshal(alert)
	c.Data["json"]= string(res)
	log.Println(c.Data["json"])
	c.ServeJSON()
}







