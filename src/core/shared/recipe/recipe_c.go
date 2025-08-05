package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func C(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"c",        // key
		"C",        // name
		"C",        // language
		"main.c",   // mainfile
		"clike",    // cmmode
		"docker.io/library/gcc:latest",
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "cd /code && gcc -o main " + notebook.GetRecipe().GetMainfile() + " && ./main"}
		},
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "cd " + notebook.GetAbsdir() + " && gcc -o main " + notebook.GetRecipe().GetMainfile() + " && ./main"}
		},
		nil,
		nil,
		recipesFS,
	)
}