package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JobExecutionInformation Contains information about the execution of a job in the Azure Batch service.
// swagger:model JobExecutionInformation
type JobExecutionInformation struct {

	// The completion time of the job.
	//
	// This property is set only if the job is in the completed state.
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// The ID of the pool to which this job is assigned.
	//
	// This element contains the actual pool where the job is assigned. When you get job details from the service, they also contain a poolInfo element, which contains the pool configuration data from when the job was added or updated. That poolInfo element may also contain a poolId element. If it does, the two IDs are the same. If it does not, it means the job ran on an auto pool, and this property contains the id of that auto pool.
	PoolID string `json:"poolId,omitempty"`

	// Details of any error encountered by the service in starting the job.
	//
	// This property is not set if there was no error starting the job.
	SchedulingError *JobSchedulingError `json:"schedulingError,omitempty"`

	// The start time of the job.
	//
	// This is the time at which the job was created.
	// Required: true
	StartTime *strfmt.DateTime `json:"startTime"`

	// A string describing the reason the job ended.
	//
	// This property is set only if the job is in the completed state. If the Batch service terminates the job, it sets the reason as follows: JMComplete – the Job Manager task completed, and killJobOnCompletion was set to true. MaxWallClockTimeExpiry – the job reached its maxWallClockTime constraint. TerminateJobSchedule – the job ran as part of a schedule, and the schedule terminated. AllTasksComplete – the job's onAllTasksComplete attribute is set to terminatejob, and all tasks in the job are complete. TaskFailed – the job's onTaskFailure attribute is set to performexitoptionsjobaction, and a task in the job failed with an exit condition that specified a jobAction of terminatejob. Any other string is a user-defined reason specified in a call to the 'Terminate a job' operation.
	TerminateReason string `json:"terminateReason,omitempty"`
}

// Validate validates this job execution information
func (m *JobExecutionInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedulingError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobExecutionInformation) validateSchedulingError(formats strfmt.Registry) error {

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

func (m *JobExecutionInformation) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}
