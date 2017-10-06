package virtual_machine_scale_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/compute/models"
)

// VirtualMachineScaleSetsPowerOffReader is a Reader for the VirtualMachineScaleSetsPowerOff structure.
type VirtualMachineScaleSetsPowerOffReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachineScaleSetsPowerOffReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualMachineScaleSetsPowerOffOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewVirtualMachineScaleSetsPowerOffAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachineScaleSetsPowerOffOK creates a VirtualMachineScaleSetsPowerOffOK with default headers values
func NewVirtualMachineScaleSetsPowerOffOK() *VirtualMachineScaleSetsPowerOffOK {
	return &VirtualMachineScaleSetsPowerOffOK{}
}

/*VirtualMachineScaleSetsPowerOffOK handles this case with default header values.

OK
*/
type VirtualMachineScaleSetsPowerOffOK struct {
	Payload *models.OperationStatusResponse
}

func (o *VirtualMachineScaleSetsPowerOffOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/poweroff][%d] virtualMachineScaleSetsPowerOffOK  %+v", 200, o.Payload)
}

func (o *VirtualMachineScaleSetsPowerOffOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OperationStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualMachineScaleSetsPowerOffAccepted creates a VirtualMachineScaleSetsPowerOffAccepted with default headers values
func NewVirtualMachineScaleSetsPowerOffAccepted() *VirtualMachineScaleSetsPowerOffAccepted {
	return &VirtualMachineScaleSetsPowerOffAccepted{}
}

/*VirtualMachineScaleSetsPowerOffAccepted handles this case with default header values.

Accepted
*/
type VirtualMachineScaleSetsPowerOffAccepted struct {
}

func (o *VirtualMachineScaleSetsPowerOffAccepted) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/poweroff][%d] virtualMachineScaleSetsPowerOffAccepted ", 202)
}

func (o *VirtualMachineScaleSetsPowerOffAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
