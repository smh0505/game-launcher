package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) LocateThumbnail(oldThumb, name string) string {
	// Get thumbnail file
	dir := a.GetFileDir("Locate Thumbnail File", "", []runtime.FileFilter{
        {
            DisplayName: "Images (*.png;*.jpg;*.gif;*.webp)",
            Pattern: "*.png;*jpg;*.gif;*.webp",
        },
    })
    if strings.TrimSpace(dir) == "" { return "" }

	// Create new file name
	ext := filepath.Ext(dir)
	next := fmt.Sprintf("%s_%d%s", name, time.Now().Unix(), ext)

	os.Mkdir("thumbnails", os.FileMode(0777))
	newThumb := filepath.Join("thumbnails", next)

	// Copy file
	input, _ := os.ReadFile(dir)
	os.WriteFile(newThumb, input, os.FileMode(0777))
	
	os.Remove(oldThumb)
	return newThumb
}

func (a *App) LocateExecutive(pos string) string {
	// Get executive file
	dir := a.GetFileDir("Locate Executive File", pos, []runtime.FileFilter{
        {
            DisplayName: "Executive (*.exe;*.lnk;*.html)",
            Pattern: "*.exe;*.lnk;*.html",
        },
    })
	return dir
}

func (a *App) OpenFolder(dir string) {
	exec.Command("explorer", dir).Start()
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

func (a *App) DeleteGame(dir, thumb string) {
    os.RemoveAll(dir)
    os.Remove(thumb)
}
