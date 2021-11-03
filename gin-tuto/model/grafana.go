package model

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
