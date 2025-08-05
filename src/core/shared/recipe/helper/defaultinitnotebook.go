package helper

import (
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/netgusto/nodebook/src/core/shared/types"
	"github.com/pkg/errors"
)

func defaultInitNotebook(recipe types.Recipe, notebookspath, name string, recipesFS fs.FS) error {

	dirPerms := os.FileMode(0755)
	filePerm := os.FileMode(0644)

	// Remove /src/recipes prefix from recipe.GetDir()
	recipeDir := strings.TrimPrefix(recipe.GetDir(), "/src/recipes/")
	srcPath := path.Join(recipeDir, "defaultcontent")
	destPathRoot := path.Join(notebookspath, name)
	
	if err := os.MkdirAll(destPathRoot, dirPerms); err != nil {
		return errors.Wrap(err, "defaultInitNotebook: Could not create notebook directory "+destPathRoot)
	}

	return fs.WalkDir(recipesFS, srcPath, func(pathStr string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// make path relative to srcPath
		relPath := strings.TrimPrefix(pathStr, srcPath)
		if relPath == "" {
			return nil // skip the root directory itself
		}
		
		destPath := path.Join(destPathRoot, relPath)

		if d.IsDir() {
			if err := os.MkdirAll(destPath, dirPerms); err != nil {
				return errors.Wrap(err, "defaultInitNotebook: Could not create notebook directory "+destPath)
			}
		} else {
			dir := filepath.Dir(destPath)
			if err := os.MkdirAll(dir, dirPerms); err != nil {
				return errors.Wrap(err, "defaultInitNotebook: Could not create notebook directory "+dir)
			}

			source, err := recipesFS.Open(pathStr)
			if err != nil {
				return errors.Wrap(err, "defaultInitNotebook: Could not open notebook default content file "+pathStr)
			}
			defer source.Close()

			destination, err := os.OpenFile(destPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, filePerm)
			if err != nil {
				return errors.Wrap(err, "defaultInitNotebook: Could not create notebook file "+destPath)
			}
			defer destination.Close()

			_, err = io.Copy(destination, source)
			if err != nil {
				return errors.Wrap(err, "defaultInitNotebook: Could not copy notebook default content file "+pathStr)
			}
		}

		return nil
	})
}
