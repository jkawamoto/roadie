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

// VirtualMachineScaleSetsGetReader is a Reader for the VirtualMachineScaleSetsGet structure.
type VirtualMachineScaleSetsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachineScaleSetsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualMachineScaleSetsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachineScaleSetsGetOK creates a VirtualMachineScaleSetsGetOK with default headers values
func NewVirtualMachineScaleSetsGetOK() *VirtualMachineScaleSetsGetOK {
	return &VirtualMachineScaleSetsGetOK{}
}

/*VirtualMachineScaleSetsGetOK handles this case with default header values.

OK
*/
type VirtualMachineScaleSetsGetOK struct {
	Payload *models.VirtualMachineScaleSet
}

func (o *VirtualMachineScaleSetsGetOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}][%d] virtualMachineScaleSetsGetOK  %+v", 200, o.Payload)
}

func (o *VirtualMachineScaleSetsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineScaleSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}