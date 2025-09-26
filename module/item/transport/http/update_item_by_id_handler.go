package httptransport

import (
	"mymodule/common"
	"mymodule/module/item/model"
	"mymodule/module/item/storage"
	"mymodule/module/item/uc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		uuidStr := c.Param("id")
		id, err := uuid.Parse(uuidStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var item *model.TodoItemUpdate
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStore(db)
		uc := uc.NewUpdateItemUseCase(store)

		if err := uc.UpdateItemByIdUC(c.Request.Context(), item, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessRes(true))
	}
}
