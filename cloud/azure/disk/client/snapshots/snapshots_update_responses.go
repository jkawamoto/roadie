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

// SnapshotsUpdateReader is a Reader for the SnapshotsUpdate structure.
type SnapshotsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSnapshotsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewSnapshotsUpdateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSnapshotsUpdateOK creates a SnapshotsUpdateOK with default headers values
func NewSnapshotsUpdateOK() *SnapshotsUpdateOK {
	return &SnapshotsUpdateOK{}
}

/*SnapshotsUpdateOK handles this case with default header values.

OK
*/
type SnapshotsUpdateOK struct {
	Payload *models.Snapshot
}

func (o *SnapshotsUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}][%d] snapshotsUpdateOK  %+v", 200, o.Payload)
}

func (o *SnapshotsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotsUpdateAccepted creates a SnapshotsUpdateAccepted with default headers values
func NewSnapshotsUpdateAccepted() *SnapshotsUpdateAccepted {
	return &SnapshotsUpdateAccepted{}
}

/*SnapshotsUpdateAccepted handles this case with default header values.

Accepted
*/
type SnapshotsUpdateAccepted struct {
	Payload *models.Snapshot
}

func (o *SnapshotsUpdateAccepted) Error() string {
	return fmt.Sprintf("[PATCH /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}][%d] snapshotsUpdateAccepted  %+v", 202, o.Payload)
}

func (o *SnapshotsUpdateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
