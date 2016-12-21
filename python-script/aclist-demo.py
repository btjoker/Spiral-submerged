# -*- coding: utf-8 -*-

import requests
import re

_Author = 'btjoker'


def biu(num):
    if num == '1':
        url = 'http://www.acfun.tv/v/list110/index.htm'
    else:
        url = 'http://www.acfun.tv/v/list110/index_' + num + '.htm'

    html = requests.get(url).text
    showre = re.compile('<span class="hint">(.*?)</span>')
    show = re.findall(showre, html)
    sou = re.compile('<a href=.*? target=.*? title=.*? class="title">(.*?)</a>')
    dio = re.findall(sou, html)
    soup = re.compile('<a href=.*? class="name">(.*?)</a>')
    jojo = re.findall(soup, html)

    cont = 0
    if sou or jojo:
        print('页面不存在！！！')
    while soup and jojo:
        print(dio[cont].ljust(40) + '\t作者：' + jojo[cont])
        cont += 1
        if cont >= 12:
            break

    print(show)

if __name__ == '__main__':
    num = input('第几页：')
    biu(num)
