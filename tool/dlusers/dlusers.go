package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/siongui/instago"
	"github.com/siongui/instago/download"
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

func main() {
	path := flag.String("path", "dlusers.txt", "path to downloaded users txt file")
	flag.Parse()

	usernames, err := FileToLines(*path)
	if err != nil {
		panic(err)
	}

	mgr := instago.NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))

	todonames := []string{}
	for _, username := range todonames {
		igdl.DownloadUserProfilePicUrlHd(username)
		igdl.DownloadAllPosts(username, mgr)
		usernames = append(usernames, username)
	}

	sort.Strings(usernames)
	for i, username := range usernames {
		fmt.Println(i, username)
	}
	WriteLinesToFile(usernames, "dlusers-new.txt")
}
