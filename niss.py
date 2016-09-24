# -*- coding: utf-8 -*-

import re, urllib.request, json, os, sys
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
for i in range(len(result3)):
	if not result3[i]:
		result3[i] = 0
	else:
		xz = i
data = {"configs":[{"server":result1[0],"server_port":int(result2[0]),"password":int(result3[0]),"method":result4[0],"remarks":""},{"server":result1[1],"server_port":int(result2[1]),"password":int(result3[1]),"method":result1[1],"remarks":""},{"server":result1[2],"server_port":int(result2[2]),"password":int(result3[2]),"method":result4[2],"remarks":""}],"strategy" : None,"index" : int(xz),"global" : False,"enabled" : True,"shareOverLan" : False,"isDefault" : False,"localPort" : 1080,"pacUrl" : None,"useOnlinePac" : False}

with open(r'D:\pytemp\iss\gui-config.json', 'wb') as f:
    f.write(json.dumps(data,sort_keys=False,indent=4).encode('utf-8')) 

soft = os.path.split(sys.argv[0])[0] +'\\Shadowsocks.exe'
if os.path.exists(soft):
	os.popen(soft)
	print('程序运行正常！')
else:
	print('未找到小飞机(Shadowsocks)！\n请手动将生成的json文件移动到小飞机根目录下。')