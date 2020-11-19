package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	//	"time"
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

	for _, fd := range fds1 {
		if _, ok := mname[fd.Info.Name()]; !ok {
			fmt.Println(fd.Info.Name() + " not in dir2")
			moveFile("../tobedone/", fd.Path, fd.Info)
			continue
		}
		fdorig := mname[fd.Info.Name()]
		/*
			if fd.Info.ModTime().Unix() != fdorig.Info.ModTime().Unix() {
				fmt.Println(fd.Info.Name())
			}
		*/
		if fd.Info.Size() > fdorig.Info.Size() {
			fmt.Println(fd.Info.Name(), fd.Info.Size(), fdorig.Info.Size())
			moveFile("../tobedone/", fd.Path, fd.Info)
		}

		if fd.Info.ModTime().Unix() != fdorig.Info.ModTime().Unix() {
			//if fd.Info.ModTime().Format(time.UnixDate) != fdorig.Info.ModTime().Format(time.UnixDate) {
			//if fmt.Sprint(fd.Info.ModTime()) != fmt.Sprint(fdorig.Info.ModTime()) {
			//fmt.Println(fd.Info.Name(), fd.Info.ModTime().Unix(), fdorig.Info.ModTime().Unix())
			//fmt.Println(fd.Info.Name(), fd.Info.ModTime(), fdorig.Info.ModTime())
			fmt.Println(fd.Info.Name(), fd.Info.ModTime().Unix()-fdorig.Info.ModTime().Unix())
			moveFile("../tobedone2/", fd.Path, fd.Info)
		}
	}
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
