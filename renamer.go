package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		log.Fatal("too many arguments")
		usage()
	}

	path := "."

	if len(args) == 1 {
		path = args[0]
	}

	info, err := os.Lstat(path)
	if os.IsNotExist(err) {
		log.Fatalf("file or path does not exist: %s", path)
	}

	rename(filepath.Clean(path), info)
}

func usage() {
	fmt.Printf("renamer [path]")
	fmt.Printf("rename all folders and files in [path], to a clean form")
}

// rename rename path to clean name, if path is directory, recursively
// rename all files in it
// path: make sure it exists and is pre-processed by filepath.Clean()
func rename(path string, info os.FileInfo) {
	dir, file := filepath.Split(path)

	if isHidden(file) {
		log.Printf("skip hidden file: %s", file)
		return
	}

	if newFile, changed := cleanName(file); changed {
		newPath := filepath.Join(dir, newFile)
		err := os.Rename(path, newPath)
		if err != nil {
			log.Printf("failed to rename file %s ==> %s, %v", path, newPath, err)
		} else {
			log.Printf("rename %s ==> %s", path, newPath)
			path = newPath
		}
	} else {
		log.Printf("ok: %s", path)
	}

	if !info.IsDir() {
		return
	}

	f, err := os.Open(path)
	if err != nil {
		log.Printf("failed to read file: %s, skip it", path)
		return
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		log.Printf("failed to list files in dir %s", path)
		return
	}

	for _, name := range names {
		filename := filepath.Join(path, name)
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			log.Printf("failed to lstat file: %s", filename)
			continue
		}

		rename(filename, fileInfo)
	}

}

func isHidden(path string) bool {
	return len(path) > 1 && strings.HasPrefix(path, ".")
}

func cleanName(path string) (string, bool) {
	newPath, changed := path, false

	if strings.ContainsAny(path, "-_[]") {
		newPath = strings.Replace(path, "-", ".", -1)
		newPath = strings.Replace(newPath, "_", ".", -1)
		newPath = strings.Replace(newPath, "[", ".", -1)
		newPath = strings.Replace(newPath, "]", ".", -1)
		changed = true
	}

	return newPath, changed
}
