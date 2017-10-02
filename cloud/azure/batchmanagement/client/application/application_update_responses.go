package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/batchmanagement/models"
)

// ApplicationUpdateReader is a Reader for the ApplicationUpdate structure.
type ApplicationUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewApplicationUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewApplicationUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationUpdateNoContent creates a ApplicationUpdateNoContent with default headers values
func NewApplicationUpdateNoContent() *ApplicationUpdateNoContent {
	return &ApplicationUpdateNoContent{}
}

/*ApplicationUpdateNoContent handles this case with default header values.

The operation was successful.
*/
type ApplicationUpdateNoContent struct {
}

func (o *ApplicationUpdateNoContent) Error() string {
	return fmt.Sprintf("[PATCH /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}][%d] applicationUpdateNoContent ", 204)
}

func (o *ApplicationUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApplicationUpdateDefault creates a ApplicationUpdateDefault with default headers values
func NewApplicationUpdateDefault(code int) *ApplicationUpdateDefault {
	return &ApplicationUpdateDefault{
		_statusCode: code,
	}
}

/*ApplicationUpdateDefault handles this case with default header values.

Error response describing why the operation failed.
*/
type ApplicationUpdateDefault struct {
	_statusCode int

	Payload *models.CloudError
}

// Code gets the status code for the application update default response
func (o *ApplicationUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}][%d] Application_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
