fn add_one(integer:&mut i32) -> i32{
    *integer += 1; //dereference / แก้ที่เนื้อหาของตัวแปร
    return *integer;
    
    
}

fn main() {
    let mut a:i32 = 10;
    println!("{a}");
    add_one(&mut a);

    println!("{a}");
    
}
