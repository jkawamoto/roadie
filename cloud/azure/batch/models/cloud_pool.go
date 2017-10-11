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

// CloudPool A pool in the Azure Batch service.
// swagger:model CloudPool
type CloudPool struct {

	// Whether the pool is resizing.
	//
	// Possible values are: steady – The pool is not resizing. There are no changes to the number of nodes in the pool in progress. A pool enters this state when it is created and when no operations are being performed on the pool to change the number of dedicated nodes. resizing - The pool is resizing; that is, compute nodes are being added to or removed from the pool. stopping - The pool was resizing, but the user has requested that the resize be stopped, but the stop request has not yet been completed.
	AllocationState string `json:"allocationState,omitempty"`

	// The time at which the pool entered its current allocation state.
	AllocationStateTransitionTime strfmt.DateTime `json:"allocationStateTransitionTime,omitempty"`

	// The list of application packages to be installed on each compute node in the pool.
	ApplicationPackageReferences []*ApplicationPackageReference `json:"applicationPackageReferences"`

	// The time interval at which to automatically adjust the pool size according to the autoscale formula.
	//
	// This property is set only if the pool automatically scales, i.e. enableAutoScale is true.
	AutoScaleEvaluationInterval strfmt.Duration `json:"autoScaleEvaluationInterval,omitempty"`

	// A formula for the desired number of compute nodes in the pool.
	//
	// This property is set only if the pool automatically scales, i.e. enableAutoScale is true.
	AutoScaleFormula string `json:"autoScaleFormula,omitempty"`

	// The results and errors from the last execution of the autoscale formula.
	//
	// This property is set only if the pool automatically scales, i.e. enableAutoScale is true.
	AutoScaleRun *AutoScaleRun `json:"autoScaleRun,omitempty"`

	// The list of certificates to be installed on each compute node in the pool.
	//
	// For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of remoteuser, a certs directory is created in the user's home directory (e.g., /home/<user-name>/certs) where certificates are placed.
	CertificateReferences []*CertificateReference `json:"certificateReferences"`

	// The cloud service configuration for the pool.
	//
	// This property and virtualMachineConfiguration are mutually exclusive and one of the properties must be specified.
	CloudServiceConfiguration *CloudServiceConfiguration `json:"cloudServiceConfiguration,omitempty"`

	// The creation time of the pool.
	CreationTime strfmt.DateTime `json:"creationTime,omitempty"`

	// The number of compute nodes currently in the pool.
	CurrentDedicated int32 `json:"currentDedicated,omitempty"`

	// The display name for the pool.
	//
	// The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
	DisplayName string `json:"displayName,omitempty"`

	// The ETag of the pool.
	//
	// This is an opaque string. You can use it to detect whether the pool has changed between requests. In particular, you can be pass the ETag when updating a pool to specify that your changes should take effect only if nobody else has modified the pool in the meantime.
	ETag string `json:"eTag,omitempty"`

	// Whether the pool size should automatically adjust over time.
	//
	// If true, the autoScaleFormula property must be set. If false, the targetDedicated property must be set.
	EnableAutoScale bool `json:"enableAutoScale,omitempty"`

	// Whether the pool permits direct communication between nodes.
	//
	// This imposes restrictions on which nodes can be assigned to the pool. Specifying this value can reduce the chance of the requested number of nodes to be allocated in the pool.
	EnableInterNodeCommunication bool `json:"enableInterNodeCommunication,omitempty"`

	// A string that uniquely identifies the pool within the account.
	//
	// The ID can contain any combination of alphanumeric characters including hyphens and underscores, and cannot contain more than 64 characters. It is common to use a GUID for the id.
	ID string `json:"id,omitempty"`

	// The last modified time of the pool.
	//
	// This is the last time at which the pool level data, such as the targetDedicated or enableAutoscale settings, changed. It does not factor in node-level changes such as a compute node changing state.
	LastModified strfmt.DateTime `json:"lastModified,omitempty"`

	// The maximum number of tasks that can run concurrently on a single compute node in the pool.
	MaxTasksPerNode int32 `json:"maxTasksPerNode,omitempty"`

	// A list of name-value pairs associated with the pool as metadata.
	Metadata []*MetadataItem `json:"metadata"`

	// The network configuration for the pool.
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty"`

	// Details of any error encountered while performing the last resize on the pool.
	//
	// This property is set only if an error occurred during the last pool resize, and only when the pool allocationState is Steady.
	ResizeError *ResizeError `json:"resizeError,omitempty"`

	// The timeout for allocation of compute nodes to the pool.
	//
	// This is the timeout for the most recent resize operation. (The initial sizing when the pool is created counts as a resize.) The default value is 15 minutes.
	ResizeTimeout strfmt.Duration `json:"resizeTimeout,omitempty"`

	// A task specified to run on each compute node as it joins the pool.
	StartTask *StartTask `json:"startTask,omitempty"`

	// The current state of the pool.
	//
	// Possible values are: active – The pool is available to run tasks subject to the availability of compute nodes. deleting – The user has requested that the pool be deleted, but the delete operation has not yet completed. upgrading – The user has requested that the operating system of the pool's nodes be upgraded, but the upgrade operation has not yet completed (that is, some nodes in the pool have not yet been upgraded). While upgrading, the pool may be able to run tasks (with reduced capacity) but this is not guaranteed.
	State string `json:"state,omitempty"`

	// The time at which the pool entered its current state.
	StateTransitionTime strfmt.DateTime `json:"stateTransitionTime,omitempty"`

	// Utilization and resource usage statistics for the entire lifetime of the pool.
	Stats *PoolStatistics `json:"stats,omitempty"`

	// The desired number of compute nodes in the pool.
	//
	// This property is not set if enableAutoScale is true. It is required if enableAutoScale is false.
	TargetDedicated int32 `json:"targetDedicated,omitempty"`

	// How the Batch service distributes tasks between compute nodes in the pool.
	TaskSchedulingPolicy *TaskSchedulingPolicy `json:"taskSchedulingPolicy,omitempty"`

	// The URL of the pool.
	URL string `json:"url,omitempty"`

	// The virtual machine configuration for the pool.
	//
	// This property and cloudServiceConfiguration are mutually exclusive and one of the properties must be specified.
	VirtualMachineConfiguration *VirtualMachineConfiguration `json:"virtualMachineConfiguration,omitempty"`

	// The size of virtual machines in the pool. All virtual machines in a pool are the same size.
	//
	// For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (http://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
	VMSize string `json:"vmSize,omitempty"`
}

