package model

import "time"

type Teammate struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Location       string `json:"location"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	PlayingPositon string `json:"playing_position"`
	Age            string `json:"age"`
}

type Opponent struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Age      string `json:"age"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type BookingData struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Date      string `json:"date"`
	UserID    int    `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type Booking struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID int    `json:"user_id"`
	//Date      time.Time `json:"date"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type BookFutsal struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	Futsal string `json:"futsal"`
}

type Futsal struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Price    string `json:"price"`
	Location string `json:"location"`
}

type Response struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	IsAdmin bool   `json:"is_admin"`
}

type FutsalUsecaseInterface interface {
	GetBookingsByDateTimeAndName(date, time, futsal string) (*[]BookFutsal, error)
	SaveBookings(data BookFutsal) error
	GetAllBookings() (*[]BookFutsal, error)
	GetBookingsByName(name string) (*[]BookFutsal, error)
	GetBookingsByUser(id int) (*[]BookFutsal, error)
	DeleteBookings(id int) (int64, error)
	GetBookingsByDateAndName(date, name string) (*[]BookFutsal, error)
	SaveFutsal(futsal Futsal) error
	GetAllFutsal() (*[]Futsal, error)
	DeleteFutsal(id int) (int64, error)
	// SaveTeam(team Teammate) (*Teammate, error)
}

type FutsalRepositoryInterface interface {
	GetBookingsByDateTimeAndName(date, time, futsal string) (*[]BookFutsal, error)
	SaveBookings(data BookFutsal) error
	GetAllBookings() (*[]BookFutsal, error)
	GetBookingsByName(name string) (*[]BookFutsal, error)
	GetBookingsByUser(id int) (*[]BookFutsal, error)
	DeleteBookings(id int) (int64, error)
	GetBookingsByDateAndName(date, name string) (*[]BookFutsal, error)
	SaveFutsal(futsal Futsal) error
	GetAllFutsal() (*[]Futsal, error)
	DeleteFutsal(id int) (int64, error)
	// SaveTeam(team Teammate) (*Teammate, error)
}
