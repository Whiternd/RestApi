package main

import "github.com/gin-gonic/gin"

func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmpoyee)
	router.PUT("/employee/:id", handler.UpdateEmpoyee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run()
}
