#include <stdio.h>

int main() {

    char name[30];
    int num = 100;
    float height = 160.5;
    char grade = 'D';

    
    printf("Your name: ");
    scanf("%29s", name);

    printf("Your name is: %s\n", name);
    printf("Number: %d\n", num);
    printf("Height: %.f\n", height);
    printf("Grade: %c\n", grade);

    return 0;
}
