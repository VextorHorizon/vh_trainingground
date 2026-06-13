fn change_bool(status: &mut bool){
    println!("{status}");
    *status = true // *status = เนื้อหาของ status
}

fn main(){
    let mut booling = false;
    let longtiju_booling = &booling;
    // println!("{longtiju_booling}");
    change_bool(&mut booling);  
    println!("{longtiju_booling}");
    println!("{booling}");
}


// borrowing concept นี่หว่า concept คือ ใครจะยืม(ชี้ไปที่ตัวแปรต้นทาง)ไปอ่านก็ได้ แต่ว่ามีคนแก้ได้แค่คนเดียว!(&mut) 
// คนนั้นคนที่ขอยืมไปเพื่อแก้ คนที่ 1 คนอื่นจะไปแก้หรือทำอะไรระหว่างคนที่ 1 แก้ไม่ได้!