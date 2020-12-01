package services

import (
	"ms-user-data-management/models"
	"ms-user-data-management/repo"
	"errors"
	"fmt"
)

type PendidikanServiceStruct struct {
	pendidikanRepo repo.PendidikanRepoInterface
}

type PendidikanServiceInterface interface {
	AddPendidikan(pendidikan *models.Pendidikan) (*models.Pendidikan, error)
	GetAllPendidikan () (*[]models.Pendidikan, error)
	GetPendidikanByID(id int) (*models.Pendidikan, error)
	UpdatePendidikan(id int, pendidikan *models.Pendidikan) (*models.Pendidikan, error)
	DeletePendidikan(id int) error
}

func CreatePendidikanServiceImpl(pendidikanRepo repo.PendidikanRepoInterface) PendidikanServiceInterface {
	return &PendidikanServiceStruct{pendidikanRepo}
}

func (u *PendidikanServiceStruct) AddPendidikan(pendidikan *models.Pendidikan) (*models.Pendidikan, error) {
	data, err := u.pendidikanRepo.AddPendidikan(pendidikan)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *PendidikanServiceStruct) GetAllPendidikan() (*[]models.Pendidikan, error) {
	data, err := u.pendidikanRepo.GetAllPendidikan()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}

func (u *PendidikanServiceStruct) GetPendidikanByID(id int) (*models.Pendidikan, error) {
	result, err := u.pendidikanRepo.GetPendidikanByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *PendidikanServiceStruct) UpdatePendidikan(id int, pendidikan *models.Pendidikan) (*models.Pendidikan, error) {
	_, err := u.pendidikanRepo.GetPendidikanByID(id)
	if err != nil {
		return nil, errors.New("ID does not exist")
	}
	result, err := u.pendidikanRepo.UpdatePendidikan(id, pendidikan)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *PendidikanServiceStruct) DeletePendidikan(id int) error {
	_, err := u.pendidikanRepo.GetPendidikanByID(id)
	if err != nil {
		return errors.New("ID does not exist")
	}

	err = u.pendidikanRepo.DeletePendidikan(id)
	if err != nil {
		return err
	}

	return nil
}
