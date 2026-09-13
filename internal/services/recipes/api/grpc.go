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
// Recipes service's gRPC implementation
package api

import (
	"context"

	"github.com/diyerdo/diyerdo/internal/services/recipes/core"
	"github.com/diyerdo/proto/gen/go/proto/recipes/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server represents the Recipes service gRPC server
type server struct {
	recipes.UnimplementedRecipesServiceServer
	core *core.RecipesCore
}

// newServer returns a new instance of the Recipes service gRPC server
func newServer(core *core.RecipesCore) (*server, error) {
	if core == nil {
		return nil, status.Error(codes.InvalidArgument, "received a `nil` core")
	}

	return &server{core: core}, nil
}

// GetRecipeForEquipment returns the recipe for the given equipment
func (t *server) GetRecipeForEquipment(ctx context.Context, request *recipes.GetRecipeForEquipmentRequest) (*recipes.GetRecipeForEquipmentResponse, error) {
	if err := validateGetRecipeForEquipmentRequest(request); err != nil {
		return nil, err
	}

	recipesList, err := t.core.GetRecipeForEquipment(request.EquipmentId)
	if err != nil {
		return nil, err
	}

	return &recipes.GetRecipeForEquipmentResponse{Recipes: recipesList}, nil
}

// GetRecipeForResource returns the recipe for the given resource
func (t *server) GetRecipeForResource(ctx context.Context, request *recipes.GetRecipeForResourceRequest) (*recipes.GetRecipeForResourceResponse, error) {
	if err := validateGetRecipeForResourceRequest(request); err != nil {
		return nil, err
	}

	recipesList, err := t.core.GetRecipeForResource(request.ResourceId)
	if err != nil {
		return nil, err
	}

	return &recipes.GetRecipeForResourceResponse{Recipes: recipesList}, nil
}

// validateGetRecipeForEquipmentRequest is a helper function that
// validates the GetRecipeForEquipmentRequest RPC ensuring:
//   - the request is not `nil`
//   - the equipment_id is valid (>= 1)
func validateGetRecipeForEquipmentRequest(request *recipes.GetRecipeForEquipmentRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "received a `nil` request")
	}

	if request.EquipmentId < 1 {
		return status.Errorf(codes.InvalidArgument, "equipment_id cannot be `%d`", request.EquipmentId)
	}

	return nil
}

// validateGetRecipeForResourceRequest is a helper function that
// validates the GetRecipeForResourceRequest RPC ensuring:
//   - the request is not `nil`
//   - the resource_id is valid (>= 1)
func validateGetRecipeForResourceRequest(request *recipes.GetRecipeForResourceRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "received a `nil` request")
	}

	if request.ResourceId < 1 {
		return status.Errorf(codes.InvalidArgument, "resource_id cannot be `%d`", request.ResourceId)
	}

	return nil
}
