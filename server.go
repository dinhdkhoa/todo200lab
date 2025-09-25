package main

import (
	"log"
	"mymodule/common"
	"net/http"
	"os"

	itemmodel "mymodule/module/item/model"
	httptransport "mymodule/module/item/transport/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DB_CONN")
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	if dsn == "" {
		log.Fatal("DB_CONN environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get underlying sql.DB")
	}
	defer sqlDB.Close()

	log.Println("Successfully connected to the database!", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("", ListItems(db))
			items.POST("", httptransport.CreateNewItem(db))
			// items.GET("/:id")
		}

	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(port)
}

func ListItems(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()

		var res []itemmodel.TodoItem
		db = db.Table(itemmodel.TodoItem{}.TableName()).Where("status <> (?)", "Deleted")

		if err := db.Select("id").Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Select("*").Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&res).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessPagingRes(res, paging, nil))
	}
}
