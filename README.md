# Product API

RESTful API untuk mengelola produk dan kategori menggunakan Go (Golang) dengan Fiber framework dan PostgreSQL database.

## Fitur

- ✅ CRUD operations untuk Categories
- ✅ CRUD operations untuk Products
- ✅ Transaction/Checkout system dengan validasi stok
- ✅ Report summary transaksi harian
- ✅ Relasi antara Product dan Category
- ✅ Health check endpoint
- ✅ PostgreSQL database dengan foreign key constraints

## Teknologi

- **Language**: Go 1.25
- **Framework**: Fiber v2
- **Database**: PostgreSQL
- **ORM**: database/sql (native Go)

## Setup

### Prerequisites

- Go 1.25 atau lebih tinggi
- PostgreSQL database
- Git

### Installation

1. Clone repository:

```bash
git clone https://github.com/fadilahonespot/product-api
cd product-api
```

2. Install dependencies:

```bash
go mod download
```

3. Setup environment variables:
   Buat file `.env` di root directory:

```env
PORT=8080
DB_CONN=postgres://username:password@host:port/database?sslmode=require
```

4. Setup database:
   Jalankan migrasi database secara berurutan:

```bash
# Connect ke PostgreSQL database
psql -h <host> -U <username> -d <database>

# Jalankan migrasi
\i migrations/001_create_categories_table.sql
\i migrations/002_create_products_table.sql
\i migrations/003_create_transactions_table.sql
```

Atau menggunakan psql command line:

```bash
psql $DB_CONN -f migrations/001_create_categories_table.sql
psql $DB_CONN -f migrations/002_create_products_table.sql
psql $DB_CONN -f migrations/003_create_transactions_table.sql
```

5. Run application:

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

---

## Docker Setup

### Prerequisites

- Docker dan Docker Compose terinstall

### Build dan Run dengan Docker

1. **Build Docker image:**

```bash
docker build -t product-api .
```

2. **Run dengan Docker:**

```bash
docker run -p 8080:8080 \
  -e PORT=8080 \
  -e DB_CONN="postgres://username:password@host:port/database?sslmode=require" \
  product-api
```

3. **Run dengan Docker Compose (Recommended):**

```bash
# Pastikan file .env sudah ada dengan konfigurasi yang benar
docker-compose up -d

# Lihat logs
docker-compose logs -f

# Stop container
docker-compose down
```

**Note:** Pastikan file `.env` sudah dikonfigurasi dengan benar sebelum menjalankan docker-compose.

### Docker Commands

```bash
# Build image
docker build -t product-api .

# Run container
docker run -p 8080:8080 --env-file .env product-api

# Stop container
docker stop $(docker ps -q --filter ancestor=product-api)

# Remove container dan image
docker rmi product-api

# View logs
docker-compose logs -f app
```

---

## API Endpoints

### Health Check

#### GET /health

Cek status server

**Response:**

```json
{
  "status": "ok",
  "message": "Server is running",
  "version": "1.0.0",
  "timestamp": "2026-02-01T10:00:00Z"
}
```

---

## Category Endpoints

### Get All Categories

#### GET /api/category

Mendapatkan semua kategori

**Response:** `200 OK`

```json
[
  {
    "id": 1,
    "name": "Electronics",
    "description": "Electronic devices and gadgets"
  },
  {
    "id": 2,
    "name": "Clothing",
    "description": "Apparel and accessories"
  }
]
```

**Error Response:** `500 Internal Server Error`

```json
{
  "message": "Failed to get categories"
}
```

---

### Get Category by ID

#### GET /api/category/:id

Mendapatkan kategori berdasarkan ID

**Parameters:**

- `id` (path parameter) - ID kategori

**Response:** `200 OK`

```json
{
  "id": 1,
  "name": "Electronics",
  "description": "Electronic devices and gadgets"
}
```

**Error Responses:**

`400 Bad Request`

