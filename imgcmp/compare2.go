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

func IsInDir2(fds1, fds2 []FileData) {
	mpath := make(map[string]FileData)
	mname := make(map[string]FileData)

	for _, fd := range fds2 {
		if _, ok := mpath[fd.Path]; ok {
			panic("path the same. impossible")
		}
		if _, ok := mname[fd.Info.Name()]; ok {
			fmt.Println(fd.Info.Name())
			panic("name the same. panic")
		}
		mpath[fd.Path] = fd
		mname[fd.Info.Name()] = fd
	}

	count := 0
	for _, fd := range fds1 {
		fdte, ok := mname[fd.Info.Name()]
		if ok {
			if fdte.Info.Size() == fd.Info.Size() && fdte.Info.ModTime() == fd.Info.ModTime() {
				continue
			}
			/*
				file1, err := os.Open(fd.Path)
				if err != nil {
					panic(err)
				}
				file2, err := os.Open(fd.Path)
				if err != nil {
					panic(err)
				}
				defer file1.Close()
				defer file2.Close()
			*/

			count++
			// name the same but size or modtime not the same
			fmt.Println(count, "same name:", fdte.Info.Name(), fdte.Info.Size(), fd.Info.Name(), fd.Info.Size())
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
