from bs4 import BeautifulSoup
import requests

url = requests.get("https://pantip.com/topic/42592792")
soup = BeautifulSoup(url.content, "html.parser")
# soup = BeautifulSoup(url.content)
data = soup.find_all("h1", {"class":"display-post-title"})
if data is None:
    print("Coun't find data")
else:
    for item in data:
        print(item.text.strip()) #ในกรณีที่ตัวที่ข้อมูล(h1) มีหลายอัน มันจะวนดู ถ้าเป็น text เฉยๆไม่ได้ เพราะข้อมูลของ find_all() เป็น list คือมันแถม [] มาด้วยนั้นล่ะ



# print(soup)
