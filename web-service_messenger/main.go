package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	pgxDB *pgx.Conn
)

type User struct {
	ID       int64  `json:id`
	Username string `json:username`
	Password string `json:password`
}

type Chat struct {
	ID      int64 `json:id`
	Creator int64 `json:creator`
	Guest   int64 `json:guest`
}

type Message struct {
	ID     int64  `json:id`
	Chat   int64  `json:chat`
	Sender int64  `json:sender`
	Text   string `json:text`
}

func setupDatabase() {
	dsn := "user=postgres password=123456 dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("failed to parse config:", err)
	}

	pgxPool, err1 := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err1 != nil {
		log.Fatal("failed to connect to the database:", err1)
	}
	pgxConn, err2 := pgxPool.Acquire(context.Background())
	if err2 != nil {
		log.Fatal("failed to connect to the database:", err2)
	}
	pgxDB = pgxConn.Conn()
	log.Println("Database connection established successfully")
}

func main() {
	setupDatabase()

	router := gin.Default()

	router.POST("/login", loginUser)

	router.POST("/chats", postChat)
	router.GET("/chats/:id", getChatById)

	router.POST("/messages", postMessage)
	router.DELETE("/messages/:id", deleteMessage)
	router.PUT("/messages/:id", editMessage)

	router.Run("localhost:8080")
}

func loginUser(c *gin.Context) {
	var newUser User
	err := c.BindJSON(&newUser)
	if err != nil {
		logrus.Error(errors.Wrap(err, "loginUser BindJSON"))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	db.Create(&newUser)
	c.JSON(http.StatusCreated, newUser)
}

func postChat(c *gin.Context) {
	var newChat Chat
	err := c.BindJSON(&newChat)
	if err != nil {
		logrus.Error(errors.Wrap(err, "postChat BindJSON"))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	db.Create(&newChat)
	c.JSON(http.StatusCreated, newChat)
}

func getChatById(c *gin.Context) {
	var chat Chat
	id := c.Param("id")
	result := db.First(&chat, id)
	if result.Error != nil {
		logrus.Error(errors.Wrap(result.Error, "getChatById"))
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}
	c.JSON(http.StatusOK, chat)
}

func postMessage(c *gin.Context) {
	var newMessage Message
	err := c.BindJSON(&newMessage)
	if err != nil {
		logrus.Error(errors.Wrap(err, "postMessage BindJSON"))
		c.JSON(http.StatusBadRequest, err.Error())
	}
	err1 := pgxDB.QueryRow(context.Background(),
		"INSERT INTO messages (chat, sender, text) VALUES ($1, $2, $3) RETURNING id", newMessage.Chat, newMessage.Sender, newMessage.Text).Scan(&newMessage.ID)

	if err1 != nil {
		logrus.Error("postMessage Error executing query:", err1)
		c.JSON(http.StatusBadRequest, err1.Error())
		return
	}
	c.JSON(http.StatusCreated, newMessage)
}

func deleteMessage(c *gin.Context) {
	id := c.Param("id")

	_, err := pgxDB.Exec(context.Background(), "DELETE FROM messages WHERE id=$1", id)
	if err != nil {
		logrus.Error("deleteMessage Error executing query:", err)
		c.JSON(http.StatusBadRequest, err.Error)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func editMessage(c *gin.Context) {
	id := c.Param("id")
	var updatedMessage Message
	err := c.BindJSON(&updatedMessage)
	if err != nil {
		logrus.Error("editMessage Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, err.Error)
		return
	}

	err = pgxDB.QueryRow(context.Background(),
		"UPDATE messages SET chat=$1, sender=$2, text=$3 WHERE id=$4 RETURNING id, chat, sender, text",
		updatedMessage.Chat, updatedMessage.Sender, updatedMessage.Text, id).Scan(&updatedMessage.ID, &updatedMessage.Chat, &updatedMessage.Sender, &updatedMessage.Text)

	if err != nil {
		logrus.Error("editMessage Error executing query:", err)
		c.JSON(http.StatusInternalServerError, err.Error)
		return
	}

	c.JSON(http.StatusOK, updatedMessage)
}
