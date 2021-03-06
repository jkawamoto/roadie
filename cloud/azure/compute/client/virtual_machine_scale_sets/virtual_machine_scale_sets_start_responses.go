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

// VirtualMachineScaleSetsStartReader is a Reader for the VirtualMachineScaleSetsStart structure.
type VirtualMachineScaleSetsStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachineScaleSetsStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualMachineScaleSetsStartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewVirtualMachineScaleSetsStartAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachineScaleSetsStartOK creates a VirtualMachineScaleSetsStartOK with default headers values
func NewVirtualMachineScaleSetsStartOK() *VirtualMachineScaleSetsStartOK {
	return &VirtualMachineScaleSetsStartOK{}
}

/*VirtualMachineScaleSetsStartOK handles this case with default header values.

OK
*/
type VirtualMachineScaleSetsStartOK struct {
	Payload *models.OperationStatusResponse
}

func (o *VirtualMachineScaleSetsStartOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/start][%d] virtualMachineScaleSetsStartOK  %+v", 200, o.Payload)
}

func (o *VirtualMachineScaleSetsStartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OperationStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualMachineScaleSetsStartAccepted creates a VirtualMachineScaleSetsStartAccepted with default headers values
func NewVirtualMachineScaleSetsStartAccepted() *VirtualMachineScaleSetsStartAccepted {
	return &VirtualMachineScaleSetsStartAccepted{}
}

/*VirtualMachineScaleSetsStartAccepted handles this case with default header values.

Accepted
*/
type VirtualMachineScaleSetsStartAccepted struct {
}

func (o *VirtualMachineScaleSetsStartAccepted) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/start][%d] virtualMachineScaleSetsStartAccepted ", 202)
}

func (o *VirtualMachineScaleSetsStartAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
