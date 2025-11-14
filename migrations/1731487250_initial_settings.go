package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		settings := app.Settings()

		settings.Meta.AppName = "Wails PocketBase"
		settings.Meta.AppURL = "http://wails.localhost/"
		// TODO - Hide controls to 'true' on build
		settings.Meta.HideControls = false

		settings.SMTP.Enabled = false
		settings.SMTP.Host = ""
		settings.SMTP.Port = 0
		settings.SMTP.Username = ""
		settings.SMTP.Password = ""
		settings.SMTP.TLS = false

		settings.Logs.MaxDays = 7
		settings.Logs.LogAuthId = false
		settings.Logs.LogIP = false

		return app.Save(settings)
	}, nil)
}
