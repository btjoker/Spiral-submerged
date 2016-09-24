# -*- coding: utf-8 -*-

_Ahore = btjoker

import urllib.request, threading, time, os
from lxml import etree

gurl = 'http://aosabook.org/en/500L/'
dirspath = r'D:\moe\pyxss\dm'

# 创建一个用量控制器，阈值为5
semaphore = threading.BoundedSemaphore(5)

# 过滤列表
dirs = os.listdir(dirspath)
invfile = ['archives.html', 'authors.html', 'categories.html', 'tags.html']

# 抓取500line网页页面中的所有hrel, 读取本地500line中的文件名与比较,获取本地不存在的链接
with urllib.request.urlopen(gurl) as f:
	html = etree.HTML(f.read())
	if dirs:
		poi = [gurl + str(x) for x in html.xpath('//@href') if x.endswith('.html') and str(x) not in invfile]
		htmlurl = [gurl+x for x in poi if gurl+x not in dirs]
	else:
		htmlurl = [gurl + str(x) for x in html.xpath('//@href') if x.endswith('.html') and str(x) not in invfile]

# 保存文件
def xget(url):
	semaphore.acquire()
	time.sleep(1)
	urllib.request.urlretrieve(url, os.path.join(dirspath,url[url.rfind('/')+1:]))
	print('%s Download Success!!' % url)
	semaphore.release()

for i in htmlurl:
	t = threading.Thread(target=xget, args=(i,))
	t.start()