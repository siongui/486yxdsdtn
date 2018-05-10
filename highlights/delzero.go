package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func DeleteZeroSizeFile(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if info.Mode().IsRegular() && info.Size() == 0 {
			fmt.Println("Removing", path)
			os.Remove(path)
		}
		return nil
	})
}

func main() {
	DeleteZeroSizeFile("Instagram/")
}
