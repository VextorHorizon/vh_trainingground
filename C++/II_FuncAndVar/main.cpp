#include <iostream> 

int twice(int n){
    return n * 2;
}

int main(){
    int x = twice(6);
    std::cout << x;
}

