package actool

import (
    "log"
    "net/http"
    "net/url"
)

// Get 带随机UA的get请求
func Get(url string) (*http.Response, error) {
    client := new(http.Client)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    req.Header.Add(GetUA())
    resp, err := client.Do(req)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    return resp, nil
}

// DlGet 代理发送get请求
func DlGet(u, d string ) (*http.Response, error) {
    proxy := func(_ *http.Request) (*url.URL, error) {
        return url.Parse(d)
    }
    transport := &http.Transport{Proxy: proxy}
    client := &http.Client{Transport: transport}
    resp, err := client.Get(u)
    return resp, err
}