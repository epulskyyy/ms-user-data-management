package models

import "github.com/jinzhu/gorm"

type Pendidikan struct {
	gorm.Model
	Nama 			string 		`gorm:"not null" json:"nama"  validate:"required"`
	IsDelete      	int			`gorm:"not null;size:2" json:"is_delete"`
}

func (u *Pendidikan) TableName() string {
	return "tb_pendidikan"
}
