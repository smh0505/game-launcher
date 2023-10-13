package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Enums
const (
    noFile int = 0
    noPswd int = 1
)

func (a *App) TestGame(dir, pswd string) map[string]any {
    // Get archive file
    newDir := ""
    if strings.TrimSpace(dir) == "" {
        newDir = a.GetFileDir("Install Game", 
                            "",
                            "Archive Files (*.zip;*.rar;*.7z)",
                            "*.zip;*.rar;*.7z")
        if strings.TrimSpace(newDir) == "" { 
            return map[string]any{
                "dir": newDir,
                "success": false,
                "error": noFile,
            }
        }
    } else {
        newDir = dir
    }

    // Password Test
    if err := exec.Command("7z", "t", newDir, fmt.Sprintf("-p%s", pswd)).Run(); err != nil {
        return map[string]any{
            "dir": newDir,
            "success": false,
            "error": noPswd,
        }
    }

    return a.InstallGame(newDir, pswd)
}

func (a *App) InstallGame(dir, pswd string) map[string]any {
	// Extract file name and generate new directory
	file := strings.TrimSuffix(filepath.Base(dir), filepath.Ext(dir))
	newDir, _ := filepath.Abs(filepath.Join("games", file))

	// Run 7z command
    exec.Command("7z", "x", dir, fmt.Sprintf("-p%s", pswd), fmt.Sprintf("-o%s", newDir), "-aoa").Run()
	
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

	return map[string]any{
        "success": true,
		"path": newDir, 
		"name": file,
		"link": exe[0],
	}
}
