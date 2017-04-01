package application_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/network/models"
)

// ApplicationGatewaysCreateOrUpdateReader is a Reader for the ApplicationGatewaysCreateOrUpdate structure.
type ApplicationGatewaysCreateOrUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationGatewaysCreateOrUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewApplicationGatewaysCreateOrUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewApplicationGatewaysCreateOrUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewApplicationGatewaysCreateOrUpdateOK creates a ApplicationGatewaysCreateOrUpdateOK with default headers values
func NewApplicationGatewaysCreateOrUpdateOK() *ApplicationGatewaysCreateOrUpdateOK {
	return &ApplicationGatewaysCreateOrUpdateOK{}
}

/*ApplicationGatewaysCreateOrUpdateOK handles this case with default header values.

ApplicationGatewaysCreateOrUpdateOK application gateways create or update o k
*/
type ApplicationGatewaysCreateOrUpdateOK struct {
	Payload *models.ApplicationGateway
}

func (o *ApplicationGatewaysCreateOrUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}][%d] applicationGatewaysCreateOrUpdateOK  %+v", 200, o.Payload)
}

func (o *ApplicationGatewaysCreateOrUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationGatewaysCreateOrUpdateCreated creates a ApplicationGatewaysCreateOrUpdateCreated with default headers values
func NewApplicationGatewaysCreateOrUpdateCreated() *ApplicationGatewaysCreateOrUpdateCreated {
	return &ApplicationGatewaysCreateOrUpdateCreated{}
}

/*ApplicationGatewaysCreateOrUpdateCreated handles this case with default header values.

ApplicationGatewaysCreateOrUpdateCreated application gateways create or update created
*/
type ApplicationGatewaysCreateOrUpdateCreated struct {
	Payload *models.ApplicationGateway
}

func (o *ApplicationGatewaysCreateOrUpdateCreated) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}][%d] applicationGatewaysCreateOrUpdateCreated  %+v", 201, o.Payload)
}

func (o *ApplicationGatewaysCreateOrUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
