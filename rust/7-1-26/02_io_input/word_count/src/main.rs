
use std::io;


fn main() {

    let mut input = String::new();

    io::stdin().read_line(&mut input).expect("NONONONONO");
    let clean_input = input.trim();

    let count = input.trim().split_whitespace().count();
    // let alphabet_in_input = clean_input.len(); 
    //แอ้ดแอ้ด ผิดจ้า clean_input ค่าของมันตอนนี้คือ bytes ของ String ของตัวแปร input

    println!("input: {clean_input} and {count} words");
}
 