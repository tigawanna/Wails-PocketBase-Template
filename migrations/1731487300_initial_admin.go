package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		admin, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
		if err != nil {
			return err
		}

		record := core.NewRecord(admin)

		record.Set("email", "admin@wails.app")
		record.Set("password", "adminpassword")

		return app.Save(record)
	}, nil)
}
