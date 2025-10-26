#include <stdio.h>

int main() {

    int x, y;
    printf("ใส่เลขสองตัว: ");
    scanf("%d %d", &x, &y);

    printf("ผลบวก: %d\n", x + y);
    printf("ผลลบ: %d\n", x - y);

    printf("ผลคูณ: %d\n", x * y);

    if (y != 0) {
        printf("ผลหาร: %d ", x / y);
        if (x % y > 0) {
            printf("เศษ: %d\n", x % y);
            }
    }
    else {
        printf("หารด้วยศูนย์ไม่ได้");
    }

    return 0;

}