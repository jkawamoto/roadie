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

// ImageDataDisk Describes a data disk.
// swagger:model ImageDataDisk
type ImageDataDisk struct {

	// The Virtual Hard Disk.
	BlobURI string `json:"blobUri,omitempty"`

	// The caching type.
	Caching string `json:"caching,omitempty"`

	// The initial disk size in GB for blank data disks, and the new desired size for existing OS and Data disks.
	DiskSizeGB int32 `json:"diskSizeGB,omitempty"`

	// The logical unit number.
	// Required: true
	Lun *int32 `json:"lun"`

	// The managedDisk.
	ManagedDisk *SubResource `json:"managedDisk,omitempty"`

	// The snapshot.
	Snapshot *SubResource `json:"snapshot,omitempty"`
}

// Validate validates this image data disk
func (m *ImageDataDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaching(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLun(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManagedDisk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSnapshot(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var imageDataDiskTypeCachingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","ReadOnly","ReadWrite"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageDataDiskTypeCachingPropEnum = append(imageDataDiskTypeCachingPropEnum, v)
	}
}

const (
	// ImageDataDiskCachingNone captures enum value "None"
	ImageDataDiskCachingNone string = "None"
	// ImageDataDiskCachingReadOnly captures enum value "ReadOnly"
	ImageDataDiskCachingReadOnly string = "ReadOnly"
	// ImageDataDiskCachingReadWrite captures enum value "ReadWrite"
	ImageDataDiskCachingReadWrite string = "ReadWrite"
)

// prop value enum
func (m *ImageDataDisk) validateCachingEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, imageDataDiskTypeCachingPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ImageDataDisk) validateCaching(formats strfmt.Registry) error {

	if swag.IsZero(m.Caching) { // not required
		return nil
	}

	// value enum
	if err := m.validateCachingEnum("caching", "body", m.Caching); err != nil {
		return err
	}

	return nil
}

func (m *ImageDataDisk) validateLun(formats strfmt.Registry) error {

	if err := validate.Required("lun", "body", m.Lun); err != nil {
		return err
	}

	return nil
}

func (m *ImageDataDisk) validateManagedDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagedDisk) { // not required
		return nil
	}

	if m.ManagedDisk != nil {

		if err := m.ManagedDisk.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ImageDataDisk) validateSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.Snapshot) { // not required
		return nil
	}

	if m.Snapshot != nil {

		if err := m.Snapshot.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
