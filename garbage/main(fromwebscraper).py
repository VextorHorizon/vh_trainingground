import requests #ใช้ request.get(url)
from bs4 import BeautifulSoup 
import argparse #รับ input จาก user
import os #เอารูปที่โหลดไปลง folder ที่เลือกไว้
import pathlib #อันนี้ auto หา folder

path = "C:\\Users\\copte\\Desktop\\Coding\\roadtor15\\webimagescraper\\target_img" #Hardcode!
if not os.path.exists(path):
    os.makedirs(path)

# parser = argparse.ArgumentParser()
# parser.add_argument('--url', type=str)
# args = parser.parse_args()

# userinput_urltarget = args.url
# # print(userinput_urltarget)


target_url = requests.get("https://ctmusicshop.com/product/cort-kx707-evertune-open-pore-black-electric-guitar/") #Hardcode!
soup = BeautifulSoup(target_url.content, "html.parser")
for pic in soup.find_all('img', width=True, height=True):
    # print(pic['width'], pic['height'])
    if pic['width'] and pic['height'] == 500:
        print(pic['width'], pic['height'])
# print(target_data)
# url
# path
# bs4 หา type <img>
         # ข้อมูลตั้งต้นที่จะใช้สำหรับไปดึงรูปได้ มี url ของเว็บ, สิ่งที่ต้องไปดึง(รูปอะ) คือมันต้องไปหา type ชะมะ แล้วก้ไปดึงออกมา เดะนะ แปลว่ามันก้ดึง icon นู่นนี่ออกมาด้วยหมดเลย

#os ชี้ว่าให้
# if __name__ == "__main__"():
