package model

type StRule struct {
}

type stSirentRuleCommon struct {
	Producer string            `json:"producer" binding:"required" example:"wiss or wingo or kiss or more"`
	Tags     map[string]string `json:"tag" binding:"required" example:"MU,SOGB,"`
	Region   string            `json:"region" example:"empty(Don't Care All) or KR, GP, JP, SEA"`
}

type StSirentRule struct {
	stSirentRuleCommon
	Start string `json:"start"`
	End   string `json:"end"`
}

type StSirentRuleRoutine struct {
	stSirentRuleCommon
	Starttime string `json:"start" binding:"required" example:"15:00 ex) use 24H"`
	Endtime   string `json:"end" binding:"required" example:"15:00 ex) use 24H"`
	Weekly    bool   `json:"weekly" default:"false" example:"true or false"`
	Day       string `json:"day" example:"Mon or TUE ..."`
	Date      int    `json:"date" example:"15"`
}
type StSirentRuleOnce struct {
	stSirentRuleCommon
	StartDatetime string `json:"start:" binding:"required" example:"2021-11-04 15:00:02 * use 24H"`
	EndDatetime   string `json:"endtime" binding:"required" example:"2021-11-04 15:00:02 * use 24H"`
}
