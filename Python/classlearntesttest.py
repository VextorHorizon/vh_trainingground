class Dog:
    def __init__(self, name, age, gender):
        self.name = name
        self.age = age
        self.gender = gender

    def dog_is_barking(self):
        print(f"Da {self.name} is barking!! WOF WOF!!")

    def user_ask_dog_name_and_age(self):
        if self.gender == 'female':
            print(f"This dog is {self.name} and she is {self.age}")

        elif self.gender == 'male':
            print(f"This dog is {self.name} and he is {self.age}")

sausage = Dog('Sausage', 4, 'female')
sausage.age = 3

sausage.user_ask_dog_name_and_age(), sausage.dog_is_barking()
