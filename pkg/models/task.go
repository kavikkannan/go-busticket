package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kavikkannan/go-task/pkg/config"
)

var db *gorm.DB

type User struct {
	
	Id uint   ` json:"id"`
	Userid        uint `json:"userid"`
	Busid      uint `json:"busid"`
}
type Bookdetails struct{
	Id uint   `json:"id"`
	Userid uint `json:"userid"`
	Busid      uint `json:"busid"`
	Seatnumber        uint `json:"seatnumber"`
} 
type Login struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

type Bus struct {
	Id       uint   `json:"id"`
	Busname string `json:"busname"`
	From	string `json:"from"`
	To string `json:"to"`
	Date    string `json:"date" `
	Duration string  `json:"Duration"`
}

type Busseat struct {
	Id       uint   `json:"id"`
	Busid        uint   `json:"busid"`
	Seatnumber   uint   `json:"seatnumber"`
	Bookedstatus bool   `json:"bookedstatus"`
}
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Bookdetails{})
	db.AutoMigrate(&Bus{})
	db.AutoMigrate(&Busseat{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func (d *Bus) CreateBus() *Bus {
	db.NewRecord(d)
	db.Create(&d)
	return d
}
func (e *Bookdetails) Createbookdetails() *Bookdetails {
	db.NewRecord(e)
	db.Create(&e)
	return e
}
func GetAllUser() []User {
	var Users []User
	db.Find(&Users)
	return Users
}


func GetAllBus() []Bus {
	var Users []Bus
	db.Find(&Users)
	return Users
}

func GetAllEmploye() []Login {
	var Users []Login
	db.Find(&Users)
	return Users
}
func GetAllTicketById(ID int64) ([]Busseat, error) {
    var userDetails []Busseat
    if err := db.Where("busid = ?", ID).Find(&userDetails).Error; err != nil {
        return nil, err
    }

    return userDetails, nil
}

func GetBookedUserById(ID int64) ([]User, error) {
    var userDetails []User
    if err := db.Where("userid = ?", ID).Find(&userDetails).Error; err != nil {
        return nil, err
    }

    return userDetails, nil
}

func GetAllBookedTicketById(ID int64) ([]Bookdetails, error) {
    var userDetails []Bookdetails
    if err := db.Where("busid = ?", ID).Find(&userDetails).Error; err != nil {
        return nil, err
    }

    return userDetails, nil
}
func GetUserByEmail(email string) (*User, *gorm.DB) {
    var getUser User
    db := db.Where("email = ?", email).Find(&getUser)
    return &getUser, db
}

func GetEmployeByEmail(email string) (*Login, *gorm.DB) {
    var getUser Login
    db := db.Where("email = ?", email).Find(&getUser)
    return &getUser, db
}
func GetTicketById(ID int64) (*Busseat, *gorm.DB) {
    var getUser Busseat
    db := db.Where("busid=?", ID).Find(&getUser)
    return &getUser, db
}

func GetSeatByID(ID int64) (*Busseat, *gorm.DB) {
    var getUser Busseat
    db := db.Where("ID=?", ID).Find(&getUser)
    return &getUser, db
}
func GetBusByID(ID int64) (*Bus, *gorm.DB) {
    var getUser Bus
    db := db.Where("ID=?", ID).Find(&getUser)
    return &getUser, db
}
func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}
