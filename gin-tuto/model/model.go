package model

type stNotifyCommon struct {
	From  string   `json:"from" binding:"required" example:"from-id"`
	To    []string `json:"to" binding:"required" example:"to-destination"`
	Title string   `json:"title" binding:"required" example:"title"`
}
type stNotifyCommonEx struct {
	Producer  string   `json:"producer" binding:"required" example:"wiss or wingo or kiss or more"`
	Tag       []string `json:"tag" binding:"required" example:"MU,SOGB,"`
	Region    string   `json:"region" example:"empty(Don't Care) or KR/GP/JP/SEA"`
	ForceSend bool     `json:"ignorefilter" default:"false" example:"bool = ture,false 알림 필터를 무시 force send"`
	stNotifyCommon
}
type StNotifyMail struct {
	stNotifyCommon
	Content stmailContent `json:"content"`
}
type stmailContent struct {
	Text string `json:"text" example:"content text in notify mail"`
}

type StNotifyTeams struct {
	stNotifyCommon
	Content stteamsContent `json:"content"`
}
type stteamsContent struct {
	Text string `json:"text" example:"content text in notify teams"`
}
