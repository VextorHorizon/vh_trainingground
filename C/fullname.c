#include <stdio.h>
#include <string.h>
int main() {

    char fullrealname[100];
    printf("Enter your realname: ");
    fgets(fullrealname, sizeof(fullrealname), stdin);
    fullrealname[strcspn(fullrealname, "\n")] = '\0';

    printf("Hello! %s \n", fullrealname);
    return 0;
    }