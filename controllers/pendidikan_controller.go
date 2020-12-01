package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ms-user-data-management/models"
	"ms-user-data-management/services"
	"ms-user-data-management/utils"
	"net/http"
	"strconv"
)

type PendidikanController struct {
	pendidikanService services.PendidikanServiceInterface
}

func CreatePendidikanController(router *gin.RouterGroup,pendidikanService services.PendidikanServiceInterface) {
	inDB := PendidikanController{pendidikanService}

	router.POST("/add-pendidikan", inDB.addPendidikan)
	router.GET("/get-pendidikans", inDB.getPendidikan)
	router.GET("/get-pendidikan/:id", inDB.getPendidikanbyId)
	router.PUT("/update-pendidikan/:id", inDB.updatePendidikan)
	router.DELETE("/delete-pendidikan/:id", inDB.deletePendidikan)

}

func (u *PendidikanController) addPendidikan (c *gin.Context) {
	var pendidikan models.Pendidikan
	err := c.ShouldBindJSON(&pendidikan)
	if err != nil {
		utils.BadRequest(c,"Error when decoder data from body with error")
		fmt.Printf("[PendidikanController.addPendidikan] Error when decoder data from body with error : %v\n", err)
		return
	}
	errValidation := validator.New().Struct(&pendidikan)
	if errValidation != nil {
		utils.BadRequest(c,"Data is required")
		return
	}
	result, err := u.pendidikanService.AddPendidikan(&pendidikan)
	if err != nil {
		utils.BadRequest(c,err.Error())
		fmt.Printf("[PendidikanController.addPendidikan] Error when request data to usecase with error: %v\n", err)
		return
	}

	models.Result(c,201,"created", result)

}

func (u *PendidikanController) getPendidikan(c *gin.Context) {

	result, err := u.pendidikanService.GetAllPendidikan()
	if err != nil {
		utils.BadRequest(c,err.Error())
		fmt.Printf("[PendidikanController.getPendidikan] Error when request data to usecase with error: %v\n", err)
		return
	}
	models.Result(c,http.StatusOK,"", result)

}

func (u *PendidikanController) getPendidikanbyId(c *gin.Context) {
	idParam := c.Param("id")
	id,_ := strconv.Atoi(idParam)

	result, err := u.pendidikanService.GetPendidikanByID(id)
	if err != nil {
		utils.NotFound(c,idParam+" is not found!")
		fmt.Printf("[PendidikanController.getPendidikan] Error when request data to usecase with error: %v\n", err)
		return
	}
	models.Result(c,http.StatusOK,"", result)

}

func (u *PendidikanController) updatePendidikan(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))

	var pendidikan models.Pendidikan

	err := c.ShouldBindJSON(&pendidikan)
	if err != nil {
		utils.BadRequest(c,"Error when decoder data from body with error")
		fmt.Printf("[PendidikanController.addPendidikan] Error when decoder data from body with error : %v\n", err)
		return
	}
	errValidation := validator.New().Struct(&pendidikan)
	if errValidation != nil {
		utils.BadRequest(c,"Data is required")
		return
	}
	result, err := u.pendidikanService.UpdatePendidikan(id, &pendidikan)
	if err != nil {
		utils.NotFound(c,err.Error())
		fmt.Printf("[PendidikanController.updatePendidikan] Error when request data to usecase with error: %v\n", err)
		return
	}

	models.Result(c,http.StatusOK,"Updated", result)
}

func (u *PendidikanController) deletePendidikan(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	err := u.pendidikanService.DeletePendidikan(id)
	if err != nil {
		utils.NotFound(c,err.Error())
		fmt.Printf("[PendidikanController.deletePendidikan] Error when request data to usecase with error: %v\n", err)
		return
	}
	models.Result(c,http.StatusOK,"Delete Success", nil)
}
