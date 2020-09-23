package service

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/tshubham7/adventure-island/models"
)

// GameService ...
type GameService interface {
	// add/create user
	CreateUser(name string) (*models.User, error)

	// add points to user
	AddPoints(userID string) (int32, error)

	// user list rank wise
	UserByRank(limit int32) []*models.User
	// get user rank
	UserRank(userID string) (int, error)
}

type game struct {
	mutex sync.Mutex
}

// JoinedUsers ...
// this will be sorted array
var JoinedUsers = []*models.User{}

// UserMap ...
var UserMap = map[string]*models.User{}

// getNewUserID ...
func getNewUserID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}

// NewGameService ...
// service constructor
func NewGameService() GameService {
	return &game{}
}

// CreateUser ...
func (g game) CreateUser(name string) (*models.User, error) {
	id := getNewUserID()
	usr := &models.User{ID: id, Name: name}
	JoinedUsers = append(JoinedUsers, usr)

	// adding user to map
	UserMap[usr.ID] = usr
	return usr, nil
}

// UserRank ...
// get user's rank
func (g game) UserRank(userID string) (int, error) {
	usr := UserMap[userID]
	if usr == nil {
		return 0, errors.New("user not found")
	}

	cp := usr.Points //current points
	in := binarySearch(JoinedUsers, 0, len(JoinedUsers)-1, cp, usr.ID)
	return in, nil
}

// UserByRank
// sorted users list by points
func (g game) UserByRank(n int32) []*models.User {
	if int(n) < len(JoinedUsers) {
		return JoinedUsers[:n]
	}
	return JoinedUsers
}

// AddPoints
// adding point to user
func (g game) AddPoints(userID string) (int32, error) {
	g.mutex.Lock()
	usr := UserMap[userID]
	if usr == nil {
		return 0, errors.New("user not found")
	}

	cp := usr.Points //current points
	in := binarySearch(JoinedUsers, 0, len(JoinedUsers)-1, cp, usr.ID)
	if in == -1 {
		return 0, errors.New("user not found")
	}

	// adding points
	usr.Points += 10

	// getting user to the new position
	JoinedUsers = shift(JoinedUsers, in, usr.Points)

	defer g.mutex.Unlock()
	return usr.Points, nil
}
