package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

var _currentDirectory = ""
var _workingDirectory = ""

func main() {
	argsFlagParser()
	checkFiles()

	fileName := "0C685917-3058-4E01-8317-A4D78FB189A2.jpg"
	fileStat, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	createDirectory(strconv.Itoa(fileStat.ModTime().Year()))
	createDirectory(fileStat.ModTime().Month().String())
	os.Rename(fileName, path.Join(_currentDirectory, fileName))
}

func checkFiles() {
	var files []string
	err := filepath.Walk(_workingDirectory, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}

func createDirectory(folderName string) {
	_currentDirectory = path.Join(_currentDirectory, folderName)
	fmt.Println(_currentDirectory)
	_, err := os.Stat(_currentDirectory)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(_currentDirectory, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

func argsFlagParser() {
	_workingDirectory := flag.String("folderPath", "", "path of the place that photos includes")
}
