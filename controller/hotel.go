package controller

import (
	"net/http"

	"github.com/eduardoarndt/hotel-service/data"
	"github.com/eduardoarndt/hotel-service/domain"
	"github.com/gin-gonic/gin"
)

func CreateHotel(c *gin.Context) {
	var newHotel domain.Hotel

	if err := c.BindJSON(&newHotel); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	data.Hotels = append(data.Hotels, newHotel)
	c.IndentedJSON(http.StatusOK, newHotel)
}

func ReadAllHotels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Hotels)
}

func ReadHotelById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range data.Hotels {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	returnNotFound(c)
}

func UpdateHotel(c *gin.Context) {
	id := c.Param("id")

	var updateHotel domain.Hotel
	err := c.BindJSON(&updateHotel)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	for index := range data.Hotels {
		if data.Hotels[index].ID == id {
			data.Hotels[index] = updateHotel
			c.IndentedJSON(http.StatusOK, data.Hotels[index])
			return
		}
	}

	returnNotFound(c)
}

func DeleteHotel(c *gin.Context) {
	id := c.Param("id")

	for index := range data.Hotels {
		if data.Hotels[index].ID == id {
			data.Hotels = removeIndex(data.Hotels, index)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "ok"})
			return
		}
	}

	returnNotFound(c)
}

func removeIndex(hotels []domain.Hotel, index int) []domain.Hotel {
	return append(hotels[:index], hotels[index+1:]...)
}

func returnNotFound(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "hotel not found"})
}
