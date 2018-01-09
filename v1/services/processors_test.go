package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Processors", func() {
	Describe("processResponse", func() {
		Context("When response is not ok", func() {
			It("should track error", func() {

			})
		})
		Context("When response is ok", func() {
			Context("and failed on deserialize response", func() {
				It("should return serialization error", func() {

				})
			})
			Context("and success on deserialize response", func() {
				It("should switch response body for inner data", func() {

				})
			})
		})
	})
})
