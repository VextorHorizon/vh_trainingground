import random as rd

while True:
    number_one = rd.randint(1, 10)
    number_two = rd.randint(1, 10)
    try:
        ans = int(input(f"The Question is... What is {number_one} x {number_two}?\n: "))
    except ValueError:
        print("Invalid input. Please enter a number.")
        continue

    print("Correct!" if ans == number_one * number_two else f"Wrong! The answer is {number_one * number_two}")
    
    option_to_play = input("Do you want to play again? (Y/N)\n: ").strip().lower()
    if option_to_play not in ['y', 'n']:
        print("Invalid option. Exiting the game.")
        break

    if  option_to_play == 'n':
        break