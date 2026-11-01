#include <iostream> // that is file from library file.lib in my computer

int twice(int n){
    return n * 2;
}

int main(){
    int x = twice(6);
    std::cout << x;
}

