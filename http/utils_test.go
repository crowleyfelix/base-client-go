package http

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {

	PDescribe("SwitchBody", func() {

	})

	Describe("ToMap", func() {
		obj := struct {
			FieldString string `param:"field_string"`
			FieldInt    int    `param:"field_int"`
		}{"value", 1}

		It("should turn object into map", func() {
			Expect(ToMap(obj)).To(Equal(map[string]string{
				"field_string": "value",
				"field_int":    "1",
			}))
		})
	})
})
