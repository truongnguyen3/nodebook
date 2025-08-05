package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/alecthomas/kingpin"
	"github.com/netgusto/nodebook/src/core"
	"github.com/pkg/errors"
)

//go:embed dist/frontend
var frontendFS embed.FS

//go:embed src/recipes
var recipesFS embed.FS

// GetFrontendFS returns the embedded frontend filesystem
func GetFrontendFS() embed.FS {
	return frontendFS
}

// GetRecipesFS returns the embedded recipes filesystem  
func GetRecipesFS() embed.FS {
	return recipesFS
}

func main() {

	webCmd := kingpin.Command("web", "Web")
	webCmdPath := webCmd.Arg("notebookspath", "path to notebooks").Default(".").ExistingDir()
	webCmdDocker := webCmd.Flag("docker", "Use docker").Bool()
	webCmdPort := webCmd.Flag("port", "HTTP port").Default("8000").Int()
	webCmdBindAddress := webCmd.Flag("bindaddress", "Bind address").Default("127.0.0.1").String()

	cliCmd := kingpin.Command("cli", "cli")
	cliCmdPath := cliCmd.Arg("notebookspath", "path to notebooks").Default(".").ExistingDir()
	cliCmdDocker := cliCmd.Flag("docker", "Use docker").Bool()

	args := os.Args[1:]
	if len(args) == 0 || (args[0] != webCmd.FullCommand() && args[0] != cliCmd.FullCommand()) {
		args = append([]string{"web"}, args...)
	}
	selected, err := kingpin.CommandLine.Parse(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch selected {
	case webCmd.FullCommand():
		absPath, err := absolutizePath(*webCmdPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Create a subfilesystem for the frontend (removing "dist/" prefix)
		frontendSubFS, err := fs.Sub(frontendFS, "dist/frontend")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		core.WebRun(absPath, *webCmdDocker, *webCmdBindAddress, *webCmdPort, frontendSubFS, recipesFS)
	case cliCmd.FullCommand():
		absPath, err := absolutizePath(*cliCmdPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		core.CliRun(absPath, *cliCmdDocker, recipesFS)
	}
}

func absolutizePath(notebooksPath string) (string, error) {
	absBookPath, err := filepath.Abs(notebooksPath)
	if err != nil {
		return "", errors.Wrapf(err, "Could not determine absolute path for \"%s\"", notebooksPath)
	}

	return absBookPath, nil
}
