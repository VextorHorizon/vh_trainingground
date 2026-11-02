#include <fstream>
#include <cstdio>

int main(){ 
    std::ifstream f("hi.txt", std::ios::binary); // input file system, ios = input output stream, binary = raw binary
    char c;
    while (f.get(c)){ //from f, put, in c
        std::printf("%02x ", c);
    }
}