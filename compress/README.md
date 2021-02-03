# 压缩算法脚本测试

## 使用

```golang

// 压缩算法性能分析
func testCompress() {
 currentDirName, _ := os.Getwd()
 logDirName := currentDirName + "/log/access.log"
 compress.ReadLineTest(logDirName)
}
```
