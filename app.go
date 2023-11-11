package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/skratchdot/open-golang/open"
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

    bytes, err := os.ReadFile("config.yml")
    if err != nil { return }

    var config struct{
        Width int
        Height int
        IsMaximized bool
    }
    yaml.Unmarshal(bytes, &config)

    if config.IsMaximized {
        runtime.WindowMaximise(a.ctx)
    } else {
        runtime.WindowSetSize(a.ctx, config.Width, config.Height)
    }
}

func (a *App) shutdown(ctx context.Context) bool {
    var config struct{
        Width int
        Height int
        IsMaximized bool
    }
    config.Width, config.Height = runtime.WindowGetSize(a.ctx)
    config.IsMaximized = runtime.WindowIsMaximised(a.ctx)

    bytes, err := yaml.Marshal(config)
    if err != nil {
        log.Fatal(err)
    }
    os.WriteFile("config.yml", bytes, os.FileMode(0777))
    return false
}

func (a *App)GetFileDir(title, defaultDir string, filters []runtime.FileFilter) string {
	dir, _ := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: title,
		DefaultDirectory: defaultDir,
		Filters: filters,
	})
	return dir
}

func (a *App) OpenImageFile(prev string) string {
	// Get image file
	dir := a.GetFileDir("Open Background Image", "", []runtime.FileFilter{
        {
            DisplayName: "Images (*.png;*.jpg;*.gif;*.webp)",
            Pattern: "*.png;*jpg;*.gif;*.webp",
        },
    })
	if strings.TrimSpace(dir) == "" { return "" }

	// Create new file name
	output := filepath.Ext(dir)
	next := fmt.Sprintf("background_%d%s", time.Now().Unix(), output)

	// Copy file
	input, _ := os.ReadFile(dir)
	os.WriteFile(next, input, os.FileMode(0777))

	os.Remove(prev)
	return next
}

func (a *App) RenameFolder(name, oldDir, oldExe string) map[string]string {
	fileDir, _ := filepath.Rel(oldDir, oldExe)

	newDir, _ := filepath.Abs(filepath.Join("games", name))
	os.Rename(oldDir, newDir)

	newExe, _ := filepath.Abs(filepath.Join(newDir, fileDir))

	return map[string]string{
		"path": newDir,
		"name": name,
		"link": newExe,
	}
}

func (a *App) Start(dir string) {
	open.Start(dir)
}
