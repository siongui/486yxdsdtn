# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

default: fmt
	@echo "\033[92mFormat utime and get filename ...\033[0m"
	@go run utime2rfc3339.go -utime=$(utime) -name=$(name)

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
