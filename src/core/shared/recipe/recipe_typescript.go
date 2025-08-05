package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func Typescript(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"typescript", // key
		"TypeScript", // name
		"TypeScript", // language
		"index.ts",   // mainfile
		"javascript", // cmmode
		"docker.io/netgusto/typescript-compiler:latest",
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "cd /code && tsc " + notebook.GetRecipe().GetMainfile() + " && node " + notebook.GetRecipe().GetMainfile()[:len(notebook.GetRecipe().GetMainfile())-3] + ".js"}
		},
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "cd " + notebook.GetAbsdir() + " && tsc " + notebook.GetRecipe().GetMainfile() + " && node " + notebook.GetRecipe().GetMainfile()[:len(notebook.GetRecipe().GetMainfile())-3] + ".js"}
		},
		nil,
		nil,
		recipesFS,
	)
}