interface ServerConfig {
    readonly port: number; // readonly คือการล็อกค่า
    host: string;
    protocol: "http" | "https" // | คือ นี่หรือนั้น
    timeout?: number; // ? คือมีหรือไม่มีก็ได้
}

const myServer: ServerConfig = {
    port: 8080,
    host: "localhost",
    protocol: "https"

}

console.log(myServer)