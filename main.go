package main

import (
	"github.com/eduardoarndt/hotel-service/controller"
	"github.com/eduardoarndt/hotel-service/data"
	"github.com/gin-gonic/gin"
)

func main() {
	postgresConnect()
	serverStart()
}

func postgresConnect() {
	data.Connect()
}

func serverStart() {
	router := gin.Default()
	router.POST("/hotel", controller.CreateHotel)
	router.GET("/hotel", controller.ReadAllHotels)
	router.GET("/hotel/:id", controller.ReadHotelById)
	router.PATCH("/hotel/:id", controller.UpdateHotel)
	router.DELETE("/hotel/:id", controller.DeleteHotel)
	router.Run()
}
