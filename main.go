package main

import (
	"database/sql"
	"futsal/controller"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	}
	db = controller.ConnectDB()
	server := controller.NewServer(db)
	controller.MigrateDB(db)
	http.ListenAndServe(":8000", server.Router)
}
