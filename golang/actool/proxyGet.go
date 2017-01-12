package actool

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
)

// ProxyList 西刺代理获取
func ProxyList() []string {
    proxy := make([]string, 101)
    resp, _ := Get("http://www.xicidaili.com/nn/")
    defer resp.Body.Close()
    doc, _ := goquery.NewDocumentFromReader(resp.Body)

    doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
        ip := s.Find("td").Eq(1).Text()
        port := s.Find("td").Eq(2).Text()
        proxy[i] = fmt.Sprintf("http://%s:%s", ip, port)
    })
    return proxy[1:]
}