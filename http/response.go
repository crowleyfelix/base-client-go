package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/crowleyfelix/base-client-go/errors"
)

//Response is the struct wrapper of http response struct
type Response struct {
	*http.Response
}

//Ok verifies if status code received represents a success
func (r *Response) Ok() bool {
	if r.Response.StatusCode < 200 || r.Response.StatusCode >= 400 {
		return false
	}
	return true
}

//JSON unmarshals the body into a struct
func (r *Response) JSON(obj interface{}) errors.Error {
	blob, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return errors.NewCallout(err.Error())
	}

	err = json.Unmarshal(blob, obj)
	if err != nil {
		return errors.NewSerializing(err.Error())
	}
	return nil
}

//Text returns de plain text of the body
func (r *Response) Text() (*string, errors.Error) {
	blob, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, errors.NewCallout(err.Error())
	}

	text := string(blob)

	return &text, nil
}
