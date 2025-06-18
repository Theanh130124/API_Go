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

func ListItem(db *gorm.DB) func(*gin.Context) {

	return func(context *gin.Context) {
		var paging common.Paging

		if err := context.ShouldBind(&paging); err != nil { // có lỗi
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process() // method cua minh

		var filter entity.Filter                            //truyen status done hoac doing ...
		if err := context.ShouldBind(&filter); err != nil { // có lỗi
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := mysql.NewSQLStorage(db)
		biz := business.NewListItemBiz(store)

		data, err := biz.ListItem(context.Request.Context(), &filter, &paging)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, common.NewSuccessResponse(paging, data, filter))
	}

}
