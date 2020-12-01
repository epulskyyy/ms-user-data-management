package repo

import (
	"ms-user-data-management/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ProfesiRepoStruct struct {
	db *gorm.DB
}

type ProfesiRepoInterface interface {
	AddProfesi(profesi *models.Profesi) (*models.Profesi, error)
	GetAllProfesi () (*[]models.Profesi, error)
	GetProfesiByID(id int) (*models.Profesi, error)
	UpdateProfesi(id int, profesi *models.Profesi) (*models.Profesi, error)
	DeleteProfesi(id int) error
}

func CreateProfesiRepoImpl(db *gorm.DB) ProfesiRepoInterface {
	return &ProfesiRepoStruct{db}
}

func (u *ProfesiRepoStruct) AddProfesi(profesi *models.Profesi) (*models.Profesi, error) {
	tx := u.db.Begin()
	err := tx.Debug().Create(profesi).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("ERROR DISINI YA")
		return nil, err
	}

	tx.Commit()
	return profesi, nil
}

func (u *ProfesiRepoStruct) GetAllProfesi() (*[]models.Profesi, error) {
	var profesi []models.Profesi
	tx := u.db.Begin()
	err := tx.Debug().Where("is_delete = 0").Find(&profesi).Error
	if err != nil {
		fmt.Println("Error di GetAllProfesi")
		return nil,err
	}

	return &profesi, nil
}

func (u *ProfesiRepoStruct) GetProfesiByID(id int) (*models.Profesi, error) {
	tx := u.db.Begin()
	data := models.Profesi{}

	err := tx.Debug().Where("id = ? and is_delete = 0", id).Find(&data).Error
	if err != nil {
		fmt.Println("Error di GetAllProfesi")
		return nil,err
	}

	return &data, nil
}

func (u *ProfesiRepoStruct) UpdateProfesi(id int, profesi *models.Profesi) (*models.Profesi, error) {
	tx := u.db.Begin()

	err := u.db.Debug().Model(&profesi).Where("id = ? and is_delete = 0", id ).Update(profesi).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[profesiRepo.Update] Error when query update data with error: %w", err)
	}

	tx.Commit()
	return profesi, nil
}

func (u *ProfesiRepoStruct) DeleteProfesi(id int) error {
	data := models.Profesi{}

	err := u.db.Debug().Model(&data).Where("id=?", id).Update("is_delete", 1).Error
	if err != nil {
		return fmt.Errorf("[profesiRepo.Delete] Error when query delete data with error: %w", err)
	}

	return nil
}
