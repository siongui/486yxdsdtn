# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

default: fmt
	@echo "\033[92mRun tool...\033[0m"
	@go run dlusers.go -path=$(IG_BACKUP_DIR)/dlusers.txt

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go

install:
	go get -u github.com/siongui/instago/download
