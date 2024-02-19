package service

import (
	"fmt"

	"github.com/Le0nar/susuwatari/internal/user"
)

type Serivce struct {
	Users []user.User
}

func NewService() *Serivce {
	return &Serivce{Users: make([]user.User, 0)}
}

func (s *Serivce) AddUser(name string) {
	user := user.User{Name: name}

	s.Users = append(s.Users, user)

	fmt.Printf("s.Users: %v\n", s.Users)
}

func (s *Serivce) RemoveUser() {

}

func (s *Serivce) ChangePosition(name string, direction string) {

}
