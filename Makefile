APP=bpm-counter
ALL_PACKAGES=$(shell go list ./...)

APP_EXECUTABLE="./out/$(APP)"

fmt:
	GO111MODULE=on go fmt $(ALL_PACKAGES)

imports:
	GO111MODULE=on goimports -w -local github.com ./

dep:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

clean:
	GO111MODULE=on go clean
	rm -rf out/

compile: clean
	mkdir -p out
	GO111MODULE=on go build -o $(APP_EXECUTABLE)

run: compile
	$(APP_EXECUTABLE)

install-gotest:
	GO111MODULE=off go get github.com/rakyll/gotest

test: install-gotest
	GO111MODULE=on gotest -p=1 -mod=readonly $(UNIT_TEST_PACKAGES)

test-cov: install-gotest
	GO111MODULE=on gotest -p=1 -mod=readonly -covermode=count -coverprofile=out/coverage.cov $(UNIT_TEST_PACKAGES)
