package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
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

func isExpired(t time.Time) bool {
	return time.Now().Sub(t) > 24*time.Hour
}

func deleteExpired(filepaths []string) {
	for _, filepath := range filepaths {
		info, err := os.Stat(filepath)
		if err != nil {
			panic(err)
		}
		if isExpired(info.ModTime()) {
			fmt.Println("remove", filepath)
			err = os.Remove(filepath)
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	path := flag.String("path", "uploaded.txt", "path to uploaded.txt")
	flag.Parse()

	lines, err := FileToLines(*path)
	if err != nil {
		panic(err)
	}

	deleteExpired(lines)
}
