package model

import (
	"errors"

	"webzen.com/notifyhandler/util"
)

type StGrafanaAlert struct {
	DashboardID int64         `json:"dashboardId"`
	EvalMatches []stEvalMatch `json:"evalMatches"`
	ImageURL    string        `json:"imageUrl"`
	Message     string        `json:"message"`
	OrgID       int64         `json:"orgId"`
	PanelID     int64         `json:"panelId"`
	RuleID      int64         `json:"ruleId"`
	RuleName    string        `json:"ruleName"`
	RuleURL     string        `json:"ruleUrl"`
	State       string        `json:"state"`
	stGrafanaAlertTags
	Title string `json:"title"`
}

type stEvalMatch struct {
	Value  int64           `json:"value"`
	Metric string          `json:"metric"`
	Tags   stEvalMatchTags `json:"tags"`
}

type stEvalMatchTags struct {
}

type stGrafanaAlertTags struct {
	Tags map[string]string `json:"tags"`
}

func (v *StGrafanaAlert) ToEMail() (*StNotifyEMail, error) {
	s := new(StNotifyEMail)
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
}

func (v *StGrafanaAlert) ToTeams() (*StNotifyTeams, error) {
	s := new(StNotifyTeams)
	s.Title = v.Title
	s.Content.Text = v.Message
	to, exist := v.Tags["Touri"]
	if exist {
		s.Touri = to
	} else {
		return nil, errors.New("to tag not found in grafana alert struct teams destination requeied to tag")
	}
	return s, nil
}
