package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db* gorm.DB
)

func Connect(){
	d, err:= gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/busticket?parseTime=true")
	if err != nil{
		panic(err)
	}
	db=d
}

func GetDB() *gorm.DB{
	return db
}
