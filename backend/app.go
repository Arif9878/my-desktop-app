package backend

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/Arif9878/my-desktop-app/backend/internal"
)

// App struct
type App struct {
	ctx context.Context
	db  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// OnStartup is called when the app starts. The context is saved
// so we can call the runtime methods.
func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
	a.db = internal.InitDB()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// CreateItem creates a new item in the database.
func (a *App) CreateItem(name string) error {
	item := &internal.Item{Name: name}
	result := a.db.Create(item)
	return result.Error
}

// GetItems retrieves all items from the database.
func (a *App) GetItems() ([]internal.Item, error) {
	var items []internal.Item
	result := a.db.Find(&items)
	return items, result.Error
}
