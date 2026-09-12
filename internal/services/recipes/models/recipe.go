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
// # Recipe model declaration
//
// The base model declaration can be found at
// https://github.com/diyerdo/proto/blob/main/proto/recipes/v1/recipes.proto
package models

import (
	"github.com/diyerdo/proto/gen/go/proto/recipes/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RecipeItem represents an ingredient item in a recipe as declared in the proto file
type RecipeItem struct {
	Id       int32
	Subtype  string
	Quantity int32
}

// NewRecipeItemDefault returns a new RecipeItem with default values
func NewRecipeItemDefault() *RecipeItem {
	return &RecipeItem{}
}

// NewRecipeItem returns a new RecipeItem instance
//
// Returns an error if one of the parameters is invalid
func NewRecipeItem(id int32, subtype string, quantity int32) (*RecipeItem, error) {
	if id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be `%d`", id)
	}

	if len(subtype) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "subtype cannot be empty")
	}

	if quantity < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "quantity cannot be `%d`", quantity)
	}

	return &RecipeItem{
		Id:       id,
		Subtype:  subtype,
		Quantity: quantity,
	}, nil
}

// ToProto converts the RecipeItem model to the protobuf RecipeItem
func (t *RecipeItem) ToProto() *recipes.RecipeItem {
	return &recipes.RecipeItem{
		Id:       t.Id,
		Subtype:  t.Subtype,
		Quantity: t.Quantity,
	}
}

// Recipe represents a recipe containing recipe items as declared in the proto file
type Recipe struct {
	Items []*RecipeItem
}

// NewRecipeDefault returns a new Recipe with default values
func NewRecipeDefault() *Recipe {
	return &Recipe{}
}

// NewRecipe returns a new Recipe instance
//
// Returns an error if items is nil
func NewRecipe(items []*RecipeItem) (*Recipe, error) {
	if items == nil {
		return nil, status.Error(codes.InvalidArgument, "items cannot be `nil`")
	}

	return &Recipe{
		Items: items,
	}, nil
}

// ToProto converts the Recipe model to the protobuf Recipe
func (t *Recipe) ToProto() *recipes.Recipe {
	protoItems := make([]*recipes.RecipeItem, len(t.Items))
	for i, item := range t.Items {
		if item != nil {
			protoItems[i] = item.ToProto()
		}
	}

	return &recipes.Recipe{
		Items: protoItems,
	}
}
