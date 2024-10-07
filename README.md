# Merchant-Bank API

## Deskripsi

Merchant-Bank API adalah backend service yang menyediakan fungsionalitas untuk login, melakukan pembayaran, dan logout bagi pelanggan dan merchant. Semua aktivitas dicatat dalam history log.

## Cara Menjalankan Aplikasi

1. **Clone Repository**

   ```sh
   git clone https://github.com/spacdust/merchant-bank-api.git
   cd merchant-bank-api

   ```

2. **Install Dependensi**

   ```sh
   go get ./...

   ```

3. **Build**

   ```sh
   go build -o merchant-bank-api.exe ./cmd

   ```

4. **Jalankan Aplikasi**

   ```sh
   ./merchant-bank-api.exe
   ```

## Endpoints API

**Login**
POST http://localhost:8080/login
    ```sh
    {
    "id": "cust1",
    "password": "password123"
    }
    ```

**Payment**
POST http://localhost:8080/payment
    ```sh
    {
        "from_id": "cust1",
        "to_id": "cust2",
        "amount": 100.0
    }
    ```

**Logout**
POST http://localhost:8080/logout
