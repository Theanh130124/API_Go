package main

import (
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
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

// Nhận json body từ sv vào structure (binh thuong gorm khong can ghi no tu scan)

// "-" la khong cho truyen tren body
type TodoItemCreation struct {
	Id          int    `json:"-" gorm:"id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	//Status      string `json:"status" gorm:"column:status"`
}

//Dùng chung

func (TodoItem) TableName() string {
	return "todo_items" // ten bang
}

//Hàm để GORM insert struct xuống db

// func(reciver) function_name kieuDL
func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName() // ten bang luc nay cung la todo_items
}

func CreateItem(db *gorm.DB) func(context *gin.Context) {

	return func(context *gin.Context) {
		var data TodoItemCreation
		//ShouldBind dang any form data hay json cung duoc
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}

func GetItem(db *gorm.DB) func(*gin.Context) {

	return func(context *gin.Context) {
		var data = TodoItem{}                        //Nhan data
		id, err := strconv.Atoi(context.Param("id")) //lay params
		if err != nil {                              // neu null
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//db.Where("id = ?", id).First(&data)  hoăc
		data.Id = id
		if err := db.First(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	}

}

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err) // in ra lỗi 1 -> dừng chương trình
	}
	fmt.Println(db)

	//now := time.Now().UTC()
	//
	//item := TodoItem{
	//	Id:          1,
	//	Title:       "Hello World",
	//	Description: "This is a test",
	//	CreateAt:    &now, // lấy con trỏ
	//	UpdateAt:    nil,  // nil là null
	//}

	//jsonData, err := json.Marshal(item)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(jsonData))
	//
	////Jsonstr du lieu tra ve
	//jsonStr := "{\"id\":1,\"title\":\"Hello World\",\"description\":\"This is a test\",\"status\":\"\",\"create_at\":\"2025-05-19T12:52:30.0840114Z\",\"update_at\":null}"
	//
	//var item2 TodoItem
	//
	//if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//Cac API GOlang
	//CRUD
	//POST /api/items -> C
	//GET /api/items -> R
	//GET /api/items/:id
	//PUT || PATCH  /api/items/:id -> U
	//DELETED  /api/items/:id   -> D

	//Url

	r := gin.Default()

	api := r.Group("/api")

	{
		items := api.Group("/items")

		items.POST("", CreateItem(db))

		items.GET("")
		items.GET("/:id", GetItem(db))
		items.PATCH("/:id")
		items.DELETE("/:id")

	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "TEST", // gin.H -> gin.map("key" : value)
		})
	})
	r.Run(":8000")
	//gin.SetMode(gin.ReleaseMode)  -> tắt debug bằng False
}
