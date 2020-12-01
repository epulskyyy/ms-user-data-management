package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Nama 			string 		`gorm:"not null" json:"nama" validate:"required,required"`
	TanggalLahir 	string 		`gorm:"not null" json:"tanggal_lahir" validate:"required,required"`
	NoKtp 			int 		`gorm:"not null" json:"no_ktp" validate:"required,required"`
	IsDelete      	int			`gorm:"not null;size:2" json:"is_delete"`
	ProfesiID 		int			`gorm:"not null;json:profesi_id" validate:"required,required"`
	Profesi			Profesi		`gorm:"foreignkey:ProfesiID;references:id"`
	PendidikanID 	int 		`gorm:"not null;json:pendidikan_id" validate:"required,required"`
	Pendidikan 		Pendidikan	`gorm:"foreignkey:PendidikanID;references:id"`
}
func (u *User) TableName() string {
	return "tb_user"
}

type UserProfesiPendidikan struct {

}

type UserPaging struct {
	Page      int    `json:"page"`
	PerPage   int    `json:"per_page"`
	Search    string `json:"search"`
	TotalPage int    `json:"total_page"`
	TotalData int    `json:"total_data"`
	Data      []User `json:"data"`
}
