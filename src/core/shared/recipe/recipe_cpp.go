package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func Cpp(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"cpp",      // key
		"Cpp",     // name
		"C++", // language
		"main.cpp", // mainfile
		"clike",   // cmmode
		"docker.io/library/gcc:latest",
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "cd /code && g++ -o main " + notebook.GetRecipe().GetMainfile() + " && ./main"}
		},
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "cd " + notebook.GetAbsdir() + " && g++ -o main " + notebook.GetRecipe().GetMainfile() + " && ./main"}
		},
		nil,
		nil,
		recipesFS,
	)
}
