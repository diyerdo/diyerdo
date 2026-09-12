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
// Tests for recipes models
package models

import (
	"testing"
)

func TestNewRecipeItem(t *testing.T) {
	tests := []struct {
		name      string
		id        int32
		subtype   string
		quantity  int32
		expectErr bool
	}{
		{
			name:      "valid recipe item",
			id:        1234,
			subtype:   "resources",
			quantity:  10,
			expectErr: false,
		},
		{
			name:      "invalid id zero",
			id:        0,
			subtype:   "resources",
			quantity:  10,
			expectErr: true,
		},
		{
			name:      "invalid id negative",
			id:        -1,
			subtype:   "resources",
			quantity:  10,
			expectErr: true,
		},
		{
			name:      "empty subtype",
			id:        1234,
			subtype:   "",
			quantity:  10,
			expectErr: true,
		},
		{
			name:      "invalid quantity zero",
			id:        1234,
			subtype:   "resources",
			quantity:  0,
			expectErr: true,
		},
		{
			name:      "invalid quantity negative",
			id:        1234,
			subtype:   "resources",
			quantity:  -5,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := NewRecipeItem(tt.id, tt.subtype, tt.quantity)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if item == nil {
				t.Fatalf("expected non-nil item")
			}

			if item.Id != tt.id {
				t.Errorf("expected id %d, got %d", tt.id, item.Id)
			}
			if item.Subtype != tt.subtype {
				t.Errorf("expected subtype %q, got %q", tt.subtype, item.Subtype)
			}
			if item.Quantity != tt.quantity {
				t.Errorf("expected quantity %d, got %d", tt.quantity, item.Quantity)
			}

			proto := item.ToProto()
			if proto.Id != item.Id || proto.Subtype != item.Subtype || proto.Quantity != item.Quantity {
				t.Errorf("proto conversion mismatch")
			}
		})
	}
}

func TestNewRecipeItemDefault(t *testing.T) {
	item := NewRecipeItemDefault()
	if item == nil {
		t.Fatalf("expected non-nil default item")
	}
	if item.Id != 0 || item.Subtype != "" || item.Quantity != 0 {
		t.Errorf("expected zero values for default item")
	}
}

func TestNewRecipe(t *testing.T) {
	t.Run("valid recipe", func(t *testing.T) {
		item, err := NewRecipeItem(1, "resources", 5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		recipe, err := NewRecipe([]*RecipeItem{item})
		if err != nil {
			t.Fatalf("unexpected error creating recipe: %v", err)
		}

		if len(recipe.Items) != 1 {
			t.Fatalf("expected 1 item, got %d", len(recipe.Items))
		}

		proto := recipe.ToProto()
		if len(proto.Items) != 1 {
			t.Fatalf("expected 1 proto item, got %d", len(proto.Items))
		}
		if proto.Items[0].Id != 1 || proto.Items[0].Subtype != "resources" || proto.Items[0].Quantity != 5 {
			t.Errorf("proto conversion mismatch")
		}
	})

	t.Run("nil items slice", func(t *testing.T) {
		_, err := NewRecipe(nil)
		if err == nil {
			t.Fatalf("expected error for nil items, got nil")
		}
	})

	t.Run("empty items slice", func(t *testing.T) {
		recipe, err := NewRecipe([]*RecipeItem{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(recipe.Items) != 0 {
			t.Errorf("expected empty items slice")
		}
		proto := recipe.ToProto()
		if len(proto.Items) != 0 {
			t.Errorf("expected empty proto items slice")
		}
	})
}

func TestNewRecipeDefault(t *testing.T) {
	recipe := NewRecipeDefault()
	if recipe == nil {
		t.Fatalf("expected non-nil default recipe")
	}
	if recipe.Items != nil {
		t.Errorf("expected nil items for default recipe")
	}
}
