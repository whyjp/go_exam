package model

type sTNotifyCommon struct {
	From  string   `json:"from" binding:"required" example:"from-id"`
	To    []string `json:"to" binding:"required" example:"to-destination"`
	Title string   `json:"title" binding:"required" example:"title"`
}
type sTNotifyCommonEx struct {
	Producer  string   `json:"producer" binding:"required" example:"wiss or wingo or kiss or more"`
	Tag       []string `json:"tag" binding:"required" example:"MU,SOGB,"`
	Region    string   `json:"region" example:"empty(Don't Care) or KR/GP/JP/SEA"`
	ForceSend bool     `json:"ignorefilter" default:"false" example:"bool = ture,false 알림 필터를 무시 force send"`
	sTNotifyCommon
}
type STNotifyMail struct {
	sTNotifyCommon
	Content mailContent `json:"content"`
}
type mailContent struct {
	Text string `json:"text" example:"content text in notify mail"`
}

type STNotifyTeams struct {
	sTNotifyCommon
	Content teamsContent `json:"content"`
}
type teamsContent struct {
	Text string `json:"text" example:"content text in notify teams"`
}
