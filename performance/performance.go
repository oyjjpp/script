package performance

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"runtime/pprof"
	"runtime/trace"
	"strconv"
	"time"
	"unsafe"
)

// 一段常用操作的性能分析
type Performance struct{}

func (p *Performance) V0() {
	data := map[string]int64{}
	for i := int64(0); i < 1000000; i++ {
		value := time.Now().Unix()
		key := strconv.FormatInt(i, 10) + "_" + strconv.FormatInt(value, 10)
		data[key] = value
	}
	fmt.Println("data length = ", len(data))
}

// 1、map初始化的时候指定容量，去除grow带来的耗时
// 2、strconv.FormatInt 改为 strconv.AppendInt,减少对象分配
func (p *Performance) V1() {
	nums := int64(1000000)
	data := make(map[string]int64, nums)
	key := make([]byte, 0)
	for i := int64(0); i < nums; i++ {
		value := time.Now().Unix()
		// 改用appendInt，去掉strconv内部[]byte转string的开销
		key = key[:0]
		key = strconv.AppendInt(key, i, 10)
		key = append(key, '_')
		key = strconv.AppendInt(key, value, 10)
		dataKey := string(key)

		data[dataKey] = value
	}
	fmt.Println("data length = ", len(data))
}

// 1、统一申请key的内存
// 2、直接将[]byte转为string
func (p *Performance) V2() {
	nums := int64(1000000)
	data := make(map[string]int64, nums)

	// 计算key长度，申请存下所有key的[]byte
	keyLen := int64(len(strconv.FormatInt(nums, 10)) + 1 + 10)
	totalLen := keyLen * nums
	key := make([]byte, totalLen)

	for i := int64(0); i < nums; i++ {
		value := time.Now().Unix()

		// 计算当前循环key的位置
		pos := i * keyLen
		b := key[pos:pos]
		b = strconv.AppendInt(b, i, 10)
		b = append(b, '_')
		b = strconv.AppendInt(b, value, 10)
		// 直接将[]byte转为string
		data[*(*string)(unsafe.Pointer(&b))] = value
	}
	fmt.Println("data length = ", len(data))
}

var (
	ver     string
	isRecod string
)

func init() {
	flag.StringVar(&ver, "ver", "0", "version")
	flag.StringVar(&isRecod, "isRecod", "n", "recod log")
}

func TestPerformance() {
	flag.Parse()
	fmt.Println(ver, isRecod)

	object := reflect.ValueOf(&Performance{})
	methodName := "V" + ver
	m := object.MethodByName(methodName)
	if isRecod == "Y" {
		currentDirName, _ := os.Getwd()
		logDirName := currentDirName + "/log/" + methodName + "_trace.out"
		// 收集trace信息
		traceFile, err := os.Create(logDirName)
		if err != nil {
			panic(err.Error())
		}
		err = trace.Start(traceFile)
		if err != nil {
			panic("start trace fail :" + err.Error())
		}
		defer trace.Stop()

		// 收集CPU信息
		logDirName = currentDirName + "/log/" + methodName + "_cpu.out"
		cpuFile, err := os.Create(logDirName)
		if err != nil {
			panic(err.Error())
		}
		defer cpuFile.Close()
		err = pprof.StartCPUProfile(cpuFile)
		if err != nil {
			panic("StartCPUProfile fail :" + err.Error())
		}
		defer pprof.StopCPUProfile()

		// 收集内存信息
		logDirName = currentDirName + "/log/" + methodName + "_mem.out"
		memFile, err := os.Create(logDirName)
		if err != nil {
			panic(err.Error())
		}
		defer pprof.WriteHeapProfile(memFile)
	}
	start := time.Now()
	m.Call(make([]reflect.Value, 0))
	fmt.Println(methodName, time.Now().Sub(start))
}
