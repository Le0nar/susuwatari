package service

import (
	"fmt"

	"github.com/Le0nar/susuwatari/internal/user"
)

type Serivce struct {
	Users map[string]user.User
}

func NewService() *Serivce {
	return &Serivce{Users: make(map[string]user.User)}
}

func (s *Serivce) AddUser(name string) []user.User {
	user := user.User{Name: name}

	s.Users[name] = user

	fmt.Printf("user: %v\n", user)
}

func (s *Serivce) ChangePosition(name string, direction string) {
	user := s.Users[name]

	switch direction {
	case "top":
		user.Position.Y++
	case "right":
		user.Position.X++
	case "bottom":
		user.Position.Y--
	case "left":
		user.Position.X--
	}

	s.Users[name] = user

	fmt.Printf("user: %v\n", user)
}

func (s *Serivce) ChangePosition(name string, direction string) []user.User {
	user := s.Users[name]

	switch direction {
	case "top":
		user.Position.Y++
	case "right":
		user.Position.X++
	case "bottom":
		user.Position.Y--
	case "left":
		user.Position.X--
	}

	s.Users[name] = user

	return getUsers(s.Users)
}

// TODO: rename
func getUsers(usersMap map[string]user.User) []user.User {
	usersSlice := make([]user.User, 0, len(usersMap))

	for _, value := range usersMap {
		usersSlice = append(usersSlice, value)
	}

	return usersSlice
}
