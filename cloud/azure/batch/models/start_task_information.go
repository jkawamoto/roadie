package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// StartTaskInformation Information about a start task running on a compute node.
// swagger:model StartTaskInformation
type StartTaskInformation struct {

	// The time at which the start task stopped running.
	//
	// This is the end time of the most recent run of the start task, if that run has completed (even if that run failed and a retry is pending). This element is not present if the start task is currently running.
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// The exit code of the program specified on the start task command line.
	//
	// This property is set only if the start task is in the completed state. In general, the exit code for a process reflects the specific convention implemented by the application developer for that process. If you use the exit code value to make decisions in your code, be sure that you know the exit code convention used by the application process. However, if the Batch service terminates the start task (due to timeout, or user termination via the API) you may see an operating system-defined exit code.
	ExitCode int32 `json:"exitCode,omitempty"`

	// The most recent time at which a retry of the task started running.
	//
	// This element is present only if the task was retried (i.e. retryCount is nonzero). If present, this is typically the same as startTime, but may be different if the task has been restarted for reasons other than retry; for example, if the compute node was rebooted during a retry, then the startTime is updated but the lastRetryTime is not.
	LastRetryTime strfmt.DateTime `json:"lastRetryTime,omitempty"`

	// The number of times the task has been retried by the Batch service.
	//
	// The task is retried if it exits with a nonzero exit code, up to the specified MaxTaskRetryCount.
	// Required: true
	RetryCount *int32 `json:"retryCount"`

	// Any error encountered scheduling the start task.
	SchedulingError *TaskSchedulingError `json:"schedulingError,omitempty"`

	// The time at which the start task started running.
	//
	// This value is reset every time the task is restarted or retried (that is, this is the most recent time at which the start task started running).
	// Required: true
	StartTime *strfmt.DateTime `json:"startTime"`

	// The state of the start task on the compute node.
	//
	// Possible values are: running – The start task is currently running. completed – The start task has exited with exit code 0, or the start task has failed and the retry limit has reached, or the start task process did not run due to scheduling errors.
	// Required: true
	State *string `json:"state"`
}

// Validate validates this start task information
func (m *StartTaskInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRetryCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSchedulingError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StartTaskInformation) validateRetryCount(formats strfmt.Registry) error {

	if err := validate.Required("retryCount", "body", m.RetryCount); err != nil {
		return err
	}

	return nil
}

func (m *StartTaskInformation) validateSchedulingError(formats strfmt.Registry) error {

	if swag.IsZero(m.SchedulingError) { // not required
		return nil
	}

	if m.SchedulingError != nil {

		if err := m.SchedulingError.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *StartTaskInformation) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

var startTaskInformationTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["running","completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		startTaskInformationTypeStatePropEnum = append(startTaskInformationTypeStatePropEnum, v)
	}
}

const (
	// StartTaskInformationStateRunning captures enum value "running"
	StartTaskInformationStateRunning string = "running"
	// StartTaskInformationStateCompleted captures enum value "completed"
	StartTaskInformationStateCompleted string = "completed"
)

// prop value enum
func (m *StartTaskInformation) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, startTaskInformationTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StartTaskInformation) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}
