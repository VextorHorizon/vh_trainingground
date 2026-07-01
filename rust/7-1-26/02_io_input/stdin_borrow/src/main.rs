use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("What the hell is that?");
    let input = input.trim();
    println!("You said: {input}");
    let address_input = &input;
    println!("{address_input}");
}
