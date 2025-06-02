package controllers

import (
	"net/http"
	"rest-api/models"
	"rest-api/services"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
    matchingService services.MatchingService
}

func NewOrderController(s services.MatchingService) *OrderController {
    return &OrderController{matchingService: s}
}

func (c *OrderController) PlaceOrder(ctx *gin.Context) {
    var order models.Order
    if err := ctx.ShouldBindJSON(&order); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    order.Timestamp = time.Now().Unix()
	fmt.Println(order)

    matched := c.matchingService.PlaceOrder(order)
    ctx.JSON(http.StatusOK, gin.H{"matched_trades": matched})
}
