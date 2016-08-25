# -*- coding: utf-8 -*-

import re
import urllib.request
import json

_Author = "btjoker"

url = 'http://www.ishadowsocks.org/'
html = urllib.request.urlopen(url).read()
soup = html.decode('utf-8', 'ignore')

fwq = re.compile('<h4>.*?服务器地址:(.*?)</h4>')
dk = re.compile('<h4>端口:(.*?)</h4>')
pwd = re.compile('<h4>.*?密码:(.*?)</h4>')
jm = re.compile('<h4>加密方式:(.*?)</h4>')
result1 = re.findall(fwq, soup)
result2 = re.findall(dk, soup)
result3 = re.findall(pwd, soup)
result4 = re.findall(jm, soup)

data = {"configs":[{"server":result1[0],"server_port":int(result2[0]),"password":int(result3[0]),"method":result4[0],"remarks":""},{"server":result1[1],"server_port":int(result2[1]),"password":int(result3[1]),"method":result1[1],"remarks":""},{"server":result1[2],"server_port":int(result2[2]),"password":int(result3[2]),"method":result4[2],"remarks":""}],"strategy" : None,"index" : 0,"global" : False,"enabled" : True,"shareOverLan" : False,"isDefault" : False,"localPort" : 1080,"pacUrl" : None,"useOnlinePac" : False}

with open('gui-config.json', 'wb') as f:
    f.write(json.dumps(data,sort_keys=False,indent=4).encode('utf-8')) 