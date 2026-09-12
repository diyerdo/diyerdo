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
// diyerdo's equipment item lookup service's core implementation's entrypoint
package core

import (
	"github.com/diyerdo/diyerdo/internal/shared/utils"
	"github.com/diyerdo/proto/gen/go/proto/equipments/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EquipmentsCore struct{}

// NewEquipmentsCore returns a new instance of EquipmentsCore
func NewEquipmentsCore() (*EquipmentsCore, error) {
	return &EquipmentsCore{}, nil
}

// GetEquipmentFromNameAndCategory returns an equipment item from the given name and category
func (t *EquipmentsCore) GetEquipmentFromNameAndCategory(name string, category *equipments.EquipmentCategory) ([]*equipments.Equipment, error) {
	categoryStr, err := equipmentCategoryToString(category)
	if err != nil {
		return nil, err
	}

	if !validCategory(categoryStr) {
		return nil, status.Error(codes.InvalidArgument, "received an `unspecified` category")
	}

	dodugo, err := utils.NewDodugoWrapper()
	if err != nil {
		return nil, err
	}

	items, err := dodugo.GetEquipmentByNameAndCategory(name, *categoryStr)
	if err != nil {
		return nil, err
	}

	convertedItems, err := dodugo.DodugoItemsToEquipments(items)
	if err != nil {
		return nil, err
	}

	return convertedItems, nil
}

// validateCategory tells if the category is unspecified ; in that case it would
// be considered as invalid and thus return false
//
// Returns true otherwise
func validCategory(category *string) bool {
	if category == nil {
		return false
	}

	return !(*category == "unspecified")
}

// equipmentCategoryToString converts an equipment category variant to its string
// representation
func equipmentCategoryToString(category *equipments.EquipmentCategory) (*string, error) {
	if category == nil {
		return nil, status.Error(codes.InvalidArgument, "received a `nil` category")
	}

	str := ""
	switch *category {
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_AMULET:
		str = "amulet"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_AXE:
		str = "axe"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_BELT:
		str = "belt"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_BOOTS:
		str = "boots"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_BOW:
		str = "bow"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_CLOAK:
		str = "cloak"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_DAGGER:
		str = "dagger"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_DOFUS:
		str = "dofus"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_DRAGOTURKEY:
		str = "dragoturkey"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_HAMMER:
		str = "hammer"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_HAT:
		str = "hat"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_LANCE:
		str = "lance"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_MAGIC_WEAPON:
		str = "magic-weapon"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_PET:
		str = "pet"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_PETSMOUNT:
		str = "petsmount"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_PICKAXE:
		str = "pickaxe"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_PRYSMARADITE:
		str = "prysmaradite"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_RHINEETLE:
		str = "rhineetle"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_RING:
		str = "ring"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_SCYTHE:
		str = "scythe"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_SEEMYOOL:
		str = "seemyool"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_SHIELD:
		str = "shield"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_SHOVEL:
		str = "shovel"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_STAFF:
		str = "staff"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_SWORD:
		str = "sword"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_TOOL:
		str = "tool"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_TROPHY:
		str = "trophy"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_WAND:
		str = "wand"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_WEAPON:
		str = "weapon"
	case equipments.EquipmentCategory_EQUIPMENT_CATEGORY_UNSPECIFIED:
		fallthrough
	default:
		str = "unspecified"
	}

	return &str, nil
}
