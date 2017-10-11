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

// VirtualMachinesCreateOrUpdateReader is a Reader for the VirtualMachinesCreateOrUpdate structure.
type VirtualMachinesCreateOrUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachinesCreateOrUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualMachinesCreateOrUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewVirtualMachinesCreateOrUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachinesCreateOrUpdateOK creates a VirtualMachinesCreateOrUpdateOK with default headers values
func NewVirtualMachinesCreateOrUpdateOK() *VirtualMachinesCreateOrUpdateOK {
	return &VirtualMachinesCreateOrUpdateOK{}
}

/*VirtualMachinesCreateOrUpdateOK handles this case with default header values.

OK
*/
type VirtualMachinesCreateOrUpdateOK struct {
	Payload *models.VirtualMachine
}

func (o *VirtualMachinesCreateOrUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}][%d] virtualMachinesCreateOrUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualMachinesCreateOrUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualMachinesCreateOrUpdateCreated creates a VirtualMachinesCreateOrUpdateCreated with default headers values
func NewVirtualMachinesCreateOrUpdateCreated() *VirtualMachinesCreateOrUpdateCreated {
	return &VirtualMachinesCreateOrUpdateCreated{}
}

/*VirtualMachinesCreateOrUpdateCreated handles this case with default header values.

Created
*/
type VirtualMachinesCreateOrUpdateCreated struct {
	Payload *models.VirtualMachine
}

func (o *VirtualMachinesCreateOrUpdateCreated) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}][%d] virtualMachinesCreateOrUpdateCreated  %+v", 201, o.Payload)
}

func (o *VirtualMachinesCreateOrUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
