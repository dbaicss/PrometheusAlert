package model


type Labels struct{
	Alertname string `json:"alertname"`
	Instance string `json:"instance"`
}
type Annotations struct{
	Description string `json:"description"`
	Summary string `json:"summary"`
	Level string `json:"level"`  //2019年2月15日 19:03:07 增加告警级别
	Mobile string `json:"mobile"` //2019年2月25日 19:09:23 增加手机号支持
	Ddurl string `json:"ddurl"` //2019年3月12日 20:33:38 增加多个钉钉告警支持
}

type Alerts struct {
	Labels Labels `json:"labels"`
	Annotations Annotations `json:"annotations"`
	StartsAt string `json:"startsAt"`
	EndsAt string `json:"endsAt"`
	GeneratorUrl string `json:"generatorURL"` //prometheus 告警返回地址
}
type Prometheus struct {
	Status string
	Alerts []Alerts
	Externalurl string `json:"externalURL"` //alertmanage 返回地址
}
