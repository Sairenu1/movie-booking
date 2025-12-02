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
    git clone [https://github.com/Sairenu1/movie-booking-api.git](https://github.com/Sairenu1/movie-booking-api.git)
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

# 🚀 Postman Testing Guide

This guide details how to test the Movie Booking API endpoints using **Postman**.

**Base URL:** `http://localhost:8080`

---

## 1. Create a Booking (POST)
Creates a new active booking.

* **Method:** `POST`
* **URL:** `http://localhost:8080/bookings`
* **Headers:**
    * `Content-Type`: `application/json`
* **Body (raw JSON):**
    ```json
    {
        "id": "101",
        "movie": "Dune 2",
        "movie_number": "MOV-009",
        "seat": "K5",
        "user": "Paul"
    }
    ```
* **Expected Status:** `201 Created`

---

## 2. Get All Active Bookings (GET)
Retrieves the list of all valid bookings.

* **Method:** `GET`
* **URL:** `http://localhost:8080/bookings`
* **Body:** None
* **Expected Status:** `200 OK`

---

## 3. Update Seat Details (PATCH)
Updates specific fields (e.g., seat) without changing the rest of the booking.

* **Method:** `PATCH`
* **URL:** `http://localhost:8080/bookings/101`
* **Headers:**
    * `Content-Type`: `application/json`
* **Body (raw JSON):**
    ```json
    {
        "seat": "K6"
    }
    ```
* **Expected Status:** `200 OK`

---

## 4. Full Update (PUT)
Replaces the entire booking. **Warning:** Missing fields will be set to empty strings.

* **Method:** `PUT`
* **URL:** `http://localhost:8080/bookings/101`
* **Headers:**
    * `Content-Type`: `application/json`
* **Body (raw JSON):**
    ```json
    {
        "id": "101",
        "movie": "Dune 2",
        "movie_number": "MOV-009",
        "seat": "VIP-1",
        "user": "Paul Atreides"
    }
    ```
* **Expected Status:** `200 OK`

---

## 5. Cancel Booking (DELETE)
Performs a "Soft Delete" (marks `is_active` as false).

* **Method:** `DELETE`
* **URL:** `http://localhost:8080/bookings/101`
* **Body:** None
* **Expected Status:** `204 No Content`

---

## 🧪 Testing API Limitations

### Test 1: Data Persistence (Memory Reset)
Since this API uses In-Memory storage, follow these steps to verify that data is lost on restart:

1.  Run the **Create Booking** request (Step 1 above).
2.  Run **Get All Active Bookings** to confirm ID `101` exists.
3.  **Stop the Go Server** (Ctrl+C in terminal).
4.  **Start the Go Server** again (`go run main.go`).
5.  Run **Get All Active Bookings** in Postman.
6.  **Result:** ID `101` will be gone (only hardcoded seed data remains).

### Test 2: Invalid ID Handling
Verify how the API handles non-existent IDs.

1.  **Method:** `GET`
2.  **URL:** `http://localhost:8080/bookings/99999`
3.  **Expected Status:** `404 Not Found`
   
⚠️ Current Limitations 
Persistence: Currently uses an In-Memory store. Data resets when the server stops. Future updates will integrate PostgreSQL.

Concurrency: Not thread-safe for high-load environments. Future updates will implement sync.Mutex.
