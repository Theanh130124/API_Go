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
	Id          int        `json:"id"` //that ra khong can dien gorm -> vì nó đã tự map đúng name json
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreateAt    *time.Time `json:"create_at"`            //hiện time null json nếu không có
	UpdateAt    *time.Time `json:"update_at ,omitempty"` // omitempty bỏ nil  , bỏ false , bỏ qua chuoi rong
}

// Nhận json body từ sv vào structure (binh thuong gorm khong can ghi no tu scan)

type TodoItemCreation struct {
	//Du lieu trong postman se la cac truong nay
	Id          int    `json:"-" gorm:"id"` //Tao khong can truyen id nen ->  "-" la khong truyen gi
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	//Status      string `json:"status" gorm:"column:status"`
}

type TodoItemUpdate struct {
	Title       string  `json:"title"`
	Description *string `json:"description"` //vì gorm sẽ dưa vào chuỗi rỗng hoặc khác new để (bỏ qua update ) khi là con trỏ thì sẽ trỏ tới 1 giá trị new -> gorm sẽ update dù chơi chuỗi rỗng
	Status      string  `json:"status"`
}

//Dùng chung

func (TodoItem) TableName() string {
	return "todo_items" // ten bang sẽ tác động đến
}

//Hàm để GORM insert struct xuống db

// func(reciver) function_name kieuDL
func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName() // ten bang luc nay cung la todo_items
	//ToDoItem tác động lên bảng todo_items nên mình chỉ cần return TableName(vì thằng này cũng cần bảng todo_items)
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}

func CreateItem(db *gorm.DB) func(context *gin.Context) {

	return func(context *gin.Context) {
		var data TodoItemCreation // truyen struct can lam viec
		//ShouldBind dạng any gửi form-data hay json cũng được
		if err := context.ShouldBind(&data); err != nil { // có lỗi
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&data).Error; err != nil { //Create truyền dữ liệu phải là &data (rút con trỏ để dữ liêu tác động thật vào db)
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"data": data.Id, //đã tạo data với Id.
		})
	}
}

func GetItem(db *gorm.DB) func(*gin.Context) {

	return func(context *gin.Context) {
		var data = TodoItem{} //Nhan data
		//err != null ->  co loi xay ra
		id, err := strconv.Atoi(context.Param("id")) //lay params id
		if err != nil {                              // neu null
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//db.Where("id = ?", id).First(&data)  hoăc
		data.Id = id // tim kiem o dk id (Thay vì mình viet như nay minh se viet Where trong kia cai nao cũng đc )

		if err := db.First(&data).Error; err != nil { //db.First (tim kiem 1 hang theo id)
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	}

}
func UpdateItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data = TodoItemUpdate{} //Nhan data
		//err != null ->  co loi xay ra
		id, err := strconv.Atoi(context.Param("id")) //lay params id
		if err != nil {                              // neu null
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//ShouldBind dạng any gửi form-data hay json cũng được
		if err := context.ShouldBind(&data); err != nil { // có lỗi
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//db.Where("id = ?", id).First(&data)  hoăc
		//data.Id = id                                 //Thay vì viết data.Id = id
		//Updates -> truyền gì thì update cái đó (không truyền k update)
		if err := db.Where("id =?", id).Updates(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": true,
		})

	}
}

// Xóa luôn khỏi todo
func DeleteItem(db *gorm.DB) func(*gin.Context) {

	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id")) //lay params id
		if err != nil {                              // neu null
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Do mình không có struct nào để quét nên phải chỉ rõ table nào nó làm việc
		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil { //db.First (tim kiem 1 hang theo id)
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": true,
		})

	}

}

// Thât ra là update nhưng trang thái Deleted
func DeleteShortItem(db *gorm.DB) func(*gin.Context) {

	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id")) //lay params id
		if err != nil {                              // neu null
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Do mình không có struct nào để quét nên phải chỉ rõ table nào nó làm việc (map có key là string và value la interface (là gì cũng đc)
		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": "Deleted"}).Error; err != nil { //db.First (tim kiem 1 hang theo id)
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": true,
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
		items.PATCH("/:id", UpdateItem(db))
		items.DELETE("/:id", DeleteItem(db))

	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "TEST", // gin.H -> gin.map("key" : value)
		})
	})
	r.Run(":8000")
	//gin.SetMode(gin.ReleaseMode)  -> tắt debug bằng False
}
