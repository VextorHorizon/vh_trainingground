#include <print>
#include <cstdint>

int main(){
    uint32_t crc = 0xFFFFFFFF; // 32bit = 4,294,967,295 or 0x00000000 -> 0xFFFFFFFF. 32 bit = 4 byte | CRC32 = 32 bit = 4 byte

    std::println("0x{:08X}", crc ^ 'a');
}