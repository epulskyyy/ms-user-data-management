package services

import (
	"ms-user-data-management/models"
	"ms-user-data-management/repo"
	"errors"
	"fmt"
)

type UserServiceStruct struct {
	userRepo repo.UserRepoInterface
}

type UserServiceInterface interface {
	AddPendidikan(user *models.User) (*models.User, error)
	GetAllUser (limit,offset,search string) (*models.UserPaging, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
	//DataIsRequired(user *models.User) error
}

func CreateUserServiceImpl(userRepo repo.UserRepoInterface) UserServiceInterface {
	return &UserServiceStruct{userRepo}
}

func (u *UserServiceStruct) AddPendidikan(user *models.User) (*models.User, error) {
	data, err := u.userRepo.AddUser(user)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserServiceStruct) GetAllUser(limit,offset,search string) (*models.UserPaging, error) {
	data, err := u.userRepo.GetAllUser(limit,offset,search)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}

func (u *UserServiceStruct) GetUserByID(id int) (*models.User, error) {
	result, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserServiceStruct) UpdateUser(id int, user *models.User) (*models.User, error) {
	_, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("studentID does not exist")
	}
	result, err := u.userRepo.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserServiceStruct) DeleteUser(id int) error {

	_, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return errors.New("ID User does not exist")
	}
	err = u.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func DataIsRequired (user *models.User) error{
	if user.Nama =="" || user.TanggalLahir == "" || user.NoKtp == 0|| user.ProfesiID ==0 || user.PendidikanID == 0{
		return errors.New("data is required")
	}

	return nil
}