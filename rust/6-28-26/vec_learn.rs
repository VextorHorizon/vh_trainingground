fn main() {
    let mut nums: Vec<i32> = vec![1,2,3];
    nums.push(2);

    println!("{:?}", nums); //debug format

    let blah = String::from("43");
    println!("{:p}", &blah); // debug pointer value
    println!("{}", &blah); // rust เดินตาม memory address เพราะว่า {} แบบปกติต้องการแสดงข้อมูลของค่า blah มันเลย deref ตอนต้น
}
