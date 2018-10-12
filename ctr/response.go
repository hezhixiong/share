package ctr

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func WriteJson(w http.ResponseWriter, code int, data interface{}) error {
	if data == nil {
		data = make(map[string]interface{})
	}

	resp := Response{
		Code:      code,
		Msg:       getCodeMsg(code),
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
	retJson, err := json.Marshal(&resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(retJson); err != nil {
		return err
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func getCodeMsg(code int) string {
	if msg, ok := codeMsg[code]; ok {
		return msg
	} else {
		return "undefined error code"
	}
}
