use std::io;

fn main() {
    let mut user_input = String::new();

    io::stdin().read_line(&mut user_input).expect("Can't read cli");
    let user_input = user_input.trim(); // เปลี่ยน user_input เป็น &str ทำให้ได้แค่ read_only เด้อ
    // เขียน read_only เพราะว่ามันก้ได้แค่อ่านจริงๆ แบบ เราจะไปแก้ค่า &str(pointer + length) ยังไงอะนะ

    let names = vec![
        String::from("Vextor"),
        String::from("Horizon"),
        String::from("VextorHorizon"),
    ];

    // for name in &names{
    //     println!("{name}")
    // }

    let slice_location = &names[..];
    // &name[0..3] รู้ขนาดที่แน่ชัด
    // &name[..] เอาแม่งมาหมดเลย
    // &numbers[2..]; ไม่รู้ตัวท้าย อยากได้ตั้งแต่ index 2 ไปจนจบ:
    // &numbers[..3]; ไม่รู้ตัวแรก อยากได้ถึง index 3:

    // println!("{:?}", slice_location);

    let mut name_check_found = false;

    for name in slice_location{
        // println!("{name}");
        if user_input == name{
            name_check_found = true;
            println!("Your name matched the data! {name}!");
        }
    } 

    println!("{}",name_check_found);
    if !name_check_found {
        println!("Your name didn't matched the data");
    }

    // !name_check_found มีความหมายว่า not name_check_found หรือ name_check_found == false นั้นหล่ะ
}


// ในกรณีของ ตัวแปรที่เก็บค่าเป็น &str (มีข้อมูลอยู่ใน heap) ใน &str จะมี pointer + length ซึ่งเก็บไว้ใน stack  

//RAM
//┌─────────────────┐
//│   code / text   │  ← machine code (instructions)
//├─────────────────┤
//│   data / rodata │  ← string literals อยู่ตรงนี้ (read-only) ค่าฝังในโปรแกรมสมบูรณ์ ยืดหยุ่นไม่ได้ ทำอะไรไม่ได้ นอกจากหยิบมาอ่าน
//├─────────────────┤
//│   heap          │  ← String::from(...) อยู่ตรงนี้
//├─────────────────┤
//│   stack         │  ← local variables, &str struct
//└─────────────────┘