package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "wpb/migrations"
)

var PBApp *pocketbase.PocketBase

func PocketBase() {

	// Create a New PB instance
	PBApp = pocketbase.NewWithConfig(pocketbase.Config{
		HideStartBanner: true,
		DefaultDev:      false,
		DefaultDataDir:  "./wpb/data",
	})

	// Not Needed, but can be used if app users want to extend the app (eg - extentions, plugins, etc)
	jsvm.MustRegister(PBApp, jsvm.Config{
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
	migratecmd.MustRegister(PBApp, PBApp.RootCmd, migratecmd.Config{
		Automigrate:  true,
		TemplateLang: "go",
		Dir:          "./wpb/migrations",
	})

	// Also not needed, but can be used to add a static built app (MPA, SPA, etc) to extend the main app
	// basically the pb_public folder to serve frontend in pocketbase
	// the main app UI will be the ./frontend app, the ./wps/public app can be used for UI extentions, plugins, etc
	PBApp.OnServe().BindFunc(func(e *core.ServeEvent) error {
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

	// Start the app
	if err := PBApp.Bootstrap(); err != nil {
		log.Fatal(err)
	}

	// Serve the app
	if err := apis.Serve(PBApp, apis.ServeConfig{
		HttpAddr:        "127.0.0.1:8080",
		ShowStartBanner: false,
		AllowedOrigins:  []string{"*", "http://wails.localhost:8080/"},
	}); err != nil {
		log.Fatal(err)
	}
}
