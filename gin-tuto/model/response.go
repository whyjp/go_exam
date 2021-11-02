package model

type StResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int64  `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}