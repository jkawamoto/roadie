package compute_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/batch/models"
)

// PoolRemoveNodesReader is a Reader for the PoolRemoveNodes structure.
type PoolRemoveNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoolRemoveNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewPoolRemoveNodesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPoolRemoveNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPoolRemoveNodesAccepted creates a PoolRemoveNodesAccepted with default headers values
func NewPoolRemoveNodesAccepted() *PoolRemoveNodesAccepted {
	return &PoolRemoveNodesAccepted{}
}

/*PoolRemoveNodesAccepted handles this case with default header values.

The request to the Batch service was successful.
*/
type PoolRemoveNodesAccepted struct {
	/*The OData ID of the resource to which the request applied.
	 */
	DataServiceID string
	/*The ETag HTTP response header. This is an opaque string. You can use it to detect whether the resource has changed between requests. In particular, you can pass the ETag to one of the If-Modified-Since, If-Unmodified-Since, If-Match or If-None-Match headers.
	 */
	ETag string
	/*The time at which the resource was last modified.
	 */
	LastModified string
	/*The client-request-id provided by the client during the request. This will be returned only if the return-client-request-id parameter was set to true.
	 */
	ClientRequestID string
	/*This header uniquely identifies the request that was made and can be used for troubleshooting the request. If a request is consistently failing and you have verified that the request is properly formulated, you may use this value to report the error to Microsoft. In your report, include the value of this header, the approximate time that the request was made, the Batch account against which the request was made, and the region that account resides in.
	 */
	RequestID string
}

func (o *PoolRemoveNodesAccepted) Error() string {
	return fmt.Sprintf("[POST /pools/{poolId}/removenodes][%d] poolRemoveNodesAccepted ", 202)
}

func (o *PoolRemoveNodesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header DataServiceId
	o.DataServiceID = response.GetHeader("DataServiceId")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response header client-request-id
	o.ClientRequestID = response.GetHeader("client-request-id")

	// response header request-id
	o.RequestID = response.GetHeader("request-id")

	return nil
}

// NewPoolRemoveNodesDefault creates a PoolRemoveNodesDefault with default headers values
func NewPoolRemoveNodesDefault(code int) *PoolRemoveNodesDefault {
	return &PoolRemoveNodesDefault{
		_statusCode: code,
	}
}

/*PoolRemoveNodesDefault handles this case with default header values.

The error from the Batch service.
*/
type PoolRemoveNodesDefault struct {
	_statusCode int

	Payload *models.BatchError
}

// Code gets the status code for the pool remove nodes default response
func (o *PoolRemoveNodesDefault) Code() int {
	return o._statusCode
}

func (o *PoolRemoveNodesDefault) Error() string {
	return fmt.Sprintf("[POST /pools/{poolId}/removenodes][%d] Pool_RemoveNodes default  %+v", o._statusCode, o.Payload)
}

func (o *PoolRemoveNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
