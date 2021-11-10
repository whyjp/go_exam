package model

import "webzen.com/notifyhandler/util"

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
	Cc string `json:"cc,omitempty" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
}
type StUniversalProducerTeams struct {
	stUniversalProducer
	Touri string `json:"touri" binding:"required" example:"http://xxx.x.xx.xxx.x."`
}

func (v *StUniversalProducerMail) ToMail() (*StNotifyMail, error) {
	s := new(StNotifyMail)
	s.Title = v.Title
	s.Content.Text = v.Content
	s.From = v.From
	s.To = util.StringsToArray(v.To)
	s.Cc = util.StringsToArray(v.Cc)
	return s, nil
}

func (v *StUniversalProducerTeams) ToTeams() (*StNotifyTeams, error) {
	s := new(StNotifyTeams)
	s.Title = v.Title
	s.Content.Text = v.Content
	s.From = v.From
	s.Touri = v.Touri
	return s, nil
}
