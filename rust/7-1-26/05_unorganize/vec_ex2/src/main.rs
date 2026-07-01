use std::io;

fn main() {
    let mut userinput = String::new();

    io::stdin().read_line(&mut userinput).expect("Error");

    let clean:i32 = userinput.trim().parse().unwrap();

    // .unwrap() คือ
    // Ok(value)  → เอา value ออกมาให้
    // Err(...)   → panic ทิ้งโปรแกรม

    let mut Ro: Vec<i32> = vec![1,2,3,4];
    Ro.push(clean);
    println!("{:?}", Ro);
}
