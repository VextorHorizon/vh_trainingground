fn main() {
    let mut word = String::from("HIHI");

    let _a = &mut word; 
    
    
    // ให้นึกภาพว่า varable เป็นสมุด 1 เล่ม a ไปยืมดูสมุดแบบห่างๆคือ let a = &word; (ดูเฉยๆ)
    // การประกาศ &mut คือการเอาสมุดไปถือจริงๆเพื่อแก้ ไม่มีใครถือสมุดได้ นอกจากคนที่ยืม และ คนที่ยืมจะคืน


    let _b = &mut word;

    println!("{_a}")
}
