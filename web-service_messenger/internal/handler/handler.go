package handler

import (
	"net/http"
	"strconv"

	"github.com/EMZemskova/server_go/internal/chat"
	"github.com/EMZemskova/server_go/internal/message"
	"github.com/EMZemskova/server_go/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type handler struct {
	userProvider    user.UserProvider
	chatProvider    chat.ChatProvider
	messageProvider message.MessageProvider
}

func New() *handler {
	return &handler{
		userProvider:    userProvider,
		chatProvider:    chatProvider,
		messageProvider: messageProvider,
	}
}

func (h *handler) LoginUser(c *gin.Context) {
	var newUser user.User
	err := c.BindJSON(&newUser)
	if err != nil {
		logrus.Error(errors.Wrap(err, "loginUser BindJSON"))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.userProvider.Create(newUser)

	if err != nil {
		logrus.Error(errors.Wrap(err, "loginUser Add to DB"))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (h *handler) PostChat(c *gin.Context) {
	var newChat chat.Chat
	err := c.BindJSON(&newChat)
	if err != nil {
		logrus.Error(errors.Wrap(err, "postChat BindJSON"))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.chatProvider.Create(newChat)

	if err != nil {
		logrus.Error(errors.Wrap(err, "postChat Add to DB"))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (h *handler) GetChatById(c *gin.Context) {
	var chat chat.Chat
	findID := c.Param("id")

	id, err := strconv.Atoi(findID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "getChatById invalid ID format"))
		c.JSON(http.StatusBadRequest, "Invalid ID format")
		return
	}

	chat, err1 := h.chatProvider.Get(int64(id))

	if err1 != nil {
		logrus.Error(errors.Wrap(err1, "getChatById Find Chat"))
		c.JSON(http.StatusBadRequest, err1.Error())
		return
	}

	c.JSON(http.StatusOK, chat)
}

func (h *handler) PostMessage(c *gin.Context) {
	var newMessage message.Message
	err := c.BindJSON(&newMessage)
	if err != nil {
		logrus.Error(errors.Wrap(err, "postMessage BindJSON"))
		c.JSON(http.StatusBadRequest, err.Error())
	}
	id, err1 := h.messageProvider.Create(newMessage)

	if err1 != nil {
		logrus.Error("postMessage Error executing query:", err1)
		c.JSON(http.StatusBadRequest, err1.Error())
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (h *handler) GetMessagebyID(c *gin.Context) {
	var message message.Message
	findID := c.Param("id")

	id, err := strconv.Atoi(findID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "GetMessagebyID invalid ID format"))
		c.JSON(http.StatusBadRequest, "Invalid ID format")
		return
	}

	message, err1 := h.messageProvider.Get(int64(id))

	if err1 != nil {
		logrus.Error(errors.Wrap(err1, "GetMessagebyID Find Chat"))
		c.JSON(http.StatusBadRequest, err1.Error())
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *handler) EditMessage(c *gin.Context) {
	var newMessage message.Message
	ID := c.Param("id")
	err := c.BindJSON(&newMessage)
	if err != nil {
		logrus.Error(errors.Wrap(err, "EditMessage BindJSON"))
		c.JSON(http.StatusBadRequest, err.Error())
	}

	id, err1 := strconv.Atoi(ID)
	if err1 != nil {
		logrus.Error(errors.Wrap(err1, "EditMessage invalid ID format"))
		c.JSON(http.StatusBadRequest, "Invalid ID format")
		return
	}

	newMessage, err = h.messageProvider.Edit(newMessage, int64(id))

	if err1 != nil {
		logrus.Error("EditMessage Error executing query:", err1)
		c.JSON(http.StatusBadRequest, err1.Error())
		return
	}
	c.JSON(http.StatusCreated, newMessage)
}

func (h *handler) DeleteMessage(c *gin.Context) {
	ID := c.Param("id")

	id, err1 := strconv.Atoi(ID)
	if err1 != nil {
		logrus.Error(errors.Wrap(err1, "DeleteMessage invalid ID format"))
		c.JSON(http.StatusBadRequest, "Invalid ID format")
		return
	}
	err := h.messageProvider.Delete(int64(id))

	if err != nil {
		logrus.Error("DeleteMessage Error executing query:", err)
		c.JSON(http.StatusBadRequest, err.Error)
		return
	}

	c.JSON(http.StatusOK, nil)
}
