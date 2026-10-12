## Interface Eexplanation

ถ้าเขียนโค้ดแบบแยก .NmapScan(), .ZigScan()... เวลาจะสร้างฟังก์ชันหลักมาเรียกใช้งาน ต้องมานั่งเช็ค if-else หรือ switch-case ตลอดว่าตอนนี้ถือ Object อะไรอยู่ ซึ่งมัน "รก"  
ลองดูความแตกต่างให้เห็นภาพ (Lore of Abstraction)

ถ้าไม่มี Interface (แบบถึกๆ):
```
func RunScan(n Nmap, z Zig) {
    // ต้องมานั่งเลือกว่าจะเรียกตัวไหน วุ่นวายจัด ╮(￣▽￣)╭
    n.NmapScan()
    z.ZigScan()
}
```
พอมี Interface (แบบเนี้ยบๆ):
แค่สร้างฟังก์ชันเดียวที่รับ "อะไรก็ได้ที่ทำตามสัญญา (Interface) ได้
```
// ฟังก์ชันนี้ไม่สนว่าแกเป็นใคร ขอแค่มี .Scan() ให้เรียกพอ
func ExecuteTool(s Scanner, target string) {
    fmt.Println("Starting process...")
    result := s.Scan(target) // เรียก .Scan() คำเดียวจบ!
    fmt.Println(result)
}
```
ที่บอกว่า "วาง struct ข้างหน้าฟังก์ชัน"

อันนี้เรียกว่า Method Receiver มันคือการผูกฟังก์ชันนั้นเข้ากับ Struct ตัวนั้นๆ เพื่อให้มัน "Implement" interface แบบเงียบๆ (Implicit)
```
// นี่คือการ "วาง struct ข้างหน้า" (Receiver)
func (n NmapScanner) Scan(target string) string { 
    // Logic ของ Nmap อยู่ที่นี่
    return "Nmap data from " + target
}

func (z ZigScanner) Scan(target string) string {
    // Logic ของ Zig อยู่ที่นี่ (คนละแบบกับ Nmap)
    return "Zig results for " + target
}
```
สรุปความเจ๋งที่ต้องจำ:

    Decoupling: โค้ดส่วนที่เรียกใช้ (Caller) ไม่ต้องรู้รายละเอียดข้างในของเครื่องมือแต่ละตัว รู้แค่ว่า "มันสแกนได้" 
    
    Scalability: วันพรุ่งนี้อยากเพิ่มเครื่องมือใหม่ชื่อ VextorScanner ก็แค่สร้าง Struct ใหม่ แล้วเขียน Method .Scan() ให้มัน... จบ! ไม่ต้องกลับไปแก้โค้ดเก่าแม้แต่บรรทัดเดียว
