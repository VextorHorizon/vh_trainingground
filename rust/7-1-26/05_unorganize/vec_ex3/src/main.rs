use std::io;

fn main() {
    let mut user_data: Vec<String> = vec![];


    for _i in 0..3{
        
            let mut name = String::new();
            io::stdin().read_line(&mut name).expect("error");

            let clean_name = name.trim(); // เรียบร้อย clean_name กลายเป็น &str ไปแล้ว
            user_data.push(clean_name.to_string());
        
    }


    println!("{:?}", user_data)
    
}
