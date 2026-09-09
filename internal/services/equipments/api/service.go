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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EquipmentsService represents the equipments service
type EquipmentsService struct {
	Server *server
}

// NewEquipmentsService creates a new equipments service instance
func NewEquipmentsService() (*EquipmentsService, error) {
	server := newServer()
	if server == nil {
		return nil, status.Error(codes.Internal, "failed to instanciate gRPC server ; gRPC server is `nil`")
	}

	return &EquipmentsService{Server: server}, nil
}
