ENVFILE=.env
HOOKSPATH=hooks
LINTFILES=$(shell git diff-index --cached --name-only HEAD | xargs printf -- '--include=%s\n')

.PHONY: dep setup test coverage mocks .env env

dep:
	@go get -t -v ./... \
	github.com/onsi/ginkgo/ginkgo \
	github.com/onsi/gomega/...  \
	github.com/axw/gocov/gocov \
	github.com/vektra/mockery/.../ \
	github.com/alecthomas/gometalinter

setup: dep
	@cp -f $(HOOKSPATH)/** .git/hooks
	@find . -name $(ENVFILE) | grep -q "." || cp $(ENVFILE).example $(ENVFILE)	

check: setup
	@gometalinter ./... --fast $(LINTFILES)

deepcheck: setup
	@gometalinter ./... $(LINTFILES)

fullcheck: setup
	gometalinter ./...

test: setup
	@ginkgo -gcflags=-l ./...	

integrationtest: setup
	@ginkgo -gcflags=-l -tags=integration ./...

coverage: setup
	@gocov test ./... | gocov report

mocks: setup
	@scripts/build_mocks
