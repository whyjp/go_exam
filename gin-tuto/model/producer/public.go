package model

type stPublicProducer struct {
	Producer string `json:"producer" binding:"required" example:"wiss or wingo or kiss or more"`
	From     string `json:"from" binding:"required" example:""`
	To       string `json:"to" binding:"required" example:""`
	Title    string `json:"title" binding:"required" example:""`
	Content  string `json:"content" binding:"required" example:""`
}

type StPublicProducerMail struct {
	stPublicProducer
	Cc string `json:"cc" binding:"required" example:""`
}
type StPublicProducerTeams struct {
	stPublicProducer
}