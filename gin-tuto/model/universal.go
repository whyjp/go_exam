package model

type stUniversalProducer struct {
	Producer string            `json:"producer" binding:"required" example:"wiss or wingo or kiss or more"`
	From     string            `json:"from" binding:"required" example:""`
	Title    string            `json:"title" binding:"required" example:""`
	Content  string            `json:"content" binding:"required" example:""`
	Tags     map[string]string `json:"tags" example:"game:#MUA2, region:KR"`
}

type StUniversalProducerMail struct {
	stUniversalProducer
	To string `json:"to" binding:"required" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
	Cc string `json:"cc" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
}
type StUniversalProducerTeams struct {
	stUniversalProducer
	Touri string `json:"touri" binding:"required" example:"http://xxx.x.xx.xxx.x."`
}
