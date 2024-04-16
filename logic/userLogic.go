package logic

import (
	"authenticaiton-authorization/entity"
	"authenticaiton-authorization/helper"
	"authenticaiton-authorization/repo"
	"authenticaiton-authorization/utils"
	"log"
)

type UserLogic struct {
	*AbstractLogic
}

var userRepo = repo.UserRepoPsql{}

func (logic *UserLogic) Create(model entity.IModel) (*string, error) {
	var user entity.User

	if _user, ok := model.(*entity.User); ok {
		user = *_user
	}
	err := user.ValidationModel()
	if err != nil {
		return nil, err
	}
	err = userRepo.ExistUser(user)
	if err != nil {
		return nil, err
	}
	user.Password, err = helper.GeneratePassword(user.Password)
	if err != nil {
		log.Println("Not Generated Password")
		return nil, err
	}
	newUserId, err := userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	newUser, err := userRepo.FindById(*newUserId)
	if err != nil {
		return nil, err
	}
	token, err := helper.CreateJWTToken(newUser)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (logic *UserLogic) Login(model entity.IModel) (*string, error) {
	var user entity.User
	if _user, ok := model.(*entity.User); ok {
		user = *_user
	}
	logUser, err := userRepo.GetUserByUsernameOrEmail(user)
	if err != nil {
		return nil, err
	}
	checkUser, err := helper.ComparePasswordAndHash(*user.Password, *logUser.Password)
	if !checkUser {
		return nil, &utils.BusinessException{
			Message: "Wrong Password",
		}
	}
	token, err := helper.CreateJWTToken(logUser)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
