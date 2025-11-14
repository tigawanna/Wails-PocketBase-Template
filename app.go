package main

import (
	"context"
	"log"

	"github.com/pocketbase/pocketbase/core"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved so we can call the runtime methods
// Also starts up pocketbase
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go PocketBase()
}

// Terminates pocketbase, close connections and clears up resorces
func (a *App) shutdown(ctx context.Context) {
	if PBApp != nil {
		PBApp.ResetBootstrapState()
	}
}

// TODO LIST METHODS
type ListItem struct {
	ID      string `json:"id"`
	Data    string `json:"data"`
	State   bool   `json:"state"`
	Updated string `json:"updated"`
}

func (a *App) GetTodoList() []ListItem {
	records := []ListItem{}
	PBApp.DB().
		Select("id", "data", "state", "updated").
		From("todo_list").
		OrderBy("updated DESC").
		All(&records)
	log.Println("Todo List fetched successfully")
	return records
}

func (a *App) AddTodo(data string) {
	collection, err := PBApp.FindCollectionByNameOrId("todo_list")
	if err != nil {
		log.Println("Error getting collection")
	}
	record := core.NewRecord(collection)
	record.Set("data", data)
	record.Set("state", false)
	err = PBApp.Save(record)
	if err != nil {
		log.Println("Error saving record")
	}
	log.Println("Todo added successfully")
}

func (a *App) DeleteTodo(id string) {
	record, err := PBApp.FindRecordById("todo_list", id)
	if err != nil {
		log.Println("Error finding record")
	}
	err = PBApp.Delete(record)
	if err != nil {
		log.Println("Error deleting record")
	}
	log.Println("Todo deleted successfully")
}

func (a *App) UpdateTodo(id string, state bool) {
	record, err := PBApp.FindRecordById("todo_list", id)
	if err != nil {
		log.Println("Error finding record")
	}
	record.Set("state", state)
	err = PBApp.Save(record)
	if err != nil {
		log.Println("Error updating record")
	}
}
