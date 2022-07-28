package errors

import (
	"encoding/json"
	"net/http"
)

const contentTypeJson = "application/json"

type jsonResponse struct {
	Data  interface{}      `json:"data"`
	Error *jsonResponseErr `json:"error"`
}

type jsonResponseErr struct {
	Text string `json:"text"`
	Code uint32 `json:"code"`
}

func Response(data interface{}, err error, httpStatuses map[uint32]map[string]interface{}) (int, string, []byte) {
	var errResp *jsonResponseErr

	status := http.StatusOK

	if err != nil {
		var text string

		code, _ := GetMark(err)

		text, status = getMarkData(code, httpStatuses)

		errResp = &jsonResponseErr{
			Text: text,
			Code: code,
		}

		data = nil
	}

	resp := jsonResponse{
		Data:  data,
		Error: errResp,
	}

	body, err := json.Marshal(resp)
	if err != nil {
		return http.StatusInternalServerError, contentTypeJson, []byte(`{"data": null, "error": {"text": "failed encode data", "code": -1}}`)
	}

	return status, contentTypeJson, body
}

func Json(w http.ResponseWriter, data interface{}, err error, httpStatuses map[uint32]map[string]interface{}) {
	status, contentType, body := Response(data, err, httpStatuses)

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	w.Write(body)
}
