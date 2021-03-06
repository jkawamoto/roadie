package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/resource/models"
)

// DeploymentsCreateOrUpdateReader is a Reader for the DeploymentsCreateOrUpdate structure.
type DeploymentsCreateOrUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeploymentsCreateOrUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeploymentsCreateOrUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewDeploymentsCreateOrUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeploymentsCreateOrUpdateOK creates a DeploymentsCreateOrUpdateOK with default headers values
func NewDeploymentsCreateOrUpdateOK() *DeploymentsCreateOrUpdateOK {
	return &DeploymentsCreateOrUpdateOK{}
}

/*DeploymentsCreateOrUpdateOK handles this case with default header values.

OK - Returns information about the deployment, including provisioning status.
*/
type DeploymentsCreateOrUpdateOK struct {
	Payload *models.DeploymentExtended
}

func (o *DeploymentsCreateOrUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}][%d] deploymentsCreateOrUpdateOK  %+v", 200, o.Payload)
}

func (o *DeploymentsCreateOrUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeploymentsCreateOrUpdateCreated creates a DeploymentsCreateOrUpdateCreated with default headers values
func NewDeploymentsCreateOrUpdateCreated() *DeploymentsCreateOrUpdateCreated {
	return &DeploymentsCreateOrUpdateCreated{}
}

/*DeploymentsCreateOrUpdateCreated handles this case with default header values.

Created - Returns information about the deployment, including provisioning status.
*/
type DeploymentsCreateOrUpdateCreated struct {
	Payload *models.DeploymentExtended
}

func (o *DeploymentsCreateOrUpdateCreated) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}][%d] deploymentsCreateOrUpdateCreated  %+v", 201, o.Payload)
}

func (o *DeploymentsCreateOrUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
