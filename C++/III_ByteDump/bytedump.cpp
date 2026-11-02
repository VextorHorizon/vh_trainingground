#include <fstream>
#include <cstdio>

int main(){
    std::ifstream f("hi.txt", std::ios::binary); //open the stream file hi.txt through f variable stream in raw binary

    char c;
    while (f.get(c)){
        std::printf("%02x ",c);
    }
}