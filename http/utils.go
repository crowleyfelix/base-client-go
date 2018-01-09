package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

func SwitchBody(resp Response, data []byte) Response {
	resp.ClearInternalBuffer()
	buf := ioutil.NopCloser(bytes.NewBuffer(data))
	resp.RawResponse().Body = buf

	return resp
}

func ToMap(obj interface{}) map[string]string {
	var mapString map[string]string

	data, _ := json.Marshal(obj)
	json.Unmarshal(data, &mapString)
	return mapString
}
