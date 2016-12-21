package filemanage

import (
    "os"
    "io"
	"path/filepath"
    "strings"
	"fmt"
)

// Copy Copy(目标路径, 源文件))
func Copy(dstPath, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil { return }
    defer src.Close()

    dstName := filenameG(dstPath, srcName)
    
    dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Printf("File: %s Copy Error!", dstName)
        return
    }
    defer dst.Close()
    return io.Copy(dst, src)
}

// Move Move(目标路径, 源文件)
func Move(dstPath, srcName string) (err error) {
    _, err = Copy(dstPath, srcName)
    if err != nil {
        fmt.Println()
        return
    }
    err = os.Remove(srcName)
    if err != nil {
        fmt.Println("原文件占用中,无法删除!")
    }
    return
}


func filenameG(dstPath, srcPath string) (path string) {
    _, file := filepath.Split(srcPath)
    if strings.HasSuffix(dstPath, "\\") {
        path = dstPath + file
        return
    }
    path = (dstPath + "\\") + file
    return
}