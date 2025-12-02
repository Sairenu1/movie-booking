# 🎬 Movie Booking API Service

![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-green)
![Status](https://img.shields.io/badge/Status-Active-success)

A lightweight, RESTful API service built with Go (Golang) to manage movie theater bookings. This project demonstrates clean REST architecture, **Soft Delete** implementation, and dependency-free routing using the standard library.

## 🚀 Features

* **CRUD Operations:** Create, Read, Update, and Delete bookings.
* **Soft Delete System:** Cancelled bookings are marked as "Inactive" rather than being permanently erased, preserving audit history.
* **Flexible Updates:** Supports both `PUT` (Full Replace) and `PATCH` (Partial Update).
* **Auto-Timestamping:** Automatically records creation time.
* **Zero Dependencies:** Built entirely using Go's standard `net/http` library.

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Router:** `http.NewServeMux` (Go Standard Library)
* **Database:** In-Memory (Slice-based for rapid prototyping)

## 📦 Getting Started

### Prerequisites
* **Go 1.22 or higher** is required (due to the usage of `r.PathValue`).

### Installation & Run
1.  **Clone the repository**
    ```bash
    git clone [https://github.com/yourusername/movie-booking-api.git](https://github.com/yourusername/movie-booking-api.git)
    cd movie-booking-api
    ```

2.  **Run the server**
    ```bash
    go run main.go
    ```

3.  The server will start at: `http://localhost:8080`

## 📡 API Endpoints

### Base URL: `http://localhost:8080`

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| **GET** | `/bookings` | Get all **active** bookings |
| **GET** | `/bookings/{id}` | Get a specific booking by ID |
| **POST** | `/bookings` | Create a new booking |
| **PATCH**| `/bookings/{id}` | Partially update a booking (e.g., change seat only) |
| **PUT** | `/bookings/{id}` | Fully update a booking (replace details) |
| **DELETE**| `/bookings/{id}` | Cancel a booking (**Soft Delete**) |

---

## 🧪 Testing with cURL

You can test the API directly from your terminal using these commands:

**1. Create a Booking**
```bash
curl -X POST http://localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -d '{"id": "101", "movie": "Dune 2", "movie_number": "MOV-009", "seat": "K5", "user": "Paul"}'
curl http://localhost:8080/bookings
curl -X PATCH http://localhost:8080/bookings/101 \
  -H "Content-Type: application/json" \
  -d '{"seat": "K6"}'
curl -X DELETE http://localhost:8080/bookings/101
{
  "id": "1",
  "movie": "Inception",
  "movie_number": "MOV-001",
  "seat": "A1",
  "user": "Alice",
  "is_active": true,       // System Controlled
  "created_at": "2023..."  // System Controlled
}
⚠️ Current Limitations 
Persistence: Currently uses an In-Memory store. Data resets when the server stops. Future updates will integrate PostgreSQL.

Concurrency: Not thread-safe for high-load environments. Future updates will implement sync.Mutex.
