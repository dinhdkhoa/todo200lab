package httptransport

import (
	"mymodule/common"
	"mymodule/module/item/model"
	"mymodule/module/item/storage"
	"mymodule/module/item/uc"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItems(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			model.Filters
		}
		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		queryString.Process()

		store := storage.NewSQLStore(db)
		uc := uc.NewListItemsUseCase(store)

		res, err := uc.ListItemsUC(c.Request.Context(), &queryString.Filters, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessPagingRes(res, queryString.Paging, queryString.Filters))
	}
}
