.PHONY: test-carvel-release
test-carvel-release:
	go clean -testcache
	SETUP_LOCAL_REGISTRY=true SETUP_KIND_CLUSTER=true TEST_CARVEL_RELEASE=true go test ./... -run TestReleaseOfCarvelPackage

.PHONY: test-carvel-release-info
test-carvel-release-info:
	go clean -testcache
	CMD_LOG_LEVEL=1 SETUP_LOCAL_REGISTRY=true SETUP_KIND_CLUSTER=true TEST_CARVEL_RELEASE=true go test ./... -v -run TestReleaseOfCarvelPackage

.PHONY: test-carvel-release-debug
test-carvel-release-debug:
	go clean -testcache
	CMD_LOG_LEVEL=2 SETUP_LOCAL_REGISTRY=true SETUP_KIND_CLUSTER=true TEST_CARVEL_RELEASE=true go test ./... -v -run TestReleaseOfCarvelPackage