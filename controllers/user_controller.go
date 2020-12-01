package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ms-user-data-management/models"
	"ms-user-data-management/services"
	"ms-user-data-management/utils"
	"net/http"
	"strconv"
)

type UserController struct {
	userService services.UserServiceInterface
}

func CreateUserController(router *gin.RouterGroup,userService services.UserServiceInterface) {
	inDB := UserController{userService}

	router.POST("/add-user", inDB.addUser)
	router.GET("/get-users", inDB.getUser)
	router.GET("/get-user/:id", inDB.getUserbyId)
	router.PUT("/update-user/:id", inDB.updateUser)
	router.DELETE("/delete-user/:id", inDB.deleteUser)

}

func (u *UserController) addUser (c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.BadRequest(c,"Error when decoder data from body with error ")
		fmt.Printf("[UserController.addUser] Error when decoder data from body with error : %v\n", err)
		return
	}

	err = services.DataIsRequired(&user)
	if err != nil{
			utils.BadRequest(c,"Data is required")
			return
	}
	result, err := u.userService.AddPendidikan(&user)
	if err != nil {
		utils.BadRequest(c,err.Error())
		fmt.Printf("[UserController.addUser] Error when request data to usecase with error: %v\n", err)
		return
	}
	models.Result(c,201,"created", result)
}

func (u *UserController) getUser(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	search := c.Query("search")
	result, err := u.userService.GetAllUser(limit,offset,search)
	if err != nil {
		utils.BadRequest(c,err.Error())
		fmt.Printf("[UserController.getUser] Error when request data to usecase with error: %v\n", err)
		return
	}

	models.Result(c,http.StatusOK,"", result)
}

func (u *UserController) getUserbyId(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))

	result, err := u.userService.GetUserByID(id)
	if err != nil {
		utils.NotFound(c,err.Error())
		fmt.Printf("[UserController.getUser] Error when request data to usecase with error: %v\n", err)
		return
	}
	models.Result(c,http.StatusOK,"", result)
}

func (u *UserController) updateUser(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.BadRequest(c,"Error when decoder data from body with error ")
		fmt.Printf("[UserController.updateUser] Error when request data to usecase with error: %v\n", err)
		return
	}

	result, err := u.userService.UpdateUser(id, &user)
	if err != nil {
		utils.NotFound(c,err.Error())
		fmt.Printf("[UserController.updateUser] Error when request data to usecase with error: %v\n", err)
		return
	}
	models.Result(c,http.StatusOK,"Updated", result)
}

func (u *UserController) deleteUser(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	err := u.userService.DeleteUser(id)
	if err != nil {
		utils.NotFound(c,err.Error())
		fmt.Printf("[UserController.deleteUser] Error when request data to usecase with error: %v\n", err)
		return
	}
	models.Result(c,http.StatusOK,"Deleted", nil)
}
