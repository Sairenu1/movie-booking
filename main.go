package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// --- Model ---
type Booking struct {
	ID          string    `json:"id"`
	Movie       string    `json:"movie"`
	MovieNumber string    `json:"movie_number"` // New Field
	Seat        string    `json:"seat"`
	User        string    `json:"user"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"` // New Field
}

// --- In-memory database ---
// Updated seed data to include new fields
var bookings = []Booking{
	{
		ID:          "1",
		Movie:       "Inception",
		MovieNumber: "MOV-001",
		Seat:        "A1",
		User:        "Alice",
		IsActive:    true,
		CreatedAt:   time.Now(),
	},
	{
		ID:          "2",
		Movie:       "Avatar",
		MovieNumber: "MOV-002",
		Seat:        "B5",
		User:        "Bob",
		IsActive:    true,
		CreatedAt:   time.Now(),
	},
	{
		ID:          "3",
		Movie:       "The Dark Knight",
		MovieNumber: "MOV-003",
		Seat:        "K9",
		User:        "Bruce",
		IsActive:    true,
		CreatedAt:   time.Now(),
	},
	{
		ID:          "4",
		Movie:       "Titanic",
		MovieNumber: "MOV-004",
		Seat:        "C12",
		User:        "Rose",
		IsActive:    false,
		CreatedAt:   time.Now(),
	},
	{
		ID:          "5",
		Movie:       "Avengers: Endgame",
		MovieNumber: "MOV-005",
		Seat:        "F1",
		User:        "Tony",
		IsActive:    true,
		CreatedAt:   time.Now(),
	},
	{
		ID:          "6",
		Movie:       "The Matrix",
		MovieNumber: "MOV-006",
		Seat:        "M1",
		User:        "Neo",
		IsActive:    true,
		CreatedAt:   time.Now(),
	},
}

// --- Handlers ---

// Req 1: Create a Movie Booking API [POST]
func createBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newBooking Booking
	if err := json.NewDecoder(r.Body).Decode(&newBooking); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Set System Defaults
	newBooking.IsActive = true
	newBooking.CreatedAt = time.Now() // Set current time

	bookings = append(bookings, newBooking)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBooking)
}

// Req 2: Update (Partial and Full)
func updateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	var updateData Booking
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, booking := range bookings {
		// We only allow updating if the ID matches AND the booking is currently active
		if booking.ID == id && booking.IsActive {

			// --- PATCH (Partial Update) ---
			// Check if fields are provided; if yes, update them.
			if r.Method == http.MethodPatch {
				if updateData.Movie != "" {
					bookings[i].Movie = updateData.Movie
				}
				if updateData.MovieNumber != "" { // New Field Logic
					bookings[i].MovieNumber = updateData.MovieNumber
				}
				if updateData.Seat != "" {
					bookings[i].Seat = updateData.Seat
				}
				if updateData.User != "" {
					bookings[i].User = updateData.User
				}
				// Note: We usually do NOT update CreatedAt
			}

			// --- PUT (Full Update) ---
			// Replace all editable fields. Missing fields become empty.
			if r.Method == http.MethodPut {
				bookings[i].Movie = updateData.Movie
				bookings[i].MovieNumber = updateData.MovieNumber // New Field Logic
				bookings[i].Seat = updateData.Seat
				bookings[i].User = updateData.User
				// We preserve ID, CreatedAt, and IsActive from the original record
			}

			// Return the updated object
			json.NewEncoder(w).Encode(bookings[i])
			return
		}
	}

	http.Error(w, "Booking not found or inactive", http.StatusNotFound)
}

// Req 3: Get all bookings (Only Active)
func getBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	activeBookings := []Booking{}

	// Filter: Only append if IsActive is true
	for _, b := range bookings {
		if b.IsActive {
			activeBookings = append(activeBookings, b)
		}
	}

	json.NewEncoder(w).Encode(activeBookings)
}

// Req 4: Get booking by ID (Only Active)
func getBookingById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	for _, booking := range bookings {
		// Check ID AND check if it is Active
		if booking.ID == id && booking.IsActive {
			json.NewEncoder(w).Encode(booking)
			return
		}
	}

	http.Error(w, "Booking not found", http.StatusNotFound)
}

// Req 5: Delete the booking (Soft Delete / IsActive = false)
func deleteBooking(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	for i, b := range bookings {
		if b.ID == id {
			// Instead of removing from slice, we just mark as Inactive
			bookings[i].IsActive = false
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Booking not found", http.StatusNotFound)
}

// --- Main ---
func main() {
	router := http.NewServeMux()

	// Register Routes
	router.HandleFunc("GET /bookings", getBookings)           // Req 3
	router.HandleFunc("GET /bookings/{id}", getBookingById)   // Req 4
	router.HandleFunc("POST /bookings", createBooking)        // Req 1
	router.HandleFunc("DELETE /bookings/{id}", deleteBooking) // Req 5

	// Register Update Routes (Req 2)
	router.HandleFunc("PUT /bookings/{id}", updateBooking)
	router.HandleFunc("PATCH /bookings/{id}", updateBooking)

	println("Movie Booking Server started on port 8080...")
	http.ListenAndServe(":8080", router)
}
