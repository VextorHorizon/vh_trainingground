use std::io;

fn main(){
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Blah Blah"); // stdin เปิดช่องให้รับคีย์บอร์ด .read_line() คือหลังเปิดช่องแล้วรอ user กด enter เพื่อเอา input นั้นเข้า varable
    println!("{input}");

    let number: i32  = input
        .trim()
        .parse().expect("not number");
    // มี .trim() เพราะว่า โดยปกติแล้วจะมี \n ติดมาด้วยเสมอๆ ที่แสดงผลออกมาเป็น string แล้วไม่เห็น \n เพราะว่ามันเว้นบรรทัดให้ไปแล้ว(จะไปเห็นได้ไงฟร่ะ)
    // มี .parse() เพื่อแปลง string เป็น integer ปล. ต้องกำหนดข้างหน้าด้วยนะ
    println!("{number}")

}
