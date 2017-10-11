package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentPropertiesExtended Deployment properties with additional details.
// swagger:model DeploymentPropertiesExtended
type DeploymentPropertiesExtended struct {

	// The correlation ID of the deployment.
	// Read Only: true
	CorrelationID string `json:"correlationId,omitempty"`

	// The debug setting of the deployment.
	DebugSetting *DebugSetting `json:"debugSetting,omitempty"`

	// The list of deployment dependencies.
	Dependencies []*Dependency `json:"dependencies"`

	// The deployment mode. Possible values are Incremental and Complete.
	Mode string `json:"mode,omitempty"`

	// Key/value pairs that represent deploymentoutput.
	Outputs interface{} `json:"outputs,omitempty"`

	// Deployment parameters. Use only one of Parameters or ParametersLink.
	Parameters interface{} `json:"parameters,omitempty"`

	// The URI referencing the parameters. Use only one of Parameters or ParametersLink.
	ParametersLink *ParametersLink `json:"parametersLink,omitempty"`

	// The list of resource providers needed for the deployment.
	Providers []*Provider `json:"providers"`

	// The state of the provisioning.
	// Read Only: true
	ProvisioningState string `json:"provisioningState,omitempty"`

	// The template content. Use only one of Template or TemplateLink.
	Template interface{} `json:"template,omitempty"`

	// The URI referencing the template. Use only one of Template or TemplateLink.
	TemplateLink *TemplateLink `json:"templateLink,omitempty"`

	// The timestamp of the template deployment.
	// Read Only: true
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this deployment properties extended
func (m *DeploymentPropertiesExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDebugSetting(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDependencies(formats); err != nil {
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

	if err := m.validateProviders(formats); err != nil {
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

func (m *DeploymentPropertiesExtended) validateDebugSetting(formats strfmt.Registry) error {

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

func (m *DeploymentPropertiesExtended) validateDependencies(formats strfmt.Registry) error {

	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {

		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {

			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var deploymentPropertiesExtendedTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Incremental","Complete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentPropertiesExtendedTypeModePropEnum = append(deploymentPropertiesExtendedTypeModePropEnum, v)
	}
}

const (
	// DeploymentPropertiesExtendedModeIncremental captures enum value "Incremental"
	DeploymentPropertiesExtendedModeIncremental string = "Incremental"
	// DeploymentPropertiesExtendedModeComplete captures enum value "Complete"
	DeploymentPropertiesExtendedModeComplete string = "Complete"
)

// prop value enum
func (m *DeploymentPropertiesExtended) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deploymentPropertiesExtendedTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentPropertiesExtended) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentPropertiesExtended) validateParametersLink(formats strfmt.Registry) error {

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

func (m *DeploymentPropertiesExtended) validateProviders(formats strfmt.Registry) error {

	if swag.IsZero(m.Providers) { // not required
		return nil
	}

	for i := 0; i < len(m.Providers); i++ {

		if swag.IsZero(m.Providers[i]) { // not required
			continue
		}

		if m.Providers[i] != nil {

			if err := m.Providers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentPropertiesExtended) validateTemplateLink(formats strfmt.Registry) error {

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