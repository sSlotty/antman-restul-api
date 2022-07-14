package main

import (
	"antman/restful-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/antman/:id", services.GetCircuit)
	// r.GET("/api/antman/:circuit", h.GetCustomer)
	// r.POST("/customers", h.SaveCustomer)
	// r.PUT("/customers/:id", h.UpdateCustomer)
	// r.DELETE("/customers/:id", h.DeleteCustomer)
	return r
}

func main() {
	r := setupRouter()
	err := r.Run()
	err = http.ListenAndServe(":8000", nil)

	if err != nil {
		return
	}

	if err != nil {
		return
	}
}
