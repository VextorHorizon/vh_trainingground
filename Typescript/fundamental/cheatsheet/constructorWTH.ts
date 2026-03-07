class Guitar {
  brand: string;

  // นี่คือ Constructor!
  constructor(brandName: string) {
    this.brand = brandName; 
    console.log("กำลังสร้างกีต้าร์ยี่ห้อ " + this.brand);
  }
}

// ตอนแกพิมพ์คำว่า new... 
// มันจะไปเรียก Constructor ข้างบนมาทำงานทันที!
const myGuitar = new Guitar("Music Man");