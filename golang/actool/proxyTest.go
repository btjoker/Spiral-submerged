package actool

import (
    "net/http"
    "net/url"
	"log"
)

// ProxyText 测试
func ProxyText(data string) {
    proxy := func(_ *http.Request) (*url.URL, error) {
        return url.Parse(data)
    }
    transport := &http.Transport{Proxy: proxy}
    client := &http.Client{Transport: transport}
    _, err := client.Get("http://qq.com/")
    if err != nil {
        log.Println(err)
        return
    }
    log.Println(data)
}