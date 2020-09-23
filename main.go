package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tshubham7/adventure-island/handler"
)

func main() {

	r := gin.Default()

	uh := handler.NewUserHandler()
	r.POST("api/user", uh.AddUser())
	r.POST("api/user/:id/reward", uh.AddPoints())
	r.GET("api/user/:id/rank", uh.UserRank())
	r.GET("api/ranks", uh.RankedUsers())
	r.Run()
}
