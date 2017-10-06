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

// PoolEvaluateAutoScaleReader is a Reader for the PoolEvaluateAutoScale structure.
type PoolEvaluateAutoScaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoolEvaluateAutoScaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPoolEvaluateAutoScaleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPoolEvaluateAutoScaleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPoolEvaluateAutoScaleOK creates a PoolEvaluateAutoScaleOK with default headers values
func NewPoolEvaluateAutoScaleOK() *PoolEvaluateAutoScaleOK {
	return &PoolEvaluateAutoScaleOK{}
}

/*PoolEvaluateAutoScaleOK handles this case with default header values.

A response containing the results of the autoscale evaluation.
*/
type PoolEvaluateAutoScaleOK struct {
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

	Payload *models.AutoScaleRun
}

func (o *PoolEvaluateAutoScaleOK) Error() string {
	return fmt.Sprintf("[POST /pools/{poolId}/evaluateautoscale][%d] poolEvaluateAutoScaleOK  %+v", 200, o.Payload)
}

func (o *PoolEvaluateAutoScaleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.AutoScaleRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPoolEvaluateAutoScaleDefault creates a PoolEvaluateAutoScaleDefault with default headers values
func NewPoolEvaluateAutoScaleDefault(code int) *PoolEvaluateAutoScaleDefault {
	return &PoolEvaluateAutoScaleDefault{
		_statusCode: code,
	}
}

/*PoolEvaluateAutoScaleDefault handles this case with default header values.

The error from the Batch service.
*/
type PoolEvaluateAutoScaleDefault struct {
	_statusCode int

	Payload *models.BatchError
}

// Code gets the status code for the pool evaluate auto scale default response
func (o *PoolEvaluateAutoScaleDefault) Code() int {
	return o._statusCode
}

func (o *PoolEvaluateAutoScaleDefault) Error() string {
	return fmt.Sprintf("[POST /pools/{poolId}/evaluateautoscale][%d] Pool_EvaluateAutoScale default  %+v", o._statusCode, o.Payload)
}

func (o *PoolEvaluateAutoScaleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
