fn change_bool(status: &mut bool){
    println!("{status}");
    *status = true
}


fn main(){
    let mut booling = false;
    let longtiju_booling = &booling;
    change_bool(&mut booling); // ให้สังเกตุ & ให้ดี มันคือ pointer 
    println!("{longtiju_booling}");
    println!("{booling}");
}