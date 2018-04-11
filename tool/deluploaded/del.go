package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
)

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

func main() {
	path := flag.String("path", "uploaded.txt", "path to uploaded.txt")
	flag.Parse()
	lookupMap, err := UploadedFileToLookupMap(*path)
	if err != nil {
		fmt.Println(err)
		return
	}

	users, err := getAllUserStat()
	if err != nil {
		fmt.Println(err)
		return
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].ExpiredSize < users[j].ExpiredSize
	})
	for _, user := range users {
		fmt.Println(user.Name, user.ExpiredCount, user.ExpiredSize, user.Count, user.Size)
	}

	maxUser := users[len(users)-1]
	fmt.Println("move", maxUser.Name, "?")
	if isConfirmed() {
		moveUserFiles(maxUser)
	}

	PrintUploadedInfo("story", lookupMap)
}
