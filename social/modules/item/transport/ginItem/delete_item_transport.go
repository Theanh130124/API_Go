package ginItem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social/common"
	"social/modules/item/business"

	"social/modules/item/storage/mysql"
	"strconv"
)

// Thât ra là update nhưng trang thái Deleted
func DeleteShortItem(db *gorm.DB) func(*gin.Context) {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id")) //lay params id
		if err != nil {                              // neu null
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Do mình không có struct nào để quét nên phải chỉ rõ table nào nó làm việc (map có key là string và value la interface (là gì cũng đc)
		store := mysql.NewSQLStorage(db)
		biz := business.NewDeleteItemBiz(store)
		if err := biz.DeleteItemById(context.Request.Context(), id); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
