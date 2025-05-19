package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreateAt    *time.Time `json:"create_at"`            //hiện time null json nếu không có
	UpdateAt    *time.Time `json:"update_at ,omitempty"` // omitempty bỏ nil  , bỏ false , bỏ qua chuoi rong
}

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err) // in ra lỗi 1 -> dừng chương trình
	}
	fmt.Println(db)

	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "Hello World",
		Description: "This is a test",
		CreateAt:    &now, // lấy con trỏ
		UpdateAt:    nil,  // nil là null
	}

	jsonData, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

	//Jsonstr du lieu tra ve
	jsonStr := "{\"id\":1,\"title\":\"Hello World\",\"description\":\"This is a test\",\"status\":\"\",\"create_at\":\"2025-05-19T12:52:30.0840114Z\",\"update_at\":null}"

	var item2 TodoItem

	if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
		fmt.Println(err)
		return
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item, // gin.H -> gin.map("key" : value)
		})
	})
	r.Run(":8000")

	//gin.SetMode(gin.ReleaseMode)  -> tắt debug bằng False

}
