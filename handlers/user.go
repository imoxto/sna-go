package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name     string
	Username string
	Hash     string
	Salt     string
}

var Users = []User{
	{"kanna", "a", "as7", "alk3"},
	{"fafnir", "b", "gh4", "78j"},
	{"lucoa", "c", "rt4", "78f"},
	{"kobayashi", "k", "ac5", "s6u"},
	
	{"elulu", "e", "b77", "2s3"},
}

var ErrUserNotFound = errors.New("no users found")

func UserFind(c echo.Context) error {
	if len(Users) == 0 {
		return ErrUserNotFound
	}
	c.JSON(http.StatusOK, Users)
	return nil
}

func UserFindOne(c echo.Context) error {
	if len(Users) == 0 {
		return ErrUserNotFound
	}
	var foundUser *User
	targetUser := c.Param("username")
	for _, user := range Users {
		if user.Username == targetUser {
			foundUser = &user
			break
		}
	}
	if foundUser == nil {
		return ErrUserNotFound
	}
	c.JSON(http.StatusOK, foundUser)
	return nil
}
