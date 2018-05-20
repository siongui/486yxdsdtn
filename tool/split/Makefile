# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


split: fmt
	@echo "\033[92mSplit dir...\033[0m"
	@go run splitdir.go -max=$(max) -src=$(IG_WORKSPACE_DIR)/$(name)/posts

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
