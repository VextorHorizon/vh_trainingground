#include <fstream>
#include <cstdio>

int main(){
    std::ifstream f("hi.txt", std::ios::binary);

    char c;
    while (f.get(c)){
        printf("%02x\n", (unsigned char)c);
    }
}

// u8 (0–255)
// default is i8 (-128..127)