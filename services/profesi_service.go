package services

import (
	"errors"
	"fmt"
	"ms-user-data-management/models"
	"ms-user-data-management/repo"
)

type ProfesiServiceStruct struct {
	profesiRepo repo.ProfesiRepoInterface
}

type ProfesiServiceInterface interface {
	AddPendidikan(profesi *models.Profesi) (*models.Profesi, error)
	GetAllProfesi () (*[]models.Profesi, error)
	GetProfesiByID(id int) (*models.Profesi, error)
	UpdateProfesi(id int, profesi *models.Profesi) (*models.Profesi, error)
	DeleteProfesi(id int) error
}

func CreateProfesiServiceImpl(profesiRepo repo.ProfesiRepoInterface) ProfesiServiceInterface {
	return &ProfesiServiceStruct{profesiRepo}
}

func (u *ProfesiServiceStruct) AddPendidikan(profesi *models.Profesi) (*models.Profesi, error) {
	data, err := u.profesiRepo.AddProfesi(profesi)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *ProfesiServiceStruct) GetAllProfesi() (*[]models.Profesi, error) {
	data, err := u.profesiRepo.GetAllProfesi()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}

func (u *ProfesiServiceStruct) GetProfesiByID(id int) (*models.Profesi, error) {
	result, err := u.profesiRepo.GetProfesiByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *ProfesiServiceStruct) UpdateProfesi(id int, profesi *models.Profesi) (*models.Profesi, error) {
	_, err := u.profesiRepo.GetProfesiByID(id)
	if err != nil {
		return nil, errors.New("ID does not exist")
	}
	result, err := u.profesiRepo.UpdateProfesi(id, profesi)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *ProfesiServiceStruct) DeleteProfesi(id int) error {
	_, err := u.profesiRepo.GetProfesiByID(id)
	if err != nil {
		return errors.New("ID does not exist")
	}

	err = u.profesiRepo.DeleteProfesi(id)
	if err != nil {
		return err
	}

	return nil
}
