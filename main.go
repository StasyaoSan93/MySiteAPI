package main

import (
	"fmt"
	"os"

	"my-api/core"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	core.ConnectDatabase()

	r := gin.Default()

	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Роуты
	r.POST("/token", login)
	r.GET("/sitedata", getSiteData)
	r.GET("/infopagedata", getInfoPageData)

	// Пример защищенного роута (в идеале добавить Middleware)
	// Группируем роуты, требующие авторизации
	protected := r.Group("/")
	protected.Use(core.AuthMiddleware())
	{
		// Эти роуты теперь под защитой JWT
		protected.POST("/sitedata", createSiteData)
	}

	// Получаем порт из .env или системных переменных
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // значение по умолчанию
	}

	fmt.Printf("Сервер запущен на порту %s\n", port)

	// ВАЖНО: Gin ожидает строку в формате ":8080"
	r.Run(":" + port)
}

func login(c *gin.Context) {
	var input core.LoginSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var user core.User
	if err := core.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if !core.CheckPasswordHash(input.Password, user.HashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, _ := core.CreateToken(user.Username)
	c.JSON(http.StatusOK, gin.H{"access_token": token, "token_type": "bearer"})
}

func getSiteData(c *gin.Context) {
	var data []core.SiteData
	core.DB.Find(&data)
	c.JSON(http.StatusOK, data)
}

func getInfoPageData(c *gin.Context) {
	var data []core.InfoPageData
	core.DB.Find(&data)
	c.JSON(http.StatusOK, data)
}

func createSiteData(c *gin.Context) {
	var input core.SiteDataCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newData := core.SiteData{
		HeaderText: input.HeaderText,
		BodyText:   input.BodyText,
		ImageURL:   input.ImageURL,
		SiteURL:    input.SiteURL,
	}
	core.DB.Create(&newData)
	c.JSON(http.StatusCreated, newData)
}