```json
{
  "message": "Invalid category ID"
}
```

`404 Not Found`

```json
{
  "message": "Category not found"
}
```

---

### Create Category

#### POST /api/category

Membuat kategori baru

**Request Body:**

```json
{
  "name": "Electronics",
  "description": "Electronic devices and gadgets"
}
```

**Response:** `201 Created`

```json
{
  "id": 1,
  "name": "Electronics",
  "description": "Electronic devices and gadgets"
}
```

**Error Responses:**

`400 Bad Request`

```json
{
  "message": "Invalid request body"
}
```

`400 Bad Request`

```json
{
  "message": "Failed to create category"
}
```

---

### Update Category

#### PUT /api/category/:id

Update kategori berdasarkan ID

**Parameters:**

- `id` (path parameter) - ID kategori

**Request Body:**

```json
{
  "name": "Updated Electronics",
  "description": "Updated description"
}
```

**Response:** `200 OK`

```json
{
  "id": 1,
  "name": "Updated Electronics",
  "description": "Updated description"
}
```

**Error Responses:**

`400 Bad Request`

```json
{
  "message": "Invalid category ID"
}
```

`400 Bad Request`

```json
{
  "message": "Invalid request body"
}
```

`400 Bad Request`

```json
{
  "message": "Failed to update category"
}
```

---

### Delete Category

#### DELETE /api/category/:id

Menghapus kategori berdasarkan ID

**Parameters:**

- `id` (path parameter) - ID kategori

**Response:** `200 OK`

```json
{
  "message": "Category deleted successfully"
}
```

**Error Responses:**

`400 Bad Request`

```json
{
  "message": "Invalid category ID"
}
```

`500 Internal Server Error`

```json
{
  "message": "Failed to delete category"
}
```

**Note:** Jika kategori masih digunakan oleh produk, penghapusan akan gagal karena foreign key constraint.

---

## Product Endpoints

### Get All Products

#### GET /api/product

Mendapatkan semua produk

**Response:** `200 OK`

```json
[
  {
    "id": 1,
    "name": "Laptop",
    "price": 10000000,
    "stock": 10,
    "category_id": 1,
    "category": {
      "id": 1,
      "name": "Electronics",
      "description": "Electronic devices and gadgets"
    }
  },
  {
    "id": 2,
    "name": "T-Shirt",
    "price": 150000,
    "stock": 50,
    "category_id": 2,
    "category": {
      "id": 2,
      "name": "Clothing",
      "description": "Apparel and accessories"
    }
  }
]
```

**Error Response:** `500 Internal Server Error`

```json
{
  "message": "Failed to get products"
}
```

---

### Get Product by ID

#### GET /api/product/:id

Mendapatkan produk berdasarkan ID

**Parameters:**

- `id` (path parameter) - ID produk

**Response:** `200 OK`

```json
{
  "id": 1,
  "name": "Laptop",
  "price": 10000000,
  "stock": 10,
  "category_id": 1,
  "category": {
    "id": 1,
    "name": "Electronics",
    "description": "Electronic devices and gadgets"
  }
}
```

**Error Responses:**

`400 Bad Request`

```json
{
  "message": "Invalid product ID"
}
```

`404 Not Found`

```json
{
  "message": "Product not found"
}
```

---

### Create Product

#### POST /api/product

Membuat produk baru

**Request Body:**

```json
{
  "name": "Laptop",
  "price": 10000000,
  "stock": 10,
  "category_id": 1
}
```

**Response:** `201 Created`

```json
{
  "id": 1,
  "name": "Laptop",
  "price": 10000000,
  "stock": 10,
  "category_id": 1,
  "category": {
    "id": 1,
    "name": "Electronics",
    "description": "Electronic devices and gadgets"
  }
}
```

**Error Responses:**

`400 Bad Request`

```json
{
  "message": "Invalid request body"
}
```

`400 Bad Request`

