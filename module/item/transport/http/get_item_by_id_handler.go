package httptransport

import (
	"mymodule/common"
	"mymodule/module/item/storage"
	"mymodule/module/item/uc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetItemById(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		uuidStr := c.Param("id")
		id, err := uuid.Parse(uuidStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStore(db)
		uc := uc.NewGetItemUseCase(store)

		item, err := uc.GetItemByIdUC(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessRes(item))
	}
}
