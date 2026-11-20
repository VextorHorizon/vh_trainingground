// ANSWER KEY - I did not write this code. I read it before solving the problem.
// Task: CRC32 of the single byte 'a', expected 0xE8B7BE43 (same as Python zlib.crc32(b'a')).
// Saved 11/20/2026 to study. Not solved until I rewrite it myself from a blank file.
// Build: cl /EHsc /std:c++latest crc_a_answerkey.cpp

#include <print>
#include <cstdint>

int main() {
    uint32_t crc  = 0xFFFFFFFF;  // start: all 32 bits = 1
    uint32_t poly = 0xEDB88320;  // the CRC32 divisor (bit-reversed)

    crc = crc ^ 'a';             // 'a' is just the byte 0x61 = 01100001
    std::println("start   {:032b}", crc);

    // round 1
    if (crc & 1) crc = (crc >> 1) ^ poly;
    else         crc = crc >> 1;
    std::println("round 1 {:032b}", crc);

    // round 2
    if (crc & 1) crc = (crc >> 1) ^ poly;
    else         crc = crc >> 1;
    std::println("round 2 {:032b}", crc);

    // round 3
    if (crc & 1) crc = (crc >> 1) ^ poly;
    else         crc = crc >> 1;
    std::println("round 3 {:032b}", crc);

    // round 4
    if (crc & 1) crc = (crc >> 1) ^ poly;
    else         crc = crc >> 1;
    std::println("round 4 {:032b}", crc);

    // round 5
    if (crc & 1) crc = (crc >> 1) ^ poly;
    else         crc = crc >> 1;
    std::println("round 5 {:032b}", crc);

    // round 6
    if (crc & 1) crc = (crc >> 1) ^ poly;
    else         crc = crc >> 1;
    std::println("round 6 {:032b}", crc);

    // round 7
    if (crc & 1) crc = (crc >> 1) ^ poly;
    else         crc = crc >> 1;
    std::println("round 7 {:032b}", crc);

    // round 8
    if (crc & 1) crc = (crc >> 1) ^ poly;
    else         crc = crc >> 1;
    std::println("round 8 {:032b}", crc);

    crc = crc ^ 0xFFFFFFFF;      // end: flip every bit
    std::println("CRC = 0x{:08X}", crc);
}
