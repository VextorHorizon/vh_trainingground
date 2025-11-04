import time
import sys
import threading
import os

class Stopwatch:
    def __init__(self):
        self.running = False
        self.elapsed = 0.0
        self.start_time = None
        
    def start(self):
        if not self.running:
            self.running = True
            self.start_time = time.time() - self.elapsed
            
    def stop(self):
        if self.running:
            self.elapsed = time.time() - self.start_time
            self.running = False
            
    def reset(self):
        self.running = False
        self.elapsed = 0.0
        self.start_time = None
        
    def get_time(self):
        if self.running:
            return time.time() - self.start_time
        return self.elapsed

def clear_line():
    sys.stdout.write('\r\033[K')
    sys.stdout.flush()

def display_time(stopwatch):
    while True:
        if stopwatch.running or stopwatch.elapsed > 0:
            t = stopwatch.get_time()
            hours = int(t // 3600)
            minutes = int((t % 3600) // 60)
            seconds = int(t % 60)
            milliseconds = int((t % 1) * 100)
            
            # บันทึก cursor position
            sys.stdout.write('\033[s')
            # ไปที่บรรทัดแสดงเวลา (บรรทัดที่ 4)
            sys.stdout.write('\033[4;1H')
            # เคลียร์บรรทัด
            sys.stdout.write('\033[K')
            # แสดงเวลา
            sys.stdout.write(f'{hours:02d}:{minutes:02d}:{seconds:02d}.{milliseconds:02d}')
            # กลับไปที่ cursor position เดิม
            sys.stdout.write('\033[u')
            sys.stdout.flush()
            time.sleep(0.01)
        else:
            time.sleep(0.1)

def main():
    # เคลียร์หน้าจอ
    os.system('clear' if os.name != 'nt' else 'cls')
    
    stopwatch = Stopwatch()
    
    print("Simple Stopwatch")
    print("Commands: 's' = start, 'p' = pause/stop, 'r' = reset, 'c' = clear, 'q' = quit")
    print("-" * 50)
    print("00:00:00.00")
    print("-" * 50)
    print("")  # บรรทัดว่างสำหรับ input
    
    # เริ่ม thread สำหรับแสดงเวลา
    display_thread = threading.Thread(target=display_time, args=(stopwatch,), daemon=True)
    display_thread.start()
    
    try:
        while True:
            cmd = input("> ").strip().lower()
            
            if cmd == 's':
                stopwatch.start()
            elif cmd == 'p':
                stopwatch.stop()
            elif cmd == 'r':
                stopwatch.reset()
                # อัพเดทการแสดงผลให้เป็น 0
                sys.stdout.write('\033[s')
                sys.stdout.write('\033[4;1H')
                sys.stdout.write('\033[K')
                sys.stdout.write('00:00:00.00')
                sys.stdout.write('\033[u')
                sys.stdout.flush()
            elif cmd == 'c':
                # เคลียร์หน้าจอและวาดใหม่
                os.system('clear' if os.name != 'nt' else 'cls')
                print("Simple Stopwatch")
                print("Commands: 's' = start, 'p' = pause/stop, 'r' = reset, 'c' = clear, 'q' = quit")
                print("-" * 50)
                # แสดงเวลาปัจจุบัน
                t = stopwatch.get_time()
                hours = int(t // 3600)
                minutes = int((t % 3600) // 60)
                seconds = int(t % 60)
                milliseconds = int((t % 1) * 100)
                print(f'{hours:02d}:{minutes:02d}:{seconds:02d}.{milliseconds:02d}')
                print("-" * 50)
                print("")
            elif cmd == 'q':
                print("\nGoodbye!")
                break
                
    except KeyboardInterrupt:
        print("\n\nGoodbye!")

if __name__ == "__main__":
    main()