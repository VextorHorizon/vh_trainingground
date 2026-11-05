#include <fstream>
#include <cstdio>

int main(){ 
    std::ifstream f("hi.txt", std::ios::binary); // input file stream, ios = input output stream, binary = raw binary
    char c;
    while (f.get(c)){ //from f, put, in c
        std::printf("%02x ", (unsigned char)c);
    }
}

    /* 
    what the freak is this programming language concept, 
    it make me wanna laying down on the bed and not learning it for 60 days 
    jkjk
    */