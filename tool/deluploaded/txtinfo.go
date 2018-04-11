package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UploadedFileToLookupMap(filePath string) (m map[string]bool, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	m = make(map[string]bool)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		m[scanner.Text()] = true
	}
	err = scanner.Err()
	return
}

func PrintUploadedInfo(typ string, lookupMap map[string]bool) {
	if !(typ == "post" || typ == "story" || typ == "postlive") {
		return
	}

	var size int64 = 0
	count := 0
	for filepath, _ := range lookupMap {
		info, err := os.Stat(filepath)
		if err != nil {
			panic(err)
		}
		if isExpired(info.ModTime()) && strings.Contains(info.Name(), "-"+typ+"-") {
			size += info.Size()
			//fmt.Println(filepath)
			count++
		}
	}
	fmt.Println("total size of all uploaded, expired", typ, "files:", size/(1024*1024), "MB")
	fmt.Println("total number of all uploaded, expired", typ, "files:", count)
}
