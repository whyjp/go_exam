package model

import "webzen.com/notifyhandler/util"

type uuniversalProducer struct {
	Producer string            `json:"producer" binding:"required" example:"swagger"`
	From     string            `json:"from" binding:"required" example:"swagger webapp"`
	Title    string            `json:"title" binding:"required" example:"swagger webapp test send notify"`
	Content  string            `json:"content" binding:"required" example:"swagger webapp test send notify\nhi\nhello\ngoodbye\nim swagger notify handler"`
	Tags     map[string]string `json:"tags" example:"game:MUA2, region:KR"`
	ImageURL string            `json:"image_url,omitempty" example:"http://internal.image.url/imgname.jpg"`
}

type UniversalProducerEMail struct {
	uuniversalProducer
	To  string `json:"to" binding:"required" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
	Cc  string `json:"cc,omitempty" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
	Bcc string `json:"bcc,omitempty" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
}
type UniversalProducerTeams struct {
	uuniversalProducer
	Touri string `json:"touri" binding:"required" example:"http://xxx.x.xx.xxx.x;http://xxx.x.xx.xxx.x"`
}

func (v *UniversalProducerEMail) ToEMail() (*NotifyEMail, error) {
	s := new(NotifyEMail)
	s.Title = v.Title
	s.Content.Text = v.Content
	s.From = v.From
	s.To = util.StringsToArray(v.To)
	s.Cc = util.StringsToArray(v.Cc)
	s.Bcc = util.StringsToArray(v.Bcc)
	return s, nil
}

func (v *UniversalProducerTeams) ToTeams() (*NotifyTeams, error) {
	s := new(NotifyTeams)
	s.Title = v.Title
	s.Content.Text = v.Content
	s.From = v.From
	s.To = util.StringsToArray(v.Touri)
	return s, nil
}
