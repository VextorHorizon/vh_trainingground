let nickname:string = "VextorHorizon";
let age:number = 17;
let isPerfect:boolean = false;

let Project = "Software House";

let techStack: string[] = ["Go", "Nestjs", "Nextjs"]
let userSession: [number, string] = [101, "101"]

function calculatorResource(cpu: number, ram: number): string {
    const total = cpu + ram;
    return `Resource usage is ${total} units`
}

const ResourceStatus: string = calculatorResource(2, 4)

console.log(ResourceStatus)
