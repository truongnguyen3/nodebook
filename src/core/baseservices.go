package core

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/netgusto/nodebook/src/core/shared/recipe"
	"github.com/netgusto/nodebook/src/core/shared/service"
	pkgErrors "github.com/pkg/errors"
)

func baseServices(notebooksPath string, recipesFS fs.FS) (*service.RecipeRegistry, *service.NotebookRegistry) {
	// Recipe registry
	recipeRegistry := service.NewRecipeRegistry()
	recipe.AddRecipesToRegistry(recipeRegistry, recipesFS)

	// Notebook registry
	nbRegistry := service.NewNotebookRegistry(notebooksPath, recipeRegistry)

	// Find notebooks
	notebooks, err := nbRegistry.FindNotebooks(nbRegistry.GetNotebooksPath())
	if err != nil {
		fmt.Println(pkgErrors.Wrapf(err, "Could not find notebooks in %s", nbRegistry.GetNotebooksPath()))
		os.Exit(1)
	}

	// Register notebooks
	for _, notebook := range notebooks {
		nbRegistry.RegisterNotebook(notebook)
	}

	return recipeRegistry, nbRegistry
}
