interface ServerConfig {
    readonly port: number;
    host: string;
    protocol: "http" | "https"
    timeout?: number;
}

const myServer: ServerConfig = {
    port: 8080,
    host: "localhost",
    protocol: "https"

}

console.log(myServer)