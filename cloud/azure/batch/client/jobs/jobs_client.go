package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for jobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
JobAdd adds a job to the specified account

The Batch service supports two ways to control the work done as part of a job. In the first approach, the user specifies a Job Manager task. The Batch service launches this task when it is ready to start the job. The Job Manager task controls all other tasks that run under this job, by using the Task APIs. In the second approach, the user directly controls the execution of tasks under an active job, by using the Task APIs. Also note: when naming jobs, avoid including sensitive information such as user names or secret project names. This information may appear in telemetry logs accessible to Microsoft Support engineers.
*/
func (a *Client) JobAdd(params *JobAddParams) (*JobAddCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_Add",
		Method:             "POST",
		PathPattern:        "/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobAddReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobAddCreated), nil

}

/*
JobDelete deletes a job

Deleting a job also deletes all tasks that are part of that job, and all job statistics. This also overrides the retention period for task data; that is, if the job contains tasks which are still retained on compute nodes, the Batch services deletes those tasks' working directories and all their contents.
*/
func (a *Client) JobDelete(params *JobDeleteParams) (*JobDeleteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_Delete",
		Method:             "DELETE",
		PathPattern:        "/jobs/{jobId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobDeleteAccepted), nil

}

/*
JobDisable disables the specified job preventing new tasks from running

The Batch Service immediately moves the job to the disabling state. Batch then uses the disableTasks parameter to determine what to do with the currently running tasks of the job. The job remains in the disabling state until the disable operation is completed and all tasks have been dealt with according to the disableTasks option; the job then moves to the disabled state. No new tasks are started under the job until it moves back to active state. If you try to disable a job that is in any state other than active, disabling, or disabled, the request fails with status code 409.
*/
func (a *Client) JobDisable(params *JobDisableParams) (*JobDisableAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobDisableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_Disable",
		Method:             "POST",
		PathPattern:        "/jobs/{jobId}/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobDisableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobDisableAccepted), nil

}

/*
JobEnable enables the specified job allowing new tasks to run

When you call this API, the Batch service sets a disabled job to the enabling state. After the this operation is completed, the job moves to the active state, and scheduling of new tasks under the job resumes. The Batch service does not allow a task to remain in the active state for more than 7 days. Therefore, if you enable a job containing active tasks which were added more than 7 days ago, those tasks will not run.
*/
func (a *Client) JobEnable(params *JobEnableParams) (*JobEnableAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobEnableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_Enable",
		Method:             "POST",
		PathPattern:        "/jobs/{jobId}/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobEnableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobEnableAccepted), nil

}

/*
JobGet gets information about the specified job
*/
func (a *Client) JobGet(params *JobGetParams) (*JobGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_Get",
		Method:             "GET",
		PathPattern:        "/jobs/{jobId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobGetOK), nil

}

/*
JobGetAllJobsLifetimeStatistics gets lifetime summary statistics for all of the jobs in the specified account

Statistics are aggregated across all jobs that have ever existed in the account, from account creation to the last update time of the statistics.
*/
func (a *Client) JobGetAllJobsLifetimeStatistics(params *JobGetAllJobsLifetimeStatisticsParams) (*JobGetAllJobsLifetimeStatisticsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobGetAllJobsLifetimeStatisticsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_GetAllJobsLifetimeStatistics",
		Method:             "GET",
		PathPattern:        "/lifetimejobstats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobGetAllJobsLifetimeStatisticsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobGetAllJobsLifetimeStatisticsOK), nil

}

/*
JobList lists all of the jobs in the specified account
*/
func (a *Client) JobList(params *JobListParams) (*JobListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_List",
		Method:             "GET",
		PathPattern:        "/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobListOK), nil

}

/*
JobListFromJobSchedule lists the jobs that have been created under the specified job schedule
*/
func (a *Client) JobListFromJobSchedule(params *JobListFromJobScheduleParams) (*JobListFromJobScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobListFromJobScheduleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_ListFromJobSchedule",
		Method:             "GET",
		PathPattern:        "/jobschedules/{jobScheduleId}/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobListFromJobScheduleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobListFromJobScheduleOK), nil

}

/*
JobListPreparationAndReleaseTaskStatus lists the execution status of the job preparation and job release task for the specified job across the compute nodes where the job has run

This API returns the Job Preparation and Job Release task status on all compute nodes that have run the Job Preparation or Job Release task. This includes nodes which have since been removed from the pool.
*/
func (a *Client) JobListPreparationAndReleaseTaskStatus(params *JobListPreparationAndReleaseTaskStatusParams) (*JobListPreparationAndReleaseTaskStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobListPreparationAndReleaseTaskStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_ListPreparationAndReleaseTaskStatus",
		Method:             "GET",
		PathPattern:        "/jobs/{jobId}/jobpreparationandreleasetaskstatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobListPreparationAndReleaseTaskStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobListPreparationAndReleaseTaskStatusOK), nil

}

/*
JobPatch updates the properties of the specified job

This replaces only the job properties specified in the request. For example, if the job has constraints, and a request does not specify the constraints element, then the job keeps the existing constraints.
*/
func (a *Client) JobPatch(params *JobPatchParams) (*JobPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobPatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_Patch",
		Method:             "PATCH",
		PathPattern:        "/jobs/{jobId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobPatchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobPatchOK), nil

}

/*
JobTerminate terminates the specified job marking it as completed

When a Terminate Job request is received, the Batch service sets the job to the terminating state. The Batch service then terminates any active or running tasks associated with the job, and runs any required Job Release tasks. The job then moves into the completed state.
*/
func (a *Client) JobTerminate(params *JobTerminateParams) (*JobTerminateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobTerminateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_Terminate",
		Method:             "POST",
		PathPattern:        "/jobs/{jobId}/terminate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobTerminateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobTerminateAccepted), nil

}

/*
JobUpdate updates the properties of the specified job

This fully replaces all the updateable properties of the job. For example, if the job has constraints associated with it and if constraints is not specified with this request, then the Batch service will remove the existing constraints.
*/
func (a *Client) JobUpdate(params *JobUpdateParams) (*JobUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Job_Update",
		Method:             "PUT",
		PathPattern:        "/jobs/{jobId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JobUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
