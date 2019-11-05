package db

import "PrometheusAlert/model"

//查询alert history
func GetAlertList() (alertList []*model.Prometheus, err error) {

	sqlstr := `select id, description  from alert`
	err = DB.Select(&alertList, sqlstr)
	return
}

//新增alert record
func InsertAlertRecord(alertList *model.Prometheus) (err error) {

	sqlstr := `insert into alert(username,password) values(?,?)`
	_,err = DB.Exec(sqlstr,alertList.Externalurl,alertList.Status)
	return
}