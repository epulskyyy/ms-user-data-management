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

type ProfesiController struct {
	profesiService services.ProfesiServiceInterface
}

func CreateProfesiController(router *gin.RouterGroup,profesiService services.ProfesiServiceInterface) {
	inDB := ProfesiController{profesiService}

	router.POST("/add-profesi", inDB.addProfesi)
	router.GET("/get-profesis", inDB.getProfesi)
	router.GET("/get-profesi/:id", inDB.getProfesibyId)
	router.PUT("/update-profesi/:id", inDB.updateProfesi)
	router.DELETE("/delete-profesi/:id", inDB.deleteProfesi)
}

func (u *ProfesiController) addProfesi (c *gin.Context) {
	var profesi models.Profesi

	err := c.ShouldBindJSON(&profesi)
	if err != nil {
		utils.BadRequest(c,"Error when decoder data from body with error")
		fmt.Printf("[ProfesiController.addProfesi] Error when decoder data from body with error : %v\n", err)
		return
	}

	errValidation := validator.New().Struct(&profesi)
	if errValidation != nil {
		utils.BadRequest(c,"Data is required")
		return
	}
	result, err := u.profesiService.AddPendidikan(&profesi)
	if err != nil {
		utils.NotFound(c,err.Error())
		fmt.Printf("[ProfesiController.addProfesi] Error when request data to usecase with error: %v\n", err)
		return
	}
	models.Result(c,201,"created", result)
}

func (u *ProfesiController) getProfesi(c *gin.Context) {

	result, err := u.profesiService.GetAllProfesi()
	if err != nil {
		utils.BadRequest(c,"Error when decoder data from body with error")
		fmt.Printf("[ProfesiController.getProfesi] Error when request data to usecase with error: %v\n", err)
		return
	}

	models.Result(c,http.StatusOK,"", result)
}

func (u *ProfesiController) getProfesibyId(c *gin.Context) {
	idParam := c.Param("id")
	id,_ := strconv.Atoi(idParam)

	result, err := u.profesiService.GetProfesiByID(id)
	if err != nil {
		utils.NotFound(c,idParam+" is not found!")
		fmt.Printf("[ProfesiController.getProfesi] Error when request data to usecase with error: %v\n", err)
		return
	}

	models.Result(c,http.StatusOK,"", result)
}

func (u *ProfesiController) updateProfesi(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))

	var profesi models.Profesi

	err := c.ShouldBindJSON(&profesi)
	if err != nil {
		utils.BadRequest(c,"Error when decoder data from body with error")
		fmt.Printf("[ProfesiController.updateProfesi] Error when request data to usecase with error: %v\n", err)
		return
	}

	errValidation := validator.New().Struct(&profesi)
	if errValidation != nil {
		utils.BadRequest(c,"Data is required")
		return
	}
	result, err := u.profesiService.UpdateProfesi(id, &profesi)
	if err != nil {
		utils.NotFound(c,err.Error())
		fmt.Printf("[ProfesiController.updateProfesi] Error when request data to usecase with error: %v\n", err)
		return
	}

	models.Result(c,http.StatusOK,"Updated", result)
}

func (u *ProfesiController) deleteProfesi(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	err := u.profesiService.DeleteProfesi(id)
	if err != nil {
		utils.NotFound(c,err.Error())
		fmt.Printf("[ProfesiController.deleteProfesi] Error when request data to usecase with error: %v\n", err)
		return
	}
	models.Result(c,http.StatusOK,"Delete success", nil)
}
