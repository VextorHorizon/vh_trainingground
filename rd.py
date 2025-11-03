import random

def main():
    target = random.randint(1, 6)
    guess = input("ลองทายเลข 1-6: ")
    try:
        guess_num = int(guess)
    except ValueError:
        print("กรุณากรอกตัวเลขนะครับ")
        return

    if guess_num < 1 or guess_num > 6:
        print("ตัวเลขต้องอยู่ระหว่าง 1 ถึง 6 เท่านั้น")
    elif guess_num == target:
        print("ถูกต้องเลย! คุณทายได้เก่งมาก")
    else:
        print(f"ยังไม่ถูกนะ ลองใหม่อีกครั้งสิ เลขที่ออกคือ {target}")

if __name__ == "__main__":
    main()