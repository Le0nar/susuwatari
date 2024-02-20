package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Le0nar/susuwatari/internal/message"
	"github.com/Le0nar/susuwatari/internal/user"
	"github.com/gorilla/websocket"
)

type service interface {
	AddUser(name string) []user.User
	ChangePosition(name string, direction string) []user.User
}

type Handler struct {
	service service
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

// TODO: rename
func (h *Handler) HandleMain(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

		// TODO: mb use messageType for write and get message
		var dto message.BasicMessageDto

		err = json.Unmarshal(p, &dto)
		if err != nil {
			log.Println(err)
			return
		}

		// TODO: отправлять сообщения всем клиентам
		switch dto.Type {
		case "AddUser":
			users := h.handleAddUser(p)
			conn.WriteJSON(users)
		case "ChangePosition":
			users := h.handleChangePosition(p)
			conn.WriteJSON(users)
		}
	}
}

func (h *Handler) handleAddUser(mes []byte) []user.User {
	var dto message.AddUserDto

	err := json.Unmarshal(mes, &dto)
	if err != nil {
		log.Println(err)
		return nil
	}

	return h.service.AddUser(dto.Name)
}

func (h *Handler) handleChangePosition(mes []byte) []user.User {
	var dto message.ChangePositionDto

	err := json.Unmarshal(mes, &dto)
	if err != nil {
		log.Println(err)
		return nil
	}

	return h.service.ChangePosition(dto.Name, dto.Direction)
}
