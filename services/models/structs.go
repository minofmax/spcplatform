package models

type ResponseBody struct {
	Status uint        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}
