package virtual_machine_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/compute/models"
)

// VirtualMachineImagesListSkusReader is a Reader for the VirtualMachineImagesListSkus structure.
type VirtualMachineImagesListSkusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachineImagesListSkusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualMachineImagesListSkusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachineImagesListSkusOK creates a VirtualMachineImagesListSkusOK with default headers values
func NewVirtualMachineImagesListSkusOK() *VirtualMachineImagesListSkusOK {
	return &VirtualMachineImagesListSkusOK{}
}

/*VirtualMachineImagesListSkusOK handles this case with default header values.

OK
*/
type VirtualMachineImagesListSkusOK struct {
	Payload []*models.VirtualMachineImageResource
}

func (o *VirtualMachineImagesListSkusOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus][%d] virtualMachineImagesListSkusOK  %+v", 200, o.Payload)
}

func (o *VirtualMachineImagesListSkusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
