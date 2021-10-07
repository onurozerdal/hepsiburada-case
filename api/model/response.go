package model

type Response struct {
	UserId   string   `json:"user-id"`
	Products []string `json:"products"`
	Type     string   `json:"type"`
}
