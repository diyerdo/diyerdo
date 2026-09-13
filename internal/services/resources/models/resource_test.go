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
// Tests for resources models
package models

import (
	"testing"
)

func TestNewResource(t *testing.T) {
	validImageUrl := "https://example.com/image.png"

	tests := []struct {
		name      string
		id        int32
		rName     string
		level     int32
		imageUrl  *string
		expectErr bool
	}{
		{
			name:      "valid resource with image",
			id:        100,
			rName:     "Blé",
			level:     1,
			imageUrl:  &validImageUrl,
			expectErr: false,
		},
		{
			name:      "valid resource with nil image",
			id:        100,
			rName:     "Blé",
			level:     1,
			imageUrl:  nil,
			expectErr: false,
		},
		{
			name:      "invalid id zero",
			id:        0,
			rName:     "Blé",
			level:     1,
			imageUrl:  &validImageUrl,
			expectErr: true,
		},
		{
			name:      "invalid id negative",
			id:        -5,
			rName:     "Blé",
			level:     1,
			imageUrl:  &validImageUrl,
			expectErr: true,
		},
		{
			name:      "empty name",
			id:        100,
			rName:     "",
			level:     1,
			imageUrl:  &validImageUrl,
			expectErr: true,
		},
		{
			name:      "invalid level zero",
			id:        100,
			rName:     "Blé",
			level:     0,
			imageUrl:  &validImageUrl,
			expectErr: true,
		},
		{
			name:      "invalid level negative",
			id:        100,
			rName:     "Blé",
			level:     -1,
			imageUrl:  &validImageUrl,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := NewResource(tt.id, tt.rName, tt.level, tt.imageUrl)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if res == nil {
				t.Fatalf("expected non-nil resource")
			}

			if res.Id != tt.id {
				t.Errorf("expected id %d, got %d", tt.id, res.Id)
			}
			if res.Name != tt.rName {
				t.Errorf("expected name %s, got %s", tt.rName, res.Name)
			}
			if res.Level != tt.level {
				t.Errorf("expected level %d, got %d", tt.level, res.Level)
			}
			if res.ImageUrl != tt.imageUrl {
				t.Errorf("expected imageUrl %v, got %v", tt.imageUrl, res.ImageUrl)
			}

			proto := res.ToProto()
			if proto.Id != res.Id || proto.Name != res.Name || proto.Level != res.Level {
				t.Errorf("proto conversion field mismatch")
			}
			if tt.imageUrl != nil && proto.ImageUrl != *tt.imageUrl {
				t.Errorf("expected proto imageUrl %s, got %s", *tt.imageUrl, proto.ImageUrl)
			}
			if tt.imageUrl == nil && proto.ImageUrl != "" {
				t.Errorf("expected empty proto imageUrl, got %s", proto.ImageUrl)
			}
		})
	}
}

func TestNewResourceDefault(t *testing.T) {
	res := NewResourceDefault()
	if res == nil {
		t.Fatalf("expected non-nil default resource")
	}
	if res.Id != 0 || res.Name != "" || res.Level != 0 || res.ImageUrl != nil {
		t.Errorf("expected zero values for default resource")
	}
}
