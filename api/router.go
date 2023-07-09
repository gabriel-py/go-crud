package api

import (
	"github.com/amopromo/gochip/api/handler"
	"github.com/amopromo/gochip/middleware"
	"github.com/amopromo/gochip/repository"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, destinationHandler *handler.DestinationHandler) *gin.Engine {
	destinationHandler.SetupRoutes(r)

	return r
}

func SetupMiddleware(r *gin.Engine, apiClientRepository repository.ApiClientRepository) *gin.Engine {
	r.Use(middleware.ValidateToken(apiClientRepository))

	return r
}
