package entity

import (
	"authenticaiton-authorization/utils"
	"errors"
)

type User struct {
	AbstractEntity
	UserName    *string   `json:"username" db:"username"`
	FirstName   *string   `json:"firstName" db:"first_name"`
	LastName    *string   `json:"lastName" db:"last_name"`
	Email       *string   `json:"email" db:"email"`
	FullName    *string   `json:"fullName" db:"full_name"`
	Password    *string   `json:"password" db:"password"`
	PhoneNumber *string   `json:"phoneNumber" db:"phone_number"`
	Verify      bool      `json:"verify" db:"is_verify"`
	Roles       RoleSlice `json:"roles"`
}

func (u *User) ValidationModel() error {
	if !utils.CheckString(u.UserName) {
		return &utils.NotNullException{
			Message: "Username",
		}
	}
	if !utils.CheckString(u.Email) {
		return &utils.NotNullException{
			Message: "Username",
		}
	}
	if !utils.CheckString(u.Password) {
		return &utils.NotNullException{
			Message: "Password",
		}
	}
	if !utils.CheckString(u.FirstName) {
		return &utils.NotNullException{
			Message: "Password",
		}
	}
	if !utils.CheckString(u.LastName) {
		return &utils.NotNullException{
			Message: "Password",
		}
	}
	return nil
}
func (u *User) CompareUniqueColumns(model IModel) error {
	compareUser, ok := model.(*User)
	if !ok {
		return errors.New("model is not of type User")
	}
	if *u.UserName == *compareUser.UserName {
		return &utils.UniqueException{
			Message: "Username",
		}
	}
	if *u.Email == *compareUser.Email {
		return &utils.UniqueException{Message: "Email"}
	}
	if *u.PhoneNumber == *compareUser.PhoneNumber {
		return &utils.UniqueException{
			Message: "Phone Number",
		}
	}
	return nil
}
