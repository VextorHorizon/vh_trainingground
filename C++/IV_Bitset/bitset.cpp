#include <iostream>
#include <bitset>
#include <string>
#include <vector>
#include <print>

int main(){
    std::vector<int> v;
    v.push_back(0b0101);
    std::println("{::04b}",v);
    int x = 0b0110;
    std::cout << std::bitset<4>(x) << "\n";
};