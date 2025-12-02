package handlers

import (
	"encoding/json"
	"movie-booking/models"
	"movie-booking/store"
	"net/http"
)

// Req 1: Create
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBooking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&newBooking); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result := store.AddBooking(newBooking)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// Req 2: Update (PUT & PATCH)
func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	var updateData models.Booking
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Check if this is a PATCH request
	isPatch := (r.Method == http.MethodPatch)

	updatedBooking, success := store.UpdateBooking(id, updateData, isPatch)

	if !success {
		http.Error(w, "Booking not found or inactive", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updatedBooking)
}

// Req 3: Get All Active
func GetBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := store.GetAllActive()
	json.NewEncoder(w).Encode(data)
}

// Req 4: Get One Active
func GetBookingById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	booking, success := store.GetOneActive(id)

	if !success {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(booking)
}

// Req 5: Delete (Soft Delete)
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	success := store.DeleteBooking(id)

	if !success {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
