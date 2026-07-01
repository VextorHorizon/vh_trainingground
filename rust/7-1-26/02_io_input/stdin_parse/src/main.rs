use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Error");

    
    let int: i32 = input.trim().parse().unwrap();

    println!("{}", plusplus(int))
}

fn plusplus(input1: i32) -> i32{
    return input1 + 1 
}