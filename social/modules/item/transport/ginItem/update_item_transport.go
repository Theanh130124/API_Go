package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social/common"
	"social/modules/item/business"
	"social/modules/item/entity"
	"social/modules/item/storage/mysql"

	"strconv"
)

// Thât ra là update nhưng trang thái Deleted
func UpdateItem(db *gorm.DB) func(*gin.Context) {
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

		store := mysql.NewSQLStorage(db)
		biz := business.NewUpdateItemBiz(store)
		if err := biz.UpdateItemById(context.Request.Context(), id, &data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
