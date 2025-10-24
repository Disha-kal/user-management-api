package models

import (
	"time"
)

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age,omitempty"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

func CalculateAge(dob string) (int, error) {
	birthDate, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return 0, err
	}

	now := time.Now()
	age := now.Year() - birthDate.Year()

	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	return age, nil
}

// UserFromDB converts database User to API response
func UserFromDB(id int32, name string, dob time.Time) (UserResponse, error) {
	age, err := CalculateAge(dob.Format("2006-01-02"))
	if err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		ID:   id,
		Name: name,
		DOB:  dob.Format("2006-01-02"),
		Age:  age,
	}, nil
}

// UsersFromDB converts slice of database Users to API responses
func UsersFromDB(users []struct {
	ID   int32
	Name string
	Dob  time.Time
}) ([]UserResponse, error) {
	var response []UserResponse
	for _, user := range users {
		userResp, err := UserFromDB(user.ID, user.Name, user.Dob)
		if err != nil {
			return nil, err
		}
		response = append(response, userResp)
	}
	return response, nil
}
