# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

post:
	@echo "\033[92mDownload single post ...\033[0m"
	@go fmt singlepost.go
	@go run singlepost.go

install:
	go get -u github.com/siongui/instago/download
