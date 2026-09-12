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
// Dodugo utility wrapper
package utils

import (
	"context"

	"github.com/diyerdo/diyerdo/internal/services/equipments/models"
	"github.com/diyerdo/proto/gen/go/proto/equipments/v1"
	"github.com/dofusdude/dodugo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DodugoWrapper is a utility wrapper for dodugo
type DodugoWrapper struct {
	client *dodugo.APIClient
}

// NewDodugoWrapper creates a new DodugoWrapper instance
//
// Returns an error in case client is `nil` or client's config is `nil`
func NewDodugoWrapper() (*DodugoWrapper, error) {
	config := dodugo.NewConfiguration()
	if config == nil {
		return nil, status.Error(codes.Internal, "config is `nil`")
	}

	client := dodugo.NewAPIClient(config)
	if client == nil {
		return nil, status.Error(codes.Internal, "client is `nil`")
	}

	return &DodugoWrapper{
		client: client,
	}, nil
}

// GetEquipmentByNameAndCategory gets equipment by name and category
func (t *DodugoWrapper) GetEquipmentByNameAndCategory(name string, category string) ([]dodugo.ListItem, error) {
	items, resp, err := t.client.
		EquipmentAPI.
		GetItemsEquipmentSearch(context.Background(), "fr", "dofus3").
		Query(name).
		FilterTypeNameId([]string{category}).
		Execute()

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"an error occured while trying to retrieve item '%s' from '%s' category: %v",
			name,
			category,
			err,
		)
	}

	if resp.StatusCode != 200 {
		return nil, status.Errorf(
			codes.Internal,
			"an error occured while trying to retrieve item '%s' from '%s' category ; status code %s",
			name,
			category,
			resp.Status,
		)
	}

	if len(items) == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"no item found with name '%s' in category '%s'",
			name,
			category,
		)
	}

	return items, nil
}

// dodugoItemsToEquipments converts a slice of dodugo.ListItem to a slice of equipments.Equipment
func (t *DodugoWrapper) DodugoItemsToEquipments(items []dodugo.ListItem) ([]*equipments.Equipment, error) {
	if items == nil {
		return nil, status.Error(codes.InvalidArgument, "received a `nil` items list")
	}

	equipments := make([]*equipments.Equipment, len(items))
	for i, item := range items {
		equipment, err := models.NewEquipment(*item.AnkamaId, *item.Name, *item.Level, *item.ImageUrls.Icon)
		if err != nil {
			return nil, err
		}

		equipments[i] = equipment.ToProto()
	}

	return equipments, nil
}
