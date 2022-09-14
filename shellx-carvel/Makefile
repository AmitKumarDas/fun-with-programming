GO_TEST_ARGS ?=

.PHONY: test
test:
	go clean -testcache
	go test ./... ${GO_TEST_ARGS}
