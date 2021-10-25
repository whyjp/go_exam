package models

type StJsonTest struct {
	First  string `form:"first" json:"first" xml:"first" binding:"required"`
	Second string `form:"second" json:"second" xml:"second" binding:"required"`
}

type StMailTest struct {
	From    string `form:"from" json:"from" xml:"from" binding:"required"`
	To      string `form:"to" json:"to" xml:"to" binding:"required"`
	Title   string `form:"title" json:"title" xml:"title" binding:"required"`
	Content string `form:"content" json:"content" xml:"content" binding:"required"`
}
