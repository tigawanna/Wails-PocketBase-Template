package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": "",
			"deleteRule": "",
			"fields": [
				{
					"autogeneratePattern": "[a-z0-9]{15}",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text2918445923",
					"max": 0,
					"min": 0,
					"name": "data",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "bool2744374011",
					"name": "state",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				}
			],
			"id": "pbc_2134014111",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_6xtkjf2FJZ` + "`" + ` ON ` + "`" + `todo_list` + "`" + ` (\n  ` + "`" + `data` + "`" + `,\n  ` + "`" + `state` + "`" + `,\n  ` + "`" + `updated` + "`" + `\n)"
			],
			"listRule": "",
			"name": "todo_list",
			"system": false,
			"type": "base",
			"updateRule": "",
			"viewRule": ""
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2134014111")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
