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
// Tests for jobs core service
package core

import (
	"testing"

	"github.com/diyerdo/diyerdo/internal/shared/utils"
	"github.com/diyerdo/proto/gen/go/proto/jobs/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewJobsCore(t *testing.T) {
	c, err := NewJobsCore()
	if err != nil {
		t.Fatalf("unexpected error creating JobsCore: %v", err)
	}
	if c == nil {
		t.Fatalf("expected non-nil JobsCore")
	}
}

func TestGetJobForItemInvalidId(t *testing.T) {
	c, err := NewJobsCore()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	invalidIds := []int32{0, -1, -100}
	for _, id := range invalidIds {
		_, err := c.GetJobForItem(id)
		if err == nil {
			t.Errorf("expected error for itemId %d, got nil", id)
		}
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.InvalidArgument {
			t.Errorf("expected InvalidArgument error for id %d, got %v", id, err)
		}
	}
}

func TestGetJobsRequirementsForItemInvalidId(t *testing.T) {
	c, err := NewJobsCore()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	invalidIds := []int32{0, -1, -100}
	for _, id := range invalidIds {
		_, err := c.GetJobsRequirementsForItem(id)
		if err == nil {
			t.Errorf("expected error for itemId %d, got nil", id)
		}
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.InvalidArgument {
			t.Errorf("expected InvalidArgument error for id %d, got %v", id, err)
		}
	}
}

func TestGetJobForItemLive(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live network test in short mode")
	}

	c, err := NewJobsCore()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	jobsList, err := c.GetJobForItem(2469) // Gelano
	if err != nil {
		t.Fatalf("unexpected error calling GetJobForItem: %v", err)
	}
	if len(jobsList) != 1 || jobsList[0] != jobs.Job_JOB_JEWELLER {
		t.Errorf("expected [JOB_JEWELLER] for Gelano (2469), got %v", jobsList)
	}

	uncraftableJobs, err := c.GetJobForItem(972) // Dofus Cawotte (uncraftable)
	if err != nil {
		t.Fatalf("unexpected error calling GetJobForItem for uncraftable: %v", err)
	}
	if len(uncraftableJobs) != 0 {
		t.Errorf("expected empty jobs list for uncraftable item, got %v", uncraftableJobs)
	}
}

func TestGetJobsRequirementsForItemLive(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live network test in short mode")
	}

	c, err := NewJobsCore()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("Equipment with single job (Gelano)", func(t *testing.T) {
		reqs, err := c.GetJobsRequirementsForItem(2469) // Gelano
		if err != nil {
			t.Fatalf("unexpected error calling GetJobsRequirementsForItem: %v", err)
		}
		if len(reqs) != 1 {
			t.Fatalf("expected 1 job requirement for Gelano, got %d", len(reqs))
		}
		if reqs[0].Job != jobs.Job_JOB_JEWELLER || reqs[0].Level != 60 {
			t.Errorf("expected Jeweller Lv 60, got Job=%v Level=%d", reqs[0].Job, reqs[0].Level)
		}
	})

	t.Run("Uncraftable item (Dofus Cawotte)", func(t *testing.T) {
		reqs, err := c.GetJobsRequirementsForItem(972) // Dofus Cawotte
		if err != nil {
			t.Fatalf("unexpected error calling GetJobsRequirementsForItem for uncraftable item: %v", err)
		}
		if len(reqs) != 0 {
			t.Errorf("expected empty job requirements for uncraftable item, got %v", reqs)
		}
	})

	t.Run("Item not found", func(t *testing.T) {
		_, err := c.GetJobsRequirementsForItem(99999999)
		if err == nil {
			t.Fatalf("expected error for nonexistent item, got nil")
		}
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.NotFound {
			t.Errorf("expected NotFound error, got %v", err)
		}
	})

	t.Run("Craftable resource (Planche de Gravure)", func(t *testing.T) {
		reqs, err := c.GetJobsRequirementsForItem(16496) // Planche de Gravure
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(reqs) != 1 {
			t.Fatalf("expected 1 job requirement, got %d", len(reqs))
		}
		if reqs[0].Job != jobs.Job_JOB_LUMBERJACK || reqs[0].Level != 140 {
			t.Errorf("expected Lumberjack Lv 140, got Job=%v Level=%d", reqs[0].Job, reqs[0].Level)
		}
	})

	t.Run("Recursive craft with sub-ingredients (Outil de gravure)", func(t *testing.T) {
		reqs, err := c.GetJobsRequirementsForItem(23578) // Outil de gravure (uses Planche de Gravure which uses Tourmaline/alloy)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(reqs) != 3 {
			t.Fatalf("expected 3 job requirements (Smith, Lumberjack, Miner), got %d: %v", len(reqs), reqs)
		}

		expectedJobs := map[jobs.Job]int32{
			jobs.Job_JOB_LUMBERJACK: 140,
			jobs.Job_JOB_MINER:      50,
			jobs.Job_JOB_SMITH:      140,
		}
		for _, r := range reqs {
			expectedLevel, ok := expectedJobs[r.Job]
			if !ok {
				t.Errorf("unexpected job %v in requirements", r.Job)
				continue
			}
			if r.Level != expectedLevel {
				t.Errorf("expected job %v to have level %d, got %d", r.Job, expectedLevel, r.Level)
			}
		}
	})
}

