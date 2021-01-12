package main

import (
	"fmt"
	"github.com/sk532359025/DDD/domain/entity"
	"github.com/sk532359025/DDD/domain/valueobj"
)

func main() {
	user := &entity.UserEntity{
		Name: "测试",
		Id: 1,
		Extra: valueobj.UserExtra{Name: "测试1213"},
	}
	fmt.Println(user)
}