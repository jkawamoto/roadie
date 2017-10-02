package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/batch/models"
)

// PoolDeleteReader is a Reader for the PoolDelete structure.
type PoolDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoolDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewPoolDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPoolDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPoolDeleteAccepted creates a PoolDeleteAccepted with default headers values
func NewPoolDeleteAccepted() *PoolDeleteAccepted {
	return &PoolDeleteAccepted{}
}

/*PoolDeleteAccepted handles this case with default header values.

The request to the Batch service was successful.
*/
type PoolDeleteAccepted struct {
	/*The client-request-id provided by the client during the request. This will be returned only if the return-client-request-id parameter was set to true.
	 */
	ClientRequestID string
	/*This header uniquely identifies the request that was made and can be used for troubleshooting the request. If a request is consistently failing and you have verified that the request is properly formulated, you may use this value to report the error to Microsoft. In your report, include the value of this header, the approximate time that the request was made, the Batch account against which the request was made, and the region that account resides in.
	 */
	RequestID string
}

func (o *PoolDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /pools/{poolId}][%d] poolDeleteAccepted ", 202)
}

func (o *PoolDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header client-request-id
	o.ClientRequestID = response.GetHeader("client-request-id")

	// response header request-id
	o.RequestID = response.GetHeader("request-id")

	return nil
}

// NewPoolDeleteDefault creates a PoolDeleteDefault with default headers values
func NewPoolDeleteDefault(code int) *PoolDeleteDefault {
	return &PoolDeleteDefault{
		_statusCode: code,
	}
}

/*PoolDeleteDefault handles this case with default header values.

The error from the Batch service.
*/
type PoolDeleteDefault struct {
	_statusCode int

	Payload *models.BatchError
}

// Code gets the status code for the pool delete default response
func (o *PoolDeleteDefault) Code() int {
	return o._statusCode
}

func (o *PoolDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /pools/{poolId}][%d] Pool_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *PoolDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
