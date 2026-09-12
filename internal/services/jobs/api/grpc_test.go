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
// Tests for jobs gRPC service
package api

import (
	"testing"

	"github.com/diyerdo/proto/gen/go/proto/jobs/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewServerNilCore(t *testing.T) {
	_, err := newServer(nil)
	if err == nil {
		t.Fatalf("expected error for nil core, got nil")
	}
}

func TestValidateGetJobForItemRequest(t *testing.T) {
	tests := []struct {
		name      string
		request   *jobs.GetJobForItemRequest
		expectErr bool
	}{
		{
			name:      "nil request",
			request:   nil,
			expectErr: true,
		},
		{
			name: "invalid item_id zero",
			request: &jobs.GetJobForItemRequest{
				ItemId: 0,
			},
			expectErr: true,
		},
		{
			name: "invalid item_id negative",
			request: &jobs.GetJobForItemRequest{
				ItemId: -10,
			},
			expectErr: true,
		},
		{
			name: "valid request",
			request: &jobs.GetJobForItemRequest{
				ItemId: 2469,
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateGetJobForItemRequest(tt.request)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				st, ok := status.FromError(err)
				if !ok || st.Code() != codes.InvalidArgument {
					t.Errorf("expected InvalidArgument code, got %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
