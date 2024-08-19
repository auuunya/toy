#coding:utf-8
import requests
from urllib.parse import urlencode
from requests.exceptions import ConnectionError
import json
import re
from bs4 import BeautifulSoup
from json.decoder import JSONDecodeError
def get_page_index(offset,Keyword):
    data = {
        'autoload': 'true',
        'count': 20,
        'cur_tab': 3,
        'format': 'json',
        'keyword': Keyword,
        'offset': offset,
    }
    params = urlencode(data)
    base = 'http://www.toutiao.com/search_content/'
    url = base + '?' + params
    try:
        response = requests.get(url)
        if response.status_code == 200:
            return response.text
        return None
    except ConnectionError:
        print('Error occurred')
        return None

def get_page_parse(text):
	try:
		data = json.loads(text)
		if data and 'data' in data.keys():
			for item in data.get('data'):
				yield item.get('article_url')
	except JSONDecodeError:
		pass
def get_photo_index(url):
    try:
        response = requests.get(url)
        if response.status_code == 200:
            return response.text
        return None
    except ConnectionError:
        print('Error occurred')
        return None
def get_photo_parse(html):
    soup = BeautifulSoup(html, 'lxml')
    result = soup.select('title')
    title = result[0].get_text() if result else ''
    for t in title:
    	print(t)
    # images_pattern = re.compile('gallery: JSON.parse\("(.*)"\)', re.S)
    # result = re.search(images_pattern, html)
    
def main():
	text=get_page_index(0,'街拍')
	for url in get_page_parse(text):
		html=get_photo_index(url)
		get_photo_parse(html)
if __name__=='__main__':
#今日头条图片爬取
	main()
