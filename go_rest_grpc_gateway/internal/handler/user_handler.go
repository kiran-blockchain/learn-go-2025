package handler

import (
    "net/http"
    "context"
    "github.com/gin-gonic/gin"
    "go_rest_grpc_gateway/internal/grpcclient"
    "go_rest_grpc_gateway/pb"
)

func GetUserHandler(c *gin.Context) {
    id := c.Param("id")
    client := grpcclient.NewUserServiceClient()

    resp, err := client.GetUser(context.Background(), &pb.GetUserRequest{Id: id})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id":    resp.Id,
        "name":  resp.Name,
        "email": resp.Email,
    })
}