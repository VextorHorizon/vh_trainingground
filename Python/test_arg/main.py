"""
1.อ่านข้อมูลที่มีอยู่ใน profile.json(data เอาไว้เก็บว่า user จะเปิดอะไรแล้วมันจะไปที่ไหน) แล้วเก็บค่าเข้า่ไปใน script นี้ 
2.รับ input จาก user (python main.py -o Youtube) 
3.โค้ดรัน เปิดเว็บที่ user เขียน
"""

import json
import argparse #python main.py --open Youtube
import webbrowser 

with open ('profile.json', 'r') as profile:
    profile_data = json.load(profile)

for data in profile_data:
    name = data["name"]
    url = data["url"]

    # print(url)

def open_web():
    webbrowser.open(url)