// Validate validates this cloud pool
func (m *CloudPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocationState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateApplicationPackageReferences(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAutoScaleRun(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCertificateReferences(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCloudServiceConfiguration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworkConfiguration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResizeError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTask(formats); err != nil {
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

	if err := m.validateTaskSchedulingPolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualMachineConfiguration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cloudPoolTypeAllocationStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["steady","resizing","stopping"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudPoolTypeAllocationStatePropEnum = append(cloudPoolTypeAllocationStatePropEnum, v)
	}
}

const (
	// CloudPoolAllocationStateSteady captures enum value "steady"
	CloudPoolAllocationStateSteady string = "steady"
	// CloudPoolAllocationStateResizing captures enum value "resizing"
	CloudPoolAllocationStateResizing string = "resizing"
	// CloudPoolAllocationStateStopping captures enum value "stopping"
	CloudPoolAllocationStateStopping string = "stopping"
)

// prop value enum
func (m *CloudPool) validateAllocationStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudPoolTypeAllocationStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudPool) validateAllocationState(formats strfmt.Registry) error {

	if swag.IsZero(m.AllocationState) { // not required
		return nil
	}

	// value enum
	if err := m.validateAllocationStateEnum("allocationState", "body", m.AllocationState); err != nil {
		return err
	}

	return nil
}

func (m *CloudPool) validateApplicationPackageReferences(formats strfmt.Registry) error {

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

func (m *CloudPool) validateAutoScaleRun(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoScaleRun) { // not required
		return nil
	}

	if m.AutoScaleRun != nil {

		if err := m.AutoScaleRun.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CloudPool) validateCertificateReferences(formats strfmt.Registry) error {

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

func (m *CloudPool) validateCloudServiceConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudServiceConfiguration) { // not required
		return nil
	}

	if m.CloudServiceConfiguration != nil {

		if err := m.CloudServiceConfiguration.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CloudPool) validateMetadata(formats strfmt.Registry) error {

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

func (m *CloudPool) validateNetworkConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkConfiguration) { // not required
		return nil
	}

	if m.NetworkConfiguration != nil {

		if err := m.NetworkConfiguration.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CloudPool) validateResizeError(formats strfmt.Registry) error {

	if swag.IsZero(m.ResizeError) { // not required
		return nil
	}

	if m.ResizeError != nil {

		if err := m.ResizeError.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CloudPool) validateStartTask(formats strfmt.Registry) error {

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

var cloudPoolTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","deleting","upgrading"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudPoolTypeStatePropEnum = append(cloudPoolTypeStatePropEnum, v)
	}
}

const (
	// CloudPoolStateActive captures enum value "active"
	CloudPoolStateActive string = "active"
	// CloudPoolStateDeleting captures enum value "deleting"
	CloudPoolStateDeleting string = "deleting"
	// CloudPoolStateUpgrading captures enum value "upgrading"
	CloudPoolStateUpgrading string = "upgrading"
)

// prop value enum
func (m *CloudPool) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudPoolTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudPool) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *CloudPool) validateStats(formats strfmt.Registry) error {

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

func (m *CloudPool) validateTaskSchedulingPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskSchedulingPolicy) { // not required
		return nil
	}

	if m.TaskSchedulingPolicy != nil {

		if err := m.TaskSchedulingPolicy.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CloudPool) validateVirtualMachineConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualMachineConfiguration) { // not required
		return nil
	}

	if m.VirtualMachineConfiguration != nil {

		if err := m.VirtualMachineConfiguration.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
