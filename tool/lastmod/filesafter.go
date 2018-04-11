package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func FindFilesAfter(dir string, t time.Time) (paths []string, infos []os.FileInfo, err error) {
	err = filepath.Walk(dir, func(p string, i os.FileInfo, e error) error {
		if err != nil {
			return err
		}

		if !i.IsDir() && i.ModTime().After(t) {
			paths = append(paths, p)
			infos = append(infos, i)
		}
		return nil
	})
	return
}

func checkFile(p string, info os.FileInfo) {
	//fmt.Println(p)
	dir := path.Dir(p)
	dir = strings.Replace(dir, os.Getenv("IG_STORY_DIR"), os.Getenv("IG_INSTAGRAM_DIR"), 1)
	dir = path.Join(dir, "stories")
	//fmt.Println(dir)
	newfilename := strings.Replace(path.Base(p), "-2018", "-story-2018", 1)
	newpath := path.Join(dir, newfilename)
	//fmt.Println(newpath)

	// check if file is the same size
	fn, err := os.Stat(newpath)
	if err != nil {
		panic(p + "\n" + newpath)
	}
	if fn.Size() < info.Size() {
		fmt.Println(p + "\n" + newpath)
	}
}

func main() {
	dir := os.Getenv("IG_STORY_DIR")
	//dir := os.Getenv("IG_INSTAGRAM_DIR")
	t, err := time.Parse("2006-01-02T15:04:05-07:00", "2018-04-07T05:48:03+08:00")
	if err != nil {
		panic(err)
	}
	paths, infos, err := FindFilesAfter(dir, t)
	if err != nil {
		panic(err)
	}
	for i, _ := range paths {
		checkFile(paths[i], infos[i])
	}
}
