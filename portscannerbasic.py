import sys #lib that use to work with CLI
import socket

if len(sys.argv) == 2: #sys.argv read argument from CLI like "simpleportscanner.py[0] 192.168.1.1[1]" #2 mean it Argument not a tags
    realtargetname = sys.argv[1]
    target = socket.gethostbyname(sys.argv[1]) #input, use it in cli (simpleportscanner.py 192.168.1.1)
else:
    print("Invild amount of Argument")
    print("Syntax: python simpleportscanner.py <ip>")
    sys.exit()

print(f"Scanning {target} Port 1 to 1024 \n---------------------------------")
try:
    for port in range(1,1024):
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM) #socket.AF_INET is command use to connect to that IP address, socket.SOCK_STREAM is command to tell it to connect with TCP protocal
        s.settimeout(0.5) #timeout

        result = s.connect_ex((target, port)) #connect to target and port
        if result == 0:
            print(f"[+] Port {port} OPEN")

        s.close()
        
    print("Finish scanning")


except KeyboardInterrupt:
    print("Interrupt by user")
    sys.exit()
except socket.gaierror:
    print(f"Could not resolve the hostname {target}")
    sys.exit()

