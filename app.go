package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

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
	exec.Command(dir).Start()
}
