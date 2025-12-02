package store

import (
	"movie-booking/models"
	"time"
)

// Private storage slice
var bookings = []models.Booking{
	{ID: "1", Movie: "Inception", MovieNumber: "MOV-001", Seat: "A1", User: "Alice", IsActive: true, CreatedAt: time.Now()},
	{ID: "2", Movie: "Avatar", MovieNumber: "MOV-002", Seat: "B5", User: "Bob", IsActive: true, CreatedAt: time.Now()},
	{ID: "3", Movie: "The Dark Knight", MovieNumber: "MOV-003", Seat: "K9", User: "Bruce", IsActive: true, CreatedAt: time.Now()},
	{ID: "4", Movie: "Titanic", MovieNumber: "MOV-004", Seat: "C12", User: "Rose", IsActive: false, CreatedAt: time.Now()},
	{ID: "5", Movie: "Avengers: Endgame", MovieNumber: "MOV-005", Seat: "F1", User: "Tony", IsActive: true, CreatedAt: time.Now()},
	{ID: "6", Movie: "The Matrix", MovieNumber: "MOV-006", Seat: "M1", User: "Neo", IsActive: true, CreatedAt: time.Now()},
}

// Add a new booking
func AddBooking(b models.Booking) models.Booking {
	b.IsActive = true
	b.CreatedAt = time.Now()
	bookings = append(bookings, b)
	return b
}

// Get all active bookings
func GetAllActive() []models.Booking {
	active := []models.Booking{}
	for _, b := range bookings {
		if b.IsActive {
			active = append(active, b)
		}
	}
	return active
}

// Get one active booking
func GetOneActive(id string) (models.Booking, bool) {
	for _, b := range bookings {
		if b.ID == id && b.IsActive {
			return b, true
		}
	}
	return models.Booking{}, false
}

// Soft delete a booking
func DeleteBooking(id string) bool {
	for i, b := range bookings {
		if b.ID == id {
			bookings[i].IsActive = false
			return true
		}
	}
	return false
}

// Update booking (Handles both PUT and PATCH logic)
func UpdateBooking(id string, data models.Booking, isPatch bool) (models.Booking, bool) {
	for i, b := range bookings {
		if b.ID == id && b.IsActive {

			if isPatch {
				// PATCH Logic: Only update fields that are present
				if data.Movie != "" {
					bookings[i].Movie = data.Movie
				}
				if data.MovieNumber != "" {
					bookings[i].MovieNumber = data.MovieNumber
				}
				if data.Seat != "" {
					bookings[i].Seat = data.Seat
				}
				if data.User != "" {
					bookings[i].User = data.User
				}
			} else {
				// PUT Logic: Replace all editable fields
				bookings[i].Movie = data.Movie
				bookings[i].MovieNumber = data.MovieNumber
				bookings[i].Seat = data.Seat
				bookings[i].User = data.User
			}
			return bookings[i], true
		}
	}
	return models.Booking{}, false
}
