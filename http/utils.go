package http

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"strconv"
)

const tagName = "param"

func SwitchBody(resp Response, data []byte) Response {
	resp.ClearInternalBuffer()
	buf := ioutil.NopCloser(bytes.NewBuffer(data))
	resp.RawResponse().Body = buf

	return resp
}

//ToMap turn objects in maps
func ToMap(obj interface{}) map[string]string {

	mapString := make(map[string]string)
	t := reflect.TypeOf(obj)
	values := reflect.ValueOf(obj)

	if t != nil {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			value := values.Field(i)

			tag := field.Tag.Get(tagName)
			if tag != "" {
				if s, ok := value.Interface().(string); ok {
					mapString[tag] = s
				} else if ps, ok := value.Interface().(*string); ok && ps != nil {
					mapString[tag] = *ps
				} else if i, ok := value.Interface().(int); ok {
					mapString[tag] = strconv.Itoa(i)
				} else if pi, ok := value.Interface().(*int); ok && pi != nil {
					mapString[tag] = strconv.Itoa(*pi)
				}
			}
		}
	}

	return mapString
}
