package entity

import "github.com/sk532359025/DDD/domain/valueobj"

// 用户 实体
// 必须有唯一标识
type UserEntity struct {
	Id    int
	Phone int
	Name  string
	Extra valueobj.UserExtra
}



