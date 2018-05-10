# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


upload:
	gphotosuploader --upload Instagram/ --maxConcurrent 1

delzero:
	@echo "\033[92mDelete zero-size files in directory ...\033[0m"
	@go run delzero.go

highlight: fmt
	@echo "\033[92mDownload story highlights of all following users ...\033[0m"
	@go run example/download.go -downloadtype=highlight

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go

install:
	#go get github.com/simonedegiacomi/gphotosuploader
	go get -u github.com/siongui/instago/download
