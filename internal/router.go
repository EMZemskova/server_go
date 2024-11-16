package internal

import (
	"github.com/EMZemskova/server_go/internal/handler"
	"github.com/gin-gonic/gin"
)

func GetRouters(handle *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/login", handle.LoginUser)
	router.GET("/user/stats/:id", handle.UserStats)
	router.GET("/user/stats", handle.PeopleStats)

	router.POST("/chats", handle.PostChat)
	router.GET("/chats/:id", handle.GetChatById)
	router.PUT("/chats/:id", handle.EditChat)
	router.DELETE("/chats/:id", handle.DeleteChat)

	router.POST("/messages", handle.PostMessage)
	router.GET("/messages/:id", handle.GetMessagebyID)
	router.DELETE("/messages/:id", handle.DeleteMessage)
	router.PUT("/messages/:id", handle.EditMessage)
	return router
}
