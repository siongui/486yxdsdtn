# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


DIR=lemoine

test: fmt
	cd $(DIR)
	@echo "\033[92mRun testing code...\033[0m"
	@go test -v $(DIR)/*.go

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt $(DIR)/*.go
