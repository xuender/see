package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/xuender/kit/los"
	"github.com/xuender/viewing/app"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	const (
		width  = 1024
		height = 768
	)

	app := app.InitApp()
	los.Must0(wails.Run(&options.App{
		Title:       "Viewing",
		Width:       width,
		Height:      height,
		AssetServer: &assetserver.Options{Assets: assets},
		OnStartup:   app.Startup,
		Bind:        app.Bind,
		Menu:        app.Menu,
	}))
}
