 from ping3 import ping

import time

latency = ping('www.google.com')


if latency:

    for i in range(4):

        time.sleep(2)

        print(f"{latency * 1000:.02f} ms")

else:

    print("Can't connect") 
