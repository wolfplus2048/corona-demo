package handler

import (
	"context"
	"github.com/wolfplus2048/corona-api"
	"github.com/wolfplus2048/corona-demo/proto"
	"log"

)
func GetSessionFromCtx(ctx context.Context) corona.Session {
	sessionVal := ctx.Value("session")
	if sessionVal == nil {
		log.Print("ctx doesn't contain a session, are you calling GetSessionFromCtx from inside a remote?")
		return nil
	}
	return sessionVal.(corona.Session)
}

type Handler struct {
	
}

func (h Handler) Init() {
}

func (h Handler) AfterInit() {
}

func (h Handler) BeforeShutdown() {
}

func (h Handler) Shutdown() {
}
func (h Handler) Greeter(ctx context.Context, req *proto.Greeter) {
	log.Printf("greeter :%s", req.UserName)
	s := GetSessionFromCtx(ctx)
	s.Push("onMessage", &proto.Message{Content: "Hello, " + req.UserName})
}