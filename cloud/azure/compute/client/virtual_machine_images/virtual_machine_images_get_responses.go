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

// VirtualMachineImagesGetReader is a Reader for the VirtualMachineImagesGet structure.
type VirtualMachineImagesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachineImagesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualMachineImagesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachineImagesGetOK creates a VirtualMachineImagesGetOK with default headers values
func NewVirtualMachineImagesGetOK() *VirtualMachineImagesGetOK {
	return &VirtualMachineImagesGetOK{}
}

/*VirtualMachineImagesGetOK handles this case with default header values.

OK
*/
type VirtualMachineImagesGetOK struct {
	Payload *models.VirtualMachineImage
}

func (o *VirtualMachineImagesGetOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}][%d] virtualMachineImagesGetOK  %+v", 200, o.Payload)
}

func (o *VirtualMachineImagesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineImage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
