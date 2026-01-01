from ping3 import ping
import time
host = "www.google.com"

for i in range(4):
    latency = ping(host)
    if latency:
        print(f"{latency * 1000:.2f} ms")
        time.sleep(3)
    else:
        print("Can't Connect")
