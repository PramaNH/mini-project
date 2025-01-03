package controllers

import (
	"context"
	"net/http"

	"be-project/db"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context, queries *db.Queries) {
	data, err := queries.ListData(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostData(c *gin.Context, queries *db.Queries) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := queries.CreateData(context.Background(), db.CreateDataParams{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}
