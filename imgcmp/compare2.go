package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	//	"github.com/corona10/goimagehash"
)

type FileData struct {
	Path string
	Info os.FileInfo
}

func moveFile(todir, path string, info os.FileInfo) {
	//fmt.Println(info.Name())
	//return

	newpath := filepath.Join(todir, info.Name())
	fmt.Printf("move %q to %q\n", path, newpath)
	err := os.Rename(path, newpath)
	if err != nil {
		panic(err)
	}
}

func AllFileData(dir string) (fds []FileData, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, e)
			return e
		}
		if info.Mode().IsRegular() {
			fds = append(fds, FileData{Path: path, Info: info})
		}

		return nil
	})
	return
}

func Print2FileData(fi1, fi2 FileData) {
	fmt.Print(fi1.Info.Name(), " ", fi1.Info.Size(), fi1.Info.ModTime())
	fmt.Print(", ")
	fmt.Print(fi2.Info.Name(), " ", fi2.Info.Size(), fi2.Info.ModTime())
	fmt.Println()
}

func IsSameName(fi1, fi2 FileData) bool {
	return fi1.Info.Name() == fi2.Info.Name()
}

func IsSameSize(fi1, fi2 FileData) bool {
	return fi1.Info.Size() == fi2.Info.Size()
}

func IsSameModTime(fi1, fi2 FileData) bool {
	return fi1.Info.ModTime() == fi2.Info.ModTime()
}

func IsToKeep(fi1, fi2 FileData) bool {
	return IsSameName(fi1, fi2) && IsSameSize(fi1, fi2) && IsSameModTime(fi1, fi2)
}

func IsToMerge(fi1, fi2 FileData) bool {
	return IsSameName(fi1, fi2) && !IsSameSize(fi1, fi2) && IsSameModTime(fi1, fi2)
}

func IsInDir2(fds1, fds2 []FileData) {
	count := 0
	for _, fd1 := range fds1 {
		for _, fd2 := range fds2 {
			if IsToKeep(fd1, fd2) {
				count++
				fmt.Print(count, ": ")
				Print2FileData(fd1, fd2)
				break
			}
		}

	}
	fmt.Println("len fds1", len(fds1), "len fds2", len(fds2))
}

func main() {
	dir1 := flag.String("dir1", "dir1", "dir 1 of photos and videos")
	dir2 := flag.String("dir2", "dir2", "dir 1 of photos and videos")
	flag.Parse()

	fds1, err := AllFileData(*dir1)
	if err != nil {
		panic(err)
	}
	fds2, err := AllFileData(*dir2)
	if err != nil {
		panic(err)
	}

	IsInDir2(fds1, fds2)
}
