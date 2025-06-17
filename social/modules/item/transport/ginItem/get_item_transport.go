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

func GetItem(db *gorm.DB) func(*gin.Context) {

	return func(context *gin.Context) {

		//err != null ->  co loi xay ra
		id, err := strconv.Atoi(context.Param("id")) //lay params id
		if err != nil {                              // neu null
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := mysql.NewSQLStorage(db)
		biz := business.NewGetItemBiz(store)

		data, err := biz.GetItemById(context.Request.Context(), id) // Ham nay luon tra data,err
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}

}
