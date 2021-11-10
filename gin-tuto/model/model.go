package model

import (
	"errors"
	"reflect"

	"webzen.com/notifyhandler/util"
)

type stNotifyCommon struct {
	From  string `json:"from" binding:"required" example:"from-id"`
	Title string `json:"title" binding:"required" example:"title"`
}

type StNotifyMail struct {
	stNotifyCommon
	To      []string      `json:"to" binding:"required" example:"{yyy@xxxx.co.kr, xxx@fasdf.com}"`
	Cc      []string      `json:"cc,omitempty" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
	Content stmailContent `json:"content"`
}
type stmailContent struct {
	Text string `json:"text" example:"content text in notify mail"`
}

type StNotifyTeams struct {
	stNotifyCommon
	Touri   string         `json:"touri" binding:"required" example:"http://xxx.x.xx.xxx.x."`
	Content stteamsContent `json:"content"`
}
type stteamsContent struct {
	Text string `json:"text" example:"content text in notify teams"`
}

func (s *StNotifyMail) SetFrom(from interface{}) (*StNotifyMail, error) {
	if s != nil {
		switch v := from.(type) {
		case StGrafanaAlert:
			s.Title = v.Title
			s.Content.Text = v.Message
			to, exist := v.Tags["To"]
			if exist {
				s.To = util.StringsToArray(to)
			} else {
				return nil, errors.New("to tag not found in grafana alert struct mail destination requeied to tag")
			}
			cc, exist := v.Tags["Cc"]
			if exist {
				s.Cc = util.StringsToArray(cc)
			}
			return s, nil
		case StUniversalProducerMail:
			s.Title = v.Title
			s.Content.Text = v.Content
			s.From = v.From
			s.To = util.StringsToArray(v.To)
			s.Cc = util.StringsToArray(v.Cc)
			return s, nil
		}
	}
	str := "Can not convert NotifyMail from " + reflect.TypeOf(from).String()

	return nil, errors.New(str)
}

func (s *StNotifyTeams) SetFrom(from interface{}) (*StNotifyTeams, error) {
	if s != nil {
		switch v := from.(type) {
		case StGrafanaAlert:
			s.Title = v.Title
			s.Content.Text = v.Message
			to, exist := v.Tags["Touri"]
			if exist {
				s.Touri = to
			} else {
				return nil, errors.New("to tag not found in grafana alert struct teams destination requeied to tag")
			}
			return s, nil
		case StUniversalProducerTeams:
			s.Title = v.Title
			s.Content.Text = v.Content
			s.From = v.From
			s.Touri = v.Touri
			return s, nil
		}
	}
	str := "Can not convert NotifyTeams from " + reflect.TypeOf(from).String()

	return nil, errors.New(str)
}
