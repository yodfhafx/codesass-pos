package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yodfhafx/codesass-pos/controller"
)

func serveRoutes(r *gin.Engine) {
	// routing
	categoryController := controller.Category{}
	categoryGroup := r.Group("/categories")
	categoryGroup.GET("", categoryController.FindAll)
	categoryGroup.GET("/:id", categoryController.FindOne)
	categoryGroup.POST("", categoryController.Create)
	categoryGroup.PATCH("/:id", categoryController.Update)
	categoryGroup.DELETE("/:id", categoryController.Delete)

	productController := controller.Product{}
	productGroup := r.Group("/products")
	productGroup.GET("", productController.FindAll)
	productGroup.GET("/:id", productController.FindOne)
	productGroup.POST("", productController.Create)
	productGroup.PATCH("/:id", productController.Update)
	productGroup.DELETE("/:id", productController.Delete)

	orderController := controller.Order{}
	orderGroup := r.Group("/orders")
	orderGroup.GET("", orderController.FindAll)
	orderGroup.GET("/:id", orderController.FindOne)
	orderGroup.POST("", orderController.Create)
}
