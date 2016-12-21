package filemanage

import (
	"fmt"
    "sync"
)


//DirCopy 复制该目录下的所有文件(不包括子目录和目录)
func DirCopy(dstPath, srcPath string) (err error) {
    filelist, err := ListDir(srcPath)
    if err != nil { return }

    for _, fl := range filelist {
        _, err := Copy(dstPath, fl)
        if err != nil {
            fmt.Printf("File: %s Copy Error!", fl)
        }
    }
    return nil
}

// DirDeepCopy 复制该目录下的所有文件(包括子目录中的文件)
func DirDeepCopy(dstPath, srcPath string) (err error) {
    filelist, err := WalkDir(srcPath)
    if err != nil { return }

    for _, fl := range filelist {
        _, err := Copy(dstPath, fl)
        if err != nil {
            fmt.Printf("File: %s Copy Error!", fl)
        }
    }
    return nil
}

// Moves 接收一个文件列表, 将其剪切到指定目录下
func Moves(dstPath string, filelist []string ) (err error) {
    wg := new(sync.WaitGroup)
    for _, fi := range filelist {
        wg.Add(1)
        go func(file string) {
            defer wg.Done()
            Move(dstPath, file)
            fmt.Printf("File: %s Done.\n", file)
        }(fi)
    }
    wg.Wait()
    return nil
}