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
// # Equipment model declaration
//
// The base model declaration can be found at
// https://github.com/diyerdo/proto/blob/main/proto/equipments/v1/equipments.proto
package models

import (
	"github.com/diyerdo/proto/gen/go/proto/equipments/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Equipment represents an Equipment as declared in the proto file
type Equipment struct {
	Id       int32
	Name     string
	Level    int32
	ImageUrl *string
}

// NewEquipmentDefault returns a new Equipment with default values
func NewEquipmentDefault() *Equipment {
	return &Equipment{}
}

// NewEquipment returns a new Equipment instance
//
// Returns an error if one of the parameters is invalid
func NewEquipment(id int32, name string, level int32, imageUrl *string) (*Equipment, error) {
	if id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be `%d`", id)
	}

	if len(name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "name cannot be empty")
	}

	if level < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "level cannot be `%d`", level)
	}

	return &Equipment{
		Id:       id,
		Name:     name,
		Level:    level,
		ImageUrl: imageUrl,
	}, nil
}

// ToProto converts the Equipment model to the protobuf Equipment
func (t *Equipment) ToProto() *equipments.Equipment {
	return &equipments.Equipment{
		Id:       t.Id,
		Name:     t.Name,
		Level:    t.Level,
		ImageUrl: t.ImageUrl,
	}
}
