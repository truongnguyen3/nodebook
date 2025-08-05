package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/recipe/helper"
	"github.com/netgusto/nodebook/src/core/shared/types"
)

func Lua(recipesFS fs.FS) types.Recipe {
	return helper.StdRecipe(
		"lua",      // key
		"Lua",      // name
		"Lua",      // language
		"main.lua", // mainfile
		"lua",      // cmmode
		"docker.io/library/lua:latest",
		func(notebook types.Notebook) []string {
			return []string{"lua", "/code/" + notebook.GetRecipe().GetMainfile()}
		},
		func(notebook types.Notebook) []string {
			return []string{"lua", notebook.GetMainFileAbsPath()}
		},
		nil,
		nil,
		recipesFS,
	)
}