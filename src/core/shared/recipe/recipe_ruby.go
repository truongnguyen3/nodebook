package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func Ruby(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"ruby",     // key
		"Ruby",     // name
		"Ruby",     // language
		"main.rb",  // mainfile
		"ruby",     // cmmode
		"docker.io/library/ruby:latest",
		func(notebook types.Notebook) []string {
			return []string{"ruby", "/code/" + notebook.GetRecipe().GetMainfile()}
		},
		func(notebook types.Notebook) []string {
			return []string{"ruby", notebook.GetMainFileAbsPath()}
		},
		nil,
		nil,
		recipesFS,
	)
}