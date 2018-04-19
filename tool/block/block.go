package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/siongui/instago"
)

func FileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func WriteLinesToFile(lines []string, filename string) (err error) {
	return ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}

func saveCurrentFollowing(mgr *instago.IGApiManager) (err error) {
	users, err := mgr.GetFollowing(mgr.GetSelfId())
	if err != nil {
		return
	}

	var lines []string
	for _, user := range users {
		line := user.Username + "   " + strconv.FormatInt(user.Pk, 10)
		lines = append(lines, line)
	}

	filename := "following-" + strconv.FormatInt(time.Now().Unix(), 10) + ".txt"

	return WriteLinesToFile(lines, filename)
}

func findLastFileStartsWith(dir, prefix string) (lastFile os.FileInfo, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range files {
		if !file.Mode().IsRegular() {
			continue
		}
		if strings.HasPrefix(file.Name(), prefix) {
			if lastFile == nil {
				lastFile = file
			} else {
				if lastFile.ModTime().Before(file.ModTime()) {
					lastFile = file
				}
			}
		}
	}

	if lastFile == nil {
		err = os.ErrNotExist
		return
	}
	return
}

func main() {
	mgr := instago.NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))

	err := saveCurrentFollowing(mgr)
	if err != nil {
		panic(err)
	}

	file, err := findLastFileStartsWith(".", "following")
	if err != nil {
		panic(err)
	}
	fmt.Println(file.Name())
}
