package main

import (
	"log"
	"os"

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

	log.Println("Successfully connected to the database!")

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("", httptransport.ListItems(db))
			items.GET("/:id", httptransport.GetItemById(db))
			items.POST("", httptransport.CreateNewItem(db))
			items.PATCH("/:id", httptransport.UpdateItem(db))
			items.DELETE("/:id", httptransport.DeleteItem(db))
		}

	}

	router.Run(port)
}
