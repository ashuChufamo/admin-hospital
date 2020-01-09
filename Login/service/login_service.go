package service

import (
	"github.com/webProj/Login"
	"github.com/webProj/entity"
)

//LoginSevice implements Login.LoginService interace
type LoginService struct {
	loginRepo Login.LoginRepository
}

//NewLoginService returns new LoginService object
func NewLoginService(logRepo Login.LoginRepository) Login.LoginService {
	return &LoginService{loginRepo: logRepo}
}

func (ls *LoginService) UsersLogin(uname string, pass string) ([]entity.Profile, []error) {
	prf, errs := ls.loginRepo.UsersLogin(uname, pass)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}
