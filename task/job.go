package task

import (
	"encoding/json"
	"fmt"
)

type Job struct {
	ID          string `json:"id"`
	OwnerID     string
	SecondaryID string
	ScheduledID string

	Service string `json:"service"`
	Name    string `json:"name"`

	Status           string
	Message          string
	Data             string
	DataType         string // to help with job data conversion
	TimePartition    string
	RemainingRetries int

	Tasks string // an array of functions to call

	CreatedTimestamp string
	UpdatedTimestamp string
}

func GetJobData[T any](job *Job) (*T, error) {
	var data T
	if len(job.Data) == 0 {
		return nil, fmt.Errorf("empty job data")
	}
	err := json.Unmarshal([]byte(job.Data), &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal job:%s owner_id:%s data: %w", job.Name, job.OwnerID, err)
	}
	return &data, nil
}
