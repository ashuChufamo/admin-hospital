package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/webProj/Login"
	"github.com/webProj/entity"
)

// LoginGormRepo implements the Login.LoginRepository interface
type LoginGormRepo struct {
	conn *gorm.DB
}

// NewLoginGormRepo will create a new object of LoginGormRepo
func NewLoginGormRepo(db *gorm.DB) Login.LoginRepository {
	return &LoginGormRepo{conn: db}
}

//LoginHand will return the Profile object or an error if the username and password doesn't exist
func (lRepo *LoginGormRepo) UsersLogin(uname string, pass string) ([]entity.Profile, []error) {
	prf := []entity.Profile{}
	fmt.Println("username: ", uname)
	fmt.Println("password: ", pass)
	errs := lRepo.conn.First(&prf, 9).GetErrors()
	if errs != nil {
		return nil, errs
	}
	return prf, errs
}
