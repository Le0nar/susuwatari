package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Le0nar/susuwatari/internal/message"
	"github.com/gorilla/websocket"
)

type service interface {
	AddUser(name string)
	ChangePosition(name string, direction string)
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

		var dto message.BasicMessageDto

		err = json.Unmarshal(p, &dto)
		if err != nil {
			log.Println(err)
			return
		}

		if dto.Type == "AddUser" {
			h.handleAddUser(p)
		}

	}
}

func (h *Handler) handleAddUser(mes []byte) {
	var dto message.AddUserDto

	err := json.Unmarshal(mes, &dto)
	if err != nil {
		log.Println(err)
		return
	}

	h.service.AddUser(dto.Name)
}
