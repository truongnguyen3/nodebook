package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func Rust(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"rust",     // key
		"Rust",     // name
		"Rust",     // language
		"main.rs",  // mainfile
		"rust",     // cmmode
		"docker.io/library/rust:latest",
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "cd /code && if [ -f Cargo.toml ]; then cargo run; else rustc main.rs && ./main; fi"}
		},
		func(notebook types.Notebook) []string {
			return []string{"sh", "-c", "cd " + notebook.GetAbsdir() + " && if [ -f Cargo.toml ]; then cargo run; else rustc main.rs && ./main; fi"}
		},
		nil,
		nil,
		recipesFS,
	)
}