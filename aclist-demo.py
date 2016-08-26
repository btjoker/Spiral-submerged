# -*- coding: utf-8 -*-

import requests
import re

def biu(num='1'):
	if num == '1':
		url = 'http://www.acfun.tv/v/list110/index.htm'
	else:
		url = 'http://www.acfun.tv/v/list110/index_'+num +'.htm'

	html = requests.get(url).text

	showre = re.compile('<span class="hint">(.*?)</span>')
	show = re.findall(showre, html)

	sou = re.compile('<a href=.*? target=.*? title=.*? class="title">(.*?)</a>')
	dio = re.findall(sou, html)	

	soup = re.compile('<a href=.*? class="name">(.*?)</a>')
	jojo = re.findall(soup, html)

	cont = 0
	while True:
		print(dio[cont].ljust(40) + '\t作者：' + jojo[cont])
		cont+=1
		if cont >= 12:
			break

	print(show)

if __name__ == '__main__':
	biu(num = input('第几页：'))

