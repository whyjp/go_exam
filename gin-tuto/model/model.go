package model

type STNotiCommon struct {
	From  string   `json:"from" binding:"required" example:"from-id"`
	To    []string `json:"to" binding:"required" example:"to-destination"`
	Title string   `json:"title" binding:"required" example:"title"`
}
type STNotiCommonEx struct {
	NotifyProducer string   `json:"notioproducer" binding:"required" example:"wiss or wingo or kiss or more"`
	Tag            []string `json:"tag" binding:"required" example:"MU,SOGB,"`
	ForceSend      bool     `json:"ignorefilter" example:"bool = ture,false 알림 필터를 무시 force send"`
	STNotiCommon
}
type STNotiMail struct {
	STNotiCommon
	Content MailContent `json:"content"`
}
type MailContent struct {
	Text string `json:"text" example:"content text in notify mail"`
}

type STNotiTeams struct {
	STNotiCommon
	Content TeamsContent `json:"content"`
}
type TeamsContent struct {
	Text string `json:"text" example:"content text in notify teams"`
}
