package main

import (
	"movie-booking/handlers"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	// Register Routes
	router.HandleFunc("GET /bookings", handlers.GetBookings)
	router.HandleFunc("GET /bookings/{id}", handlers.GetBookingById)
	router.HandleFunc("POST /bookings", handlers.CreateBooking)
	router.HandleFunc("DELETE /bookings/{id}", handlers.DeleteBooking)
	router.HandleFunc("PUT /bookings/{id}", handlers.UpdateBooking)
	router.HandleFunc("PATCH /bookings/{id}", handlers.UpdateBooking)

	println("Movie Booking Server started on port 8080...")
	http.ListenAndServe(":8080", router)
}
