package models

import "github.com/jinzhu/gorm"


func Migrate(db *gorm.DB)  {
	//db.Debug().DropTableIfExists(&User{},&Profesi{},&Pendidikan{})
	db.Debug().AutoMigrate(&User{},&Profesi{},&Pendidikan{})
	//db.Model(&User{}).AddForeignKey("profesi_id", "tb_profesi(id)", "CASCADE", "CASCADE")
	//db.Model(&User{}).AddForeignKey("pendidikan_id", "tb_pendidikan(id)", "CASCADE", "CASCADE")
}
