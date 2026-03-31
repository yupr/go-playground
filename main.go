package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NOTE: ユーザーを表す構造体を定義
type User struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// データベースへの接続設定
	dsn := "dbuser:dbpassword@tcp(db:3306)/playground_db?charset=utf8mb4&parseTime=True&loc=Local"

	// GORM を使って MySQL に接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベース接続に失敗しました:", err)
	}
	fmt.Println("データベース接続に成功しました！")

	// マイグレーション
	db.AutoMigrate(&User{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// ユーザー一覧取得 API (GET)
	r.GET("/users", func(c *gin.Context) {
		users := []User{}

		// users テーブルの全件を取得
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	// 新規ユーザー作成 API (POST)
	r.POST("/users", func(c *gin.Context) {
		var newUser User

		// リクエストされた JSON の値が構造体の型に合わない場合はエラーを返す
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// データベースに保存
		db.Create(&newUser)
		c.JSON(http.StatusCreated, newUser)
	})

	r.Run(":8080")
}
