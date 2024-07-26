package handler

import (
	"github.com/gin-gonic/gin"
	"todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.POST("/sing-in", h.singIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.creatList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deletList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.creatItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItemById)
				items.DELETE("/:item_id", h.deletItem)
			}

		}
	}
	return router
}
