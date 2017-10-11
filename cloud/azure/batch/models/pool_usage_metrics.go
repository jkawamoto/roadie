package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PoolUsageMetrics Usage metrics for a pool across an aggregation interval.
// swagger:model PoolUsageMetrics
type PoolUsageMetrics struct {

	// The cross data center network egress from the pool during this interval, in GiB.
	// Required: true
	DataEgressGiB *float64 `json:"dataEgressGiB"`

	// The cross data center network ingress to the pool during this interval, in GiB.
	// Required: true
	DataIngressGiB *float64 `json:"dataIngressGiB"`

	// The end time of the aggregation interval covered by this entry.
	// Required: true
	EndTime *strfmt.DateTime `json:"endTime"`

	// The ID of the pool whose metrics are aggregated in this entry.
	// Required: true
	PoolID *string `json:"poolId"`

	// The start time of the aggregation interval covered by this entry.
	// Required: true
	StartTime *strfmt.DateTime `json:"startTime"`

	// The total core hours used in the pool during this aggregation interval.
	// Required: true
	TotalCoreHours *float64 `json:"totalCoreHours"`

	// The size of virtual machines in the pool. All VMs in a pool are the same size.
	//
	// For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (http://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
	// Required: true
	VMSize *string `json:"vmSize"`
}

// Validate validates this pool usage metrics
func (m *PoolUsageMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataEgressGiB(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDataIngressGiB(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePoolID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTotalCoreHours(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVMSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolUsageMetrics) validateDataEgressGiB(formats strfmt.Registry) error {

	if err := validate.Required("dataEgressGiB", "body", m.DataEgressGiB); err != nil {
		return err
	}

	return nil
}

func (m *PoolUsageMetrics) validateDataIngressGiB(formats strfmt.Registry) error {

	if err := validate.Required("dataIngressGiB", "body", m.DataIngressGiB); err != nil {
		return err
	}

	return nil
}

func (m *PoolUsageMetrics) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("endTime", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *PoolUsageMetrics) validatePoolID(formats strfmt.Registry) error {

	if err := validate.Required("poolId", "body", m.PoolID); err != nil {
		return err
	}

	return nil
}

func (m *PoolUsageMetrics) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *PoolUsageMetrics) validateTotalCoreHours(formats strfmt.Registry) error {

	if err := validate.Required("totalCoreHours", "body", m.TotalCoreHours); err != nil {
		return err
	}

	return nil
}

func (m *PoolUsageMetrics) validateVMSize(formats strfmt.Registry) error {

	if err := validate.Required("vmSize", "body", m.VMSize); err != nil {
		return err
	}

	return nil
}
