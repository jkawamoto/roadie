package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// JobTerminateParameter Options when terminating a job.
// swagger:model JobTerminateParameter
type JobTerminateParameter struct {

	// The text you want to appear as the job's TerminateReason. The default is 'UserTerminate'.
	TerminateReason string `json:"terminateReason,omitempty"`
}

// Validate validates this job terminate parameter
func (m *JobTerminateParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}