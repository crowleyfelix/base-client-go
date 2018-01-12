.PHONY: dep setup test coverage mocks 

dep:
	@go get -t -v ./... \
	github.com/onsi/ginkgo/ginkgo \
	github.com/onsi/gomega/...  \
	github.com/axw/gocov/gocov \
	github.com/vektra/mockery/.../ \
	github.com/alecthomas/gometalinter

setup: dep

check: setup
	@gometalinter ./... || true 

test: setup
	@ginkgo -gcflags=-l ./...	

integrationtest: setup
	@ginkgo -gcflags=-l -tags=integration ./...

coverage: setup
	@gocov test ./... | gocov report

mocks: setup
	scripts/build_mocks
