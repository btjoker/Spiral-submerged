# -*- coding: utf-8 -*-
from lxml import etree
import sqlite3
import time
import threading
import urllib.request
_Author = 'btjoker'


class Tb(object):
    def __init__(self):
        self.header = {
            'User-Agent': 'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)'
        }
        self.dbname = 'proxy.db'
        self.now = time.strftime('%Y-%m-%d')

    def getContent(self, num):
        url = 'http://www.kuaidaili.com/free/inha/' + str(num)
        try:
            req = urllib.request.Request(url=url, headers=self.header)
            html = urllib.request.urlopen(req, timeout=10)
            content = html.read()
            root = etree.HTML(content)
            result_a = root.xpath('//td[@data-title="IP"]/text()')
            result_b = root.xpath('//td[@data-title="PORT"]/text()')
            for i in zip(result_a, result_b):
                if self.isAlive(i[0], i[1]):
                    self.insert_db(self.now, i[0], i[1])
        except urllib.error.HTTPError:
            print('服务暂时不可用!')

    def isAlive(self, ip, port):
        proxy = {'http': ip + ':' + port}
        proxy_support = urllib.request.ProxyHandler(proxy)
        opener = urllib.request.build_opener(proxy_support)
        urllib.request.install_opener(opener)
        test_url = 'http://qq.com'
        req = urllib.request.Request(url=test_url, headers=self.header)
        try:
            resp = urllib.request.urlopen(req, timeout=10)
            if resp.code == 200:
                print('可用!!\t-->%s' % proxy)
                return True
            else:
                print('无效!!\t-->%s' % proxy)
                return False
        except:
            print('无效!!\t-->%s' % proxy)
            return False

    def insert_db(self, date, ip, port):
        dbname = self.dbname
        try:
            conn = sqlite3.connect(dbname)
        except:
            print('Error to open database %s' % self.dbname)
        create_tb='''
        CREATE TABLE IF NOT EXISTS PROXY
        (DATE TEXT,
        IP TEXT,
        PORT TEXT
        );
        '''
        conn.execute(create_tb)
        insert_db_cmd='''
        INSERT INTO PROXY (DATE, IP, PORT) VALUES ('%s','%s','%s');
        ''' % (date, ip, port)
        conn.execute(insert_db_cmd)
        conn.commit()
        conn.close()

    def loop(self, page=5):
        for i in range(1, page):
            ap = threading.Thread(target=self.getContent, args=(i,))
            ap.start()

if __name__ == '__main__':
    now = time.strftime('%Y-%m-%d--%H: %M: %S')
    print('开始于 %s' % now)
    run = Tb()
    run.loop()
