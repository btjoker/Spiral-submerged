package actool

import (
	"os"
    "io/ioutil"
	"log"
    "strings"
    "path/filepath"
)

// DL 下载功能, 第二个可选参数为一个文件夹名称
func DL(url string, pth ...string) {
    resp, err := Get(url)
    defer resp.Body.Close()
    if err != nil {
        log.Println(err)
        return
    }

    var file *os.File
    var fname string
    dir, _ := os.Getwd()
    tmp := strings.Split(resp.Request.URL.String(), "/")
    
    if len(pth) == 0 {
        fname = filepath.Join(dir, tmp[len(tmp) - 1])
    } else {
        dirs := filepath.Join(dir, pth[0])
        os.Mkdir(dirs, os.ModePerm)
        fname = filepath.Join(dir, pth[0], tmp[len(tmp) - 1])
    }
    if Exist(fname) {
        log.Println("已存在同名文件!")
        return
    }

    file, _ = os.Create(fname)
    content, _ := ioutil.ReadAll(resp.Body)
        
    defer file.Close()
    file.Write(content)
    log.Println(fname, "Done!")
}

// Exist 判断文件是否存在
func Exist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil || os.IsExist(err)
}