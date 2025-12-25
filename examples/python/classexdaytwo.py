class Player: #Blueprint จ้า 
    def __init__(self, codename: str) -> None: #init คือข้อมูลค่าตั้งต้นของตัว object โดยเดิม
        self.name = codename #เอา codename ที่รับ input มา ผูกกับ self.name


player_a = Player("VextorHorizon") #self.name = "VextorHorizon" ละก้ๆ self.name ก้คือ codename นะ
player_b = Player("RinRin") #self.name = "RinRin"
playera = [player_a, player_b] #เก็บไว้เป็น list โล้ด
# print(playera)
num = 0 #มาร์คล่วงหน้าไว้นับเลขในลูป
for data in playera: #ลูปเว้ย
    num += 1 #นับเลขก้มาว่ะ แบบ ถ้ามันลูป มันก้จะมา + 1 
    print(f"Player_{num} name: {data.name}") 
    #data.name คือ playera นั้นล่ะ ซึง playera มันเป็น list เวลาเป็น loop ชะ มันจะวันดูทีละตัวๆ
    #(ลูปแรก)วนไป num(0) + 1 ละก้ data.name คือ player_a นะ 
    #(ลูปสอง)วนไป num(1) + 1 ละก้ data.name คือ player_b น่ะ
    #list หมดแล้ว