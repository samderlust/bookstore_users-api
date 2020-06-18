package services

import (
	"github.com/samderlust/bookstore_users-api/utils/crypto_utils"
	"github.com/samderlust/bookstore_users-api/utils/errors"

	"github.com/samderlust/bookstore_users-api/domain/users"
)

var (
	UsersService usersServiceInterface = &userService{}
)

type userService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	FindByStatus(status string) (users.Users, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
}

//CreateUser and store it into database
func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil

}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if !isPartial {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	} else {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

func (s *userService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

func (s *userService) FindByStatus(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
