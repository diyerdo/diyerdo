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
// Tests for equipments core service
package core

import (
	"testing"

	"github.com/diyerdo/proto/gen/go/proto/equipments/v1"
)

func TestEquipmentCategoryToString(t *testing.T) {
	tests := []struct {
		name        string
		category    *equipments.EquipmentCategory
		expectedStr string
		expectErr   bool
	}{
		{
			name:      "nil category",
			category:  nil,
			expectErr: true,
		},
		{
			name:        "unspecified",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_UNSPECIFIED),
			expectedStr: "unspecified",
			expectErr:   false,
		},
		{
			name:        "amulet",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_AMULET),
			expectedStr: "amulet",
			expectErr:   false,
		},
		{
			name:        "axe",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_AXE),
			expectedStr: "axe",
			expectErr:   false,
		},
		{
			name:        "belt",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_BELT),
			expectedStr: "belt",
			expectErr:   false,
		},
		{
			name:        "boots",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_BOOTS),
			expectedStr: "boots",
			expectErr:   false,
		},
		{
			name:        "bow",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_BOW),
			expectedStr: "bow",
			expectErr:   false,
		},
		{
			name:        "cloak",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_CLOAK),
			expectedStr: "cloak",
			expectErr:   false,
		},
		{
			name:        "dagger",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_DAGGER),
			expectedStr: "dagger",
			expectErr:   false,
		},
		{
			name:        "dofus",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_DOFUS),
			expectedStr: "dofus",
			expectErr:   false,
		},
		{
			name:        "dragoturkey",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_DRAGOTURKEY),
			expectedStr: "dragoturkey",
			expectErr:   false,
		},
		{
			name:        "hammer",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_HAMMER),
			expectedStr: "hammer",
			expectErr:   false,
		},
		{
			name:        "hat",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_HAT),
			expectedStr: "hat",
			expectErr:   false,
		},
		{
			name:        "lance",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_LANCE),
			expectedStr: "lance",
			expectErr:   false,
		},
		{
			name:        "magic weapon",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_MAGIC_WEAPON),
			expectedStr: "magic-weapon",
			expectErr:   false,
		},
		{
			name:        "pet",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_PET),
			expectedStr: "pet",
			expectErr:   false,
		},
		{
			name:        "petsmount",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_PETSMOUNT),
			expectedStr: "petsmount",
			expectErr:   false,
		},
		{
			name:        "pickaxe",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_PICKAXE),
			expectedStr: "pickaxe",
			expectErr:   false,
		},
		{
			name:        "prysmaradite",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_PRYSMARADITE),
			expectedStr: "prysmaradite",
			expectErr:   false,
		},
		{
			name:        "rhineetle",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_RHINEETLE),
			expectedStr: "rhineetle",
			expectErr:   false,
		},
		{
			name:        "ring",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_RING),
			expectedStr: "ring",
			expectErr:   false,
		},
		{
			name:        "scythe",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_SCYTHE),
			expectedStr: "scythe",
			expectErr:   false,
		},
		{
			name:        "seemyool",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_SEEMYOOL),
			expectedStr: "seemyool",
			expectErr:   false,
		},
		{
			name:        "shield",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_SHIELD),
			expectedStr: "shield",
			expectErr:   false,
		},
		{
			name:        "shovel",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_SHOVEL),
			expectedStr: "shovel",
			expectErr:   false,
		},
		{
			name:        "staff",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_STAFF),
			expectedStr: "staff",
			expectErr:   false,
		},
		{
			name:        "sword",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_SWORD),
			expectedStr: "sword",
			expectErr:   false,
		},
		{
			name:        "tool",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_TOOL),
			expectedStr: "tool",
			expectErr:   false,
		},
		{
			name:        "trophy",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_TROPHY),
			expectedStr: "trophy",
			expectErr:   false,
		},
		{
			name:        "wand",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_WAND),
			expectedStr: "wand",
			expectErr:   false,
		},
		{
			name:        "weapon",
			category:    catPtr(equipments.EquipmentCategory_EQUIPMENT_CATEGORY_WEAPON),
			expectedStr: "weapon",
			expectErr:   false,
		},
		{
			name:        "unknown category",
			category:    catPtr(equipments.EquipmentCategory(999)),
			expectedStr: "unspecified",
			expectErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := equipmentCategoryToString(tt.category)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res == nil {
				t.Fatalf("expected non-nil result")
			}
			if *res != tt.expectedStr {
				t.Errorf("expected %q, got %q", tt.expectedStr, *res)
			}
		})
	}
}

func TestValidCategory(t *testing.T) {
	unspecified := "unspecified"
	amulet := "amulet"

	if validCategory(nil) {
		t.Errorf("expected nil category to be invalid")
	}
	if validCategory(&unspecified) {
		t.Errorf("expected 'unspecified' category to be invalid")
	}
	if !validCategory(&amulet) {
		t.Errorf("expected 'amulet' category to be valid")
	}
}

func catPtr(c equipments.EquipmentCategory) *equipments.EquipmentCategory {
	return &c
}
