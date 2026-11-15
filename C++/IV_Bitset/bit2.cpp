#include <print>

int main(){

    int x = 0b0110;
    int y = 0b1100;
    int z = 0b1001;

    std::println("{:04b}",x >> 1);
    std::println("{:04b}",y ^ 0b1010);
    std::println("{:04b}",z >> 1 );

}