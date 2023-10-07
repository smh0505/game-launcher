package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
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
				DisplayName: "Images (*.png;*.jpg;*.gif;*.webp)",
				Pattern: "*.png;*.jpg;*.gif;*.webp",
			},
		},
	})

	if err != nil || strings.TrimSpace(dir) == "" {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: "Error: Couldn't find the image file",
		})
		return ""
	}

	output := filepath.Ext(dir)
	input, err := os.ReadFile(dir)
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: fmt.Sprintf("Error: Couldn't load %s", filepath.Base(dir)),
		})
		return ""
	}

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
	dir, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: ".",
		Title: "Install Game",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Archive Files(*.zip;*.rar;*.7z)",
				Pattern: "*.zip;*.rar;*.7z",
			},
		},
	})

	if err != nil || strings.TrimSpace(dir) == "" {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: "Error: Couldn't find the archive",
		})
		return nil
	}

	file := filepath.Base(dir)
	ext := filepath.Ext(file)
	file = strings.TrimSuffix(file, ext)
	newDir := filepath.Join("games", file)

	cmd := exec.Command("7z", "x", dir, fmt.Sprintf(`-o%s`, newDir), "-aoa")
	if err = cmd.Run(); err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type: runtime.ErrorDialog,
			Title: "Internal Error",
			Message: fmt.Sprintf("Error: Couldn't extract %s into games folder", filepath.Base(dir)),
		})
		return nil
	}
	
	exe := []string{}
	fs.WalkDir(os.DirFS(newDir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fs.SkipDir
		}
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

	data := map[string]string{
		"path": newDir, 
		"name": file,
		"link": exe[0],
	}
	return data
}
