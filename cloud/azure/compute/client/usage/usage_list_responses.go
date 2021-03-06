package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/compute/models"
)

// UsageListReader is a Reader for the UsageList structure.
type UsageListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsageListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsageListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsageListOK creates a UsageListOK with default headers values
func NewUsageListOK() *UsageListOK {
	return &UsageListOK{}
}

/*UsageListOK handles this case with default header values.

OK
*/
type UsageListOK struct {
	Payload *models.ListUsagesResult
}

func (o *UsageListOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/usages][%d] usageListOK  %+v", 200, o.Payload)
}

func (o *UsageListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListUsagesResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
