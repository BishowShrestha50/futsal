package usecase

import (
	"errors"
	"futsal/model"
)

type UseCase struct {
	Repo model.FutsalRepositoryInterface
}

func NewUsecase(repo model.FutsalUsecaseInterface) model.FutsalUsecaseInterface {
	return &UseCase{Repo: repo}
}

func (u *UseCase) GetBookingsByDateTimeAndName(date, time, futsal string) (*[]model.BookFutsal, error) {
	bookings, err := u.Repo.GetBookingsByDateTimeAndName(date, time, futsal)
	if err != nil {
		return nil, err
	}
	if len(*bookings) != 0 {
		return nil, errors.New("time Slot Not Available")
	}
	return bookings, nil
}

func (u *UseCase) SaveBookings(data model.BookFutsal) error {
	err := u.Repo.SaveBookings(data)
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) GetAllBookings() (*[]model.BookFutsal, error) {
	data, err := u.Repo.GetAllBookings()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UseCase) GetBookingsByName(name string) (*[]model.BookFutsal, error) {
	data, err := u.Repo.GetBookingsByName(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UseCase) GetBookingsByUser(id int) (*[]model.BookFutsal, error) {
	data, err := u.Repo.GetBookingsByUser(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UseCase) DeleteBookings(id int) (int64, error) {
	data, err := u.Repo.DeleteBookings(id)
	if err != nil {
		return 0, err
	}
	return data, nil
}

func (u *UseCase) GetBookingsByDateAndName(date, name string) (*[]model.BookFutsal, error) {
	data, err := u.Repo.GetBookingsByDateAndName(date, name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UseCase) SaveFutsal(futsal model.Futsal) error {
	err := u.Repo.SaveFutsal(futsal)
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) GetAllFutsal() (*[]model.Futsal, error) {
	data, err := u.Repo.GetAllFutsal()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UseCase) DeleteFutsal(id int) (int64, error) {
	data, err := u.Repo.DeleteFutsal(id)
	if err != nil {
		return 0, err
	}
	return data, nil
}

// func (u *UseCase) SaveTeam(team model.Teammate) (*model.Teammate, error) {

// }
