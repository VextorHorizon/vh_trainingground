interface Product {
    id: number
    name: string
    price: number
    category?: string
}

type SortBy = "price" | "name"

function filterByMaxPrice(products: Product[], max: number): Product[] {
    return products.filter(p => p.price <= max)
}

function sortProducts(products: Product[], by: SortBy): Product[] {
    return [...products].sort((a, b) => 
        by === "price" ? a.price - b.price : a.name.localeCompare(b.name)
    )
}

const products: Product[] = [
    { id: 1, name: "Laptop", price: 999 },
    { id: 2, name: "Phone", price: 599, category: "Mobile" },
    { id: 3, name: "Tablet", price: 399 },
]

const cheap = filterByMaxPrice(products, 600)
const sorted = sortProducts(cheap, "price")
console.log(sorted)
