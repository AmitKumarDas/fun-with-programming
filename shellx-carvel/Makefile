.PHONY: test-carvel-release-e2e
test-carvel-release-e2e:
	go clean -testcache
	SETUP_LOCAL_REGISTRY=true SETUP_KIND_CLUSTER=true TEST_CARVEL_RELEASE=true go test ./... -run TestCarvelReleaseE2E

.PHONY: test-carvel-release-e2e-info
test-carvel-release-e2e-info:
	go clean -testcache
	CMD_LOG_LEVEL=1 SETUP_LOCAL_REGISTRY=true SETUP_KIND_CLUSTER=true TEST_CARVEL_RELEASE=true go test ./... -v -run TestCarvelReleaseE2E

.PHONY: test-carvel-release-e2e-debug
test-carvel-release-e2e-debug:
	go clean -testcache
	CMD_LOG_LEVEL=2 SETUP_LOCAL_REGISTRY=true SETUP_KIND_CLUSTER=true TEST_CARVEL_RELEASE=true go test ./... -v -run TestCarvelReleaseE2E

.PHONY: carvel-publish
carvel-publish:
	go clean -testcache
	go test ./... -run TestCarvelPublish

.PHONY: unit-test
unit-test:
	go clean -testcache
	go test ./internal/...

.PHONY: go-tidy
go-tidy:
	go mod tidy

.PHONY: go-fmt
go-fmt:
	go fmt ./...

.PHONY: go-vet
go-vet:
	go vet ./...

