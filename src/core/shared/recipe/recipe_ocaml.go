package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func Ocaml(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"ocaml",    // key
		"OCaml",    // name
		"OCaml",    // language
		"index.ml", // mainfile
		"mllike",   // cmmode
		"docker.io/library/ocaml/opam:latest",
		func(notebook types.Notebook) []string {
			return []string{"ocaml", "/code/" + notebook.GetRecipe().GetMainfile()}
		},
		func(notebook types.Notebook) []string {
			return []string{"ocaml", notebook.GetMainFileAbsPath()}
		},
		nil,
		nil,
		recipesFS,
	)
}