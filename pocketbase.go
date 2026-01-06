package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "wpb/migrations"
)

// PB App Struct with wails context
// the wails context is used for Event Listeners or other wails methods
type PocketBaseApp struct {
	pbApp    *pocketbase.PocketBase
	wailsCtx context.Context
}

// Creates a new PB App application struct
func NewPocketBaseApp() *PocketBaseApp {
	log.Println("Create PB App")
	return &PocketBaseApp{}
}

// Add Wails Context to the application struct
func (app *PocketBaseApp) SetWailsContext(ctx context.Context) {
	log.Println("Added Wails Context to PB")
	app.wailsCtx = ctx
}

// Helper to get PB app instance
func (app *PocketBaseApp) GetPB() *pocketbase.PocketBase {
	return app.pbApp
}

// Method to setup and start the PB instance
func (app *PocketBaseApp) StartPB() error {
	log.Println("Starting PB App")

	// Create a New PB instance
	app.pbApp = pocketbase.NewWithConfig(pocketbase.Config{
		HideStartBanner: true,
		DefaultDev:      false,
		DefaultDataDir:  "./wpb/data",
	})

	// Not Needed, but can be used if app users want to extend the app (eg - extentions, plugins, etc)
	jsvm.MustRegister(app.pbApp, jsvm.Config{
		HooksWatch:             true,
		HooksDir:               "./wpb/hooks",
		HooksFilesPattern:      `^.*(\.pb\.js|\.pb\.ts)$`,
		MigrationsDir:          "./wpb/migrations",
		MigrationsFilesPattern: `^.*(\.pb\.js|\.pb\.ts)$`,
		TypesDir:               "./wpb/types",
	})

	// Register new migrations
	// During dev mode, while making changes to the database, copy over the files from ./wpb/migrations to ./migrations
	// NOTE: use "js" on build and "go" on dev
	migratecmd.MustRegister(app.pbApp, app.pbApp.RootCmd, migratecmd.Config{
		Automigrate:  true,
		TemplateLang: "go",
		Dir:          "./wpb/migrations",
	})

	// Also not needed, but can be used to add a static built app (MPA, SPA, etc) to extend the main app
	// basically the pb_public folder to serve frontend in pocketbase
	// the main app UI will be the ./frontend app, the ./wps/public app can be used for UI extentions, plugins, etc
	/*
		app.pbApp.OnServe().BindFunc(func(e *core.ServeEvent) error {
			e.Router.BindFunc(func(r *core.RequestEvent) error {
				// this helps serve the extention UI inside an iframe
				r.Response.Header().Del("X-Frame-Options")
				return r.Next()
			})
			if !e.Router.HasRoute(http.MethodGet, "/{path...}") {
				e.Router.GET("/{path...}", apis.Static(os.DirFS("./wps/public"), true))
			}
			return e.Next()
		})
	*/

	// Start the app
	if err := app.pbApp.Bootstrap(); err != nil {
		log.Println("PB Bootstrapped")
		return err
	}

	// METHOD 1:
	// Serve the app using a http server
	// this is better if your PB app has to serve files to the frontend due to wails AssetServer.Handler not working with vite v5.0.0+
	// NOTE about AssetServer.Handler: This does not work with vite v5.0.0+ and wails v2 due to changes in vite. Changes are planned in v3 to support similar functionality under vite v5.0.0+. If you need this feature, stay with vite v4.0.0+. See issue 3240 for details
	go func() {
		if err := apis.Serve(app.pbApp, apis.ServeConfig{
			HttpAddr:        "127.0.0.1:8080",
			ShowStartBanner: false,
			AllowedOrigins:  []string{"*", "http://wails.localhost:8080/"},
		}); err != nil && err != http.ErrServerClosed {
			log.Println("Error Serving PB")
			log.Fatal(err)
		}
		log.Println("Serving PB")
	}()

	// METHOD 2:
	// If you just want to interact with pocketbas as a database and no files to serve
	// Bootstrap() sets up the datbase connections and other core backend functions, but will not setup cron, http server and user migrations
	/*
		if err := app.pbApp.RunAllMigrations(); err != nil {
			log.Println("Error Migrating PB")
			return err
		}
	*/

	return nil
}

// Stop PB, terminate PB server, close db connection and free resources
func (app *PocketBaseApp) StopPB() error {
	log.Println("Stopping PB App")
	if app.pbApp == nil {
		return errors.New("Cannot Stop PB, Does not exist")
	}

	evt := &core.TerminateEvent{
		App: app.pbApp,
	}

	return app.pbApp.OnTerminate().Trigger(evt, func(e *core.TerminateEvent) error {
		if err := app.pbApp.ResetBootstrapState(); err != nil {
			log.Println("Error resetting PB Bootstrap")
			return err
		}
		return nil
	})

}
