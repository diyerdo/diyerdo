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
// diyerdo's jobs lookup service's core implementation's entrypoint
package core

import (
	"github.com/diyerdo/diyerdo/internal/shared/utils"
	"github.com/diyerdo/proto/gen/go/proto/jobs/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JobsCore struct{}

// NewJobsCore returns a new instance of JobsCore
func NewJobsCore() (*JobsCore, error) {
	return &JobsCore{}, nil
}

// GetJobForItem returns the jobs required to craft the given item ID
func (t *JobsCore) GetJobForItem(itemId int32) ([]jobs.Job, error) {
	if itemId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "itemId cannot be `%d`", itemId)
	}

	dofusDBWrapper, err := utils.NewDofusDBWrapper()
	if err != nil {
		return nil, err
	}

	dofusDBRecipes, err := dofusDBWrapper.GetRecipesForResultId(itemId)
	if err != nil {
		return nil, err
	}

	if len(dofusDBRecipes) == 0 {
		return []jobs.Job{}, nil
	}

	return dofusDBRecipesToJobs(dofusDBRecipes)
}

// dofusDBRecipesToJobs converts a slice of utils.DofusDBRecipe into a slice of jobs.Job
// while deduplicating jobs and preserving their order
func dofusDBRecipesToJobs(recipes []utils.DofusDBRecipe) ([]jobs.Job, error) {
	seen := make(map[jobs.Job]bool)
	jobList := make([]jobs.Job, 0, len(recipes))

	for _, recipe := range recipes {
		job, err := dofusDBJobIdToJob(recipe.JobId)
		if err != nil {
			return nil, err
		}

		if !seen[job] {
			seen[job] = true
			jobList = append(jobList, job)
		}
	}

	return jobList, nil
}

// dofusDBJobIdToJob converts a DofusDB job ID to the protobuf Job enum
func dofusDBJobIdToJob(jobId int32) (jobs.Job, error) {
	switch jobId {
	case 2:
		return jobs.Job_JOB_LUMBERJACK, nil
	case 11:
		return jobs.Job_JOB_SMITH, nil
	case 13:
		return jobs.Job_JOB_CARVER, nil
	case 15:
		return jobs.Job_JOB_SHOEMAKER, nil
	case 16:
		return jobs.Job_JOB_JEWELLER, nil
	case 24:
		return jobs.Job_JOB_MINER, nil
	case 26:
		return jobs.Job_JOB_ALCHEMIST, nil
	case 27:
		return jobs.Job_JOB_TAILOR, nil
	case 28:
		return jobs.Job_JOB_FARMER, nil
	case 36:
		return jobs.Job_JOB_FISHERMAN, nil
	case 41:
		return jobs.Job_JOB_HUNTER, nil
	case 60:
		return jobs.Job_JOB_SHIELDMAKER, nil
	case 65:
		return jobs.Job_JOB_HANDYMAN, nil
	default:
		return jobs.Job_JOB_UNSPECIFIED, status.Errorf(codes.Internal, "unknown DofusDB job ID `%d`", jobId)
	}
}

// validJob tells if the job is unspecified ; in that case it would
// be considered as invalid and thus return false
//
// Returns true otherwise
func validJob(job *string) bool {
	if job == nil {
		return false
	}

	return !(*job == "unspecified")
}

// jobToString converts a job variant to its string representation
func jobToString(job *jobs.Job) (*string, error) {
	if job == nil {
		return nil, status.Error(codes.InvalidArgument, "received a `nil` job")
	}

	str := ""
	switch *job {
	case jobs.Job_JOB_ALCHEMIST:
		str = "alchemist"
	case jobs.Job_JOB_CARVER:
		str = "carver"
	case jobs.Job_JOB_FARMER:
		str = "farmer"
	case jobs.Job_JOB_FISHERMAN:
		str = "fisherman"
	case jobs.Job_JOB_HANDYMAN:
		str = "handyman"
	case jobs.Job_JOB_HUNTER:
		str = "hunter"
	case jobs.Job_JOB_JEWELLER:
		str = "jeweller"
	case jobs.Job_JOB_LUMBERJACK:
		str = "lumberjack"
	case jobs.Job_JOB_MINER:
		str = "miner"
	case jobs.Job_JOB_SHIELDMAKER:
		str = "shieldmaker"
	case jobs.Job_JOB_SHOEMAKER:
		str = "shoemaker"
	case jobs.Job_JOB_SMITH:
		str = "smith"
	case jobs.Job_JOB_TAILOR:
		str = "tailor"
	case jobs.Job_JOB_UNSPECIFIED:
		fallthrough
	default:
		str = "unspecified"
	}

	return &str, nil
}
