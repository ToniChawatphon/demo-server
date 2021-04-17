package main

import (
	"github.com/ToniChawatphon/demo-server/app"
)

func main() {
	app.Init()
	app.HTTPServer(app.Setting.Port)
}
