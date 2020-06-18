package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samderlust/bookstore_users-api/domain/users"
	"github.com/samderlust/bookstore_users-api/services"
	"github.com/samderlust/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	userID, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	user, err1 := services.UsersService.GetUser(userID)
	if err1 != nil {
		c.JSON(err1.Status, err1)
		return
	}

	c.JSON(http.StatusOK, user.Marhsall(c.GetHeader("X-Public") == "true"))
}

//CreateUser controller to talk to user service and create new user
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.UsersService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result.Marhsall(c.GetHeader("X-Public") == "true"))

}

func UpdateUser(c *gin.Context) {
	userID, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, err1 := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err1.Status, err1)
		return
	}
	c.JSON(http.StatusOK, result.Marhsall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userID, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"Status": "Deleted"})

}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.FindByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marhsall(c.GetHeader("X-Public") == "true"))
}

func getUserId(userId string) (int64, *errors.RestErr) {
	userID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return userID, nil
}
