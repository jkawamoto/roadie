package availability_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/compute/models"
)

// AvailabilitySetsListAvailableSizesReader is a Reader for the AvailabilitySetsListAvailableSizes structure.
type AvailabilitySetsListAvailableSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AvailabilitySetsListAvailableSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAvailabilitySetsListAvailableSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAvailabilitySetsListAvailableSizesOK creates a AvailabilitySetsListAvailableSizesOK with default headers values
func NewAvailabilitySetsListAvailableSizesOK() *AvailabilitySetsListAvailableSizesOK {
	return &AvailabilitySetsListAvailableSizesOK{}
}

/*AvailabilitySetsListAvailableSizesOK handles this case with default header values.

OK
*/
type AvailabilitySetsListAvailableSizesOK struct {
	Payload *models.VirtualMachineSizeListResult
}

func (o *AvailabilitySetsListAvailableSizesOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes][%d] availabilitySetsListAvailableSizesOK  %+v", 200, o.Payload)
}

func (o *AvailabilitySetsListAvailableSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineSizeListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
