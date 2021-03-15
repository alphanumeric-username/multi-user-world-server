package server

import (
	"fmt"

	"github.com/alphanumeric-username/muwsvr/cli"
	"github.com/alphanumeric-username/muwsvr/server/http_service"
)

//StartServer starts the Multi-User World Server HTTP and UDP services. It
func StartServer(config cli.ServerConfig) {
	outc := make(chan string)
	httpFlag, udpFlag := false, false
	httpSvc := http_service.NewHTTPServiceAPI(config)

	httpFlag = true
	fmt.Println("Starting HTTP service...")
	go http_service.StartHTTPService(&httpSvc, config, outc, &httpFlag)
	for httpFlag || udpFlag {

	}
}
