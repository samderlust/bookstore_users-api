package users

import (
	"strings"

	"github.com/samderlust/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

//User defined user entity
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

//Validate validate the input for user ENtity
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Password = strings.TrimSpace(strings.ToLower(user.Password))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
