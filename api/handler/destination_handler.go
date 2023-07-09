package handler

import (
	"net/http"

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
	destination, err := h.DestinationRepository.Create(c.Params.ByName("name"), c.Params.ByName("slug"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating destination"})
		return
	}

	c.JSON(http.StatusOK, destination)
}
