package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func Clojure(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"clojure",      // key
		"Clojure",     // name
		"Clojure", // language
		"index.clj", // mainfile
		"clojure",   // cmmode
		"docker.io/library/clojure:latest",
		func(notebook types.Notebook) []string {
			return []string{"clojure", "/code/" + notebook.GetRecipe().GetMainfile()}
		},
		func(notebook types.Notebook) []string {
			return []string{"clojure", notebook.GetMainFileAbsPath()}
		},
		nil,
		nil,
		recipesFS,
	)
}
