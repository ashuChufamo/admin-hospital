package Login

import "github.com/webProj/entity"

//LoginRepository will handle Login related Repositories
type LoginRepository interface {
	UsersLogin(uname string, pass string) ([]entity.Profile, []error)
}
