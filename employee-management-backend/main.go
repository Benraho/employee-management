package main

import (
    "github.com/gin-gonic/gin"
    "employee-management-backend/config"
    "employee-management-backend/routes"
    "log"
    "net/http"
)

func main() {
    router := gin.Default()

    // Middleware CORS
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusOK)
            return
        }
        c.Next()
    })

    config.ConnectDB()
    log.Println("Connected to MongoDB Atlas!")

    // Routes
    routes.EmployeeRoutes(router)

    router.Run(":8080")
}
