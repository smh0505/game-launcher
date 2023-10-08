package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
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

func (a *App)GetFileDir(title, defaultDir, displayName, pattern string) string {
	dir, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: title,
		DefaultDirectory: defaultDir,
		Filters: []runtime.FileFilter{
			{
				DisplayName: displayName,
				Pattern: pattern,
			},
		},
	})

	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: fmt.Sprintf("Error: %s", err.Error()),
		})
	}

	return dir
}

func (a *App) OpenImageFile(prev string) string {
	// Get image file
	dir := a.GetFileDir("Open Background Image", 
						"",
						"Images (*.png;*.jpg;*.gif;*.webp)", 
						"*.png;*jpg;*.gif;*.webp")
	if strings.TrimSpace(dir) == "" { return "" }

	// Create new file name
	output := filepath.Ext(dir)
	next := fmt.Sprintf("background_%d%s", time.Now().Unix(), output)

	// Copy file
	input, _ := os.ReadFile(dir)
	os.WriteFile(next, input, os.FileMode(0777))
	
	// Delete previous file
	if err := os.Remove(prev); err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: fmt.Sprintf("Error: %s", err.Error()),
		})
	}

	return next
}

func (a *App) InstallGame() map[string]string {
	// Get archive file
	dir := a.GetFileDir("Install Game", 
						"",
						"Archive Files (*.zip;*.rar;*.7z)", 
						"*.zip;*.rar;*.7z")

	if strings.TrimSpace(dir) == "" { return nil }

	// Extract file name and generate new directory
	file := strings.TrimSuffix(filepath.Base(dir), filepath.Ext(dir))
	newDir := filepath.Join("games", file)

	// Run 7z command
	cmd := exec.Command("7z", "x", dir, fmt.Sprintf(`-o%s`, newDir), "-aoa")
	if err := cmd.Run(); err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: fmt.Sprintf("Error: Cannot extract %s into games folder\n%s", filepath.Base(dir), err.Error()),
		})
		return nil
	}
	
	// Search for the first .exe file
	exe := []string{}
	fs.WalkDir(os.DirFS(newDir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil { return fs.SkipDir }
		if filepath.Ext(path) == ".exe" {
			exe = append(exe, filepath.Join(newDir, path))
		}
		return nil
	})
	if len(exe) == 0 {
		exe = append(exe, "")
	}

	return map[string]string{
		"path": newDir, 
		"name": file,
		"link": exe[0],
	}
}

func (a *App) SaveMetadata(data interface{}) {
	bytes, _ := yaml.Marshal(data)
	os.Mkdir("games", os.FileMode(0777))
	os.WriteFile("games/metadata.yml", bytes, os.FileMode(0777))
}

func (a *App) LoadMetadata() interface{} {
	bytes, _ := os.ReadFile("games/metadata.yml")
	var data interface{}
	yaml.Unmarshal(bytes, &data)
	return data
}

func (a *App) LocateThumbnail() string {
	// Get thumbnail file
	dir := a.GetFileDir("Locate Thumbnail File", 
						".",
						"Images (*.png;*.jpg;*.gif;*.webp)", 
						"*.png;*jpg;*.gif;*.webp")
	if strings.TrimSpace(dir) == "" { return "" }

	// Extract relative path
	curr, _ := os.Getwd()
	target, err := filepath.Rel(curr, dir)
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: "Error: Image file not located in the current location",
		})
		return ""
	}

	return target
}

func (a *App) LocateExecutive() string {
	// Get executive file
	dir := a.GetFileDir("Locate Executive File", 
						".",
						"Executive (*.exe)", 
						"*.exe")
	return dir
}

func (a *App) Start(dir string) {
	if err := exec.Command(dir).Start(); err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: fmt.Sprintf("Error: %s", err.Error()),
		})
	}
}

func (a *App) OpenFolder(dir string) {
	if err := exec.Command("explorer", dir).Start(); err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: fmt.Sprintf("Error: %s", err.Error()),
		})
	}
}
