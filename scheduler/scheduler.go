package scheduler

import (
	"context"
	"github.com/Seann-Moser/mservice/task"
)

type Scheduler interface {
	Start(ctx context.Context) error

	SaveJob(ctx context.Context, job *task.Job) error
	GetJob(ctx context.Context, id string) (*task.Job, error)
	DeleteJob(ctx context.Context, id string) error
	GetJobs(ctx context.Context, service string, status string) ([]*task.Job, error)
	CanRunJob(ctx context.Context, job *task.Job) bool

	GetScheduledJobs(ctx context.Context, service string) ([]*task.Job, error)
	SaveScheduledJob(ctx context.Context, scheduledJob *ScheduledJob) (*ScheduledJob, error)
	DeleteScheduledJob(ctx context.Context, scheduledJob *ScheduledJob) error
	GetScheduledJobsForOwner(ctx context.Context, owner string) ([]*ScheduledJob, error)
	CanScheduleJob(ctx context.Context, scheduledJob *ScheduledJob) bool
}

type ScheduledJob struct {
	//todo add fields
	// todo implement k8s job
	// todo support running custom job functions like gpt tools

}
