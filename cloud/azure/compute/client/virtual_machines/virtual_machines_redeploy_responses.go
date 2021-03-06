package virtual_machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/compute/models"
)

// VirtualMachinesRedeployReader is a Reader for the VirtualMachinesRedeploy structure.
type VirtualMachinesRedeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachinesRedeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualMachinesRedeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewVirtualMachinesRedeployAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachinesRedeployOK creates a VirtualMachinesRedeployOK with default headers values
func NewVirtualMachinesRedeployOK() *VirtualMachinesRedeployOK {
	return &VirtualMachinesRedeployOK{}
}

/*VirtualMachinesRedeployOK handles this case with default header values.

OK
*/
type VirtualMachinesRedeployOK struct {
	Payload *models.OperationStatusResponse
}

func (o *VirtualMachinesRedeployOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/redeploy][%d] virtualMachinesRedeployOK  %+v", 200, o.Payload)
}

func (o *VirtualMachinesRedeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OperationStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualMachinesRedeployAccepted creates a VirtualMachinesRedeployAccepted with default headers values
func NewVirtualMachinesRedeployAccepted() *VirtualMachinesRedeployAccepted {
	return &VirtualMachinesRedeployAccepted{}
}

/*VirtualMachinesRedeployAccepted handles this case with default header values.

Accepted
*/
type VirtualMachinesRedeployAccepted struct {
}

func (o *VirtualMachinesRedeployAccepted) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/redeploy][%d] virtualMachinesRedeployAccepted ", 202)
}

func (o *VirtualMachinesRedeployAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
