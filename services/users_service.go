package services

import (
	"github.com/samderlust/bookstore_users-api/utils/errors"

	"github.com/samderlust/bookstore_users-api/domain/users"
)

//CreateUser and store it into database
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil

}
