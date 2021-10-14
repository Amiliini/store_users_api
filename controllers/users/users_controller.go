package users

import (
	"fmt"
	"kauppa/domain/users"
	"kauppa/services"
	"kauppa/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}

	//All of the below is replaced by simple shouldBindJson from gin
	/*
		fmt.Println(user)
		bytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			//TODO:handle error

			return
		}
		if err := json.Unmarshal(bytes, &user); err != nil {
			//TODO: handle json error
			println(err.Error())
			return
		}
	*/
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)

		return
	}
	c.JSON(http.StatusCreated, result)
	fmt.Println(user)
}

//mikä tässä mättää??
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented, "implement me!")
//}
