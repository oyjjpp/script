package compress

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	DOMAIN_NAME = ""
	DEBUG_INFO  = ""
)

// curl执行命令
func curl(link string) (float64, float64) {
	command := `curl -H "Accept-Encoding:deflate,gzip,br" -o /dev/null -s -w "%{time_total} %{size_download} %{http_code} \t\n" `
	command = command + `"` + link + `"`
	cmd := exec.Command("/usr/bin/sh", "-c", command)
	if out, err := cmd.Output(); err == nil {
		rs := strings.Split(string(out), " ")
		data1, _ := strconv.ParseFloat(rs[0], 64)
		data2, _ := strconv.ParseFloat(rs[1], 64)
		return data1, data2
	}
	return 0, 0
}

func ReadLineTest(filePth string) {
	f, err := os.Open(filePth)
	if err != nil {
		return
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)

	for {
		line, err := bfRd.ReadString('\n')
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			continue
		}
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		var data AccessLog
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			continue
		}
		requestUrl := strings.Split(data.Request, " ")
		if len(requestUrl) < 1 {
			continue
		}
		link := DOMAIN_NAME + requestUrl[1]
		if data.RequestMethod == "POST" {
			link = link + "&" + data.PostData
		}
		link = link + DEBUG_INFO
		log.Println(link)
	}
}

// ReadLine
// 按行读取文件
// @param filePth 文件名称
func ReadLine(filePth string) {
	f, err := os.Open(filePth)
	if err != nil {
		return
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)

	var timeSum, byteSum float64

	for {
		line, err := bfRd.ReadString('\n')
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			continue
		}
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		var data AccessLog
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			continue
		}
		requestUrl := strings.Split(data.Request, " ")
		if len(requestUrl) < 1 {
			continue
		}
		link := DOMAIN_NAME + requestUrl[1]
		if data.RequestMethod == "POST" {
			link = link + "&" + data.PostData
		}
		link = link + DEBUG_INFO
		curTime, curByte := curl(link)
		timeSum = timeSum + curTime
		byteSum = byteSum + curByte/1024
	}
	log.Println(timeSum, byteSum)
}

// 自定义nginx日志
type AccessLog struct {
	RemoteAddr        string `json:"remote_addr"`
	Timestamp         string `json:"timestamp"`
	RemoteUser        string `json:"remote_user"`
	BodyBytesSent     int    `json:"body_bytes_sent"`
	BytesSent         int    `json:"bytes_sent"`
	RequestTime       string `json:"request_time"`
	Status            int    `json:"status"`
	Request           string `json:"request"`
	RequestPath       string `json:"request_path"`
	RequestMethod     string `json:"request_method"`
	PostData          string `json:"post_data"`
	HTTPReferrer      string `json:"http_referrer"`
	HTTPXForwardedFor string `json:"http_x_forwarded_for"`
	HTTPUserAgent     string `json:"http_user_agent"`
	HTTPXAllocateIdc  string `json:"http_x_allocate_idc"`
	HTTPXScheme       string `json:"http_x_scheme"`
	ServerName        string `json:"server_name"`
	Extend            string `json:"extend"`
}
