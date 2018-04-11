package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

var maxFilesPerDir = 1500

func getTargetDir(groupCount int, srcDir string) string {
	return path.Join(srcDir, "dir"+strconv.Itoa(groupCount))
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func moveFilesToDir(filesGroup [][]os.FileInfo, srcDir string) {
	for i, group := range filesGroup {
		targetDir := getTargetDir(i, srcDir)
		createDirIfNotExist(targetDir)
		for _, file := range group {
			filepath := path.Join(srcDir, file.Name())
			targetFilepath := path.Join(targetDir, file.Name())
			fmt.Println("move", filepath, "to", targetFilepath)
			os.Rename(filepath, targetFilepath)
		}
	}
}

func makeFilesGroup(srcDir string) [][]os.FileInfo {
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		panic(err)
	}

	count := 0
	var filesGroup [][]os.FileInfo
	filesGroup = append(filesGroup, []os.FileInfo{})
	groupCount := 0
	for _, file := range files {
		if file.Mode().IsRegular() {
			filesGroup[groupCount] = append(filesGroup[groupCount], file)
			count++
			if count == maxFilesPerDir {
				count = 0
				filesGroup = append(filesGroup, []os.FileInfo{})
				groupCount++
			}
		}
	}

	return filesGroup
}

func main() {
	srcDir := flag.String("src", "", "source dir to split")
	max := flag.Int("max", 1500, "max number of files in split dir")
	flag.Parse()

	maxFilesPerDir = *max
	filesGroup := makeFilesGroup(*srcDir)
	moveFilesToDir(filesGroup, *srcDir)
}
