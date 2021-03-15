package http_service

import (
	"fmt"

	"github.com/alphanumeric-username/muwsvr/chat"
	"github.com/alphanumeric-username/muwsvr/cli"
	"github.com/alphanumeric-username/muwsvr/server/user"
)

//HTTPServiceAPI stores the current state of the HTTP Service
type HTTPServiceAPI struct {
	Version     string
	ServerChat  chat.Chat
	ChatEnabled bool
	Users       []user.User
	MaxUsers    int
}

//NewHTTPServiceAPI creates a new HTTPServiceAPI object and returns it
func NewHTTPServiceAPI(config cli.ServerConfig) HTTPServiceAPI {
	var svc HTTPServiceAPI

	svc.Version = cli.Version
	svc.ChatEnabled = config.EnableChat
	if config.EnableChat {
		svc.ServerChat = chat.NewChat(config.MaxMessages)
	}
	svc.Users = make([]user.User, 0)
	svc.MaxUsers = config.MaxUsers

	return svc
}

//StartHTTPService starts a http server listening on port set in server config file. It should run in it's own goroutine
func StartHTTPService(svc *HTTPServiceAPI, config cli.ServerConfig, outc chan<- string, serviceIsRunning *bool) {

	defer func() { *serviceIsRunning = false }()

	svr := createHTTPServer(svc)
	svr.Addr = fmt.Sprintf("%s:%d", config.IP, config.HTTPPort)

	err := svr.ListenAndServe()
	if err != nil {
		outc <- fmt.Sprintf("Error on HTTP service start up: %s", err.Error())
	}
}
