class House {
    
    lightHouse(){

        if (light == true){
            console.log("The light is on");
        } else {
            console.log("The light is off");
        }
    }
}

let light: boolean = false;

const myHouse = new House();

myHouse.lightHouse();