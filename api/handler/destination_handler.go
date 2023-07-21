package handler

import (
	"net/http"

	"github.com/amopromo/gochip/model"
	"github.com/amopromo/gochip/repository"

	"github.com/gin-gonic/gin"
)

type DestinationHandler struct {
	DestinationRepository repository.DestinationRepository
}

func (h *DestinationHandler) SetupRoutes(router *gin.Engine) {
	router.GET("/destinations", h.GetDestinations)
	router.POST("/destinations", h.CreateDestinations)
}

func (h *DestinationHandler) GetDestinations(c *gin.Context) {
	destinations, err := h.DestinationRepository.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving destinations"})
		return
	}

	c.JSON(http.StatusOK, destinations)
}

func (h *DestinationHandler) CreateDestinations(c *gin.Context) {
	var Dest model.Destination
	if errBind := c.BindJSON(&Dest); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error binding json"})
		return
	}
	destination, err := h.DestinationRepository.Create(Dest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, destination)
}
