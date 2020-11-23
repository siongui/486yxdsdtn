package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/corona10/goimagehash"
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

func removeFile(name string) {
	fmt.Printf("remove %q\n", name)
	err := os.Remove(name)
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

	// not used
	if IsSameName(fi1, fi2) && !IsSameSize(fi1, fi2) && IsSameModTime(fi1, fi2) {
		isSame, err := IsSamePhoto(fi1.Path, fi2.Path)
		if err != nil {
			panic(err)
		}
		return isSame
	}
	return false
}

func KeepFirstDeleteSecond(fkeep, fdelete FileData) {
	moveFile("../../keep/", fkeep.Path, fkeep.Info)
	removeFile(fdelete.Path)
}

func MergeTwoFile(fi1, fi2 FileData) {
	if fi1.Info.Size() > fi2.Info.Size() {
		KeepFirstDeleteSecond(fi1, fi2)
	} else {
		KeepFirstDeleteSecond(fi2, fi1)
	}
}

func IsSamePhoto(p1, p2 string) (isSame bool, err error) {
	f1, err := os.Open(p1)
	if err != nil {
		return
	}
	defer f1.Close()
	f2, err := os.Open(p2)
	if err != nil {
		return
	}
	defer f2.Close()

	img1, err := jpeg.Decode(f1)
	if err != nil {
		return
	}
	img2, err := jpeg.Decode(f2)
	if err != nil {
		return
	}

	width, height := 16, 16
	h1, err := goimagehash.ExtPerceptionHash(img1, width, height)
	if err != nil {
		return
	}
	h2, err := goimagehash.ExtPerceptionHash(img2, width, height)
	if err != nil {
		return
	}

	d, err := h1.Distance(h2)
	if err != nil {
		return
	}
	if d == 0 {
		isSame = true
	}
	return
}

func IsInDir2(fds1, fds2 []FileData) {
	count := 0
	for _, fd1 := range fds1 {
		for _, fd2 := range fds2 {
			if IsToMerge(fd1, fd2) {
				count++
				fmt.Print(count, ": ")
				Print2FileData(fd1, fd2)
				//MergeTwoFile(fd1, fd2)
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
