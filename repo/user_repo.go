package repo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"ms-user-data-management/models"
	"strconv"
	"strings"
)

type UserRepoStruct struct {
	db *gorm.DB
}

type UserRepoInterface interface {
	AddUser(user *models.User) (*models.User, error)
	GetAllUser (limit,offset,search string) (*models.UserPaging, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

func CreateUserRepoImpl(db *gorm.DB) UserRepoInterface {
	return &UserRepoStruct{db}
}

func (u *UserRepoStruct) AddUser(user *models.User) (*models.User, error) {
	tx := u.db.Begin()

	err := tx.Debug().Create(user).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("ERROR DISINI YA")
		return nil, err
	}

	tx.Commit()
	return user, nil
}
func (u *UserRepoStruct) GetAllUser(limit,offset,search string) (*models.UserPaging, error) {
	var userPaging models.UserPaging
	var user []models.User
	var user2 []models.User
	page, _ := strconv.Atoi(offset)
	perPage, _ := strconv.Atoi(limit)
	newOfset := (page-1)*perPage
	Text := strings.Replace(" "+search+ " ", " ", "%", -1)
	tx := u.db.Begin()
	err := tx.Debug().Preload("Profesi").Preload("Pendidikan").Where(" nama LIKE ? AND is_delete = 0",Text).Limit(limit).Offset(newOfset).Find(&user).Error
	if err != nil {
		fmt.Println("Error di GetAllUser")
		return nil,err
	}

	var count int
	errC := tx.Where("is_delete = 0").Find(&user2).Count(&count).Error
	if errC != nil {
		fmt.Println("Error di count")
		return nil, errC
	}
	input:=float64(count) / float64(perPage)
	totalpage:=TotalPage(input)
	userPaging.Search = search
	userPaging.Page=page
	userPaging.PerPage=perPage
	userPaging.TotalData=count
	userPaging.TotalPage=totalpage
	userPaging.Data = user
	return &userPaging, nil
}

func (u *UserRepoStruct) GetUserByID(id int) (*models.User, error) {
	tx := u.db.Begin()
	var data models.User

	err := tx.Debug().Where("id = ? and is_delete = 0", id).Preload("Profesi").Preload("Pendidikan").Find(&data).Error
	if err != nil {
		fmt.Println("Error di get user by id")
		return nil,err
	}

	return &data, nil
}

func (u *UserRepoStruct) UpdateUser(id int, user *models.User) (*models.User, error) {
	tx := u.db.Begin()

	err := u.db.Debug().Model(&user).Where("id = ? AND is_delete=0", id).Update(user).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[StudentRepo.Update] Error when query update data with error: %w", err)
	}

	tx.Commit()
	return user, nil
}

func (u *UserRepoStruct) DeleteUser(id int) error {
	data := models.User{}
	err := u.db.Debug().Model(&data).Where("id=? and is_delete = 0", id).Update("is_delete", 1).Error
	if err != nil {
		return fmt.Errorf("[TabunganRepo.Delete] Error when query delete data with error: %w", err)
	}

	return nil
}


func TotalPage(input_num float64) int {
	s := strconv.FormatFloat(input_num, 'f', -1, 64)
	fmt.Println(s)
	z:= strings.Split(s, string('.'))
	nilai1,_ := strconv.Atoi(z[0])
	var nilai int
	if nilai1 ==0 {
		fmt.Println(nilai1)
		nilai=nilai1
	}
	if len(z)>=2 {
		nilai = nilai1+ 1
	}else{
		nilai=nilai1
	}
	return nilai
}