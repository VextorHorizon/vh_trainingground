fn main() {
    let x = "Hello"; //&str ชี้ไปที่ read-only data ที่บันทึกว่า Hello ใน binary
    // let mut y = x;

    let y = String::from("Halloween"); //เป็นค่า String 
    println!("{x}");
    println!("{y}");
    // y = 3;
    // println!("{y}")
    not_main();
    print_things(53)
}

fn not_main(){
    println!("From not main!");
}

fn print_things(thing: i32){
    println!("things from thing is: {thing}");
}