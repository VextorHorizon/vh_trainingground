#include <print>

int main(){
    uint32_t byte = 0xFFFFFFFF;

    std::println("{:08X}", 'a' ^ byte);
}