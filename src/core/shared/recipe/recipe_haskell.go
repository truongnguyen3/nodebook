package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func Haskell(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"haskell",      // key
		"Haskell",     // name
		"Haskell", // language
		"main.hs", // mainfile
		"haskell",   // cmmode
		"docker.io/library/haskell:latest",
		func(notebook types.Notebook) []string {
			return []string{"runhaskell", "/code/" + notebook.GetRecipe().GetMainfile()}
		},
		func(notebook types.Notebook) []string {
			return []string{"runhaskell", notebook.GetMainFileAbsPath()}
		},
		nil,
		nil,
		recipesFS,
	)
}
