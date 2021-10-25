package models

type StJsonTest struct {
	First  string `form:"first" json:"first" xml:"first" binding:"required"`
	Second string `form:"second" json:"second" xml:"second" binding:"required"`
}
