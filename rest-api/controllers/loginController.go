package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "rest-api/models"
    "rest-api/utils"
)

func SignupHandler(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Username string `json:"username"`
            Email    string `json:"email"`
            Password string `json:"password"`
        }

        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }

        hash, _ := utils.HashPassword(req.Password)
        user := models.User{
            Username:     req.Username,
            Email:        req.Email,
            PasswordHash: hash,
        }

        if err := db.Create(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "User created"})
    }
}

func LoginHandler(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }

        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }

        var user models.User
        if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        token, _ := utils.GenerateJWT(uint(user.ID))
        c.JSON(http.StatusOK, gin.H{"token": token})
    }
}
