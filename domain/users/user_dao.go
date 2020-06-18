package users

import (
	"fmt"

	"github.com/samderlust/bookstore_users-api/datasources/mysql/users_db"
	"github.com/samderlust/bookstore_users-api/logger"
	"github.com/samderlust/bookstore_users-api/utils/errors"
	"github.com/samderlust/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users(firstName, lastName,email, password, status) VALUES(?,?,?,?,?);"
	queryGetUser          = "SELECT * FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET firstName=?, lastName=?,email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?"
	queryFindUserByStatus = "SELECT id, firstName, lastName, email, status FROM users WHERE status=?"
)

//Save user into database
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	insertRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Status)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	userID, err := insertRes.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.ID = userID
	return nil
}

//Get user from database
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error prepare get user stament", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.DateCreated); err != nil {
		logger.Error("error trying to get user by id", err)
		return errors.NewInternalServerError("database user")
	}

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil

}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(&user.ID); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil

}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())

	}
	defer rows.Close()
	result := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Status, &user.ID); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return result, nil
}
