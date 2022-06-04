package handler

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (h *Handler) InitRouts() *echo.Echo { //echo.Route
	router := echo.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)
			lists.GET("/", h.GetAllLists)
			lists.GET("/:id", h.GetListById)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteList)

			items := api.Group("/items")
			{
				items.POST("/", h.CreateItem)
				items.GET("/", h.GetAllItems)
				items.GET("/:id", h.GetItemById)
				items.PUT("/:id", h.UpdateItem)
				items.DELETE("/:id", h.DeleteItem)
			}
		}
	}
	return router
}
