package helper

import (
	"encoding/json"
	"net/http"
)

func WriteStatuMsgResp(w http.ResponseWriter, msg []string, status string) error {
	rspdata := struct {
		Status string
		Msg    []string `json:"msgs"`
	}{
		Status: status,
		Msg:    msg,
	}

	return WriteJsonResp(w, rspdata)
}

func WriteJsonResp(w http.ResponseWriter, resp interface{}) error {
	blobresp, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	w.Write(blobresp)

	return nil
}
