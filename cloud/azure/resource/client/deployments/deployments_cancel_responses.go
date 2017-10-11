package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeploymentsCancelReader is a Reader for the DeploymentsCancel structure.
type DeploymentsCancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeploymentsCancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeploymentsCancelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeploymentsCancelNoContent creates a DeploymentsCancelNoContent with default headers values
func NewDeploymentsCancelNoContent() *DeploymentsCancelNoContent {
	return &DeploymentsCancelNoContent{}
}

/*DeploymentsCancelNoContent handles this case with default header values.

No Content
*/
type DeploymentsCancelNoContent struct {
}

func (o *DeploymentsCancelNoContent) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}/cancel][%d] deploymentsCancelNoContent ", 204)
}

func (o *DeploymentsCancelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}