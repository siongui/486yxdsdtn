package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

type nameSize struct {
	name string
	size int64
}

// Create directory if it does not exist
func createDirIfNotExist(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return
}

func isDirExist(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func isConfirmed() bool {
	var s string

	fmt.Printf("(y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}

func isExpired(t time.Time) bool {
	return time.Now().Sub(t) > 24*time.Hour
}

func calculateDirSize() (dirsize int64) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			if isExpired(file.ModTime()) {
				dirsize += file.Size()
			}
		}
	}
	return
}

func moveMaxSizeUser(ns nameSize) {
	fmt.Println("move", ns.name, "?")
	if !isConfirmed() {
		return
	}

	movedUserDir := path.Join(os.Getenv("IG_WORKSPACE_DIR"), ns.name)
	err := createDirIfNotExist(movedUserDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	userpostsdir := path.Join(os.Getenv("IG_INSTAGRAM_DIR"), ns.name, "posts")
	if isDirExist(userpostsdir) {
		fmt.Println(userpostsdir, "exists")

		newpath := path.Join(os.Getenv("IG_WORKSPACE_DIR"), ns.name, "posts")
		fmt.Println("moving ", userpostsdir, " to ", newpath, " ...")
		err := os.Rename(userpostsdir, newpath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = os.Chdir(ns.name)
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	targetdir := path.Join(os.Getenv("IG_WORKSPACE_DIR"), ns.name, "story")
	err = createDirIfNotExist(targetdir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			newpath := path.Join(targetdir, file.Name())
			fmt.Println("moving ", file.Name(), " to ", newpath, " ...")
			err := os.Rename(file.Name(), newpath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Fprintf(os.Stdout, "https://www.instagram.com/%s/\n", ns.name)
	return
}

func main() {
	err := os.Chdir(os.Getenv("IG_STORY_DIR"))
	if err != nil {
		fmt.Println(err)
		return
	}

	dirs, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	var nameSizeArray []nameSize
	for _, dir := range dirs {
		err = os.Chdir(dir.Name())
		if err != nil {
			fmt.Println(err)
			return
		}

		size := calculateDirSize()
		//size2, err := calculateUserDirSize(dir.Name())
		//if err == nil {
		//	size += size2
		//}

		ns := nameSize{dir.Name(), size}
		nameSizeArray = append(nameSizeArray, ns)

		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	sort.Slice(nameSizeArray, func(i, j int) bool {
		return nameSizeArray[i].size < nameSizeArray[j].size
	})

	for _, ns := range nameSizeArray {
		fmt.Println(ns)
	}

	moveMaxSizeUser(nameSizeArray[len(nameSizeArray)-1])
}
