// diyerdo backend implementation
// Copyright (C) 2026 DrLarck
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//
// ---
//
// diyerdo's equipment recipe lookup service's core implementation's entrypoint
package core

import (
	"github.com/diyerdo/diyerdo/internal/services/recipes/models"
	"github.com/diyerdo/diyerdo/internal/shared/utils"
	"github.com/diyerdo/proto/gen/go/proto/recipes/v1"
	"github.com/dofusdude/dodugo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecipesCore struct{}

// NewRecipesCore returns a new instance of RecipesCore
func NewRecipesCore() (*RecipesCore, error) {
	return &RecipesCore{}, nil
}

// GetRecipeForEquipment returns the recipe for the given equipment ID
func (t *RecipesCore) GetRecipeForEquipment(equipmentId int32) ([]*recipes.Recipe, error) {
	if equipmentId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "equipmentId cannot be `%d`", equipmentId)
	}

	dodugoWrapper, err := utils.NewDodugoWrapper()
	if err != nil {
		return nil, err
	}

	dodugoRecipes, err := dodugoWrapper.GetEquipmentRecipe(equipmentId)
	if err != nil {
		return nil, err
	}

	if len(dodugoRecipes) == 0 {
		return []*recipes.Recipe{}, nil
	}

	recipeModel, err := dodugoRecipesToRecipeModel(dodugoRecipes)
	if err != nil {
		return nil, err
	}

	return []*recipes.Recipe{recipeModel.ToProto()}, nil
}

// dodugoRecipesToRecipeModel converts a slice of dodugo.Recipe into a models.Recipe
func dodugoRecipesToRecipeModel(items []dodugo.Recipe) (*models.Recipe, error) {
	recipeItems := make([]*models.RecipeItem, len(items))
	for i, item := range items {
		if item.ItemAnkamaId == nil || item.ItemSubtype == nil || item.Quantity == nil {
			return nil, status.Error(codes.Internal, "dodugo recipe item contains nil fields")
		}

		recipeItem, err := models.NewRecipeItem(*item.ItemAnkamaId, *item.ItemSubtype, *item.Quantity)
		if err != nil {
			return nil, err
		}

		recipeItems[i] = recipeItem
	}

	return models.NewRecipe(recipeItems)
}
