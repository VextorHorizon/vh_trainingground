// ก๊อปค่าเข้า → แก้ก๊อป → return ออกมาให้คนเรียกรับเอง

fn double_it(mut integer: i32) -> i32{
    integer = integer * 2;
    return integer
} // ถ้าไม่ได้กำหนด return type: มัน expect ()

fn main() {

    let mut int = 21;
    int = double_it(int);
    println!("{int}");
}

// แก้ค่าเดิม แก้ของจริงๆเลย

fn double_it(integer:&mut i32){
    *integer = *integer * 2; 
    //แก้ของจริงไปแล้ว ไม่ต้อง return 
} // ถ้าไม่ได้กําหนด return type: มัน expect ()

fn main() {

    let mut int = 21;
    double_it(&mut int);
    println!("{int}");
}