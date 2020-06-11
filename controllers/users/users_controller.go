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
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		e := errors.NewBadRequestError("invalid user id")
		c.JSON(e.Status, e)
		return
	}

	result, err1 := services.GetUser(userID)
	if err1 != nil {
		c.JSON(err1.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

//CreateUser controller to talk to user service and create new user
func CreateUser(c *gin.Context) {
	var user users.User

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(err)
	// fmt.Println(string(bytes))
	// if err != nil {
	// 	//TODO: hanlde error
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func SeachUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement  me serac")

}
