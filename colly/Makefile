# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


ig:
	go run src/github.com/gocolly/colly/_examples/instagram/instagram.go $(name)

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go

install:
	@echo "\033[92mInstall colly (Go) ...\033[0m"
	go get -u github.com/gocolly/colly
