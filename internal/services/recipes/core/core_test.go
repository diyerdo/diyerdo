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
// Tests for recipes core service
package core

import (
	"testing"

	"github.com/dofusdude/dodugo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewRecipesCore(t *testing.T) {
	c, err := NewRecipesCore()
	if err != nil {
		t.Fatalf("unexpected error creating RecipesCore: %v", err)
	}
	if c == nil {
		t.Fatalf("expected non-nil RecipesCore")
	}
}

func TestGetRecipeForEquipmentInvalidId(t *testing.T) {
	c, err := NewRecipesCore()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	invalidIds := []int32{0, -1, -100}
	for _, id := range invalidIds {
		_, err := c.GetRecipeForEquipment(id)
		if err == nil {
			t.Errorf("expected error for equipmentId %d, got nil", id)
		}
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.InvalidArgument {
			t.Errorf("expected InvalidArgument error for id %d, got %v", id, err)
		}
	}
}

func TestDodugoRecipesToRecipeModel(t *testing.T) {
	id := int32(100)
	subtype := "resources"
	qty := int32(4)

	t.Run("valid items", func(t *testing.T) {
		items := []dodugo.Recipe{
			{
				ItemAnkamaId: &id,
				ItemSubtype:  &subtype,
				Quantity:     &qty,
			},
		}

		recipe, err := dodugoRecipesToRecipeModel(items)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if recipe == nil {
			t.Fatalf("expected non-nil recipe")
		}
		if len(recipe.Items) != 1 {
			t.Fatalf("expected 1 item, got %d", len(recipe.Items))
		}
		if recipe.Items[0].Id != id || recipe.Items[0].Subtype != subtype || recipe.Items[0].Quantity != qty {
			t.Errorf("item fields do not match")
		}
	})

	t.Run("nil ItemAnkamaId", func(t *testing.T) {
		items := []dodugo.Recipe{
			{
				ItemAnkamaId: nil,
				ItemSubtype:  &subtype,
				Quantity:     &qty,
			},
		}

		_, err := dodugoRecipesToRecipeModel(items)
		if err == nil {
			t.Fatalf("expected error for nil ItemAnkamaId, got nil")
		}
	})

	t.Run("nil ItemSubtype", func(t *testing.T) {
		items := []dodugo.Recipe{
			{
				ItemAnkamaId: &id,
				ItemSubtype:  nil,
				Quantity:     &qty,
			},
		}

		_, err := dodugoRecipesToRecipeModel(items)
		if err == nil {
			t.Fatalf("expected error for nil ItemSubtype, got nil")
		}
	})

	t.Run("nil Quantity", func(t *testing.T) {
		items := []dodugo.Recipe{
			{
				ItemAnkamaId: &id,
				ItemSubtype:  &subtype,
				Quantity:     nil,
			},
		}

		_, err := dodugoRecipesToRecipeModel(items)
		if err == nil {
			t.Fatalf("expected error for nil Quantity, got nil")
		}
	})

	t.Run("invalid item value", func(t *testing.T) {
		invalidId := int32(-1)
		items := []dodugo.Recipe{
			{
				ItemAnkamaId: &invalidId,
				ItemSubtype:  &subtype,
				Quantity:     &qty,
			},
		}

		_, err := dodugoRecipesToRecipeModel(items)
		if err == nil {
			t.Fatalf("expected error for invalid item values, got nil")
		}
	})
}
