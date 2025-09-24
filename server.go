package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id          uuid.UUID     `json:"id" gorm:"column:id"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (TodoItem) TableName() string { return "\"200lab\".todo_items" }

type TodoItemCreation struct {
	Id          int    `json:"id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }

type Paging struct {
	Page    int    `json:"page" form:"page"`
	Limit   int    `json:"limit" form:"limit"`
	Total   int64  `json:"total" form:"-"`
	Keyword string `json:"-" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 1 {
		p.Page = 1
	}
	if p.Limit <= 1 {
		p.Limit = 1
	}
	if p.Limit >= 100 {
		p.Limit = 100
	}
}

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
			items.GET("/", ListItems(db))
			items.GET("/:id")
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

		var paging Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()

		var res []TodoItem
		db = db.Table(TodoItem{}.TableName()).Where("status <> (?)", "Deleted")

		if err := db.Select("id").Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Select("*").Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&res).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res, "paging": paging})
	}
}
