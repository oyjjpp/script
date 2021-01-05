package main

import (
	"os"

	_ "github.com/oyjjpp/script/basics"
	"github.com/oyjjpp/script/compress"
	"github.com/oyjjpp/script/performance"
)

func main() {
	testPerformance()
}

// 压缩算法性能分析
func testCompress() {
	currentDirName, _ := os.Getwd()
	logDirName := currentDirName + "/log/access.log"
	compress.ReadLineTest(logDirName)
}

func testPerformance() {
	performance.TestPerformance()
}
