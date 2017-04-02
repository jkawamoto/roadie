package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/disk/models"
)

// SnapshotsListReader is a Reader for the SnapshotsList structure.
type SnapshotsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSnapshotsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSnapshotsListOK creates a SnapshotsListOK with default headers values
func NewSnapshotsListOK() *SnapshotsListOK {
	return &SnapshotsListOK{}
}

/*SnapshotsListOK handles this case with default header values.

OK
*/
type SnapshotsListOK struct {
	Payload *models.SnapshotList
}

func (o *SnapshotsListOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots][%d] snapshotsListOK  %+v", 200, o.Payload)
}

func (o *SnapshotsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
