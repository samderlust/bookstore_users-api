package users

import (
	"fmt"

	"github.com/samderlust/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestErr {
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewNotFoundError(fmt.Sprintf("user %d already exists", user.ID))
	}
	userDB[user.ID] = user
	return nil
}

func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.LastName = result.LastName
	user.FirstName = result.FirstName
	return nil
}
