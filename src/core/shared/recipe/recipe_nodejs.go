package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func NodeJS(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"nodejs",   // key
		"Node.js",  // name
		"NodeJS",   // language
		"index.js", // mainfile
		"javascript", // cmmode
		"docker.io/library/node:latest",
		func(notebook types.Notebook) []string {
			return []string{"node", "/code/" + notebook.GetRecipe().GetMainfile()}
		},
		func(notebook types.Notebook) []string {
			return []string{"node", notebook.GetMainFileAbsPath()}
		},
		nil,
		nil,
		recipesFS,
	)
}