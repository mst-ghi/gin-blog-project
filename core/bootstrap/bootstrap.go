package bootstrap

import (
	"blog/core"
	"blog/core/config"
	"blog/core/engine"
	"blog/database"
)

func Serve() {
	config.InitializeAndSet()
	database.InitializeAndConnect()

	core.Initialize()

	engine.Initialize()
	engine.RegisterMiddlewares()
	engine.RegisterRoutes()
	engine.Serve()
}
