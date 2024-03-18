package models

type Record struct {
	ID     string      `json:"id"`
	Data   interface{} `json:"data"`
	Action string      `json:"action"`
}
