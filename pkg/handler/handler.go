package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ukurysheva/sneaker-q/pkg/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	// auth := router.Group("/auth")
	// {
	// 	auth.POST("/sign-up", h.signUp)
	// 	auth.POST("/sign-in", h.signIn)
	// }
	api := router.Group("/api")
	{
		models := router.Group("/model")
		{
			models.GET("/shop/:shop", h.getShopModels)
			models.GET("/search", h.searchModels)
			// lists.GET("/:id", h.getListById)

		}
		api.GET("/connect", h.connect)
	}

	return router
}
