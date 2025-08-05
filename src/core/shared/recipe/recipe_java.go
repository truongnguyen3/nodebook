package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func Java(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"java",      // key
		"Java",     // name
		"Java", // language
		"main.java", // mainfile
		"clike",   // cmmode
		"docker.io/library/openjdk:latest",
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "cd /code && javac " + notebook.GetRecipe().GetMainfile() + " && java main"}
		},
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "cd " + notebook.GetAbsdir() + " && javac " + notebook.GetRecipe().GetMainfile() + " && java main"}
		},
		nil,
		nil,
		recipesFS,
	)
}
