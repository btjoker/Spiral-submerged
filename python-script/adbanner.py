# -*- coding: utf-8 -*-

import urllib.request
import threading
import time
from os.path import exists
from os import mkdir

_Author = 'btjoker'
# 重复次数计数
ccount = 0
# 错误计数
ecount = 0
# 下载完成计数
rcount = 0
url = 'http://cover.acfunwiki.org/cover.php'
dirs = r'D:\ac'
if not exists(dirs):
    mkdir(dirs)
    dirs = 'D:\\ac\\'


def download(url):
    global ccount
    global ecount
    global rcount
    try:
        content = urllib.request.urlopen(url)
        filename = content.url.split('/')[-1]
        if not exists(dirs + filename):
            with open(dirs + filename, 'wb') as f:
                f.write(content.read())
                rcount += 1
            print('%s done!!!' % filename)
        else:
            ccount += 1
    except:
        ecount += 1
    finally:
        time.sleep(1)


if __name__ == '__main__':
    while 1:
        for i in range(4):
            t = threading.Thread(target=download, args=(url,))
            t.start()
        if count >= 500:
            break
    else:
        print('本次下载：\n\t完成: %s\n\t失败: %s' % (rcount, ecount))
