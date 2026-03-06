class House {
    // 1. กำหนด Property ให้ชัดเจน (Encapsulation)
    private isLightOn: boolean;

    // 2. ใช้ Constructor ในการกำหนดค่าเริ่มต้น
    constructor(initialLightStatus: boolean = false) {
        this.isLightOn = initialLightStatus;
    }

    // 3. Method ที่จัดการ Logic ภายใน
    public checkLightStatus(): void {
        if (this.isLightOn) {
            console.log("The light is on (๑˃ᴗ˂)ﻭ");
        } else {
            console.log("The light is off ╮(￣▽￣)╭");
        }
    }

    // แถม: Method สำหรับเปลี่ยนสถานะ (Setter)
    public toggleLight(): void {
        this.isLightOn = !this.isLightOn;
    }
}

// การใช้งาน
const myHouse = new House(false);
myHouse.checkLightStatus();
myHouse.toggleLight();
myHouse.checkLightStatus();