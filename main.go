package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

var _currentDirectory = ""

func main() {
	fileName := "0C685917-3058-4E01-8317-A4D78FB189A2.jpg"
	fileStat, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	createDirectory(strconv.Itoa(fileStat.ModTime().Year()))
	createDirectory(fileStat.ModTime().Month().String())
	os.Rename(fileName, path.Join(_currentDirectory, fileName))
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
