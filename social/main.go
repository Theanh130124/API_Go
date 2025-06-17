package main

import (
	"social/common"
	"social/modules/item/entity"
	"social/modules/item/transport/ginItem"

	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetItem(db *gorm.DB) func(*gin.Context) {

	return func(context *gin.Context) {
		var data = entity.TodoItem{} //Nhan data
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

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}

}
func UpdateItem(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		var data = entity.TodoItemUpdate{} //Nhan data
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

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

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
		if err := db.Table(entity.TodoItem{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil { //db.First (tim kiem 1 hang theo id)
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

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
		if err := db.Table(entity.TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": "Deleted"}).Error; err != nil { //db.First (tim kiem 1 hang theo id)
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}

func ListItem(db *gorm.DB) func(context *gin.Context) {

	return func(context *gin.Context) {
		var paging common.Paging

		if err := context.ShouldBind(&paging); err != nil { // có lỗi
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()
		//Nếu không truyền gì mặc định là page 1 và limit = 10
		var res []entity.TodoItem // là các slide Todo( hay là List Item)

		//Do cái deleteShort vẫn còn nên mình cần lọc ra

		// <> là khác
		db = db.Where("status <> ?", "Deleted")

		//Find để tìm nhiều dòng dữ liệu và Order theo id desc (giảm -> mới trc)

		if err := db.Table(entity.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).Find(&res).Error; err != nil { //Create truyền dữ liệu phải là &data (rút con trỏ để dữ liêu tác động thật vào db)
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, common.NewSuccessResponse(paging, res, nil))
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

		items.POST("", ginItem.CreateItem(db))

		items.GET("", ListItem(db))
		items.GET("/:id", GetItem(db))
		items.PATCH("/:id", UpdateItem(db))
		//Gin sẽ phân tích:
		//
		///deleteShort/1 → trùng với pattern /:id (vì :id là wildcard).
		//
		//Nó hiểu "deleteShort" là id, và gọi DeleteItem(db).
		//
		//Kết quả:
		//
		//Hàm DeleteItem(db) được gọi với id = "deleteShort" → lỗi strconv.Atoi("deleteShort") → lỗi 400 hoặc 404.
		//
		//Hàm DeleteShortItem sẽ không bao giờ được gọi, vì route /:id đã match trước rồi.
		items.DELETE("/deleteShort/:id", DeleteShortItem(db)) // phai dat tren  items.DELETE("/:id", DeleteItem(db))
		items.DELETE("/:id", DeleteItem(db))

	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "TEST", // ginItem.H -> ginItem.map("key" : value)
		})
	})
	r.Run(":8000")
	//ginItem.SetMode(ginItem.ReleaseMode)  -> tắt debug bằng False
}
