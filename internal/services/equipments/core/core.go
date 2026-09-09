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

type EquipmentItemLookupCore struct{}

// / NewEquipmentItemLookupCore returns a new instance of EquipmentItemLookupCore
func NewEquipmentItemLookupCore() (*EquipmentItemLookupCore, error) {
	return &EquipmentItemLookupCore{}, nil
}
