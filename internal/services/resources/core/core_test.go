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
// Tests for resources core service
package core

import (
	"testing"

	"github.com/dofusdude/dodugo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewResourcesCore(t *testing.T) {
	c, err := NewResourcesCore()
	if err != nil {
		t.Fatalf("unexpected error creating ResourcesCore: %v", err)
	}
	if c == nil {
		t.Fatalf("expected non-nil ResourcesCore")
	}
}

func TestGetResourceFromIdInvalidId(t *testing.T) {
	c, err := NewResourcesCore()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	invalidIds := []int32{0, -1, -100}
	for _, id := range invalidIds {
		_, err := c.GetResourceFromId(id)
		if err == nil {
			t.Errorf("expected error for resourceId %d, got nil", id)
		}
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.InvalidArgument {
			t.Errorf("expected InvalidArgument error for id %d, got %v", id, err)
		}
	}
}

func TestDodugoResourceToResource(t *testing.T) {
	ankamaId := int32(289)
	name := "Blé"
	level := int32(1)
	imageUrl := "https://example.com/ble.png"

	t.Run("valid resource with sd image", func(t *testing.T) {
		sd := dodugo.NewNullableString(&imageUrl)
		item := &dodugo.Resource{
			AnkamaId: &ankamaId,
			Name:     &name,
			Level:    &level,
			ImageUrls: &dodugo.Images{
				Sd: *sd,
			},
		}

		res, err := dodugoResourceToResource(item)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res == nil {
			t.Fatalf("expected non-nil resource")
		}
		if res.Id != ankamaId || res.Name != name || res.Level != level || res.ImageUrl != imageUrl {
			t.Errorf("field mismatch: got %+v", res)
		}
	})

	t.Run("valid resource with nil sd image", func(t *testing.T) {
		sd := dodugo.NewNullableString(nil)
		item := &dodugo.Resource{
			AnkamaId: &ankamaId,
			Name:     &name,
			Level:    &level,
			ImageUrls: &dodugo.Images{
				Sd: *sd,
			},
		}

		res, err := dodugoResourceToResource(item)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res == nil {
			t.Fatalf("expected non-nil resource")
		}
		if res.Id != ankamaId || res.Name != name || res.Level != level || res.ImageUrl != "" {
			t.Errorf("field mismatch: got %+v", res)
		}
	})

	t.Run("nil item", func(t *testing.T) {
		_, err := dodugoResourceToResource(nil)
		if err == nil {
			t.Fatalf("expected error for nil item, got nil")
		}
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.InvalidArgument {
			t.Errorf("expected InvalidArgument error, got %v", err)
		}
	})

	t.Run("invalid item value", func(t *testing.T) {
		invalidId := int32(0)
		sd := dodugo.NewNullableString(&imageUrl)
		item := &dodugo.Resource{
			AnkamaId: &invalidId,
			Name:     &name,
			Level:    &level,
			ImageUrls: &dodugo.Images{
				Sd: *sd,
			},
		}

		_, err := dodugoResourceToResource(item)
		if err == nil {
			t.Fatalf("expected error for invalid item values, got nil")
		}
	})
}
