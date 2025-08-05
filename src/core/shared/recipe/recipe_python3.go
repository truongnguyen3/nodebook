package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func Python3(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"python3",  // key
		"Python 3", // name
		"Python",   // language
		"main.py",  // mainfile
		"python",   // cmmode
		"docker.io/library/python:latest",
		func(notebook types.Notebook) []string {
			return []string{"python", "/code/" + notebook.GetRecipe().GetMainfile()}
		},
		func(notebook types.Notebook) []string {
			return []string{"python", notebook.GetMainFileAbsPath()}
		},
		nil,
		nil,
		recipesFS,
	)
}