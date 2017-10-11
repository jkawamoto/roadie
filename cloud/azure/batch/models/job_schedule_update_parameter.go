package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JobScheduleUpdateParameter The set of changes to be made to a job schedule.
// swagger:model JobScheduleUpdateParameter
type JobScheduleUpdateParameter struct {

	// Details of the jobs to be created on this schedule.
	//
	// Updates affect only jobs that are started after the update has taken place. Any currently active job continues with the older specification.
	// Required: true
	JobSpecification *JobSpecification `json:"jobSpecification"`

	// A list of name-value pairs associated with the job schedule as metadata.
	//
	// If you do not specify this element, it takes the default value of an empty list; in effect, any existing metadata is deleted.
	Metadata []*MetadataItem `json:"metadata"`

	// The schedule according to which jobs will be created.
	//
	// If you do not specify this element, it is equivalent to passing the default schedule: that is, a single job scheduled to run immediately.
	// Required: true
	Schedule *Schedule `json:"schedule"`
}

// Validate validates this job schedule update parameter
func (m *JobScheduleUpdateParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobSpecification(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobScheduleUpdateParameter) validateJobSpecification(formats strfmt.Registry) error {

	if err := validate.Required("jobSpecification", "body", m.JobSpecification); err != nil {
		return err
	}

	if m.JobSpecification != nil {

		if err := m.JobSpecification.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *JobScheduleUpdateParameter) validateMetadata(formats strfmt.Registry) error {

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

func (m *JobScheduleUpdateParameter) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule); err != nil {
		return err
	}

	if m.Schedule != nil {

		if err := m.Schedule.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}