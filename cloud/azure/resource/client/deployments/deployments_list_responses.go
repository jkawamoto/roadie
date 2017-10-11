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

// DeploymentsListReader is a Reader for the DeploymentsList structure.
type DeploymentsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeploymentsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeploymentsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeploymentsListOK creates a DeploymentsListOK with default headers values
func NewDeploymentsListOK() *DeploymentsListOK {
	return &DeploymentsListOK{}
}

/*DeploymentsListOK handles this case with default header values.

OK - Returns an array of deployments.
*/
type DeploymentsListOK struct {
	Payload *models.DeploymentListResult
}

func (o *DeploymentsListOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/][%d] deploymentsListOK  %+v", 200, o.Payload)
}

func (o *DeploymentsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}