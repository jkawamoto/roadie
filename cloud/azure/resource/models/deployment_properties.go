package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentProperties Deployment properties.
// swagger:model DeploymentProperties
type DeploymentProperties struct {

	// The debug setting of the deployment.
	DebugSetting *DebugSetting `json:"debugSetting,omitempty"`

	// The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental mode, resources are deployed without deleting existing resources that are not included in the template. In Complete mode, resources are deployed and existing resources in the resource group that are not included in the template are deleted. Be careful when using Complete mode as you may unintentionally delete resources.
	// Required: true
	Mode *string `json:"mode"`

	// Name and value pairs that define the deployment parameters for the template. You use this element when you want to provide the parameter values directly in the request rather than link to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
	Parameters interface{} `json:"parameters,omitempty"`

	// The URI of parameters file. You use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
	ParametersLink *ParametersLink `json:"parametersLink,omitempty"`

	// The template content. You use this element when you want to pass the template syntax directly in the request rather than link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the template property, but not both.
	Template interface{} `json:"template,omitempty"`

	// The URI of the template. Use either the templateLink property or the template property, but not both.
	TemplateLink *TemplateLink `json:"templateLink,omitempty"`
}

// Validate validates this deployment properties
func (m *DeploymentProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDebugSetting(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParametersLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTemplateLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentProperties) validateDebugSetting(formats strfmt.Registry) error {

	if swag.IsZero(m.DebugSetting) { // not required
		return nil
	}

	if m.DebugSetting != nil {

		if err := m.DebugSetting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debugSetting")
			}
			return err
		}
	}

	return nil
}

var deploymentPropertiesTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Incremental","Complete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentPropertiesTypeModePropEnum = append(deploymentPropertiesTypeModePropEnum, v)
	}
}

const (
	// DeploymentPropertiesModeIncremental captures enum value "Incremental"
	DeploymentPropertiesModeIncremental string = "Incremental"
	// DeploymentPropertiesModeComplete captures enum value "Complete"
	DeploymentPropertiesModeComplete string = "Complete"
)

// prop value enum
func (m *DeploymentProperties) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deploymentPropertiesTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentProperties) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentProperties) validateParametersLink(formats strfmt.Registry) error {

	if swag.IsZero(m.ParametersLink) { // not required
		return nil
	}

	if m.ParametersLink != nil {

		if err := m.ParametersLink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parametersLink")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentProperties) validateTemplateLink(formats strfmt.Registry) error {

	if swag.IsZero(m.TemplateLink) { // not required
		return nil
	}

	if m.TemplateLink != nil {

		if err := m.TemplateLink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("templateLink")
			}
			return err
		}
	}

	return nil
}