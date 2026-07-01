use std::io;

fn main() {

    let mut input_a = String::new();
    let mut input_b = String::new();


    println!("First Number: "); io::stdin().read_line(&mut input_a).expect("Error");
    println!("Second Number: "); io::stdin().read_line(&mut input_b).expect("Error");

    let parse_inputa: i32 = input_a.trim().parse().expect("Error, not integer");
    let parse_inputb: i32 = input_b.trim().parse().expect("Error, not integer");


    let output = plus_plus(parse_inputa, parse_inputb);

    println!("{output}");
}

fn plus_plus(value_a: i32, value_b: i32) -> i32 {
    value_a + value_b
}