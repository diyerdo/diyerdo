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
// Equipments service's implementation
package api

import (
	"github.com/diyerdo/diyerdo/internal/services/equipments/core"
)

// EquipmentsService represents the equipments service
type EquipmentsService struct {
	Server *server
}

// NewEquipmentsService creates a new equipments service instance
func NewEquipmentsService() (*EquipmentsService, error) {
	core, err := core.NewEquipmentsCore()
	if err != nil {
		return nil, err
	}

	server, err := newServer(core)
	if err != nil {
		return nil, err
	}

	return &EquipmentsService{Server: server}, nil
}
