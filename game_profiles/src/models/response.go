package models

type Response struct {
	Err  bool        `json:"err"`
	Data interface{} `json:"data"`
}
