We implemented raw TCP server + HTTP parsing manually. This will help you understand how HTTP actually works under the hood.

# Project Structure
ecommerce-app/<br>
│── main.go             # Entry point of app <br>
│── server/ <br>
│     └── server.go     # Low-level TCP server (handles HTTP parsing)<br>
│── handlers/ <br>
│     └── product.go    # Product CRUD logic <br>
│── models/ <br>
│     └── product.go    # Product struct and in-memory store <br>


# How to Run <br>
go mod init ecommerce-app <br>
go run main.go <br>

# Testing with curl <br>
1. Create: curl.exe -X POST http://localhost:3000/products -H "Content-Type: application/json" -d '{\"name\":\"Laptop\",\"price\":1200}'
 <br>

2. Get: curl http://localhost:8080/products
<br>
3. Update: curl.exe -X PUT http://localhost:3000/products/3 -H "Content-Type: application/json" -d '{\"name\":\"Laptop\",\"price\":1599}'
<br>
4. Delete: curl.exe -X DELETE http://localhost:3000/products/1
<br>
