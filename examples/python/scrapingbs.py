from bs4 import BeautifulSoup
import requests

url = requests.get("https://pantip.com/topic/42592792")
soup = BeautifulSoup(url.content, "html.parser")
# soup = BeautifulSoup(url.content)
data = soup.find("h1", {"class":"display-post-title"})
if data is None:
    print("Coun't find data")
else:
    print(data.text)


# print(soup)
