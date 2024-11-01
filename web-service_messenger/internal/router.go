package internal

import (
	"github.com/EMZemskova/server_go/internal/handler"
	"github.com/gin-gonic/gin"
)

func GetRouters() *gin.Engine {
	router := gin.Default()
	handler := handler.New()

	router.POST("/login", handler.LoginUser)

	router.POST("/chats", handler.PostChat)
	router.GET("/chats/:id", handler.GetChatById)

	router.POST("/messages", handler.PostMessage)
	router.GET("/messages/:id", handler.GetMessagebyID)
	router.DELETE("/messages/:id", handler.DeleteMessage)
	router.PUT("/messages/:id", handler.EditMessage)
	return router
}
