package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tshubham7/adventure-island/service.go"
)

type user struct {
	game service.GameService
}

// UserHandler ...
type UserHandler interface {
	// add/join user
	AddUser() gin.HandlerFunc

	// add points to user
	AddPoints() gin.HandlerFunc

	// get user's rank
	UserRank() gin.HandlerFunc

	// list of ranked users
	RankedUsers() gin.HandlerFunc
}

// NewUserHandler ...
func NewUserHandler() UserHandler {
	gm := service.NewGameService()
	return &user{gm}
}

type createUser struct {
	Name string `json:"name"`
}

// AddUser ...
func (u user) AddUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		var req createUser
		err := c.Bind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid or missing params",
				"error":   err.Error(),
			})
		}

		user, _ := u.game.CreateUser(req.Name)
		c.JSON(http.StatusOK, user)
	}
}

// AddPoints ...
func (u user) AddPoints() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params",
				"error":   "found blank url path param",
			})
		}
		p, err := u.game.AddPoints(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid paramas",
				"error":   "user not found with provided id",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message":     "successfully added",
			"totalPoints": p,
		})
	}
}

// UserRank ...
func (u user) UserRank() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params",
				"error":   "found blank url path param",
			})
		}
		r, err := u.game.UserRank(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid paramas",
				"error":   "user not found with provided id",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"rank": r,
		})
	}
}

// RankedUsers ...
func (u user) RankedUsers() gin.HandlerFunc {

	return func(c *gin.Context) {
		l := c.Query("limit")
		lt, err := strconv.Atoi(l)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params",
				"error":   err.Error(),
			})
		}
		u := u.game.UserByRank(int32(lt))
		c.JSON(http.StatusOK, gin.H{
			"limit": lt,
			"users": u,
		})
	}
}
