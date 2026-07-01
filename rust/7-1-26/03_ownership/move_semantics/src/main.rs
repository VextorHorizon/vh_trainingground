fn main() {
    let a = String::from("Hello");

    println!("{:p}", &a);

    println!("{:p}", a.as_ptr());

    let c = a;

    println!("{:p}", &c);

    println!("{:p}", c.as_ptr());
}


// .as_ptr() เอาง่ายๆคือการขอดู memory address ของ heap ใน stack ของค่านั้นๆ แบบ ดูเนื้อหาของ c(stack) แล้วเอาค่า memory address ที่เก็บไว้ใน c

// สรุปง่ายๆแล้วการ moved คือการที่เปลี่ยนผู้ถือโฉนด(ค่าใน stack) จากคนหนึ่งไปยังอีกคนนึง แค่นั้นเลย ค่าใน heap ยังคงเหมือนเดิม แต่เปลี่ยนผู้ถือใน stack