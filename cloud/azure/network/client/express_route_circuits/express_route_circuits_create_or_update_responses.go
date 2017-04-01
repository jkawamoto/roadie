package express_route_circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/network/models"
)

// ExpressRouteCircuitsCreateOrUpdateReader is a Reader for the ExpressRouteCircuitsCreateOrUpdate structure.
type ExpressRouteCircuitsCreateOrUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExpressRouteCircuitsCreateOrUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExpressRouteCircuitsCreateOrUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewExpressRouteCircuitsCreateOrUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExpressRouteCircuitsCreateOrUpdateOK creates a ExpressRouteCircuitsCreateOrUpdateOK with default headers values
func NewExpressRouteCircuitsCreateOrUpdateOK() *ExpressRouteCircuitsCreateOrUpdateOK {
	return &ExpressRouteCircuitsCreateOrUpdateOK{}
}

/*ExpressRouteCircuitsCreateOrUpdateOK handles this case with default header values.

ExpressRouteCircuitsCreateOrUpdateOK express route circuits create or update o k
*/
type ExpressRouteCircuitsCreateOrUpdateOK struct {
	Payload *models.ExpressRouteCircuit
}

func (o *ExpressRouteCircuitsCreateOrUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}][%d] expressRouteCircuitsCreateOrUpdateOK  %+v", 200, o.Payload)
}

func (o *ExpressRouteCircuitsCreateOrUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExpressRouteCircuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExpressRouteCircuitsCreateOrUpdateCreated creates a ExpressRouteCircuitsCreateOrUpdateCreated with default headers values
func NewExpressRouteCircuitsCreateOrUpdateCreated() *ExpressRouteCircuitsCreateOrUpdateCreated {
	return &ExpressRouteCircuitsCreateOrUpdateCreated{}
}

/*ExpressRouteCircuitsCreateOrUpdateCreated handles this case with default header values.

ExpressRouteCircuitsCreateOrUpdateCreated express route circuits create or update created
*/
type ExpressRouteCircuitsCreateOrUpdateCreated struct {
	Payload *models.ExpressRouteCircuit
}

func (o *ExpressRouteCircuitsCreateOrUpdateCreated) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}][%d] expressRouteCircuitsCreateOrUpdateCreated  %+v", 201, o.Payload)
}

func (o *ExpressRouteCircuitsCreateOrUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExpressRouteCircuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
