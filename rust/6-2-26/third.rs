fn main() {
    let result = plusplus(2, 3);
    println!("{result}");

    let resultstr = returnstring();
    println!("{resultstr}")
}

fn plusplus(x: i32, y: i32) -> i32 {
    x + y
}

fn returnstring() -> String {
    let x = "Hi";
    println!("{:?}", std::any::type_name_of_val(&x));
    "Hi".to_string()
}
