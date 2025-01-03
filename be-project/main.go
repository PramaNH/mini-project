package main

import (
	"be-project/controllers"
	"be-project/db"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func setupRoutes(r *gin.Engine, queries *db.Queries) {
	r.POST("/login", controllers.Login)

	r.GET("/data", func(c *gin.Context) {
		controllers.GetData(c, queries)
	})
	r.POST("/data", func(c *gin.Context) {
		controllers.PostData(c, queries)
	})
}

func seedData(queries *db.Queries) {
	// Seed users
	_, err := queries.CreateUser(context.Background(), db.CreateUserParams{
		Username: "testuser",
		Password: "testpassword",
	})
	if err != nil {
		// Abaikan error jika user sudah ada
		fmt.Println("User seed already exists or failed:", err)
	}

	// Seed data entries
	_, err = queries.CreateData(context.Background(), db.CreateDataParams{
		Name:  "Test Name",
		Email: "test@example.com",
	})
	if err != nil {
		// Abaikan error jika data sudah ada
		fmt.Println("Data seed already exists or failed:", err)
	}

	fmt.Println("Database seeded successfully")
}


func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DB_URL")
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer dbConn.Close() // Ensure the connection is closed when the app exits

	// Tes database connection
	if err = dbConn.Ping(); err != nil {
		log.Fatal("Database connection is not alive:", err)
	}
	
	
	// SQLC queries
	queries := db.New(dbConn)
	log.Println("Database connected successfully and queries initialized")
	controllers.InitDB(queries)
	// Buat Gin server
	r := gin.Default()
	
	seedData(queries)
	
	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	setupRoutes(r, queries)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
