package aggregates

import "github.com/sk532359025/DDD/domain/entity"


type Member struct {

	User *entity.UserEntity

	Log *entity.UserLogEntity
}

