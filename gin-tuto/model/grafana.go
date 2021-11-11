package model

import (
	"errors"

	"webzen.com/notifyhandler/util"
)

type GrafanaAlert struct {
	DashboardID int64       `json:"dashboardId"`
	EvalMatches []evalMatch `json:"evalMatches"`
	ImageURL    string      `json:"imageUrl"`
	Message     string      `json:"message"`
	OrgID       int64       `json:"orgId"`
	PanelID     int64       `json:"panelId"`
	RuleID      int64       `json:"ruleId"`
	RuleName    string      `json:"ruleName"`
	RuleURL     string      `json:"ruleUrl"`
	State       string      `json:"state"`
	GrafanaAlertTags
	Title string `json:"title"`
}

type evalMatch struct {
	Value  int64         `json:"value"`
	Metric string        `json:"metric"`
	Tags   evalMatchTags `json:"tags"`
}

type evalMatchTags struct {
}

type GrafanaAlertTags struct {
	Tags map[string]string `json:"tags"`
}

func (v *GrafanaAlert) ToEMail() (*NotifyEMail, error) {
	s := new(NotifyEMail)
	s.Title = v.Title
	s.Content.Text = v.Message
	to, exist := v.Tags["to"]
	if exist {
		s.To = util.StringsToArray(to)
	} else {
		return nil, errors.New("to tag not found in grafana alert struct mail destination requeied to tag")
	}
	cc, exist := v.Tags["cc"]
	if exist {
		s.Cc = util.StringsToArray(cc)
	}
	return s, nil
}

func (v *GrafanaAlert) ToTeams() (*NotifyTeams, error) {
	s := new(NotifyTeams)
	s.Title = v.Title
	s.Content.Text = v.Message
	to, exist := v.Tags["to"]
	if exist {
		s.To = util.StringsToArray(to)
	} else {
		return nil, errors.New("to tag not found in grafana alert struct teams destination requeied to tag")
	}
	return s, nil
}
