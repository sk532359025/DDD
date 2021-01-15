package main

import "github.com/sk532359025/DDD/server/router"

func main() {
	router := router.Router()
	router.Run("127.0.0.1:8080")

}
