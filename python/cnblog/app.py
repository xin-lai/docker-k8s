# 关于BeautifulSoup，请阅读官方文档：https://beautifulsoup.readthedocs.io/zh_CN/v4.4.0/#id52
from bs4 import BeautifulSoup
import os
import sys
import requests
import time
import re

url = "https://www.cnblogs.com/codelove/default.html?page={page}"

#已完成的页数序号，初时为0
page = 0
while True:
    page += 1
    request_url = url.format(page=page)
    response = requests.get(request_url)
    #使用BeautifulSoup的html5lib解析器解析HTML（兼容性最好）
    html = BeautifulSoup(response.text,'html5lib')

    #获取当前HTML的所有的博客元素
    blog_list = html.select(".forFlow .day")

    # 循环在读不到新的博客时结束
    if not blog_list:
        break

    print("fetch: ", request_url)

    for blog in blog_list:
        # 获取标题
        title = blog.select(".postTitle a")[0].string
        print('--------------------------'+title+'--------------------------');

        # 获取博客链接
        blog_url = blog.select(".postTitle a")[0]["href"]
        print(blog_url);

        # 获取博客日期
        date = blog.select(".dayTitle a")[0].get_text()
        print(date)

        # 获取博客简介
        des = blog.select(".postCon > div")[0].get_text()
        print(des)

        print('-------------------------------------------------------------------------------------');

