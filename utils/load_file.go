package utils

import (
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func LoadFile(filepath string) string {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		panic("Could not find Caller of util.ReadFile")
	}

	absolutePath := path.Join(path.Dir(filename), filepath)

	fileContent, err := os.ReadFile(absolutePath)

	if err != nil {
		log.Fatal(err)
	}

	fileStringContent := string(fileContent)

	return strings.TrimRight(fileStringContent, "\n")
}
