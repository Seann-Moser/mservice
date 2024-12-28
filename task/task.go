package task

import "context"

type Task interface {
	GetName() string
	Run(ctx context.Context, job *Job) (interface{}, error)
}
