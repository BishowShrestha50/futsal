package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"futsal/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	util "futsal/usecase/utils"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (server *Server) RegisterOpponent(w http.ResponseWriter, r *http.Request) {
	stmt, err := server.DB.Prepare("INSERT INTO opponent(name, location, email, phone, age) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	opponent := model.Opponent{}

	json.Unmarshal(body, &opponent)
	_, err = stmt.Exec(opponent.Name, opponent.Location, opponent.Email, opponent.Phone, opponent.Age)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
}

func (server *Server) RegisterTeam(w http.ResponseWriter, r *http.Request) {
	stmt, err := server.DB.Prepare("INSERT INTO team(name, location, email, phone, playing_position, age) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	teamMate := model.Teammate{}
	json.Unmarshal(body, &teamMate)
	_, err = stmt.Exec(teamMate.Name, teamMate.Location, teamMate.Email, teamMate.Phone, teamMate.PlayingPositon, teamMate.Age)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
}

func (server *Server) GetOpponent(w http.ResponseWriter, r *http.Request) {
	var datas []model.Opponent
	result, err := server.DB.Query("SELECT *  from opponent")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var data model.Opponent
		err := result.Scan(&data.ID, &data.Name, &data.Location, &data.Email, &data.Phone, &data.Age)
		if err != nil {
			panic(err.Error())
		}
		datas = append(datas, data)
	}
	fmt.Println(datas)
	json.NewEncoder(w).Encode(datas)
}

func (server *Server) GetTeam(w http.ResponseWriter, r *http.Request) {
	var datas []model.Teammate
	result, err := server.DB.Query("SELECT *  from team")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var data model.Teammate
		err := result.Scan(&data.ID, &data.Name, &data.Location, &data.Email, &data.Phone, &data.PlayingPositon, &data.Age)
		if err != nil {
			panic(err.Error())
		}
		datas = append(datas, data)
	}
	fmt.Println(datas)
	json.NewEncoder(w).Encode(datas)
}

func (server *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var users = make(map[string]model.User)
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		var newUser model.User
		json.Unmarshal(body, &newUser)
		logrus.Println(newUser)
		if err != nil {
			log.Println("Error decoding JSON:", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if newUser.Username == "" || newUser.Password == "" {
			http.Error(w, "Please provide both username and password", http.StatusBadRequest)
			return
		}

		if _, exists := users[newUser.Username]; exists {
			http.Error(w, "Username already exists. Please choose a different username", http.StatusBadRequest)
			return
		}
		rows, err := server.DB.Query("SELECT * FROM users where username = ? OR email = ?", newUser.Username, newUser.Email)
		if err != nil {
			logrus.Error("aaa", err)
			util.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		defer rows.Close()
		if rows.Next() {
			util.ERROR(w, http.StatusUnprocessableEntity, errors.New("username or email already taken"))
			return
		}
		_, err = server.DB.Exec("INSERT INTO users (username, password,email,phone) VALUES (?, ?,?,?)", newUser.Username, newUser.Password, newUser.Email, newUser.Phone)
		if err != nil {
			logrus.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users[newUser.Username] = newUser
		util.JSON(w, http.StatusCreated, users)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func (server *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginUser model.User
	err := decoder.Decode(&loginUser)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Query the database for the user
	row := server.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", loginUser.Username)
	var dbUser model.User
	err = row.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if loginUser.Password != dbUser.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := util.GenerateToken(strconv.Itoa(dbUser.ID))
	if err != nil {
		logrus.Error(err)
	}
	tokenDecode, err := util.DecodeToken(token)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("tok", tokenDecode)
	// Successful login
	response := model.Response{
		Message: "Login successful",
		//	"user":    dbUser,
		Token: token,
	}
	if loginUser.Username == "admin" {
		logrus.Info("jb", loginUser.Username)
		response.IsAdmin = true
	} else {

		logrus.Info("jb", loginUser.Username)
		response.IsAdmin = false
	}
	//json.NewEncoder(w).Encode(response)
	util.JSON(w, http.StatusOK, response)
}

func (server *Server) BookFutsals(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
	token := bearerToken[6:]
	userID, _ := util.DecodeToken(token)
	id, _ := strconv.ParseInt(userID, 10, 64)
	var booking model.BookFutsal
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err)
		return
	}
	json.Unmarshal(body, &booking)
	booking.UserID = int(id)
	_, err = server.UseCase.GetBookingsByDateTimeAndName(booking.Date, booking.Time, booking.Futsal)
	if err != nil {
		util.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// Perform booking logic here
	err = server.UseCase.SaveBookings(booking)
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	util.JSON(w, http.StatusOK, booking)
}

func (server *Server) GetFutsalBookings(w http.ResponseWriter, r *http.Request) {
	data, err := server.UseCase.GetAllBookings()
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	util.JSON(w, http.StatusOK, data)
}

func (server *Server) GetBookingsByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	data, err := server.UseCase.GetBookingsByName(name)
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	util.JSON(w, http.StatusOK, data)
}

func (server *Server) GetUserBookings(w http.ResponseWriter, r *http.Request) {

	bearerToken := r.Header.Get("Authorization")
	token := bearerToken[6:]
	userID, _ := util.DecodeToken(token)
	id, _ := strconv.ParseInt(userID, 10, 64)
	datas, err := server.UseCase.GetBookingsByUser(int(id))
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	util.JSON(w, http.StatusOK, datas)
}

func (server *Server) DeleteFutsalBookings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Time Slot Not Available")
		return
	}
	rowsAffected, err := server.UseCase.DeleteBookings(int(pid))
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	logrus.Info("Deleted %d row(s)\n", rowsAffected)
	util.JSON(w, http.StatusOK, rowsAffected)
}

func (server *Server) FutsalSlot(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	futsal := r.URL.Query().Get("futsal")
	logrus.Info("data", date, futsal)
	slot := []string{"6-7 AM", "7-8 AM", "8-9 AM", "9-10 AM", "10-11 AM", "11-12 AM", "12-1 PM", "1-2 PM", "2-3 PM", "3-4 PM", "4-5 PM", "5-6 PM", "6-7 PM", "7-8 PM", "8-9 PM"}
	datas, err := server.UseCase.GetBookingsByDateAndName(date, futsal)
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	logrus.Info("futsal-total", datas)
	// Remove duplicate elements from arr1
	availableSlot := util.RemoveDuplicates(slot, *datas)
	logrus.Info("slot", availableSlot)
	util.JSON(w, http.StatusOK, availableSlot)
}

func (server *Server) RegisterFutsal(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	futsal := model.Futsal{}
	json.Unmarshal(body, &futsal)
	err = server.UseCase.SaveFutsal(futsal)
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	util.JSON(w, http.StatusOK, "Futsal registered")
}

func (server *Server) GetAllFutsal(w http.ResponseWriter, r *http.Request) {
	datas, err := server.UseCase.GetAllFutsal()
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	util.JSON(w, http.StatusOK, datas)
}

func (server *Server) DeleteFutsal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		util.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	rowsAffected, err := server.UseCase.DeleteFutsal(int(pid))
	if err != nil {
		util.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	logrus.Info("Deleted %d row(s)\n", rowsAffected)
	util.JSON(w, http.StatusOK, rowsAffected)
}
