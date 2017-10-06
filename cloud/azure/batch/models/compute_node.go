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

// ComputeNode A compute node in the Batch service.
// swagger:model ComputeNode
type ComputeNode struct {

	// An identifier which can be passed when adding a task to request that the task be scheduled close to this compute node.
	AffinityID string `json:"affinityId,omitempty"`

	// The time at which this compute node was allocated to the pool.
	AllocationTime strfmt.DateTime `json:"allocationTime,omitempty"`

	// The list of certificates installed on the compute node.
	//
	// For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of remoteuser, a certs directory is created in the user's home directory (e.g., /home/<user-name>/certs) where certificates are placed.
	CertificateReferences []*CertificateReference `json:"certificateReferences"`

	// The list of errors that are currently being encountered by the compute node.
	Errors []*ComputeNodeError `json:"errors"`

	// The ID of the compute node.
	//
	// Every node that is added to a pool is assigned a unique ID. Whenever a node is removed from a pool, all of its local files are deleted, and the ID is reclaimed and could be reused for new nodes.
	ID string `json:"id,omitempty"`

	// The IP address that other compute nodes can use to communicate with this compute node.
	//
	// Every node that is added to a pool is assigned a unique IP address. Whenever a node is removed from a pool, all of its local files are deleted, and the IP address is reclaimed and could be reused for new nodes.
	IPAddress string `json:"ipAddress,omitempty"`

	// The time at which the compute node was started.
	//
	// This property may not be present if the node state is unusable.
	LastBootTime strfmt.DateTime `json:"lastBootTime,omitempty"`

	// The list of tasks that are currently running on the compute node.
	RecentTasks []*TaskInformation `json:"recentTasks"`

	// The total number of currently running job tasks on the compute node. This includes Job Preparation, Job Release, and Job Manager tasks, but not the pool start task.
	RunningTasksCount int32 `json:"runningTasksCount,omitempty"`

	// Whether the compute node is available for task scheduling.
	//
	// Possible values are: enabled – Tasks can be scheduled on the node. disabled – No new tasks will be scheduled on the node. Tasks already running on the node may still run to completion. All nodes start with scheduling enabled.
	SchedulingState string `json:"schedulingState,omitempty"`

	// The task specified to run on the compute node as it joins the pool.
	StartTask *StartTask `json:"startTask,omitempty"`

	// Runtime information about the execution of the start task on the compute node.
	StartTaskInfo *StartTaskInformation `json:"startTaskInfo,omitempty"`

	// The current state of the compute node.
	State string `json:"state,omitempty"`

	// The time at which the compute node entered its current state.
	StateTransitionTime strfmt.DateTime `json:"stateTransitionTime,omitempty"`

	// The total number of job tasks completed on the compute node. This includes Job Preparation, Job Release and Job Manager tasks, but not the pool start task.
	TotalTasksRun int32 `json:"totalTasksRun,omitempty"`

	// The total number of job tasks which completed successfully (with exitCode 0) on the compute node. This includes Job Preparation, Job Release, and Job Manager tasks, but not the pool start task.
	TotalTasksSucceeded int32 `json:"totalTasksSucceeded,omitempty"`

	// The URL of the compute node.
	URL string `json:"url,omitempty"`

	// The size of the virtual machine hosting the compute node.
	//
	// For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (http://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
	VMSize string `json:"vmSize,omitempty"`
}

// Validate validates this compute node
func (m *ComputeNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateReferences(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecentTasks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSchedulingState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTask(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTaskInfo(formats); err != nil {
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

func (m *ComputeNode) validateCertificateReferences(formats strfmt.Registry) error {

	if swag.IsZero(m.CertificateReferences) { // not required
		return nil
	}

	for i := 0; i < len(m.CertificateReferences); i++ {

		if swag.IsZero(m.CertificateReferences[i]) { // not required
			continue
		}

		if m.CertificateReferences[i] != nil {

			if err := m.CertificateReferences[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ComputeNode) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {

		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {

			if err := m.Errors[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ComputeNode) validateRecentTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.RecentTasks) { // not required
		return nil
	}

	for i := 0; i < len(m.RecentTasks); i++ {

		if swag.IsZero(m.RecentTasks[i]) { // not required
			continue
		}

		if m.RecentTasks[i] != nil {

			if err := m.RecentTasks[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var computeNodeTypeSchedulingStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computeNodeTypeSchedulingStatePropEnum = append(computeNodeTypeSchedulingStatePropEnum, v)
	}
}

const (
	// ComputeNodeSchedulingStateEnabled captures enum value "enabled"
	ComputeNodeSchedulingStateEnabled string = "enabled"
	// ComputeNodeSchedulingStateDisabled captures enum value "disabled"
	ComputeNodeSchedulingStateDisabled string = "disabled"
)

// prop value enum
func (m *ComputeNode) validateSchedulingStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, computeNodeTypeSchedulingStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ComputeNode) validateSchedulingState(formats strfmt.Registry) error {

	if swag.IsZero(m.SchedulingState) { // not required
		return nil
	}

	// value enum
	if err := m.validateSchedulingStateEnum("schedulingState", "body", m.SchedulingState); err != nil {
		return err
	}

	return nil
}

func (m *ComputeNode) validateStartTask(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTask) { // not required
		return nil
	}

	if m.StartTask != nil {

		if err := m.StartTask.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputeNode) validateStartTaskInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTaskInfo) { // not required
		return nil
	}

	if m.StartTaskInfo != nil {

		if err := m.StartTaskInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var computeNodeTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["idle","rebooting","reimaging","running","unusable","creating","starting","waitingforstarttask","starttaskfailed","unknown","leavingpool","offline"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computeNodeTypeStatePropEnum = append(computeNodeTypeStatePropEnum, v)
	}
}

const (
	// ComputeNodeStateIDLE captures enum value "idle"
	ComputeNodeStateIDLE string = "idle"
	// ComputeNodeStateRebooting captures enum value "rebooting"
	ComputeNodeStateRebooting string = "rebooting"
	// ComputeNodeStateReimaging captures enum value "reimaging"
	ComputeNodeStateReimaging string = "reimaging"
	// ComputeNodeStateRunning captures enum value "running"
	ComputeNodeStateRunning string = "running"
	// ComputeNodeStateUnusable captures enum value "unusable"
	ComputeNodeStateUnusable string = "unusable"
	// ComputeNodeStateCreating captures enum value "creating"
	ComputeNodeStateCreating string = "creating"
	// ComputeNodeStateStarting captures enum value "starting"
	ComputeNodeStateStarting string = "starting"
	// ComputeNodeStateWaitingforstarttask captures enum value "waitingforstarttask"
	ComputeNodeStateWaitingforstarttask string = "waitingforstarttask"
	// ComputeNodeStateStarttaskfailed captures enum value "starttaskfailed"
	ComputeNodeStateStarttaskfailed string = "starttaskfailed"
	// ComputeNodeStateUnknown captures enum value "unknown"
	ComputeNodeStateUnknown string = "unknown"
	// ComputeNodeStateLeavingpool captures enum value "leavingpool"
	ComputeNodeStateLeavingpool string = "leavingpool"
	// ComputeNodeStateOffline captures enum value "offline"
	ComputeNodeStateOffline string = "offline"
)

// prop value enum
func (m *ComputeNode) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, computeNodeTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ComputeNode) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}
