#coding:utf-8
import requests
import re
from requests.exceptions import RequestException
import json
#响应网站
def get_one_page(url):
    try:
        response=requests.get(url)
        if response.status_code==200:
            return response.text
        return None
    except RequestException:
        return None
#正则获取匹配内容
def get_link(html):
    patten=re.compile('<dd>.*?board-index.*?>(\d+)</i>.*?data-src="(.*?)".*?name">'\
                  '<a.*?>(.*?)</a>.*?star">(.*?)</p>.*?releasetime">(.*?)</p>'\
                  '.*?integer">(.*?)</i>.*?fraction">(.*?)</i>.*?',re.S)
    items=re.findall(patten,html)
    for item in items:
        yield {
            'index':item[0],
            'image':item[1],
            'title':item[2],
            'star':item[3].strip()[3:],
            'time':item[4][5:],
            'score':item[5]+item[6]
            }
#文件保存，注意编码问题
def save_content(content):
    with open('movie.txt','a',encoding='utf-8') as f:
        f.write(json.dumps(content,ensure_ascii=False)+'\n')
        f.close()
#nums int 传递的页数
def main(nums):
    for num in range(nums):
        url='http://maoyan.com/board/4?offset='+str(num*10)
        html=get_one_page(url)
        links=get_link(html)
        for link in links:
            save_content(link)
if __name__=='__main__':
    main(10)
