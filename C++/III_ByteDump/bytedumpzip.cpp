#include <fstream>
#include <print>

int main(){
    std::ifstream f("hi.zip", std::ios::binary);

    char c;
    int count = 1;

    while (f.get(c)){
        unsigned char u = c;
        char shown = (u >= 32 && u <= 126) ? c : '.';
        // std::println("Index: {}, Byte: {:02x}, Meaning: {}",count,(unsigned char)c,shown);
        // count++;

        std::printf("%c",shown);
    }
}

// u8 (0–255)
// default is i8 (-128..127)

/*
PK..........d]................hi.txt...      ← 1. แฟ้ม: ป้าย + ชื่อ + เนื้อไฟล์ (ก)
PK............d]......(จุดยาว).....hi.txt    ← 2. สารบัญ: นี่แหละ ls ของจริง
PK..........4...'.....                       ← 3. ท้ายไฟล์: บอกว่าสารบัญอยู่ตรงไหน
*/