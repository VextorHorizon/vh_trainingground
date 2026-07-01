use std::io;

fn main() {
    let mut varvar = String::new();

    let mama = &mut varvar;

    io::stdin().read_line(&mut varvar).expect("123123");

    println!("{:?}", varvar) // อันนี้ได้ ไม่มีปัญหา 

    // println!("{:?}", mama) ไม่ได้ มันมีปัญหา

    // สรุปคือ borrowing drop (borrow's lifetime) เมื่อนับบรรทัดสุดท้ายของตัวที่ยืม เราจะเรียก mama ในตอนแรก แล้วตัวอื่นมาต่อได้ ถ้าไม่มีบรรทัดที่ตัว mama จะใช้ ตอนที่บรรทัดที่ 8 ถืออยู่
    // compiler อ่านทั้งหมดและตรวจสอบความถูกต้องก่อนจะรัน
    // drop mama ได้ เพราะว่า ไม่ได้มีการเรียกใช้ mama ก่อนหรือหลังการยืมตัวถัดไป
    // แต่ถ้ามีการเรียก แล้ว mama ยังใช้อยู่ = ไม่ได้ 
    // core concept คือ borrow ยืม(เพื่อแก้)ได้ตัวเดียว ถ้า mama ใช้ varvar อยู่ คนอื่นไม่มีสิทธิ์ใช้จนกว่า mama จะไม่ได้ใช้ varvar แล้ว

    
}
