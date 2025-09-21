package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Luiz"},
	{ID: 2, Name: "Fernando"},
}

func main() {
	server := setupRouter()
	server.Run(":8000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, users)
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for _, u := range users {
			if id == strconv.Itoa(u.ID) {
				ctx.JSON(200, u)
				return
			}
		}

		ctx.JSON(404, gin.H{"error": "user not found"})
	})

	r.POST("/users", func(ctx *gin.Context) {
		var newUser User

		if err := ctx.ShouldBindJSON(&newUser); err != nil {
			ctx.JSON(400, gin.H{"error": "invalid request"})
			return
		}

		users = append(users, newUser)
		ctx.JSON(201, newUser)
	})

	r.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
