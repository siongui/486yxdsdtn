package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func GetExistingPaths(paths []string) (eps []string) {
	for _, path := range paths {
		// check if file exist
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// file not exists
			fmt.Println(path, "not exist")
		} else {
			eps = append(eps, path)
		}
	}
	return
}

func WriteLinesToFile(lines []string, filename string) (err error) {
	return ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}

func main() {
	path := flag.String("path", "uploaded.txt", "path to uploaded.txt")
	out := flag.String("out", "uploaded-new.txt", "path to new uploaded.txt")
	flag.Parse()
	lines, err := FileToLines(*path)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = WriteLinesToFile(GetExistingPaths(lines), *out)
	if err == nil {
		fmt.Println("new uploaded.txt created!")
	} else {
		fmt.Println(err)
		return
	}
}
