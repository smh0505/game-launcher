package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App)GetFileDir(title, displayName, pattern, errMsg string) string {
	dir, _ := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: title,
		Filters: []runtime.FileFilter{
			{
				DisplayName: displayName,
				Pattern: pattern,
			},
		},
	})

	if strings.TrimSpace(dir) == "" {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: errMsg,
		})
	}

	return dir
}

func (a *App) OpenImageFile() string {
	// Get image file
	dir := a.GetFileDir("Open Background Image", 
						"Images (*.png;*.jpg;*.gif;*.webp)", 
						"*.png;*jpg;*.gif;*.webp", 
						"Error: Couldn't find the image file")
	if strings.TrimSpace(dir) == "" { return "" }

	// Extract extension
	output := filepath.Ext(dir)

	// Read file
	input, err := os.ReadFile(dir)
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: fmt.Sprintf("Error: Couldn't load %s", filepath.Base(dir)),
		})
		return ""
	}

	// Copy file
	err = os.WriteFile("background" + output, input, os.FileMode(0777))
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: fmt.Sprintf("Error: Couldn't save %s", filepath.Base(dir)),
		})
		return ""
	}

	return "background" + output
}

func (a *App) InstallGame() map[string]string {
	// Get archive file
	dir := a.GetFileDir("Install Game", 
						"Archive Files (*.zip;*.rar;*.7z)", 
						"*.zip;*.rar;*.7z", 
						"Error: Couldn't find the archive file")

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
			Message: fmt.Sprintf("Error: Couldn't extract %s into games folder", filepath.Base(dir)),
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
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: "Error: Couldn't find any executable file",
		})
		exe = append(exe, "")
	}

	return map[string]string{
		"path": newDir, 
		"name": file,
		"link": exe[0],
	}
}

func (a *App) SaveMetadata(data interface{}) {
	bytes, err := yaml.Marshal(data)
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: "Error: Cannot encode the metadata",
		})
		return
	}

	if err = os.WriteFile("games/metadata.yml", bytes, os.FileMode(0777)); err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: "Error: Cannot write the metadata into a file",
		})
	}
}

func (a *App) LoadMetadata() interface{} {
	bytes, err := os.ReadFile("games/metadata.yml")
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: "Error: Cannot read the metadata",
		})
		return nil
	}

	var data interface{}
	if err = yaml.Unmarshal(bytes, &data); err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: "Error: Cannot decode the metadata",
		})
	}
	return data
}
