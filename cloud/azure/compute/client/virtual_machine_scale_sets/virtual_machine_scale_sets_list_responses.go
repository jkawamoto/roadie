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

// VirtualMachineScaleSetsListReader is a Reader for the VirtualMachineScaleSetsList structure.
type VirtualMachineScaleSetsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachineScaleSetsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualMachineScaleSetsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachineScaleSetsListOK creates a VirtualMachineScaleSetsListOK with default headers values
func NewVirtualMachineScaleSetsListOK() *VirtualMachineScaleSetsListOK {
	return &VirtualMachineScaleSetsListOK{}
}

/*VirtualMachineScaleSetsListOK handles this case with default header values.

OK
*/
type VirtualMachineScaleSetsListOK struct {
	Payload *models.VirtualMachineScaleSetListResult
}

func (o *VirtualMachineScaleSetsListOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets][%d] virtualMachineScaleSetsListOK  %+v", 200, o.Payload)
}

func (o *VirtualMachineScaleSetsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineScaleSetListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}