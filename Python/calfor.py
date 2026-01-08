def plus(a, b):
    return a + b

def minus(a, b):
    return a - b 

def multiply(a, b):
    return a * b

def subdivide(a, b):
    return a / b

user_option = input("Choose option\n(Plus=1, Minus=2, Multiply=3, Subdivide=4)\n: ")

if user_option == '1':
    firstnum = int(input(": "))
    secnum = int(input(": "))
    result = plus(firstnum, secnum)
    print(f"{result}")

elif user_option == '2':
    firstnum = int(input(": "))
    secnum = int(input(": "))
    result = minus(firstnum, secnum)
    print(f"{result}")

elif user_option == '3':
    firstnum = int(input(": "))
    secnum = int(input(": "))
    result = multiply(firstnum, secnum)
    print(f"{result}")

elif user_option == '4':
    firstnum = int(input(": "))
    secnum = int(input(": "))
    result = subdivide(firstnum, secnum)
    print(f"{result}")



'''
Too much function and too much if else
'''