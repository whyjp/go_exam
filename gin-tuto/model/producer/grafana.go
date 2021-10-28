package model

type GrafanaAlert struct {
	DashboardID int64            `json:"dashboardId"`
	EvalMatches []EvalMatch      `json:"evalMatches"`
	ImageURL    string           `json:"imageUrl"`
	Message     string           `json:"message"`
	OrgID       int64            `json:"orgId"`
	PanelID     int64            `json:"panelId"`
	RuleID      int64            `json:"ruleId"`
	RuleName    string           `json:"ruleName"`
	RuleURL     string           `json:"ruleUrl"`
	State       string           `json:"state"`
	Tags        GrafanaAlertTags `json:"tags"`
	Title       string           `json:"title"`
}

type EvalMatch struct {
	Value  int64         `json:"value"`
	Metric string        `json:"metric"`
	Tags   EvalMatchTags `json:"tags"`
}

type EvalMatchTags struct {
}

type GrafanaAlertTags struct {
	TagName string `json:"tag name"`
}
