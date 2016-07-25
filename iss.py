# -*- coding: utf-8 -*-
_Author = "btjoker"
import scrapy
from bs4 import BeautifulSoup
import json
import win32api

class IssSpider(scrapy.Spider):
    name = "iss"
    allowed_domains = ["www.ishadowsocks.org"]
    start_urls = (
        'http://www.ishadowsocks.org/',
    )

    def parse(self, response):
    	html_doc = response.body
    	soup = BeautifulSoup(html_doc, 'lxml')
    	num = 0
    	info = []
    	ap = soup.select('.col-lg-4 > h4')
        for i in ap:
			if not i.string:
				num+=1
			if num ==3:
				break
			if i.string:
				if i.string.find(':') > -1:
					info.append(i.string[i.string.find(':') + 1:])
        data = {"configs":[{"server":info[0],"server_port":int(info[1]),"password":int(info[2]),"method":info[3],"remarks":""},{"server":info[4],"server_port":int(info[5]),"password":int(info[6]),"method":info[7],"remarks":""},{"server":info[8],"server_port":int(info[9]),"password":int(info[10]),"method":info[11],"remarks":""}],"strategy" : None,"index" : 0,"global" : False,"enabled" : True,"shareOverLan" : False,"isDefault" : False,"localPort" : 1080,"pacUrl" : None,"useOnlinePac" : False}
        with open('gui-config.json', 'w') as f:
        	f.write(json.dumps(data,sort_keys=False,indent=4).encode('utf-8'))
        win32api.ShellExecute(0, 'open', 'Shadowsocks.exe', '','',1)
        pass