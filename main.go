package main

import (
	"os"

	"github.com/oyjjpp/script/compress"
	"github.com/oyjjpp/script/performance"
)

// var (
// 	ver     int
// 	isRecod string
// )

// func init() {
// 	flag.IntVar(&ver, "ver", 0, "version")
// 	flag.StringVar(&isRecod, "isRecod", "n", "recod log")
// }

func main() {
	// flag.Parse()
	// fmt.Println(ver, isRecod)
	performance.TestPerformance()
}

// 压缩算法性能分析
func testCompress() {
	currentDirName, _ := os.Getwd()
	logDirName := currentDirName + "/log/access.log"
	compress.ReadLineTest(logDirName)
}
