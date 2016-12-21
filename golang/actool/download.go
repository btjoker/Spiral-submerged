package actool

import (
	"os"
    "io/ioutil"
	"log"
    "strings"
    de "net/url"
)

// DL 下载
func DL(url string) {
    doc, err := Get(url)
    if err != nil {
        log.Println(err)
        return
    }
    defer doc.Body.Close()

    content, _ := ioutil.ReadAll(doc.Body)
    pares, _ := de.Parse(url)
    one := strings.Split(pares.Path, "/")
    name := one[len(one) - 1]

    // 判断path中是否存在 . 号来判断是否是正常文件
    if strings.Count(name, ".") != 1 {
        name+=".html"
    }
    file, _ := os.Create(name)
    defer file.Close()
    file.Write(content)
}