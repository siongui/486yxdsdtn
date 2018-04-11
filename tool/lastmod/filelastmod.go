package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func FindLastModifiedFileBefore(dir string, t time.Time) (path string, info os.FileInfo, err error) {
	isFirst := true
	min := 0 * time.Second
	err = filepath.Walk(dir, func(p string, i os.FileInfo, e error) error {
		if err != nil {
			return err
		}

		if !i.IsDir() && i.ModTime().Before(t) {
			if isFirst {
				isFirst = false
				path = p
				info = i
				min = t.Sub(i.ModTime())
			}
			if diff := t.Sub(i.ModTime()); diff < min {
				path = p
				min = diff
				info = i
			}
		}
		return nil
	})
	return
}

func main() {
	dir := os.Getenv("IG_STORY_DIR")
	//dir := os.Getenv("IG_INSTAGRAM_DIR")
	path, info, err := FindLastModifiedFileBefore(dir, time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	fmt.Println(info)
}
