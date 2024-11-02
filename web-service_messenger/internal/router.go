package internal

import (
	"github.com/EMZemskova/server_go/internal/handler"
	"github.com/gin-gonic/gin"
)

func GetRouters(handle *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/login", handle.LoginUser)

	router.POST("/chats", handle.PostChat)
	router.GET("/chats/:id", handle.GetChatById)

	router.POST("/messages", handle.PostMessage)
	router.GET("/messages/:id", handle.GetMessagebyID)
	router.DELETE("/messages/:id", handle.DeleteMessage)
	router.PUT("/messages/:id", handle.EditMessage)
	return router
}
