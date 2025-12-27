import requests
from bs4 import BeautifulSoup
import argparse
import os

class Imagescaper:
    def __init__(self, target_web_url: str): 
        self.url = target_web_url
         # ข้อมูลตั้งต้นที่จะใช้สำหรับไปดึงรูปได้ มี url ของเว็บ, สิ่งที่ต้องไปดึง(รูปอะ) คือมันต้องไปหา type ชะมะ แล้วก้ไปดึงออกมา เดะนะ แปลว่ามันก้ดึง icon นู่นนี่ออกมาด้วยหมดเลย