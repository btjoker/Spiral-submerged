# -*- coding: utf-8 -*-

import sys
import requests
import re
import time

'''
接口返回json样式
__GetZoneResult_ = {
    mts:'1306388', 手机号前6位：网段(139)+地区识别号(0916)
    province:'江苏', 省份
    catName:'中国联通', 运营商
    telString:'13063881681', 手机号码
    areaVid:'30511', 区域代码
    ispVid:'137815084', 运营商代码
    carrier:'江苏联通' 地区运营商
}
'''
_Author = 'btjoker'
if len(sys.argv) == 1:
    temp = input('请输入手机号')
    if len(temp) >= 11 and temp.isalnum():
        phone = temp
    else:
        print('请输入正确的手机号！')
elif len(sys.argv) == 2:
    if len(sys.argv[1]) >= 11 and sys.argv[1].isalnum():
        phone = sys.argv[1]
    else:
        print('请输入正确的手机号！')
else:
    print('本脚本只接受一个命令行参数。')

start = time.time()
url = 'http://tcc.taobao.com/cc/json/mobile_tel_segment.htm?tel='
mts = re.compile('mts:\'(.*?)\'')
telString = re.compile('telString:\'(.*?)\'')
carrier = re.compile('carrier:\'(.*?)\'')
# province = re.compile('province:\'(.*?)\'')

html = requests.get(url+phone).text
if len(html) > 50:
    a = re.findall(telString, html)
    b = re.findall(mts, html)
    c = re.findall(carrier, html)
    # d = re.findall(province, html)
    # print('号码：%s\n号段：%s\n所属：%s\n地区：%s' % (a, b, c, d))
    print('号码：%s\n号段：%s\n所属：%s' % (a, b, c))
    print('总耗时：%.2f 秒' % (time.time() - start))
