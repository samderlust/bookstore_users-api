package users

import "encoding/json"

type PublicUser struct {
	ID int64 `json:"id"`

	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
}

func (user *User) Marhsall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}

func (users Users) Marhsall(isPublic bool) interface{} {
	res := make([]interface{}, len(users))
	for i, u := range users {
		res[i] = u.Marhsall(isPublic)

	}
	return res
}
