package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NOTE: ユーザーを表す構造体を定義
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/users", func(c *gin.Context) {
		users := []User{
			{ID: 1, Name: "John", Age: 20},
			{ID: 2, Name: "Jane", Age: 21},
			{ID: 3, Name: "Jim", Age: 22},
		}

		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080")
}
