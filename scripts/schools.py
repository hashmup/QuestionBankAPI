import urllib2
from bs4 import BeautifulSoup

html = urllib2.urlopen("http://example.com")

soup = BeautifulSoup(html)
