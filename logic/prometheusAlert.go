package logic

import (
	"PrometheusAlert/model"
	"PrometheusAlert/dao/db"
)
//查询alert history
func GetAlertList() (alertList []*model.Prometheus, err error) {
	alertList, err = db.GetAlertList()
	if err != nil {
		return nil, err
	}
	return
}

//新增alert history
func InsertAlertListRecord(alertList *model.Prometheus) (err error) {
	err = db.InsertAlertRecord(alertList)
	if err != nil {
		return err
	}
	return
}