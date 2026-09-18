# RESTful API Golang - Category Management

Dokumentasi pembuatan aplikasi **RESTful API CRUD sederhana** menggunakan Golang. Proyek ini berfokus pada pemahaman konsep dasar arsitektur RESTful API, validasi data, koneksi database, serta pengamanan endpoint menggunakan API-Key.

---

## 📌 Ruang Lingkup Proyek

* **Tujuan Utama:** Mempelajari implementasi RESTful API di Go (fokus pada arsitektur & *best practices*).
* **Entitas Data:** `Category`
    * `id` (`number` / `int`)
    * `name` (`string`)
* **Fitur Utama (CRUD):**
    * **Create Category** (Membuat kategori baru)
    * **Get Category** (Mengambil data kategori berdasarkan ID)
    * **List Category** (Mengambil seluruh data kategori)
    * **Update Category** (Memperbarui data kategori)
    * **Delete Category** (Menghapus data kategori)
* **Keamanan:** Setiap *endpoint* dilindungi menggunakan Autentikasi berbasis **API-Key** pada *Header*.

---

## 📦 Dependency (Pustaka Pihak Ketiga)

Proyek ini menggunakan beberapa dependensi utama untuk mendukung fungsionalitas database, *routing*, dan validasi data:

| Library | Deskripsi / Kegunaan |
| :--- | :--- |
| **`go-sql-driver/mysql`** | Driver resmi untuk komunikasi dari Golang ke database MySQL. |
| **`julienschmidt/httprouter`** | Router HTTP berkecepatan tinggi & minimalis untuk menangani routing RESTful API. |
| **`go-playground/validator/v10`** | Library validasi data request (misal: memastikan field `name` tidak kosong). |

### Perintah Instalasi

Jalankan perintah berikut di terminal kamu:

```bash
# Driver MySQL
go get -u [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

# HttpRouter untuk Routing API
go get [github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

# Validator v10 untuk Validasi Input Struct
go get [github.com/go-playground/validator/v10](https://github.com/go-playground/validator/v10)
```