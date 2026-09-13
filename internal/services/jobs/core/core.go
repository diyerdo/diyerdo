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
	"sort"

	"github.com/diyerdo/diyerdo/internal/shared/utils"
	"github.com/diyerdo/proto/gen/go/proto/jobs/v1"
	"github.com/dofusdude/dodugo"
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

// GetJobsRequirementsForItem returns all the jobs and maximum required levels involved in crafting the given item
func (t *JobsCore) GetJobsRequirementsForItem(itemId int32) ([]*jobs.JobRequirements, error) {
	if itemId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "itemId cannot be `%d`", itemId)
	}

	dodugoWrapper, err := utils.NewDodugoWrapper()
	if err != nil {
		return nil, err
	}

	// First find the item as an Equipment; if not found, find as a Resource
	item, err := resolveItem(dodugoWrapper, itemId, "")
	if err != nil {
		return nil, err
	}

	// If the item exists but has no craft recipe, return empty requirements
	if len(item.recipe) == 0 {
		return []*jobs.JobRequirements{}, nil
	}

	maxLevelByJob := make(map[jobs.Job]int32)
	visited := make(map[int32]bool)
	visited[itemId] = true

	// Job(s) for the target item itself
	targetJobs, err := t.GetJobForItem(itemId)
	if err != nil {
		return nil, err
	}
	for _, job := range targetJobs {
		if item.level > maxLevelByJob[job] {
			maxLevelByJob[job] = item.level
		}
	}

	// Recursively traverse ingredients
	if err := t.traverseRecipeIngredients(item.recipe, visited, maxLevelByJob, dodugoWrapper); err != nil {
		return nil, err
	}

	// Deterministic sorting by Job enum value
	result := make([]*jobs.JobRequirements, 0, len(maxLevelByJob))
	for job, level := range maxLevelByJob {
		result = append(result, &jobs.JobRequirements{
			Job:   job,
			Level: level,
		})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Job < result[j].Job
	})

	return result, nil
}

// resolvedItem holds the level and recipe of a resolved item
type resolvedItem struct {
	level  int32
	recipe []dodugo.Recipe
}

// resolveItem attempts to find an item by ID, trying Equipment first and Resource second
// (or using the subtype hint when available)
func resolveItem(dodugoWrapper *utils.DodugoWrapper, itemId int32, subtypeHint string) (*resolvedItem, error) {
	if subtypeHint == "resources" {
		res, err := dodugoWrapper.GetResource(itemId)
		if err == nil && res != nil {
			return &resolvedItem{
				level:  res.GetLevel(),
				recipe: res.GetRecipe(),
			}, nil
		}
		if status.Code(err) != codes.NotFound {
			return nil, err
		}

		equip, err := dodugoWrapper.GetEquipment(itemId)
		if err == nil && equip != nil {
			return &resolvedItem{
				level:  equip.GetLevel(),
				recipe: equip.GetRecipe(),
			}, nil
		}
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "no equipment or resource found with id '%d'", itemId)
		}
		return nil, err
	}

	// Default: Equipment first, Resource fallback
	equip, err := dodugoWrapper.GetEquipment(itemId)
	if err == nil && equip != nil {
		return &resolvedItem{
			level:  equip.GetLevel(),
			recipe: equip.GetRecipe(),
		}, nil
	}
	if status.Code(err) != codes.NotFound {
		return nil, err
	}

	res, err := dodugoWrapper.GetResource(itemId)
	if err == nil && res != nil {
		return &resolvedItem{
			level:  res.GetLevel(),
			recipe: res.GetRecipe(),
		}, nil
	}
	if status.Code(err) == codes.NotFound {
		return nil, status.Errorf(codes.NotFound, "no equipment or resource found with id '%d'", itemId)
	}
	return nil, err
}

// traverseRecipeIngredients recursively inspects ingredients and collects jobs and levels
func (t *JobsCore) traverseRecipeIngredients(
	ingredients []dodugo.Recipe,
	visited map[int32]bool,
	maxLevelByJob map[jobs.Job]int32,
	dodugoWrapper *utils.DodugoWrapper,
) error {
	for _, ing := range ingredients {
		subId := ing.GetItemAnkamaId()
		if subId < 1 || visited[subId] {
			continue
		}
		visited[subId] = true

		subItem, err := resolveItem(dodugoWrapper, subId, ing.GetItemSubtype())
		if err != nil {
			if status.Code(err) == codes.NotFound {
				continue
			}
			return err
		}

		// If the ingredient has a recipe, it is craftable
		if len(subItem.recipe) > 0 {
			subJobs, err := t.GetJobForItem(subId)
			if err != nil {
				return err
			}
			for _, job := range subJobs {
				if subItem.level > maxLevelByJob[job] {
					maxLevelByJob[job] = subItem.level
				}
			}

			if err := t.traverseRecipeIngredients(subItem.recipe, visited, maxLevelByJob, dodugoWrapper); err != nil {
				return err
			}
		}
	}

	return nil
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
