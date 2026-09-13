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
// Jobs service's gRPC implementation
package api

import (
	"context"

	"github.com/diyerdo/diyerdo/internal/services/jobs/core"
	"github.com/diyerdo/proto/gen/go/proto/jobs/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server represents the Jobs service gRPC server
type server struct {
	jobs.UnimplementedJobsServiceServer
	core *core.JobsCore
}

// newServer returns a new instance of the Jobs service gRPC server
func newServer(core *core.JobsCore) (*server, error) {
	if core == nil {
		return nil, status.Error(codes.InvalidArgument, "received a `nil` core")
	}

	return &server{core: core}, nil
}

// GetJobForItem returns the jobs required to craft the given item
func (t *server) GetJobForItem(ctx context.Context, request *jobs.GetJobForItemRequest) (*jobs.GetJobForItemResponse, error) {
	if err := validateGetJobForItemRequest(request); err != nil {
		return nil, err
	}

	jobsList, err := t.core.GetJobForItem(request.ItemId)
	if err != nil {
		return nil, err
	}

	return &jobs.GetJobForItemResponse{Jobs: jobsList}, nil
}

// GetJobsRequirementsForItem returns the job requirements to craft the given item
func (t *server) GetJobsRequirementsForItem(ctx context.Context, request *jobs.GetJobsRequirementsForItemRequest) (*jobs.GetJobsRequirementsForItemResponse, error) {
	if err := validateGetJobsRequirementsForItemRequest(request); err != nil {
		return nil, err
	}

	jobRequirements, err := t.core.GetJobsRequirementsForItem(request.ItemId)
	if err != nil {
		return nil, err
	}

	return &jobs.GetJobsRequirementsForItemResponse{JobRequirements: jobRequirements}, nil
}

// validateGetJobForItemRequest is a helper function that
// validates the GetJobForItemRequest RPC ensuring:
//   - the request is not `nil`
//   - the item_id is valid (>= 1)
func validateGetJobForItemRequest(request *jobs.GetJobForItemRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "received a `nil` request")
	}

	if request.ItemId < 1 {
		return status.Errorf(codes.InvalidArgument, "item_id cannot be `%d`", request.ItemId)
	}

	return nil
}

// validateGetJobsRequirementsForItemRequest is a helper function that
// validates the GetJobsRequirementsForItemRequest RPC ensuring:
//   - the request is not `nil`
//   - the item_id is valid (>= 1)
func validateGetJobsRequirementsForItemRequest(request *jobs.GetJobsRequirementsForItemRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "received a `nil` request")
	}

	if request.ItemId < 1 {
		return status.Errorf(codes.InvalidArgument, "item_id cannot be `%d`", request.ItemId)
	}

	return nil
}
