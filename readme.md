# Echo Products API

**Echo Products API** is a RESTful API built in **Go** using the **Echo** web framework and **GORM** for PostgreSQL.  
It provides endpoints for managing products and demonstrates a clean project structure with controllers, models, and migrations.

## 📦 Features

- 🧱 Clean project structure (`config`, `controller`, `model`, `migrations`)
- 🚀 RESTful CRUD endpoints for products
- 📡 Built with **Echo** — fast and minimal HTTP framework
- 📜 GORM integration for PostgreSQL
- 📦 Go modules for dependency management

---

## 🧠 Requirements

Before running the project, ensure you have:

- Go 1.20+ installed
- PostgreSQL installed and running
- Environment variables for database configuration

---

## 🛠 Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/mbchavez27/echo-products-api.git
   cd echo-products-api
   ```
2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Configure database**

   Edit `config/db.go` with your PostgreSQL credentials:

   ```go
   dsn := "host=localhost user=your_user password=your_pass dbname=echo_products_api port=5432 sslmode=disable TimeZone=Asia/Jakarta"
   ```

---

## 🚀 Running the Server

```bash
go run server.go
```

By default, the API will start on port `:8080`.

---

## 📍 API Endpoints

### **Products**

| Method | Endpoint        | Description                |
| ------ | --------------- | -------------------------- |
| GET    | `/products`     | List all products          |
| GET    | `/products/:id` | Get a product by ID        |
| POST   | `/products`     | Create a new product       |
| PUT    | `/products/:id` | Update an existing product |
| DELETE | `/products/:id` | Delete a product           |

---

## 🧪 Example Requests

### Create a Product

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{ "name": "T-Shirt", "price": 19.99, "ratings": 4.5, "description": "Comfortable cotton t-shirt", "category": "Clothing" }'
```

### Get All Products

```bash
curl http://localhost:8080/products
```

### Get Product by ID

```bash
curl http://localhost:8080/products/1
```

### Update a Product

```bash
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{ "name": "T-Shirt Updated", "price": 24.99, "ratings": 4.8, "description": "Updated description", "category": "Clothing" }'
```

### Delete a Product

```bash
curl -X DELETE http://localhost:8080/products/1
```

---

## 🧩 Project Structure

```
.
├── config/           # Database and app configuration
├── controller/       # Handler functions for routes
├── migrations/       # Database migration scripts
├── model/            # Database models
├── server.go         # Application entrypoint
├── go.mod            # Go module file
└── go.sum            # Go checksum file
```

---

## 🗄 Database Migrations (PostgreSQL)

Your migrations are stored in the `migrations/` folder:

```
migrations/
├── 001_create_products_table.up.sql
└── 001_create_products_table.down.sql
```

### Install Migrate CLI

```bash
# macOS
brew install golang-migrate
# Ubuntu
sudo apt-get install migrate
```

### Run Migrations

```bash
migrate -path ./migrations -database "postgres://your_user:your_pass@localhost:5432/echo_products_api?sslmode=disable" up
```

### Rollback Migration

```bash
migrate -path ./migrations -database "postgres://your_user:your_pass@localhost:5432/echo_products_api?sslmode=disable" down
```

### Check Migration Version

```bash
migrate -path ./migrations -database "postgres://your_user:your_pass@localhost:5432/echo_products_api?sslmode=disable" version
```

---

## 📜 License

This project is licensed under the **Apache License 2.0**.
