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

is_found = False
for data in profile_data:
    name = data["name"]
    profile_url = data["url"]

    if name == args.open:
        webbrowser.open(profile_url)
        is_found = True
        sys.exit()
    # if is_found == False:
    #     continue
if is_found == False:
    print("No website that you looking for!")


    # print(profile_url)
    # print(args.url)




