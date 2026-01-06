package main

import (
	"context"
	"log"
)

// App struct
type WailsApp struct {
	wailsCtx context.Context
	pbApp    *PocketBaseApp
}

// Creates a new App application struct
func NewWailsApp(pbApp *PocketBaseApp) *WailsApp {
	log.Println("Create Wails App")
	return &WailsApp{
		pbApp: pbApp,
	}
}

// Startup is called when the app starts. The context is saved so we can call the runtime methods
// Also sets wails context to PB starts PB app instance
func (app *WailsApp) startup(ctx context.Context) {
	log.Println("Starting Wails App")
	app.wailsCtx = ctx

	app.pbApp.SetWailsContext(ctx)
	if err := app.pbApp.StartPB(); err != nil {
		log.Fatal(err)
	}
}

// Terminates PB, close connections and clears up resorces
func (app *WailsApp) shutdown(ctx context.Context) {
	if err := app.pbApp.StopPB(); err != nil {
		log.Fatal(err)
	}
	log.Println("Stopping Wails App")
}
