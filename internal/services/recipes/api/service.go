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
// Recipes service's implementation
package api

import (
	"github.com/diyerdo/diyerdo/internal/services/recipes/core"
)

// RecipesService represents the recipes service
type RecipesService struct {
	Server *server
}

// NewRecipesService creates a new recipes service instance
func NewRecipesService() (*RecipesService, error) {
	core, err := core.NewRecipesCore()
	if err != nil {
		return nil, err
	}

	server, err := newServer(core)
	if err != nil {
		return nil, err
	}

	return &RecipesService{Server: server}, nil
}