```json
{
  "message": "Category not found"
}
```

**Note:** `category_id` harus mengacu ke kategori yang sudah ada di database.

---

### Update Product

#### PUT /api/product/:id

Update produk berdasarkan ID

**Parameters:**

- `id` (path parameter) - ID produk

**Request Body:**

```json
{
  "name": "Updated Laptop",
  "price": 12000000,
  "stock": 15,
  "category_id": 1
}
```

**Response:** `200 OK`

```json
{
  "id": 1,
  "name": "Updated Laptop",
  "price": 12000000,
  "stock": 15,
  "category_id": 1,
  "category": {
    "id": 1,
    "name": "Electronics",
    "description": "Electronic devices and gadgets"
  }
}
```

**Error Responses:**

`400 Bad Request`

```json
{
  "message": "Invalid product ID"
}
```

`400 Bad Request`

```json
{
  "message": "Invalid request body"
}
```

`400 Bad Request`

```json
{
  "message": "Product not found"
}
```

`400 Bad Request`

```json
{
  "message": "Category not found"
}
```

---

### Delete Product

#### DELETE /api/product/:id

Menghapus produk berdasarkan ID

**Parameters:**

- `id` (path parameter) - ID produk

**Response:** `200 OK`

```json
{
  "message": "Product deleted successfully"
}
```

**Error Responses:**

`400 Bad Request`

```json
{
  "message": "Invalid product ID"
}
```

`500 Internal Server Error`

```json
{
  "message": "Failed to delete product"
}
```

---

## Transaction Endpoints

### Checkout (Create Transaction)

#### POST /api/checkout

Membuat transaksi baru (checkout) dengan validasi stok otomatis. Stok produk akan otomatis dikurangi setelah transaksi berhasil dibuat.

**Request Body:**

```json
{
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 2,
      "quantity": 1
    }
  ]
}
```

**Response:** `200 OK`

```json
{
  "id": 1,
  "total_amount": 20150000,
  "created_at": "2026-02-01T10:30:00Z",
  "details": [
    {
      "id": 1,
      "transaction_id": 1,
      "product_id": 1,
      "product_name": "Laptop",
      "quantity": 2,
      "subtotal": 20000000
    },
    {
      "id": 2,
      "transaction_id": 1,
      "product_id": 2,
      "product_name": "T-Shirt",
      "quantity": 1,
      "subtotal": 150000
    }
  ]
}
```

**Error Responses:**

`400 Bad Request`

```json
{
  "message": "Invalid request body"
}
```

`400 Bad Request`

```json
{
  "message": "product stock not enough"
}
```

**Note:** 
- Transaksi menggunakan database transaction untuk memastikan atomicity
- Stok produk akan otomatis dikurangi setelah transaksi berhasil
- Jika salah satu produk stok tidak cukup, seluruh transaksi akan di-rollback

---

### Get Transaction Summary (Hari Ini)

#### GET /api/report/hari-ini

Mendapatkan ringkasan transaksi untuk hari ini (dari 00:00:00 sampai 23:59:59).

**Response:** `200 OK`

```json
{
  "total_revenue": 20150000,
  "total_transaksi": 5,
  "produk_terlaris": {
    "nama": "Laptop",
    "qty_terjual": 10
  }
}
```

**Error Response:** `500 Internal Server Error`

```json
{
  "message": "Failed to get summary"
}
```

---

## Data Models

### Category

```json
{
  "id": 1,
  "name": "string",
  "description": "string"
}
```

**Fields:**

- `id` (integer) - Primary key, auto-increment
- `name` (string, required) - Nama kategori
- `description` (string, optional) - Deskripsi kategori

### Product

```json
{
  "id": 1,
  "name": "string",
  "price": 0,
  "stock": 0,
  "category_id": 1,
  "category": {
    "id": 1,
    "name": "string",
    "description": "string"
  }
}
```

**Fields:**

