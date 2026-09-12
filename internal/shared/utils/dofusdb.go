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
// DofusDB utility wrapper
package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const defaultDofusDBBaseURL = "https://api.dofusdb.fr"

// DofusDBRecipe represents the recipe data structure returned by DofusDB
type DofusDBRecipe struct {
	Id       int32 `json:"id"`
	ResultId int32 `json:"resultId"`
	JobId    int32 `json:"jobId"`
}

// dofusDBRecipesResponse represents the response envelope from DofusDB recipes endpoint
type dofusDBRecipesResponse struct {
	Total int32           `json:"total"`
	Data  []DofusDBRecipe `json:"data"`
}

// DofusDBWrapper is a utility wrapper for DofusDB API
type DofusDBWrapper struct {
	client  *http.Client
	baseURL string
}

// NewDofusDBWrapper creates a new DofusDBWrapper instance
func NewDofusDBWrapper() (*DofusDBWrapper, error) {
	return &DofusDBWrapper{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: defaultDofusDBBaseURL,
	}, nil
}

// GetRecipesForResultId retrieves recipes for a given result/item id from DofusDB
func (t *DofusDBWrapper) GetRecipesForResultId(resultId int32) ([]DofusDBRecipe, error) {
	if resultId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "resultId cannot be `%d`", resultId)
	}

	url := fmt.Sprintf("%s/recipes?resultId=%d", t.baseURL, resultId)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create request for resultId '%d': %v", resultId, err)
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "an error occurred while trying to retrieve recipes for '%d': %v", resultId, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.Internal, "failed to retrieve recipes for '%d'; status code %s", resultId, resp.Status)
	}

	var recipesResp dofusDBRecipesResponse
	if err := json.NewDecoder(resp.Body).Decode(&recipesResp); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to decode recipes response for '%d': %v", resultId, err)
	}

	return recipesResp.Data, nil
}
