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
// Jobs service's implementation
package api

import (
	"github.com/diyerdo/diyerdo/internal/services/jobs/core"
)

// JobsService represents the jobs service
type JobsService struct {
	Server *server
}

// NewJobsService creates a new jobs service instance
func NewJobsService() (*JobsService, error) {
	core, err := core.NewJobsCore()
	if err != nil {
		return nil, err
	}

	server, err := newServer(core)
	if err != nil {
		return nil, err
	}

	return &JobsService{Server: server}, nil
}
