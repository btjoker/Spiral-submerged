# -*- coding: utf-8 -*-

import urllib.request
import sys

url = sys.argv[1]
html = urllib.request.urlretrieve(url, url[url.rfind('/')+1:])

print('DownLoad succeed!!')
