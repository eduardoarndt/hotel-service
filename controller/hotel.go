package controller

import (
	"fmt"
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

	insertedId, err := data.CreateHotel(newHotel)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	newHotel.ID = fmt.Sprint(*insertedId)
	c.IndentedJSON(http.StatusOK, newHotel)
}

func ReadAllHotels(c *gin.Context) {
	hotels, err := data.GetAllHotels()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
	}
	c.IndentedJSON(http.StatusOK, hotels)
}

func ReadHotelById(c *gin.Context) {
	id := c.Param("id")

	hotel, err := data.GetHotel(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			returnNotFound(c)
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, hotel)
}

func UpdateHotel(c *gin.Context) {
	id := c.Param("id")

	var updateHotel domain.Hotel
	err := c.BindJSON(&updateHotel)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if updateHotel.ID != "" {
		if id != updateHotel.ID {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ids dont match"})
			return
		}
	} else {
		updateHotel.ID = id
	}

	err = data.UpdateHotel(updateHotel)
	if err != nil {
		fmt.Println("error updating...")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updateHotel)
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
