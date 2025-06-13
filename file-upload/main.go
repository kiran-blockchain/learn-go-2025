package main

import (
	
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve uploaded files
	r.Static("/uploads", "./uploads")

	// Upload endpoint
	r.POST("/upload", func(c *gin.Context) {
		// Parse the uploaded file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed: " + err.Error()})
			return
		}

		// Save the file to uploads/ directory
		dst := filepath.Join("uploads", file.Filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file: " + err.Error()})
			return
		}

		// Respond with file URL
		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
			"url":     "/uploads/" + file.Filename,
		})
	})

	r.Run(":8080")
}
