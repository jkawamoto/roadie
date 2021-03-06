package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JobManagerTask Specifies details of a Job Manager task.
// swagger:model JobManagerTask
type JobManagerTask struct {

	// A list of application packages that the Batch service will deploy to the compute node before running the command line.
	//
	// Application packages are downloaded and deployed to a shared directory, not the task directory. Therefore, if a referenced package is already on the compute node, and is up to date, then it is not re-downloaded; the existing copy on the compute node is used. If a referenced application package cannot be installed, for example because the package has been deleted or because download failed, the task fails with a scheduling error. This property is currently not supported on jobs running on pools created using the virtualMachineConfiguration (IaaS) property. If a task specifying applicationPackageReferences runs on such a pool, it fails with a scheduling error with code TaskSchedulingConstraintFailed.
	ApplicationPackageReferences []*ApplicationPackageReference `json:"applicationPackageReferences"`

	// The command line of the Job Manager task.
	//
	// The command line does not run under a shell, and therefore cannot take advantage of shell features such as environment variable expansion. If you want to take advantage of such features, you should invoke the shell in the command line, for example using "cmd /c MyCommand" in Windows or "/bin/sh -c MyCommand" in Linux.
	// Required: true
	CommandLine *string `json:"commandLine"`

	// Constraints that apply to the Job Manager task.
	Constraints *TaskConstraints `json:"constraints,omitempty"`

	// The display name of the Job Manager task.
	//
	// It need not be unique and can contain any Unicode characters up to a maximum length of 1024.
	DisplayName string `json:"displayName,omitempty"`

	// A list of environment variable settings for the Job Manager task.
	EnvironmentSettings []*EnvironmentSetting `json:"environmentSettings"`

	// A string that uniquely identifies the Job Manager taskwithin the job.
	//
	// The id can contain any combination of alphanumeric characters including hyphens and underscores and cannot contain more than 64 characters.
	// Required: true
	ID *string `json:"id"`

	// Whether completion of the Job Manager task signifies completion of the entire job.
	//
	// If true, when the Job Manager task completes, the Batch service marks the job as complete. If any tasks are still running at this time (other than Job Release), those tasks are terminated. If false, the completion of the Job Manager task does not affect the job status. In this case, you should either use the onAllTasksComplete attribute to terminate the job, or have a client or user terminate the job explicitly. An example of this is if the Job Manager creates a set of tasks but then takes no further role in their execution. The default value is true. If you are using the onAllTasksComplete and onTaskFailure attributes to control job lifetime, and using the job manager task only to create the tasks for the job (not to monitor progress), then it is important to set killJobOnCompletion to false.
	KillJobOnCompletion bool `json:"killJobOnCompletion,omitempty"`

	// A list of files that the Batch service will download to the compute node before running the command line.
	//
	// Files listed under this element are located in the task's working directory.
	ResourceFiles []*ResourceFile `json:"resourceFiles"`

	// Whether to run the Job Manager task in elevated mode. The default value is false.
	RunElevated bool `json:"runElevated,omitempty"`

	// Whether the Job Manager task requires exclusive use of the compute node where it runs.
	//
	// If true, no other tasks will run on the same compute node for as long as the Job Manager is running. If false, other tasks can run simultaneously with the Job Manager on a compute node. The Job Manager task counts normally against the node's concurrent task limit, so this is only relevant if the node allows multiple concurrent tasks. The default value is true.
	RunExclusive bool `json:"runExclusive,omitempty"`
}

// Validate validates this job manager task
func (m *JobManagerTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationPackageReferences(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCommandLine(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConstraints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnvironmentSettings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResourceFiles(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobManagerTask) validateApplicationPackageReferences(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplicationPackageReferences) { // not required
		return nil
	}

	for i := 0; i < len(m.ApplicationPackageReferences); i++ {

		if swag.IsZero(m.ApplicationPackageReferences[i]) { // not required
			continue
		}

		if m.ApplicationPackageReferences[i] != nil {

			if err := m.ApplicationPackageReferences[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *JobManagerTask) validateCommandLine(formats strfmt.Registry) error {

	if err := validate.Required("commandLine", "body", m.CommandLine); err != nil {
		return err
	}

	return nil
}

func (m *JobManagerTask) validateConstraints(formats strfmt.Registry) error {

	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	if m.Constraints != nil {

		if err := m.Constraints.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *JobManagerTask) validateEnvironmentSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.EnvironmentSettings); i++ {

		if swag.IsZero(m.EnvironmentSettings[i]) { // not required
			continue
		}

		if m.EnvironmentSettings[i] != nil {

			if err := m.EnvironmentSettings[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *JobManagerTask) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *JobManagerTask) validateResourceFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceFiles); i++ {

		if swag.IsZero(m.ResourceFiles[i]) { // not required
			continue
		}

		if m.ResourceFiles[i] != nil {

			if err := m.ResourceFiles[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
