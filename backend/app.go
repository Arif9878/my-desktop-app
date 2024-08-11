package backend

import (
	"context"
	"database/sql"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
	// a.db = internal.InitDB()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CreateItem(name string) error {
	_, err := a.db.Exec("INSERT INTO items (name) VALUES (?)", name)
	return err
}

func (a *App) GetItems() ([]Item, error) {
	rows, err := a.db.Query("SELECT id, name FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
