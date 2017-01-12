package main

import (
    "actool"
    "github.com/goquery"
	"log"
	"os"
    "path/filepath"
	"strconv"
    "sync"
    "strings"
)
const author = "btjoker"

var now, _ = os.Getwd() // 获取当前工作的目录

func get(index int) {
    mk := strconv.Itoa(index) // 转换字符串
    err := os.MkdirAll(`img\` + mk, os.ModePerm) // 先创建文件, 根据返回值绝对是否发起get请求
    if err != nil { 
        // log.Println(err)
        return
    }
    wg := new(sync.WaitGroup)   // 信号量锁
    indu := filepath.Join(now, `img\` + mk)  // 当前进程文件夹名
    
    url := `http://jandan.net/ooxx/page-` + mk
    log.Println(url)
    resp, _ := actool.Get(url)
    defer resp.Body.Close()
    doc, _ := goquery.NewDocumentFromReader(resp.Body)  // 解析html
    os.Chdir(indu)  //跳转到该目录下
    
    imglist := doc.Find("a.view_img_link")
    for i := 0; i < imglist.Length(); i++ {
        img, _ := imglist.Eq(i).Attr("href")
        wg.Add(1)
        go func(u string) {
            defer wg.Done()
            actool.DL(u)
        }(img)
    }
    wg.Wait()
    os.Chdir(now)   //返回运行目录
}

func begin() int {
    resp, _ := actool.Get("http://jandan.net/ooxx")
    defer resp.Body.Close()
    doc, _ := goquery.NewDocumentFromReader(resp.Body)
    next, _ := strconv.Atoi(strings.TrimSpace(doc.Find(".cp-pagenavi a").Eq(0).Text())) //获取页数
    next++
    return next
}

func main() {
    next := begin()
    for ; next > 0; next-- {
        get(next)      
    }
}