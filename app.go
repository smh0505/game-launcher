package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) OpenImageFile() string {
	dir, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: ".",
		Title: "Open Background Image",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg)",
				Pattern: "*.png;*.jpg",
			},
		},
	})

	if err != nil || strings.TrimSpace(dir) == "" {
		return ""
	}

	output := filepath.Ext(dir)
	input, err := os.ReadFile(dir)
	if err != nil {
		return ""
	}
	err = os.WriteFile("background" + output, input, os.FileMode(0777))
	if err != nil {
		return ""
	}
	return "background" + output
}
