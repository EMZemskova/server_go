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

type Handler struct {
	userProvider    user.Provider
	chatProvider    chat.Provider
	messageProvider message.Provider
}

func New(userProvider user.Provider, chatProvider chat.Provider, messageProvider message.Provider) *Handler {
	return &Handler{
		userProvider:    userProvider,
		chatProvider:    chatProvider,
		messageProvider: messageProvider,
	}
}

func (h *Handler) LoginUser(c *gin.Context) {
	var newUser user.User
	if err := c.BindJSON(&newUser); err != nil {
		logrus.Error("loginUser BindJSON ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "loginUser BindJSON"))
		return
	}
	id, err := h.userProvider.Create(newUser)
	if err != nil {
		logrus.Error("loginUser Add to DB ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "loginUser Add to DB"))
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (h *Handler) UserStats(c *gin.Context) {
	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		logrus.Error("UserStats invalid ID format ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "UserStats invalid ID format"))
		return
	}
	UserStat, err := h.userProvider.GetStat(int64(id))
	if err != nil {
		logrus.Error("UserStats error ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "UserStats error"))
		return
	}
	c.JSON(http.StatusOK, UserStat)
}

func (h *Handler) PeopleStats(c *gin.Context) {
	PeopleStats, err := h.userProvider.GetStats()
	if err != nil {
		logrus.Error("PeopleStats error ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "PeopleStats error"))
		return
	}
	c.JSON(http.StatusOK, PeopleStats)
}

func (h *Handler) PostChat(c *gin.Context) {
	var newChat chat.Chat
	if err := c.BindJSON(&newChat); err != nil {
		logrus.Error("postChat BindJSON ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "postChat BindJSON"))
		return
	}
	id, err := h.chatProvider.Create(newChat)
	if err != nil {
		logrus.Error("postChat Add to DB ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "postChat Add to DB"))
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (h *Handler) GetChatById(c *gin.Context) {
	findID := c.Param("id")
	id, err := strconv.Atoi(findID)
	if err != nil {
		logrus.Error("getChatById invalid ID format ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "getChatById invalid ID format"))
		return
	}
	chat, err := h.chatProvider.Get(int64(id))
	if err != nil {
		logrus.Error("getChatById Find Chat ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "getChatById Find Chat"))
		return
	}
	c.JSON(http.StatusOK, chat)
}

func (h *Handler) EditChat(c *gin.Context) {
	var editChat chat.Chat
	ID := c.Param("id")
	if err := c.BindJSON(&editChat); err != nil {
		logrus.Error("EditChat BindJSON ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "EditChat BindJSON"))
	}
	id, err := strconv.Atoi(ID)
	if err != nil {
		logrus.Error("EditChat invalid ID format ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "EditChat invalid ID format"))
		return
	}
	editChat, err = h.chatProvider.Edit(editChat, int64(id))
	if err != nil {
		logrus.Error("EditChat Error executing query ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "EditChat Error executing query"))
		return
	}
	c.JSON(http.StatusCreated, editChat)
}

func (h *Handler) DeleteChat(c *gin.Context) {
	var deletedChat chat.Chat
	ID := c.Param("id")
	if err := c.BindJSON(&deletedChat); err != nil {
		logrus.Error("DeleteChat BindJSON ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "DeleteChat BindJSON"))
	}
	id, err := strconv.Atoi(ID)
	if err != nil {
		logrus.Error("EditChat invalid ID format ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "EditChat invalid ID format"))
		return
	}
	id, err = h.chatProvider.Delete(deletedChat, int64(id))
	if err != nil {
		logrus.Error("DeleteChat Error executing query:", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "DeleteChat Error executing query"))
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) PostMessage(c *gin.Context) {
	var newMessage message.Message
	if err := c.BindJSON(&newMessage); err != nil {
		logrus.Error("postMessage BindJSON ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "postMessage BindJSON"))
	}
	id, err := h.messageProvider.Create(newMessage)
	if err != nil {
		logrus.Error("postMessage Error executing query:", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "Error executing query"))
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (h *Handler) GetMessagebyID(c *gin.Context) {
	findID := c.Param("id")
	id, err := strconv.Atoi(findID)
	if err != nil {
		logrus.Error("GetMessagebyID invalid ID format ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "GetMessagebyID invalid ID format"))
		return
	}
	message, err := h.messageProvider.Get(int64(id))
	if err != nil {
		logrus.Error("GetMessagebyID Find Chat ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "GetMessagebyID Find Chat"))
		return
	}
	c.JSON(http.StatusOK, message)
}

func (h *Handler) EditMessage(c *gin.Context) {
	var newMessage message.Message
	ID := c.Param("id")
	if err := c.BindJSON(&newMessage); err != nil {
		logrus.Error("EditMessage BindJSON ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "EditMessage BindJSON"))
	}
	id, err := strconv.Atoi(ID)
	if err != nil {
		logrus.Error("EditMessage invalid ID format ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "EditMessage invalid ID format"))
		return
	}
	newMessage, err = h.messageProvider.Edit(newMessage, int64(id))
	if err != nil {
		logrus.Error("EditMessage Error executing query ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "EditMessage Error executing query"))
		return
	}
	c.JSON(http.StatusCreated, newMessage)
}

func (h *Handler) DeleteMessage(c *gin.Context) {
	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		logrus.Error("DeleteMessage invalid ID format ", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "DeleteMessage invalid ID format"))
		return
	}
	err = h.messageProvider.Delete(int64(id))
	if err != nil {
		logrus.Error("DeleteMessage Error executing query:", err)
		c.JSON(http.StatusBadRequest, errors.Wrap(err, "DeleteMessage Error executing query"))
		return
	}
	c.JSON(http.StatusOK, nil)
}
