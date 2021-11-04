package model

type stNotifyCommon struct {
	From  string `json:"from" binding:"required" example:"from-id"`
	Title string `json:"title" binding:"required" example:"title"`
}

type StNotifyMail struct {
	stNotifyCommon
	To      string        `json:"to" binding:"required" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
	Cc      string        `json:"cc" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
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
