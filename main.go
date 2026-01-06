package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	log.Println("STARTING APP")

	// Create a PocketBase Instance
	pbApp := NewPocketBaseApp()

	// Create a Wails Instance and pass PB to it
	wApp := NewWailsApp(pbApp)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wpb",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        wApp.startup,
		OnShutdown:       wApp.shutdown,
		Bind: []interface{}{
			wApp,
		},
	})

	if err != nil {
		log.Println("APP CRASHED")
		log.Fatal(err)
	}
}
