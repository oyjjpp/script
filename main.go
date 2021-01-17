package main

import (
	"log"
	"os"
	"runtime"
	"sync"

	_ "github.com/oyjjpp/script/basics"
	"github.com/oyjjpp/script/compress"
	"github.com/oyjjpp/script/performance"
)

func main() {

	syncConcurrency()
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

// syncConcurrency
// 通过WaitGroup实现并发控制
func syncConcurrency() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
		}(i)
	}
	wg.Wait()
	// time.Sleep(5 * time.Second)
	log.Println("finished ", runtime.NumGoroutine())
}
