package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// Create directory if it does not exist
func createDirIfNotExist(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return
}

func isExist(dirorfilepath string) bool {
	if _, err := os.Stat(dirorfilepath); os.IsNotExist(err) {
		return false
	}
	return true
}

func moveFileToDir(filepath, dir string) (err error) {
	filename := path.Base(filepath)
	dstpath := path.Join(dir, filename)
	fmt.Println("move", filepath, "to", dstpath)
	err = os.Rename(filepath, dstpath)
	return
}

func moveUserFiles(user UserStat) (err error) {
	dstUserDir := path.Join(os.Getenv("IG_WORKSPACE_DIR"), user.Name)
	dstStoryDir := path.Join(dstUserDir, "stories")
	dstPostDir := path.Join(dstUserDir, "posts")
	dstPostliveDir := path.Join(dstUserDir, "postlives")

	// move old story
	oldStoryDir := path.Join(os.Getenv("IG_STORY_DIR"), user.Name)
	if isExist(oldStoryDir) {
		createDirIfNotExist(dstStoryDir)
		files, err := ioutil.ReadDir(oldStoryDir)
		if err != nil {
			return err
		}
		for _, file := range files {
			if !file.IsDir() {
				filepath := path.Join(oldStoryDir, file.Name())
				err = moveFileToDir(filepath, dstStoryDir)
				if err != nil {
					return err
				}
			}
		}
	}

	// move story
	storyDir := path.Join(os.Getenv("IG_INSTAGRAM_DIR"), user.Name, "stories")
	if isExist(storyDir) {
		createDirIfNotExist(dstStoryDir)
		files, err := ioutil.ReadDir(storyDir)
		if err != nil {
			return err
		}
		for _, file := range files {
			if !file.IsDir() && isExpired(file.ModTime()) {
				filepath := path.Join(storyDir, file.Name())
				err = moveFileToDir(filepath, dstStoryDir)
				if err != nil {
					return err
				}
			}
		}
	}

	// move post
	postDir := path.Join(os.Getenv("IG_INSTAGRAM_DIR"), user.Name, "posts")
	if isExist(postDir) {
		createDirIfNotExist(dstPostDir)
		files, err := ioutil.ReadDir(postDir)
		if err != nil {
			return err
		}
		for _, file := range files {
			if !file.IsDir() {
				filepath := path.Join(postDir, file.Name())
				err = moveFileToDir(filepath, dstPostDir)
				if err != nil {
					return err
				}
			}
		}
	}

	// move postlive
	postliveDir := path.Join(os.Getenv("IG_INSTAGRAM_DIR"), user.Name, "postlives")
	if isExist(postliveDir) {
		createDirIfNotExist(dstPostliveDir)
		files, err := ioutil.ReadDir(postliveDir)
		if err != nil {
			return err
		}
		for _, file := range files {
			if !file.IsDir() && isExpired(file.ModTime()) {
				filepath := path.Join(postliveDir, file.Name())
				err = moveFileToDir(filepath, dstPostliveDir)
				if err != nil {
					return err
				}
			}
		}
	}

	fmt.Println("https://www.instagram.com/" + user.Name + "/")
	return
}
