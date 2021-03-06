package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TaskStatistics Resource usage statistics for a task.
// swagger:model TaskStatistics
type TaskStatistics struct {

	// The total kernel mode CPU time (summed across all cores and all compute nodes) consumed by the task.
	// Required: true
	KernelCPUTime *strfmt.Duration `json:"kernelCPUTime"`

	// The time at which the statistics were last updated. All statistics are limited to the range between startTime and lastUpdateTime.
	// Required: true
	LastUpdateTime *strfmt.DateTime `json:"lastUpdateTime"`

	// The total gibibytes read from disk by the task.
	// Required: true
	ReadIOGiB *float64 `json:"readIOGiB"`

	// The total number of disk read operations made by the task.
	// Required: true
	ReadIOps *int64 `json:"readIOps"`

	// The start time of the time range covered by the statistics.
	// Required: true
	StartTime *strfmt.DateTime `json:"startTime"`

	// The URL of the statistics.
	// Required: true
	URL *string `json:"url"`

	// The total user mode CPU time (summed across all cores and all compute nodes) consumed by the task.
	// Required: true
	UserCPUTime *strfmt.Duration `json:"userCPUTime"`

	// The total wait time of the task. The wait time for a task is defined as the elapsed time between the creation of the task and the start of task execution. (If the task is retried due to failures, the wait time is the time to the most recent task execution.)
	// Required: true
	WaitTime *strfmt.Duration `json:"waitTime"`

	// The total wall clock time of the task.
	//
	// The wall clock time is the elapsed time from when the task started running on a compute node to when it finished (or to the last time the statistics were updated, if the task had not finished by then). If the task was retried, this includes the wall clock time of all the task retries.
	// Required: true
	WallClockTime *strfmt.Duration `json:"wallClockTime"`

	// The total gibibytes written to disk by the task.
	// Required: true
	WriteIOGiB *float64 `json:"writeIOGiB"`

	// The total number of disk write operations made by the task.
	// Required: true
	WriteIOps *int64 `json:"writeIOps"`
}

// Validate validates this task statistics
func (m *TaskStatistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKernelCPUTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastUpdateTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReadIOGiB(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReadIOps(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserCPUTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWaitTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWallClockTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWriteIOGiB(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWriteIOps(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskStatistics) validateKernelCPUTime(formats strfmt.Registry) error {

	if err := validate.Required("kernelCPUTime", "body", m.KernelCPUTime); err != nil {
		return err
	}

	return nil
}

func (m *TaskStatistics) validateLastUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdateTime", "body", m.LastUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *TaskStatistics) validateReadIOGiB(formats strfmt.Registry) error {

	if err := validate.Required("readIOGiB", "body", m.ReadIOGiB); err != nil {
		return err
	}

	return nil
}

func (m *TaskStatistics) validateReadIOps(formats strfmt.Registry) error {

	if err := validate.Required("readIOps", "body", m.ReadIOps); err != nil {
		return err
	}

	return nil
}

func (m *TaskStatistics) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *TaskStatistics) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *TaskStatistics) validateUserCPUTime(formats strfmt.Registry) error {

	if err := validate.Required("userCPUTime", "body", m.UserCPUTime); err != nil {
		return err
	}

	return nil
}

func (m *TaskStatistics) validateWaitTime(formats strfmt.Registry) error {

	if err := validate.Required("waitTime", "body", m.WaitTime); err != nil {
		return err
	}

	return nil
}

func (m *TaskStatistics) validateWallClockTime(formats strfmt.Registry) error {

	if err := validate.Required("wallClockTime", "body", m.WallClockTime); err != nil {
		return err
	}

	return nil
}

func (m *TaskStatistics) validateWriteIOGiB(formats strfmt.Registry) error {

	if err := validate.Required("writeIOGiB", "body", m.WriteIOGiB); err != nil {
		return err
	}

	return nil
}

func (m *TaskStatistics) validateWriteIOps(formats strfmt.Registry) error {

	if err := validate.Required("writeIOps", "body", m.WriteIOps); err != nil {
		return err
	}

	return nil
}
