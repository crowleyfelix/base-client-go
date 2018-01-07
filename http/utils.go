package http

import (
	"bytes"
	"io/ioutil"
)

func SwitchBody(resp Response, data []byte) Response {
	resp.ClearInternalBuffer()
	buf := ioutil.NopCloser(bytes.NewBuffer(data))
	resp.RawResponse().Body = buf

	return resp
}
