class Guitar {
  brand: string;

  constructor(brandName: string) {
    this.brand = brandName; 
    console.log("กำลังสร้างกีต้าร์ยี่ห้อ " + this.brand);
  }
}
const myGuitar = new Guitar("Music Man");
