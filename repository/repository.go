package repository

import (
	"database/sql"
	"errors"
	"futsal/model"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) model.FutsalUsecaseInterface {
	return &Repository{DB: db}
}

func (r *Repository) GetBookingsByDateTimeAndName(date, time, futsal string) (*[]model.BookFutsal, error) {
	rows, err := r.DB.Query("SELECT * FROM bookfutsal WHERE date = ? AND time = ? AND futsal = ?", date, time, futsal)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var bookings []model.BookFutsal
	for rows.Next() {
		var booking model.BookFutsal
		err := rows.Scan(&booking.ID, &booking.Name, &booking.Phone, &booking.UserID, &booking.Date, &booking.Time, &booking.Futsal)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	return &bookings, nil

}
func (r *Repository) SaveBookings(booking model.BookFutsal) error {
	_, err := r.DB.Exec("INSERT INTO bookfutsal (name, phone,user_id,date,time,futsal) VALUES (?, ?,?,?,?,?)", booking.Name, booking.Phone, booking.UserID, booking.Date, booking.Time, booking.Futsal)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllBookings() (*[]model.BookFutsal, error) {
	var datas []model.BookFutsal
	result, err := r.DB.Query("SELECT *  from bookfutsal")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var data model.BookFutsal
		err := result.Scan(&data.ID, &data.Name, &data.Phone, &data.UserID, &data.Date, &data.Time, &data.Futsal)
		if err != nil {
			panic(err.Error())
		}
		datas = append(datas, data)
	}
	return &datas, nil
}

func (r *Repository) GetBookingsByName(name string) (*[]model.BookFutsal, error) {
	var datas []model.BookFutsal
	result, err := r.DB.Query("SELECT *  from bookfutsal WHERE futsal = ?", name)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var data model.BookFutsal
		err := result.Scan(&data.ID, &data.Name, &data.Phone, &data.UserID, &data.Date, &data.Time, &data.Futsal)
		if err != nil {
			panic(err.Error())
		}
		datas = append(datas, data)
	}
	return &datas, nil
}

func (r *Repository) GetBookingsByUser(id int) (*[]model.BookFutsal, error) {
	var datas []model.BookFutsal
	result, err := r.DB.Query("SELECT *  from bookfutsal where user_id = ?", int(id))
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var data model.BookFutsal
		err := result.Scan(&data.ID, &data.Name, &data.Phone, &data.UserID, &data.Date, &data.Time, &data.Futsal)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return &datas, nil
}

func (u *Repository) DeleteBookings(id int) (int64, error) {
	// Prepare the DELETE statement
	stmt, err := u.DB.Prepare("DELETE FROM bookfutsal WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the DELETE statement
	result, err := stmt.Exec(int(id))
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil

}

func (r *Repository) GetBookingsByDateAndName(date, name string) (*[]model.BookFutsal, error) {
	result, err := r.DB.Query("SELECT *  from bookfutsal where futsal = ? AND date = ?", name, date)
	if err != nil {
		return nil, err
	}
	var totalFutsal []model.BookFutsal
	for result.Next() {
		var bookFutsal model.BookFutsal
		err := result.Scan(&bookFutsal.ID, &bookFutsal.Name, &bookFutsal.Phone, &bookFutsal.UserID, &bookFutsal.Date, &bookFutsal.Time, &bookFutsal.Futsal)
		if err != nil {
			return nil, errors.New("time slot not available")
		}
		totalFutsal = append(totalFutsal, bookFutsal)
	}
	return &totalFutsal, nil
}

func (r *Repository) SaveFutsal(futsal model.Futsal) error {
	stmt, err := r.DB.Prepare("INSERT INTO futsal(name, phone,price,location) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(futsal.Name, futsal.Phone, futsal.Price, futsal.Location)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllFutsal() (*[]model.Futsal, error) {
	var datas []model.Futsal
	result, err := r.DB.Query("SELECT *  from futsal")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var data model.Futsal
		err := result.Scan(&data.ID, &data.Name, &data.Phone, &data.Price, &data.Location)
		if err != nil {
			panic(err.Error())
		}
		datas = append(datas, data)
	}
	return &datas, err
}

func (r *Repository) DeleteFutsal(id int) (int64, error) {
	// Prepare the DELETE statement
	stmt, err := r.DB.Prepare("DELETE FROM futsal WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the DELETE statement
	result, err := stmt.Exec(int(id))
	if err != nil {
		return 0, err
	}

	// Get the number of rows affected by the DELETE statement
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// func (u *UseCase) SaveTeam(team model.Teammate) (*model.Teammate, error) {

// }
