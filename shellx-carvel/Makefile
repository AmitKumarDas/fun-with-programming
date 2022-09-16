GO_TEST_ARGS ?=

.PHONY: test-carvel-release
test-carvel-release:
	go clean -testcache
	SETUP_LOCAL_REGISTRY=true go test ./... ${GO_TEST_ARGS} -run TestReleaseOfCarvelPackage
