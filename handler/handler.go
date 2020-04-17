package handler

import (
	"context"
	"log"
)

type Handler struct {
	
}

func (h Handler) Init() {
	panic("implement me")
}

func (h Handler) AfterInit() {
	panic("implement me")
}

func (h Handler) BeforeShutdown() {
	panic("implement me")
}

func (h Handler) Shutdown() {
	panic("implement me")
}
func (h Handler) Greeter(ctx context.Context, payload []byte) {
	log.Printf("greeter:%v", payload)
}