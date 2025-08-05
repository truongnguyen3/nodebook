package recipe

import (
	"io/fs"

	"github.com/netgusto/nodebook/src/core/shared/service"
)

func AddRecipesToRegistry(recipeRegistry *service.RecipeRegistry, recipesFS fs.FS) {

	recipeRegistry.
		AddRecipe(C(recipesFS)).
		AddRecipe(Clojure(recipesFS)).
		AddRecipe(Cpp(recipesFS)).
		AddRecipe(Csharp(recipesFS)).
		AddRecipe(Elixir(recipesFS)).
		AddRecipe(Fsharp(recipesFS)).
		AddRecipe(Go(recipesFS)).
		AddRecipe(Haskell(recipesFS)).
		AddRecipe(Java(recipesFS)).
		AddRecipe(Lua(recipesFS)).
		AddRecipe(NodeJS(recipesFS)).
		AddRecipe(Ocaml(recipesFS)).
		AddRecipe(Php(recipesFS)).
		AddRecipe(Python3(recipesFS)).
		AddRecipe(R(recipesFS)).
		AddRecipe(Ruby(recipesFS)).
		AddRecipe(Rust(recipesFS)).
		AddRecipe(Swift(recipesFS)).
		AddRecipe(Typescript(recipesFS))
}
