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
// Tests for DofusDB utility wrapper
package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewDofusDBWrapper(t *testing.T) {
	wrapper, err := NewDofusDBWrapper()
	if err != nil {
		t.Fatalf("unexpected error creating DofusDBWrapper: %v", err)
	}
	if wrapper == nil {
		t.Fatalf("expected non-nil DofusDBWrapper")
	}
}

func TestGetRecipesForResultIdInvalidId(t *testing.T) {
	wrapper, err := NewDofusDBWrapper()
	if err != nil {
		t.Fatalf("unexpected error creating DofusDBWrapper: %v", err)
	}

	invalidIds := []int32{0, -1, -50}
	for _, id := range invalidIds {
		_, err := wrapper.GetRecipesForResultId(id)
		if err == nil {
			t.Errorf("expected error for id %d, got nil", id)
		}
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.InvalidArgument {
			t.Errorf("expected InvalidArgument error code for id %d, got %v", id, err)
		}
	}
}

func TestGetRecipesForResultIdSuccess(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("resultId") != "2469" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"total":1,"limit":10,"skip":0,"data":[{"id":2469,"resultId":2469,"jobId":16}]}`))
	}))
	defer mockServer.Close()

	wrapper := &DofusDBWrapper{
		client:  mockServer.Client(),
		baseURL: mockServer.URL,
	}

	recipes, err := wrapper.GetRecipesForResultId(2469)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(recipes) != 1 {
		t.Fatalf("expected 1 recipe, got %d", len(recipes))
	}
	if recipes[0].ResultId != 2469 || recipes[0].JobId != 16 {
		t.Errorf("recipe fields mismatch: got %+v", recipes[0])
	}
}

func TestGetRecipesForResultIdEmpty(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"total":0,"limit":10,"skip":0,"data":[]}`))
	}))
	defer mockServer.Close()

	wrapper := &DofusDBWrapper{
		client:  mockServer.Client(),
		baseURL: mockServer.URL,
	}

	recipes, err := wrapper.GetRecipesForResultId(972)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(recipes) != 0 {
		t.Fatalf("expected 0 recipes, got %d", len(recipes))
	}
}

func TestGetRecipesForResultIdServerError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	wrapper := &DofusDBWrapper{
		client:  mockServer.Client(),
		baseURL: mockServer.URL,
	}

	_, err := wrapper.GetRecipesForResultId(100)
	if err == nil {
		t.Fatalf("expected error on 500 status code, got nil")
	}
	st, ok := status.FromError(err)
	if !ok || st.Code() != codes.Internal {
		t.Errorf("expected Internal error code, got %v", err)
	}
}

func TestGetRecipesForResultIdInvalidJSON(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`invalid json`))
	}))
	defer mockServer.Close()

	wrapper := &DofusDBWrapper{
		client:  mockServer.Client(),
		baseURL: mockServer.URL,
	}

	_, err := wrapper.GetRecipesForResultId(100)
	if err == nil {
		t.Fatalf("expected error on invalid json, got nil")
	}
	st, ok := status.FromError(err)
	if !ok || st.Code() != codes.Internal {
		t.Errorf("expected Internal error code, got %v", err)
	}
}
