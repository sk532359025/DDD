package repos

import "github.com/sk532359025/DDD/domain/entity"

type IUserRepos interface {
	FindByName(name string) *entity.UserEntity
	SaveUser(use *entity.UserEntity) *entity.UserEntity
	UpdateUser(id int) error
	DeleteUser(id int) error
}

type IUserLogRepos interface {
	FindByUserId(UserId int) *entity.UserLogEntity
	SaveUser(log *entity.UserLogEntity) *entity.UserLogEntity
}
