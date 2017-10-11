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

// CloudJobSchedule A job schedule that allows recurring jobs by specifying when to run jobs and a specification used to create each job.
// swagger:model CloudJobSchedule
type CloudJobSchedule struct {

	// The creation time of the job schedule.
	CreationTime strfmt.DateTime `json:"creationTime,omitempty"`

	// The display name for the schedule.
	DisplayName string `json:"displayName,omitempty"`

	// The ETag of the job schedule.
	//
	// This is an opaque string. You can use it to detect whether the job schedule has changed between requests. In particular, you can be pass the ETag with an Update Job Schedule request to specify that your changes should take effect only if nobody else has modified the schedule in the meantime.
	ETag string `json:"eTag,omitempty"`

	// Information about jobs that have been and will be run under this schedule.
	ExecutionInfo *JobScheduleExecutionInformation `json:"executionInfo,omitempty"`

	// A string that uniquely identifies the schedule within the account.
	//
	// It is common to use a GUID for the id.
	ID string `json:"id,omitempty"`

	// The details of the jobs to be created on this schedule.
	JobSpecification *JobSpecification `json:"jobSpecification,omitempty"`

	// The last modified time of the job schedule.
	//
	// This is the last time at which the schedule level data, such as the job specification or recurrence information, changed. It does not factor in job-level changes such as new jobs being created or jobs changing state.
	LastModified strfmt.DateTime `json:"lastModified,omitempty"`

	// A list of name-value pairs associated with the schedule as metadata.
	//
	// The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
	Metadata []*MetadataItem `json:"metadata"`

	// The previous state of the job schedule.
	//
	// This property is not present if the job schedule is in its initial active state.
	PreviousState string `json:"previousState,omitempty"`

	// The time at which the job schedule entered its previous state.
	//
	// This property is not present if the job schedule is in its initial active state.
	PreviousStateTransitionTime strfmt.DateTime `json:"previousStateTransitionTime,omitempty"`

	// The schedule according to which jobs will be created.
	Schedule *Schedule `json:"schedule,omitempty"`

	// The current state of the job schedule.
	State string `json:"state,omitempty"`

	// The time at which the job schedule entered the current state.
	StateTransitionTime strfmt.DateTime `json:"stateTransitionTime,omitempty"`

	// The lifetime resource usage statistics for the job schedule.
	Stats *JobScheduleStatistics `json:"stats,omitempty"`

	// The URL of the job schedule.
	URL string `json:"url,omitempty"`
}

// Validate validates this cloud job schedule
func (m *CloudJobSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutionInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateJobSpecification(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePreviousState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudJobSchedule) validateExecutionInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutionInfo) { // not required
		return nil
	}

	if m.ExecutionInfo != nil {

		if err := m.ExecutionInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CloudJobSchedule) validateJobSpecification(formats strfmt.Registry) error {

	if swag.IsZero(m.JobSpecification) { // not required
		return nil
	}

	if m.JobSpecification != nil {

		if err := m.JobSpecification.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CloudJobSchedule) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	for i := 0; i < len(m.Metadata); i++ {

		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {

			if err := m.Metadata[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var cloudJobScheduleTypePreviousStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","completed","disabled","terminating","deleting"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudJobScheduleTypePreviousStatePropEnum = append(cloudJobScheduleTypePreviousStatePropEnum, v)
	}
}

const (
	// CloudJobSchedulePreviousStateActive captures enum value "active"
	CloudJobSchedulePreviousStateActive string = "active"
	// CloudJobSchedulePreviousStateCompleted captures enum value "completed"
	CloudJobSchedulePreviousStateCompleted string = "completed"
	// CloudJobSchedulePreviousStateDisabled captures enum value "disabled"
	CloudJobSchedulePreviousStateDisabled string = "disabled"
	// CloudJobSchedulePreviousStateTerminating captures enum value "terminating"
	CloudJobSchedulePreviousStateTerminating string = "terminating"
	// CloudJobSchedulePreviousStateDeleting captures enum value "deleting"
	CloudJobSchedulePreviousStateDeleting string = "deleting"
)

// prop value enum
func (m *CloudJobSchedule) validatePreviousStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudJobScheduleTypePreviousStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudJobSchedule) validatePreviousState(formats strfmt.Registry) error {

	if swag.IsZero(m.PreviousState) { // not required
		return nil
	}

	// value enum
	if err := m.validatePreviousStateEnum("previousState", "body", m.PreviousState); err != nil {
		return err
	}

	return nil
}

func (m *CloudJobSchedule) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {

		if err := m.Schedule.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var cloudJobScheduleTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","completed","disabled","terminating","deleting"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudJobScheduleTypeStatePropEnum = append(cloudJobScheduleTypeStatePropEnum, v)
	}
}

const (
	// CloudJobScheduleStateActive captures enum value "active"
	CloudJobScheduleStateActive string = "active"
	// CloudJobScheduleStateCompleted captures enum value "completed"
	CloudJobScheduleStateCompleted string = "completed"
	// CloudJobScheduleStateDisabled captures enum value "disabled"
	CloudJobScheduleStateDisabled string = "disabled"
	// CloudJobScheduleStateTerminating captures enum value "terminating"
	CloudJobScheduleStateTerminating string = "terminating"
	// CloudJobScheduleStateDeleting captures enum value "deleting"
	CloudJobScheduleStateDeleting string = "deleting"
)

// prop value enum
func (m *CloudJobSchedule) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudJobScheduleTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudJobSchedule) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *CloudJobSchedule) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {

		if err := m.Stats.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}