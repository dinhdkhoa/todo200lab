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

func CreateNewItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var item model.TodoItemCreation
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStore(db)
		uc := uc.NewCreateItemUseCase(store)

		uc.CreateItemUC(c, &item)

		c.JSON(http.StatusOK, common.NewSuccessRes(item.Id))
	}
}
