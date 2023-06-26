package controller

import (
	"net/http"

	"futsal/controller/middleware"
)

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")
}

func (server *Server) initializeRoutes() {
	server.Router.Use(middleware.CORS)

	server.setJSON("/team", server.RegisterTeam, "POST")
	server.setJSON("/team", server.GetTeam, "GET")
	server.setJSON("/opponent", server.RegisterOpponent, "POST")
	server.setJSON("/opponent", server.GetOpponent, "GET")
	server.setJSON("/register", server.RegisterHandler, "POST")
	server.setJSON("/bookings", server.BookFutsals, "POST")
	server.setJSON("/slot", server.FutsalSlot, "GET")
	server.setJSON("/bookings", server.GetFutsalBookings, "GET")
	server.setJSON("/bookings/{name}", server.GetBookingsByName, "GET")
	server.setJSON("/userbooking", server.GetUserBookings, "GET")
	server.setJSON("/bookings/{id}", server.DeleteFutsalBookings, "DELETE")
	server.setJSON("/login", server.LoginHandler, "POST")
	server.setJSON("/futsal", server.RegisterFutsal, "POST")
	server.setJSON("/futsal", server.GetAllFutsal, "GET")
	server.setJSON("/futsal/{id}", server.DeleteFutsal, "DELETE")
}
