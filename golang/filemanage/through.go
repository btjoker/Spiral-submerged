package filemanage

import (
    "os"
    "path/filepath"
	"io/ioutil"
    "strings"
)

// ListDirs 遍历给定路径下的,不包括子目录文件, 后缀如 ".jpg"
func ListDirs(dirPath string, suffix []string) (files []string, err error) {
    files = make([]string,0,200)

    dir, err := ioutil.ReadDir(dirPath) //读取文件目录,为一个列表
    if err != nil {
        return nil, err
    }

    PthSep := string(os.PathSeparator)  //平台文件分隔符
    // suffix = strings.ToUpper(suffix)    //转为大写

    for _, fi := range dir {
        if fi.IsDir(){  // 忽略目录
            continue
        } else if suffixs(fi.Name(), suffix) {
            // 匹配指定后缀文件
            files = append(files, dirPath + PthSep + fi.Name())
        }
    }
    
    return files, nil
}

// WalkDirs 遍历给定路径下的指定后缀的所有文件, 后缀如 ".jpg"
func WalkDirs(dirPath string, suffix []string) (files []string, err error) {
    files = make([]string, 0, 200)

    err = filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error {
        if fi.IsDir() {
            return nil
        } else if suffixs(fi.Name(), suffix) {
            files = append(files, filename)
        }
        return nil
    })
    return files, err
}


// 处理后缀匹配问题, 忽略大小写
func suffixs(fname string, suffix []string) bool {
    for _, str := range suffix {
        if strings.HasSuffix(strings.ToUpper(fname), strings.ToUpper(str)) {
            return true
        }
    }
    return false
}

// ListDir 获取给定目录下所有文件名(不包含子目录文件和目录)
func ListDir(dirPath string) (files []string, err error) {
    files = make([]string,0,200)

    dir, err := ioutil.ReadDir(dirPath)
    if err != nil {
        return nil, err
    }

    PthSep := string(os.PathSeparator)  //平台文件分隔符

    for _, fi := range dir {
        if fi.IsDir(){  // 忽略目录
            continue
        }
        files = append(files, dirPath + PthSep + fi.Name())
    }
    return files, nil
}

// WalkDir 获取给定目录下所有文件名
func WalkDir(dirPath string) (files []string, err error) {
    files = make([]string, 0, 200)

    err = filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error {
        if fi.IsDir() {
            return nil
        }
        files = append(files, filename)
        return nil
    })
    return files, err
}