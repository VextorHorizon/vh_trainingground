fruits = ["Banana", "Apple", "Pineapple"]
# num = 0
# for fruit in fruits:
#     num += 1
#     print(f"Number {num} is {fruit}")
# print(fruits)
for num, fruit in enumerate(fruits, start=1):
    print(f"Number {num} is {fruit}")