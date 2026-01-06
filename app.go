package main

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
)

// TODO LIST METHODS

type ListItem struct {
	ID      string `json:"id"`
	Data    string `json:"data"`
	State   bool   `json:"state"`
	Updated string `json:"updated"`
}

func (app *WailsApp) GetTodoList() []ListItem {
	records := []ListItem{}
	err := app.pbApp.GetPB().DB().
		Select("id", "data", "state", "updated").
		From("todo_list").
		OrderBy("state ASC", "updated DESC").
		All(&records)
	if err != nil {
		log.Println("Error fetching records")
		return records
	}
	log.Println("Todo List fetched successfully")
	return records
}

func (app *WailsApp) AddTodo(data string) {
	collection, err := app.pbApp.GetPB().FindCollectionByNameOrId("todo_list")
	if err != nil {
		log.Println("Error getting collection")
		return
	}
	record := core.NewRecord(collection)
	record.Set("data", data)
	record.Set("state", false)
	err = app.pbApp.GetPB().Save(record)
	if err != nil {
		log.Println("Error saving record")
		return
	}
	log.Println("Todo added successfully")
}

func (app *WailsApp) DeleteTodo(id string) {
	record, err := app.pbApp.GetPB().FindRecordById("todo_list", id)
	if err != nil {
		log.Println("Error finding record")
		return
	}
	err = app.pbApp.GetPB().Delete(record)
	if err != nil {
		log.Println("Error deleting record")
		return
	}
	log.Println("Todo deleted successfully")
}

func (app *WailsApp) UpdateTodo(id string, state bool) {
	record, err := app.pbApp.GetPB().FindRecordById("todo_list", id)
	if err != nil {
		log.Println("Error finding record")
		return
	}
	record.Set("state", state)
	err = app.pbApp.GetPB().Save(record)
	if err != nil {
		log.Println("Error updating record")
		return
	}
	log.Println("Todo updated successfully")
}
