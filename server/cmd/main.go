package main

import (
	"net/http"

	"github.com/Le0nar/susuwatari/internal/handler"
	"github.com/Le0nar/susuwatari/internal/service"
)

func main() {
	srvc := service.NewService()
	handlr := handler.NewHandler(srvc)

	http.HandleFunc("/", handlr.HandleMain)
	http.ListenAndServe(":8080", nil)
}
