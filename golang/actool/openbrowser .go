package actool

import "os/exec"


var list = []string{
    `Program Files\Mozilla Firefox\firefox.exe`,
    `Program Files\Google\Chrome\Application\chrome.exe`,
    `sougou\SogouExplorer\SogouExplorer.exe`,
    `360安全浏览器\360se6\Application\360se.exe`,
}

var drives = []string{
    `C:\`,
    `D:\`,
    `E:\`,
    `F:\`,
}

// Run 传入一个url, 然后自动打开浏览器跳转到该地址, 阻塞性质
func Run(url string) {
    for _, drive := range drives {
        for _, path := range list {
            _, err := exec.LookPath(drive + path)
            if err == nil {
                exec.Command(drive + path, url).Run()
                return
            }
        }
    }
    exec.Command(`C:\Program Files\Internet Explorer\iexplore.exe`, url).Run()
}

// Start 传入一个url, 然后自动打开浏览器跳转到该地址, 非阻塞性质
func Start(url string) {
    for _, drive := range drives {
        for _, path := range list {
            _, err := exec.LookPath(drive + path)
            if err == nil {
                exec.Command(drive + path, url).Start()
                return
            }
        }
    }
}
