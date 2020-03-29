package utils

import (
	"encoding/json"
	"log"
)

// RespMsg is general response body
type RespMsg struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data`
}

// NewRespMsg creates a new RespMsg
func NewRespMsg() *RespMsg {
	return &RespMsg{}
}

// JSONBytes converts RespMsg to json in []byte
func (resp *RespMsg) JSONBytes() []byte {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return r
}
