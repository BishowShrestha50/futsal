package controller

import (
	"database/sql"
	"fmt"
	"futsal/model"
	"os"

	"futsal/usecase"

	"futsal/repository"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	DB      *sql.DB
	Router  *mux.Router
	UseCase model.FutsalUsecaseInterface
}

func NewServer(db *sql.DB) *Server {
	repo := repository.NewRepository(db)
	useCase := usecase.NewUsecase(repo)
	server := &Server{
		DB:      db,
		Router:  mux.NewRouter(),
		UseCase: useCase,
	}
	server.initializeRoutes()
	logrus.Info("Server started...")
	return server
}

func ConnectDB() *sql.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%v:%v)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()
	return db
}

func MigrateDB(db *sql.DB) error {
	// Create the players table if it doesn't exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS teammate (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			location VARCHAR(255),
			email VARCHAR(255),
			phone VARCHAR(20),
			playing_position VARCHAR(50),
			age VARCHAR(20)
		)
	`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return err
	}

	createTableOpponent := `
	CREATE TABLE IF NOT EXISTS opponent (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		location VARCHAR(255),
		email VARCHAR(255),
		phone VARCHAR(20),
		age VARCHAR(20)
	)
`
	_, err = db.Exec(createTableOpponent)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return err
	}

	createTableTeam := `
		CREATE TABLE IF NOT EXISTS team (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			location VARCHAR(255),
			email VARCHAR(255),
			phone VARCHAR(20),
			playing_position VARCHAR(50),
			age VARCHAR(20)
		)
	`
	_, err = db.Exec(createTableTeam)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return err
	}

	createTableUser := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255),
		password VARCHAR(255),
		email VARCHAR(255),
		phone VARCHAR(255)
	)
`
	_, err = db.Exec(createTableUser)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return err
	}

	createTableBookings := `
	CREATE TABLE IF NOT EXISTS bookings (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		phone VARCHAR(255),
		user_id INT,
		start_time TIMESTAMP,
		end_time TIMESTAMP
	)
`
	_, err = db.Exec(createTableBookings)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return err
	}

	createTableBookFutsal := `
	CREATE TABLE IF NOT EXISTS bookfutsal (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		phone VARCHAR(255),
		user_id INT,
		date VARCHAR(255),
		time VARCHAR(255),
		futsal VARCHAR(255)
	)
`
	_, err = db.Exec(createTableBookFutsal)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return err
	}

	createTableFutsal := `
	CREATE TABLE IF NOT EXISTS futsal (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		phone VARCHAR(255),
		price VARCHAR(255),
		location VARCHAR(255)
	)
`
	_, err = db.Exec(createTableFutsal)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return err
	}
	return nil
}
