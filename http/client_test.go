package http

// import (
// 	"github.com/crowleyfelix/base-client-go/errors"
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// type mockInterceptor struct {
// 	err              errors.Error
// 	onRequestCalled  bool
// 	onResponseCalled bool
// }

// func (m *mockInterceptor) OnRequest(req Request) Request {
// 	m.onRequestCalled = true
// 	return req
// }
// func (m *mockInterceptor) OnResponse(resp Response) (Response, errors.Error) {
// 	m.onResponseCalled = true
// 	return resp, m.err
// }

// var _ = Describe("Client", func() {

// 	var (
// 		url          = "url"
// 		interceptors []Interceptor
// 		client          *client
// 	)

// 	BeforeEach(func() {
// 		client = NewClient("url", interceptors...).(*client)
// 	})
// 	Describe("NewClient", func() {
// 		It("should set interceptors", func() {
// 			Expect(client.interceptors).To(Equal(interceptors))
// 		})
// 	})
// 	Describe("Request", func() {

// 		var (
// 			req *Request
// 		)

// 		var (
// 			resp   *Response
// 			err    error
// 		)
// 		var (
// 			externalInterceptor *mockInterceptor
// 			localInterceptor    *mockInterceptor
// 			paramInterceptor    *mockInterceptor
// 		)

// 		//output
// 		var (
// 			actualResp Response
// 			actualErr  errors.Error
// 		)

// 		BeforeEach(func() {
// 			req = new(Request)
// 			err = nil

// 			externalInterceptor = new(mockInterceptor)
// 			localInterceptor = new(mockInterceptor)
// 			paramInterceptor = new(mockInterceptor)

// 			Interceptors = []Interceptor{externalInterceptor}
// 			client.interceptors = []Interceptor{localInterceptor}
// 		})
// 		JustBeforeEach(func() {
// 			actualResp, actualErr = client.Request(method, url, req, paramInterceptor)
// 		})

// 		Context("When Request failled", func() {
// 			BeforeEach(func() {
// 				err = errors.Build(0)
// 			})
// 			It("should call OnRequest of external interceptors", func() {
// 				Expect(externalInterceptor.onRequestCalled).To(BeTrue())
// 				Expect(externalInterceptor.onResponseCalled).To(BeFalse())
// 			})
// 			It("should call OnRequest of local interceptors", func() {
// 				Expect(localInterceptor.onRequestCalled).To(BeTrue())
// 				Expect(localInterceptor.onResponseCalled).To(BeFalse())
// 			})
// 			It("should call OnRequest of parameterized interceptors", func() {
// 				Expect(paramInterceptor.onRequestCalled).To(BeTrue())
// 				Expect(paramInterceptor.onResponseCalled).To(BeFalse())
// 			})
// 			It("should return an error", func() {
// 				Expect(actualResp).To(BeNil())
// 				Expect(actualErr).To(BeAssignableToTypeOf(errors.NewCallout()))
// 			})
// 		})
// 		Context("When Request success", func() {
// 			Context("and external interceptors return an error", func() {
// 				BeforeEach(func() {
// 					externalInterceptor.err = errors.Build(0)
// 				})
// 				It("should return an error", func() {
// 					Expect(actualErr).To(BeAssignableToTypeOf(externalInterceptor.err))
// 				})
// 				It("should not call next interceptors", func() {
// 					Expect(localInterceptor.onResponseCalled).To(BeFalse())
// 					Expect(paramInterceptor.onResponseCalled).To(BeFalse())
// 				})
// 			})
// 			Context("and local interceptors return an error", func() {
// 				BeforeEach(func() {
// 					localInterceptor.err = errors.Build(0)
// 				})
// 				It("should return an error", func() {
// 					Expect(actualErr).To(BeAssignableToTypeOf(localInterceptor.err))
// 				})
// 				It("should not call next interceptors", func() {
// 					Expect(paramInterceptor.onResponseCalled).To(BeFalse())
// 				})
// 			})
// 			Context("and parameterized interceptors return an error", func() {
// 				BeforeEach(func() {
// 					paramInterceptor.err = errors.Build(0)
// 				})
// 				It("should return an error", func() {
// 					Expect(actualErr).To(BeAssignableToTypeOf(paramInterceptor.err))
// 				})
// 			})
// 			Context("and executed with success", func() {
// 				It("should return Response", func() {
// 					Expect(actualErr).To(BeNil())
// 					Expect(actualResp).To(Equal(resp))
// 				})
// 			})
// 		})
// 	})
// 	Describe("allInterceptors", func() {
// 		var all []Interceptor

// 		BeforeEach(func() {
// 			Interceptors = []Interceptor{new(mockInterceptor)}
// 			client.interceptors = []Interceptor{new(mockInterceptor)}
// 			all = client.allInterceptors(new(mockInterceptor))
// 		})
// 		It("should return all available interceptors", func() {
// 			Expect(all).To(HaveLen(3))
// 		})
// 	})
// })
