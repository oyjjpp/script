package main

import (
	"os"

	"github.com/oyjjpp/script/compress"
)

func main() {
	currentDirName, _ := os.Getwd()
	logDirName := currentDirName + "/log/access.log"
	compress.ReadLineTest(logDirName)
}
