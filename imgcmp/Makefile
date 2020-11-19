export GOROOT=$(realpath ../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

compare: fmt
	go run compare2.go -dir1="../../dir1" -dir2="../../dir2"

fmt:
	go fmt *.go

# https://stackoverflow.com/questions/32680834/how-to-compare-images-with-go
# https://github.com/vitali-fedulov/images
# https://github.com/corona10/goimagehash
# https://stackoverflow.com/questions/843972/image-comparison-fast-algorithm
install:
	go get -u github.com/corona10/goimagehash
