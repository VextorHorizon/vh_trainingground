"""
1.อ่านข้อมูลที่มีอยู่ใน profile.json(data เอาไว้เก็บว่า user จะเปิดอะไรแล้วมันจะไปที่ไหน) แล้วเก็บค่าเข้า่ไปใน script นี้ 
2.รับ input จาก user (python main.py --url Youtube) 
3.โค้ดรัน เปิดเว็บที่ user เขียน
"""
import sys
import json
import argparse #python main.py --open Youtube
import webbrowser 

try:
    with open ('profile.json', 'r') as profile:
        profile_data = json.load(profile)
except FileNotFoundError:
    print("profile.json is not found!")
    exit()
except json.JSONDecodeError:
    print("profile.json is broken!")
    exit()

parser = argparse.ArgumentParser()
parser.add_argument('--open', type=str) 
args = parser.parse_args()

user_input = args.open

url_dic = {}
for data in profile_data:
    url_dic[data["name"]] = data["url"]

# print(url_dic)
if user_input in url_dic:
    webbrowser.open(url_dic[user_input])
else:
    print("Not found website you looking for!")




    # print(profile_url)
    # print(args.url)




