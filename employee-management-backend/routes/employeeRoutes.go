package routes

import (
	"employee-management-backend/controllers"
	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.Engine) {
	router.GET("/employees", controllers.GetEmployees)
	router.GET("/employees/:id", controllers.GetEmployeeByID)
	router.POST("/employees", controllers.CreateEmployee)
	router.PUT("/employees/:id", controllers.UpdateEmployee)
	router.DELETE("/employees/:id", controllers.DeleteEmployee)
}
