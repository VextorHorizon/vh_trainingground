use std::fs;

fn main() -> std::io::Result<()> {
    for entry in fs::read_dir(".")? {
        let dir = entry?;
        println!("{:?}", dir.path());
    }
    Ok(())
}


// ? is a 
// if err != nil {
//    return err
// } 
// in golang... but rust quite short hand