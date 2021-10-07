package model

type ProductView struct {
	Event   string `json:"event"`
	MessageId   string `json:"messageid"`
	UserId    string    `json:"userid"`
	Properties *Properties `json:"properties"`
	Context *Context `json:"context"`
}