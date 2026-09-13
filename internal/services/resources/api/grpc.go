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
// Resources service's gRPC implementation
package api

import (
	"context"

	"github.com/diyerdo/diyerdo/internal/services/resources/core"
	"github.com/diyerdo/proto/gen/go/proto/resources/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server represents the Resources service gRPC server
type server struct {
	resources.UnimplementedResourcesServiceServer
	core *core.ResourcesCore
}

// newServer returns a new instance of the Resources service gRPC server
func newServer(core *core.ResourcesCore) (*server, error) {
	if core == nil {
		return nil, status.Error(codes.InvalidArgument, "received a `nil` core")
	}

	return &server{core: core}, nil
}

// GetResourceFromId returns the resource from the given resource ID
func (t *server) GetResourceFromId(ctx context.Context, request *resources.GetResourceFromIdRequest) (*resources.GetResourceFromIdResponse, error) {
	if err := validateGetResourceFromIdRequest(request); err != nil {
		return nil, err
	}

	resource, err := t.core.GetResourceFromId(request.ResourceId)
	if err != nil {
		return nil, err
	}

	return &resources.GetResourceFromIdResponse{Resource: resource}, nil
}

// validateGetResourceFromIdRequest is a helper function that
// validates the GetResourceFromIdRequest RPC ensuring:
//   - the request is not `nil`
//   - the resource_id is valid (>= 1)
func validateGetResourceFromIdRequest(request *resources.GetResourceFromIdRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "received a `nil` request")
	}

	if request.ResourceId < 1 {
		return status.Errorf(codes.InvalidArgument, "resource_id cannot be `%d`", request.ResourceId)
	}

	return nil
}
