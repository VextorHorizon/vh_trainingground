use std::io;

fn main() {

    let mut city = String::new();

    io::stdin().read_line(&mut city).expect("Error"); //มันแถม \n มาข้างหลังด้วย

    let clean_input = city.trim(); //city_word_count = &str แล้ว .trim() ไม่ใช่การ copy เข้ามาอยู่ในตัวแปร

    // &str = ช่วงของข้อมูลใน heap — ptr กับ len แค่นั้น ไม่ได้ยืม header ของ city มา
    // &city = ยืม String ทั้งก้อน header + ทุกอย่าง


    
    // let x = city; 
    // อันนี้คือ move city เข้ามาเป็น x หรือนึกภาพตามคือ ผสานรวมร่างกันเป็นหนึ่งเดียวกัน
    // เวลาตัวแปรอื่นชี้มาที่ &city จะ err เพราะว่าไม่มี city อีกแล้ว

    let x = city.clone();
    //อันนี้เป็นการ copy city เข้ามาเก็บใน x — copy ที่เป็นการ copy ค่าจริงๆของ city เข้ามาเก็บไว้ในตัวเอง เป็น String

    let city_word_count = clean_input.len(); 

    println!("This is, {clean_input}! and word count is {city_word_count}!");
    

}

// ทำความเข้าใจ syntax มันก้เลยจะซ้ำๆนิดนึงนะ