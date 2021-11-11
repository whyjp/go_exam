package model

type notifyCommon struct {
	From     string   `json:"from" binding:"required" example:"from-id"`
	To       []string `json:"to" binding:"required" example:"{yyy@xxxx.co.kr, xxx@fasdf.com}"`
	Title    string   `json:"title" binding:"required" example:"title"`
	ImageURL string   `json:"image_url,omitempty" example:"internal image url"`
}

type NotifyEMail struct {
	notifyCommon
	Cc      []string       `json:"cc,omitempty" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
	Bcc     []string       `json:"bcc,omitempty" example:"xxx@yyyy.com;yyy@xxxx.co.kr"`
	Content stEmailContent `json:"content"`
}
type stEmailContent struct {
	Text string `json:"text" example:"content text in notify Email"`
}

type NotifyTeams struct {
	notifyCommon
	Content stteamsContent `json:"content"`
}
type stteamsContent struct {
	Text string `json:"text" example:"content text in notify teams"`
}
