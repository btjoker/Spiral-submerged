# -*- coding: utf-8 -*-

import os
import sys

_Author = "btjoker"
'''
搜索当前目录下文件，根据命令行的参数是否包含.号来判断搜索名称还是后缀。
搜索方式为精确搜索，结果为该文件的绝对路径。
'''


def ffind(name):
    dirls = os.listdir()
    filename = [y for y in dirls if os.path.splitext(y)[0] == name]
    if not filename:
        print('未找到相关文件！')
    else:
        for i in [os.path.abspath(y) for y in filename]:
            print(i)
        print('\n搜索完毕！共：%s个结果！' % len(filename))


def fsuffix(name):
    dirls = os.listdir()
    suffix = [y for y in dirls if os.path.splitext(y)[1] == name]
    if not suffix:
        print('未找到后缀名为 %s的文件' % name)
    else:
        for i in [os.path.abspath(y) for y in suffix]:
            print(i)
        print('\n搜索完毕！共：%s个结果！' % len(suffix))


if __name__ == '__main__':
    cs = sys.argv
    if len(cs) == 1:
        print('请输入文件名或带点号的后缀名！')
    elif len(cs) == 2:
        if cs[1].find('.') > -1:
            fsuffix(cs[1])
        else:
            ffind(cs[1])
