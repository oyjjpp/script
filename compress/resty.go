package compress

import (
	"log"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
)

// 全局公用的transport方便于做长连接和连接池
var transport *http.Transport

func init() {
	transport = createTransport()
}

func createTransport() *http.Transport {
	dialer := &net.Dialer{
		Timeout:   1 * time.Second,
		KeepAlive: 3 * time.Second,
	}
	return &http.Transport{
		DialContext:         dialer.DialContext,
		MaxIdleConnsPerHost: runtime.GOMAXPROCS(0) + 1,
	}
}

func newClient() *http.Client {
	return &http.Client{
		Transport: transport,
	}
}

func restyCurl(link string) {
	client := resty.NewWithClient(newClient())

	// SetHeader("Accept-Encoding", "gzip, deflate, br").
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Accept-Encoding", "br").
		Get(link)

	if err == nil {
		log.Println(resp.RawResponse.ContentLength, resp.Size())
	}
}
