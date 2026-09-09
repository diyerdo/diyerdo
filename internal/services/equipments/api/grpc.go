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
// Equipments service's gRPC implementation
package api

import (
	"context"

	"github.com/diyerdo/proto/gen/go/proto/equipments/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server represents the Equipments service gRPC server
type server struct {
	equipments.UnimplementedEquipmentsServiceServer
}

// NewServer returns a new instance of the Equipments service gRPC server
func newServer() *server {
	return &server{}
}

// GetEquipmentFromNameAndCategory returns the equipment from the given name and category
func (t *server) GetEquipmentFromNameAndCategory(ctx context.Context, request *equipments.GetEquipmentFromNameAndCategoryRequest) (*equipments.GetEquipmentFromNameAndCategoryResponse, error) {
	if err := validateGetEquipmentFromNameAndCategoryRequest(request); err != nil {
		return nil, err
	}

	return &equipments.GetEquipmentFromNameAndCategoryResponse{}, nil
}

// validateGetEquipmentFromNameAndCategoryRequest is a helper function that
// validates the GetEquipmentFromNameAndCategoryRequest RPC ensuring:
//   - the request is not `nil`
//   - the name is not empty
//   - the category is not unspecified
func validateGetEquipmentFromNameAndCategoryRequest(request *equipments.GetEquipmentFromNameAndCategoryRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "received a `nil` request")
	}

	if len(request.Name) == 0 {
		return status.Error(codes.InvalidArgument, "received an empty name")
	}

	if request.Category == equipments.EquipmentCategory_EQUIPMENT_CATEGORY_UNSPECIFIED {
		return status.Error(codes.InvalidArgument, "received an unspecified category")
	}

	return nil
}
