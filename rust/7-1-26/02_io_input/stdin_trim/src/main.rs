use std::io;


fn main() {
    let mut blah_blah = String::new();
    
    io::stdin().read_line(&mut blah_blah).expect("error");

    let clean_input = blah_blah.trim();
    println!("{clean_input}");

}