func TestDofusDBJobIdToJob(t *testing.T) {
	tests := []struct {
		name        string
		jobId       int32
		expectedJob jobs.Job
		expectErr   bool
	}{
		{
			name:        "Lumberjack",
			jobId:       2,
			expectedJob: jobs.Job_JOB_LUMBERJACK,
			expectErr:   false,
		},
		{
			name:        "Smith",
			jobId:       11,
			expectedJob: jobs.Job_JOB_SMITH,
			expectErr:   false,
		},
		{
			name:        "Carver",
			jobId:       13,
			expectedJob: jobs.Job_JOB_CARVER,
			expectErr:   false,
		},
		{
			name:        "Shoemaker",
			jobId:       15,
			expectedJob: jobs.Job_JOB_SHOEMAKER,
			expectErr:   false,
		},
		{
			name:        "Jeweller",
			jobId:       16,
			expectedJob: jobs.Job_JOB_JEWELLER,
			expectErr:   false,
		},
		{
			name:        "Miner",
			jobId:       24,
			expectedJob: jobs.Job_JOB_MINER,
			expectErr:   false,
		},
		{
			name:        "Alchemist",
			jobId:       26,
			expectedJob: jobs.Job_JOB_ALCHEMIST,
			expectErr:   false,
		},
		{
			name:        "Tailor",
			jobId:       27,
			expectedJob: jobs.Job_JOB_TAILOR,
			expectErr:   false,
		},
		{
			name:        "Farmer",
			jobId:       28,
			expectedJob: jobs.Job_JOB_FARMER,
			expectErr:   false,
		},
		{
			name:        "Fisherman",
			jobId:       36,
			expectedJob: jobs.Job_JOB_FISHERMAN,
			expectErr:   false,
		},
		{
			name:        "Hunter",
			jobId:       41,
			expectedJob: jobs.Job_JOB_HUNTER,
			expectErr:   false,
		},
		{
			name:        "Shieldmaker",
			jobId:       60,
			expectedJob: jobs.Job_JOB_SHIELDMAKER,
			expectErr:   false,
		},
		{
			name:        "Handyman",
			jobId:       65,
			expectedJob: jobs.Job_JOB_HANDYMAN,
			expectErr:   false,
		},
		{
			name:        "Unknown job ID",
			jobId:       999,
			expectedJob: jobs.Job_JOB_UNSPECIFIED,
			expectErr:   true,
		},
		{
			name:        "Zero job ID",
			jobId:       0,
			expectedJob: jobs.Job_JOB_UNSPECIFIED,
			expectErr:   true,
		},
		{
			name:        "Negative job ID",
			jobId:       -5,
			expectedJob: jobs.Job_JOB_UNSPECIFIED,
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job, err := dofusDBJobIdToJob(tt.jobId)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error for jobId %d, got nil", tt.jobId)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for jobId %d: %v", tt.jobId, err)
			}
			if job != tt.expectedJob {
				t.Errorf("expected job %v, got %v", tt.expectedJob, job)
			}
		})
	}
}

func TestDofusDBRecipesToJobs(t *testing.T) {
	t.Run("empty recipes", func(t *testing.T) {
		jobList, err := dofusDBRecipesToJobs([]utils.DofusDBRecipe{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(jobList) != 0 {
			t.Errorf("expected 0 jobs, got %d", len(jobList))
		}
	})

	t.Run("single recipe", func(t *testing.T) {
		recipes := []utils.DofusDBRecipe{
			{Id: 1, ResultId: 2469, JobId: 16},
		}
		jobList, err := dofusDBRecipesToJobs(recipes)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(jobList) != 1 || jobList[0] != jobs.Job_JOB_JEWELLER {
			t.Errorf("expected [JOB_JEWELLER], got %v", jobList)
		}
	})

	t.Run("multiple recipes deduplication", func(t *testing.T) {
		recipes := []utils.DofusDBRecipe{
			{Id: 1, ResultId: 2469, JobId: 16},
			{Id: 2, ResultId: 2469, JobId: 16},
			{Id: 3, ResultId: 2469, JobId: 27},
			{Id: 4, ResultId: 2469, JobId: 16},
		}
		jobList, err := dofusDBRecipesToJobs(recipes)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(jobList) != 2 {
			t.Fatalf("expected 2 unique jobs, got %d", len(jobList))
		}
		if jobList[0] != jobs.Job_JOB_JEWELLER || jobList[1] != jobs.Job_JOB_TAILOR {
			t.Errorf("expected [JOB_JEWELLER, JOB_TAILOR], got %v", jobList)
		}
	})

	t.Run("unknown job ID in recipe returns error", func(t *testing.T) {
		recipes := []utils.DofusDBRecipe{
			{Id: 1, ResultId: 100, JobId: 9999},
		}
		_, err := dofusDBRecipesToJobs(recipes)
		if err == nil {
			t.Fatalf("expected error for unknown job ID, got nil")
		}
	})
}

func TestValidJob(t *testing.T) {
	unspecified := "unspecified"
	alchemist := "alchemist"

	if validJob(nil) {
		t.Errorf("expected nil job to be invalid")
	}
	if validJob(&unspecified) {
		t.Errorf("expected 'unspecified' job to be invalid")
	}
	if !validJob(&alchemist) {
		t.Errorf("expected 'alchemist' job to be valid")
	}
}

func jobPtr(j jobs.Job) *jobs.Job {
	return &j
}
