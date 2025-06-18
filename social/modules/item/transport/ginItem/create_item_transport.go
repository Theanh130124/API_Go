package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social/common"
	"social/modules/item/business"
	"social/modules/item/entity"
	"social/modules/item/storage/mysql"
)

func CreateItem(db *gorm.DB) func(context *gin.Context) {

	return func(context *gin.Context) {
		var data entity.TodoItemCreation // truyen struct can lam viec
		//ShouldBind dạng any gửi form-data hay json cũng được
		if err := context.ShouldBind(&data); err != nil { // có lỗi
			context.JSON(http.StatusBadRequest, err)
			return
		}

		store := mysql.NewSQLStorage(db)
		biz := business.NewCreateItemBiz(store)
		if err := biz.CreateItem(context.Request.Context(), &data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data)) //hoac data.Id (thi chi hien id)
	}
}
