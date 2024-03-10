package main

import (
	"github.com/qPyth/mobydev-internship-admin/internal/app"
	"github.com/qPyth/mobydev-internship-admin/internal/config"
)

func main() {
	cfg := config.MustLoad()
	app.Run(cfg)
}
