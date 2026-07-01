fn main() {
    let word = String::from("HiHi");

    let word_byte = count_bytes(&word);

    println!("{}", word_byte);
    println!("{word}");





}

fn count_bytes(input_word: &String) -> usize {
    input_word.len()
}

// usize คือ unsignned size ความหมายคือ ค่า integer ที่ไม่ติดลบนั้นเอง ประมาณว่า 0 ถึงประมาณ 1.8 หมื่นล้านล้านล้าน
// i32 คือ signned size ความหมายคือ integer ที่ติดลบได้นั้นเอง ประมาณว่า -2 ล้าน ถึง +2 ล้าน

// let blah_blah = word; = move
// let blah_blah = &word; = borrowing