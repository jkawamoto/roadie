package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ResourceFile A file to be downloaded from Azure blob storage to a compute node.
// swagger:model ResourceFile
type ResourceFile struct {

	// The URL of the file within Azure Blob Storage.
	//
	// This URL must be readable using anonymous access; that is, the Batch service does not present any credentials when downloading the blob. There are two ways to get such a URL for a blob in Azure storage: include a Shared Access Signature (SAS) granting read permissions on the blob, or set the ACL for the blob or its container to allow public access.
	// Required: true
	BlobSource *string `json:"blobSource"`

	// The file permission mode attribute in octal format.
	//
	// This property applies only to files being downloaded to Linux compute nodes. It will be ignored if it is specified for a resourceFile which will be downloaded to a Windows node. If this property is not specified for a Linux node, then a default value of 0770 is applied to the file.
	FileMode string `json:"fileMode,omitempty"`

	// The location on the compute node to which to download the file, relative to the task's working directory.
	// Required: true
	FilePath *string `json:"filePath"`
}

// Validate validates this resource file
func (m *ResourceFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlobSource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFilePath(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceFile) validateBlobSource(formats strfmt.Registry) error {

	if err := validate.Required("blobSource", "body", m.BlobSource); err != nil {
		return err
	}

	return nil
}

func (m *ResourceFile) validateFilePath(formats strfmt.Registry) error {

	if err := validate.Required("filePath", "body", m.FilePath); err != nil {
		return err
	}

	return nil
}
