package batch_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/batchmanagement/models"
)

// BatchAccountUpdateReader is a Reader for the BatchAccountUpdate structure.
type BatchAccountUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchAccountUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBatchAccountUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewBatchAccountUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBatchAccountUpdateOK creates a BatchAccountUpdateOK with default headers values
func NewBatchAccountUpdateOK() *BatchAccountUpdateOK {
	return &BatchAccountUpdateOK{}
}

/*BatchAccountUpdateOK handles this case with default header values.

The operation was successful. The response contains the Batch account entity.
*/
type BatchAccountUpdateOK struct {
	Payload *models.BatchAccount
}

func (o *BatchAccountUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}][%d] batchAccountUpdateOK  %+v", 200, o.Payload)
}

func (o *BatchAccountUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchAccountUpdateDefault creates a BatchAccountUpdateDefault with default headers values
func NewBatchAccountUpdateDefault(code int) *BatchAccountUpdateDefault {
	return &BatchAccountUpdateDefault{
		_statusCode: code,
	}
}

/*BatchAccountUpdateDefault handles this case with default header values.

Error response describing why the operation failed.
*/
type BatchAccountUpdateDefault struct {
	_statusCode int

	Payload *models.CloudError
}

// Code gets the status code for the batch account update default response
func (o *BatchAccountUpdateDefault) Code() int {
	return o._statusCode
}

func (o *BatchAccountUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}][%d] BatchAccount_Update default  %+v", o._statusCode, o.Payload)
}

func (o *BatchAccountUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
