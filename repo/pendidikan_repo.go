package repo

import (
	"ms-user-data-management/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

type PendidikanRepoStruct struct {
	db *gorm.DB
}

type PendidikanRepoInterface interface {
	AddPendidikan(pendidikan *models.Pendidikan) (*models.Pendidikan, error)
	GetAllPendidikan() (*[]models.Pendidikan, error)
	GetPendidikanByID(id int) (*models.Pendidikan, error)
	UpdatePendidikan(id int, pendidikan *models.Pendidikan) (*models.Pendidikan, error)
	DeletePendidikan(id int) error
}

func CreatePendidikanRepoImpl(db *gorm.DB) PendidikanRepoInterface {
	return &PendidikanRepoStruct{db}
}

func (u *PendidikanRepoStruct) AddPendidikan(pendidikan *models.Pendidikan) (*models.Pendidikan, error) {
	tx := u.db.Begin()

	err := tx.Debug().Create(pendidikan).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("ERROR DISINI YA")
		return nil, err
	}

	tx.Commit()
	return pendidikan, nil
}

func (u *PendidikanRepoStruct) GetAllPendidikan() (*[]models.Pendidikan, error) {
	var pendidikan []models.Pendidikan
	tx := u.db.Begin()

	err := tx.Debug().Where("is_delete = 0").Find(&pendidikan).Error
	if err != nil {
		fmt.Println("Error di GetAllPendidikan")
		return nil, err
	}

	return &pendidikan, nil
}

func (u *PendidikanRepoStruct) GetPendidikanByID(id int) (*models.Pendidikan, error) {
	tx := u.db.Begin()
	data := models.Pendidikan{}

	err := tx.Debug().Where("id = ? And is_delete = 0", id).Find(&data).Error
	if err != nil {
		fmt.Println("Error di GetAllPendidikan")
		return nil, err
	}

	return &data, nil
}

func (u *PendidikanRepoStruct) UpdatePendidikan(id int, pendidikan *models.Pendidikan) (*models.Pendidikan, error) {
	tx := u.db.Begin()

	err := u.db.Debug().Model(&pendidikan).Where("id = ? AND is_delete = 0", id).Update(pendidikan).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[PendidikanRepo.Update] Error when query update data with error: %w", err)
	}

	tx.Commit()
	return pendidikan, nil
}

func (u *PendidikanRepoStruct) DeletePendidikan(id int) error {
	data := models.Pendidikan{}
	err := u.db.Debug().Model(&data).Where("id=?", id).Update("is_delete", 1).Error
	if err != nil {
		return fmt.Errorf("[PendidikanRepo.Delete] Error when query delete data with error: %w", err)
	}

	return nil
}
