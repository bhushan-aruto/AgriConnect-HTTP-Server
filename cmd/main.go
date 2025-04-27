package main

import (
	"github.com/bhushn-aruto/krushi-sayak-http-server/config"
	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/infra/server"
)

func main() {
	conf := config.LoadConfig()
	server.StartApp(conf)
}
