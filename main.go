package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fileInfo, err := ioutil.ReadDir(exPath)

	if err != nil {
		panic(err)
	}

	var fileName string
	var fileExt string
	var isPng bool
	for _, file := range fileInfo {
		fileName = file.Name()
		fileExt = filepath.Ext(fileName)
		isPng = fileExt == ".png"

		if isPng && (strings.Contains(fileName, "_index_") || strings.Contains(fileName, "frame_")) {
			err := os.Remove(fileName)
			if err != nil {
				panic(err)
			}
		}
	}
}
