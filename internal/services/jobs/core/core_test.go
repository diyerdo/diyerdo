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

func TestJobToString(t *testing.T) {
	tests := []struct {
		name        string
		job         *jobs.Job
		expectedStr string
		expectErr   bool
	}{
		{
			name:      "nil job",
			job:       nil,
			expectErr: true,
		},
		{
			name:        "unspecified",
			job:         jobPtr(jobs.Job_JOB_UNSPECIFIED),
			expectedStr: "unspecified",
			expectErr:   false,
		},
		{
			name:        "alchemist",
			job:         jobPtr(jobs.Job_JOB_ALCHEMIST),
			expectedStr: "alchemist",
			expectErr:   false,
		},
		{
			name:        "carver",
			job:         jobPtr(jobs.Job_JOB_CARVER),
			expectedStr: "carver",
			expectErr:   false,
		},
		{
			name:        "farmer",
			job:         jobPtr(jobs.Job_JOB_FARMER),
			expectedStr: "farmer",
			expectErr:   false,
		},
		{
			name:        "fisherman",
			job:         jobPtr(jobs.Job_JOB_FISHERMAN),
			expectedStr: "fisherman",
			expectErr:   false,
		},
		{
			name:        "handyman",
			job:         jobPtr(jobs.Job_JOB_HANDYMAN),
			expectedStr: "handyman",
			expectErr:   false,
		},
		{
			name:        "hunter",
			job:         jobPtr(jobs.Job_JOB_HUNTER),
			expectedStr: "hunter",
			expectErr:   false,
		},
		{
			name:        "jeweller",
			job:         jobPtr(jobs.Job_JOB_JEWELLER),
			expectedStr: "jeweller",
			expectErr:   false,
		},
		{
			name:        "lumberjack",
			job:         jobPtr(jobs.Job_JOB_LUMBERJACK),
			expectedStr: "lumberjack",
			expectErr:   false,
		},
		{
			name:        "miner",
			job:         jobPtr(jobs.Job_JOB_MINER),
			expectedStr: "miner",
			expectErr:   false,
		},
		{
			name:        "shieldmaker",
			job:         jobPtr(jobs.Job_JOB_SHIELDMAKER),
			expectedStr: "shieldmaker",
			expectErr:   false,
		},
		{
			name:        "shoemaker",
			job:         jobPtr(jobs.Job_JOB_SHOEMAKER),
			expectedStr: "shoemaker",
			expectErr:   false,
		},
		{
			name:        "smith",
			job:         jobPtr(jobs.Job_JOB_SMITH),
			expectedStr: "smith",
			expectErr:   false,
		},
		{
			name:        "tailor",
			job:         jobPtr(jobs.Job_JOB_TAILOR),
			expectedStr: "tailor",
			expectErr:   false,
		},
		{
			name:        "unknown job",
			job:         jobPtr(jobs.Job(999)),
			expectedStr: "unspecified",
			expectErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := jobToString(tt.job)
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
				t.Fatalf("expected non-nil result")
			}
			if *res != tt.expectedStr {
				t.Errorf("expected %q, got %q", tt.expectedStr, *res)
			}
		})
	}
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