- `id` (integer) - Primary key, auto-increment
- `name` (string, required) - Nama produk
- `price` (integer, required) - Harga produk
- `stock` (integer, required) - Stok produk (default: 0)
- `category_id` (integer, required) - Foreign key ke categories table
- `category` (object, optional) - Object kategori (populated saat GET)

### Transaction

```json
{
  "id": 1,
  "total_amount": 20150000,
  "created_at": "2026-02-01T10:30:00Z",
  "details": [
    {
      "id": 1,
      "transaction_id": 1,
      "product_id": 1,
      "product_name": "Laptop",
      "quantity": 2,
      "subtotal": 20000000
    }
  ]
}
```

**Fields:**

- `id` (integer) - Primary key, auto-increment
- `total_amount` (integer, required) - Total jumlah transaksi
- `created_at` (string) - Timestamp transaksi dibuat
- `details` (array, optional) - Array detail transaksi

### TransactionDetail

```json
{
  "id": 1,
  "transaction_id": 1,
  "product_id": 1,
  "product_name": "Laptop",
  "quantity": 2,
  "subtotal": 20000000
}
```

**Fields:**

- `id` (integer) - Primary key, auto-increment
- `transaction_id` (integer, required) - Foreign key ke transactions table
- `product_id` (integer, required) - Foreign key ke products table
- `product_name` (string, optional) - Nama produk (populated saat GET)
- `quantity` (integer, required) - Jumlah produk
- `subtotal` (integer, required) - Subtotal (price × quantity)

### CheckoutRequest

```json
{
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    }
  ]
}
```

**Fields:**

- `items` (array, required) - Array item yang akan di-checkout
  - `product_id` (integer, required) - ID produk
  - `quantity` (integer, required) - Jumlah produk

### SummaryResponse

```json
{
  "total_revenue": 20150000,
  "total_transaksi": 5,
  "produk_terlaris": {
    "nama": "Laptop",
    "qty_terjual": 10
  }
}
```

**Fields:**

- `total_revenue` (integer) - Total pendapatan hari ini
- `total_transaksi` (integer) - Jumlah transaksi hari ini
- `produk_terlaris` (object) - Produk terlaris hari ini
  - `nama` (string) - Nama produk
  - `qty_terjual` (integer) - Jumlah terjual

---

## Error Handling

API menggunakan HTTP status codes standar:

- `200 OK` - Request berhasil
- `201 Created` - Resource berhasil dibuat
- `400 Bad Request` - Request tidak valid
- `404 Not Found` - Resource tidak ditemukan
- `500 Internal Server Error` - Server error

Semua error response mengikuti format:

```json
{
  "message": "Error message description"
}
```

---

## Database Schema

### Categories Table

```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Products Table

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    category_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_products_category FOREIGN KEY (category_id) 
        REFERENCES categories(id) ON DELETE RESTRICT
);
```

### Transactions Table

```sql
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    total_amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Transaction Details Table

```sql
CREATE TABLE transaction_details (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id),
    quantity INT NOT NULL,
    subtotal INT NOT NULL
);
```

---

## Testing dengan cURL

### Health Check

```bash
curl http://localhost:8080/health
```

### Create Category

```bash
curl -X POST http://localhost:8080/api/category \
  -H "Content-Type: application/json" \
  -d '{"name":"Electronics","description":"Electronic devices"}'
```

### Get All Categories

```bash
curl http://localhost:8080/api/category
```

### Create Product

```bash
curl -X POST http://localhost:8080/api/product \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","price":10000000,"stock":10,"category_id":1}'
```

### Get All Products

```bash
curl http://localhost:8080/api/product
```

### Checkout (Create Transaction)

```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {"product_id": 1, "quantity": 2},
      {"product_id": 2, "quantity": 1}
    ]
  }'
```

### Get Transaction Summary (Hari Ini)

```bash
curl http://localhost:8080/api/report/hari-ini
```

---

## License

MIT License
