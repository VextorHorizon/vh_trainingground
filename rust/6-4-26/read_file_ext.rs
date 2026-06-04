use std::fs;

fn main() {
    for entry in fs::read_dir(".") {
        let dir = entry;
        println!("{:?}", dir.path());
    }
}