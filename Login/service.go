package Login

import "github.com/webProj/entity"

//LoginService will handle Login related services
type LoginService interface {
	UsersLogin(uname string, pass string) ([]entity.Profile, []error)
}
