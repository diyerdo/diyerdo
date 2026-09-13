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
// diyerdo's resource item lookup service's core implementation's entrypoint
package core

import (
	"github.com/diyerdo/diyerdo/internal/services/resources/models"
	"github.com/diyerdo/diyerdo/internal/shared/utils"
	"github.com/diyerdo/proto/gen/go/proto/resources/v1"
	"github.com/dofusdude/dodugo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ResourcesCore struct{}

// NewResourcesCore returns a new instance of ResourcesCore
func NewResourcesCore() (*ResourcesCore, error) {
	return &ResourcesCore{}, nil
}

// GetResourceFromId returns the resource for the given resource ID
func (t *ResourcesCore) GetResourceFromId(resourceId int32) (*resources.Resource, error) {
	if resourceId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "resourceId cannot be `%d`", resourceId)
	}

	dodugoWrapper, err := utils.NewDodugoWrapper()
	if err != nil {
		return nil, err
	}

	dodugoResource, err := dodugoWrapper.GetResource(resourceId)
	if err != nil {
		return nil, err
	}

	return dodugoResourceToResource(dodugoResource)
}

// dodugoResourceToResource converts a dodugo.Resource to a resources.Resource
func dodugoResourceToResource(item *dodugo.Resource) (*resources.Resource, error) {
	if item == nil {
		return nil, status.Error(codes.InvalidArgument, "received a `nil` item")
	}

	resource, err := models.NewResource(*item.AnkamaId, *item.Name, *item.Level, item.ImageUrls.Sd.Get())
	if err != nil {
		return nil, err
	}

	return resource.ToProto(), nil
}
