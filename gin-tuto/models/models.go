package models

type StJsonTest struct {
	First  string `json:"first" binding:"required" example:"first-exam"`
	Second string `json:"second" binding:"required" example:"second-exam"`
}

type STNotiCommon struct {
	From  string   `json:"from" binding:"required" example:"from-id"`
	To    []string `json:"to" binding:"required" example:"to-destination"`
	Title string   `json:"title" binding:"required" example:"title"`
}
type STNotiMail struct {
	STNotiCommon
	Content ContentMail `json:"content"`
}
type ContentMail struct {
	Text string `json:"text" example:"content text in notify mail"`
}
