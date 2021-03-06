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

// VirtualMachinesDeallocateReader is a Reader for the VirtualMachinesDeallocate structure.
type VirtualMachinesDeallocateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachinesDeallocateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualMachinesDeallocateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewVirtualMachinesDeallocateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachinesDeallocateOK creates a VirtualMachinesDeallocateOK with default headers values
func NewVirtualMachinesDeallocateOK() *VirtualMachinesDeallocateOK {
	return &VirtualMachinesDeallocateOK{}
}

/*VirtualMachinesDeallocateOK handles this case with default header values.

OK
*/
type VirtualMachinesDeallocateOK struct {
	Payload *models.OperationStatusResponse
}

func (o *VirtualMachinesDeallocateOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/deallocate][%d] virtualMachinesDeallocateOK  %+v", 200, o.Payload)
}

func (o *VirtualMachinesDeallocateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OperationStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualMachinesDeallocateAccepted creates a VirtualMachinesDeallocateAccepted with default headers values
func NewVirtualMachinesDeallocateAccepted() *VirtualMachinesDeallocateAccepted {
	return &VirtualMachinesDeallocateAccepted{}
}

/*VirtualMachinesDeallocateAccepted handles this case with default header values.

Accepted
*/
type VirtualMachinesDeallocateAccepted struct {
}

func (o *VirtualMachinesDeallocateAccepted) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/deallocate][%d] virtualMachinesDeallocateAccepted ", 202)
}

func (o *VirtualMachinesDeallocateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
