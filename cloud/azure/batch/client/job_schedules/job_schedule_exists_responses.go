package job_schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/batch/models"
)

// JobScheduleExistsReader is a Reader for the JobScheduleExists structure.
type JobScheduleExistsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobScheduleExistsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJobScheduleExistsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewJobScheduleExistsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewJobScheduleExistsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewJobScheduleExistsOK creates a JobScheduleExistsOK with default headers values
func NewJobScheduleExistsOK() *JobScheduleExistsOK {
	return &JobScheduleExistsOK{}
}

/*JobScheduleExistsOK handles this case with default header values.

A response containing headers related to the job schedule, if it exists.
*/
type JobScheduleExistsOK struct {
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

func (o *JobScheduleExistsOK) Error() string {
	return fmt.Sprintf("[HEAD /jobschedules/{jobScheduleId}][%d] jobScheduleExistsOK ", 200)
}

func (o *JobScheduleExistsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewJobScheduleExistsNotFound creates a JobScheduleExistsNotFound with default headers values
func NewJobScheduleExistsNotFound() *JobScheduleExistsNotFound {
	return &JobScheduleExistsNotFound{}
}

/*JobScheduleExistsNotFound handles this case with default header values.

The job schedule does not exist.
*/
type JobScheduleExistsNotFound struct {
}

func (o *JobScheduleExistsNotFound) Error() string {
	return fmt.Sprintf("[HEAD /jobschedules/{jobScheduleId}][%d] jobScheduleExistsNotFound ", 404)
}

func (o *JobScheduleExistsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewJobScheduleExistsDefault creates a JobScheduleExistsDefault with default headers values
func NewJobScheduleExistsDefault(code int) *JobScheduleExistsDefault {
	return &JobScheduleExistsDefault{
		_statusCode: code,
	}
}

/*JobScheduleExistsDefault handles this case with default header values.

The error from the Batch service.
*/
type JobScheduleExistsDefault struct {
	_statusCode int

	Payload *models.BatchError
}

// Code gets the status code for the job schedule exists default response
func (o *JobScheduleExistsDefault) Code() int {
	return o._statusCode
}

func (o *JobScheduleExistsDefault) Error() string {
	return fmt.Sprintf("[HEAD /jobschedules/{jobScheduleId}][%d] JobSchedule_Exists default  %+v", o._statusCode, o.Payload)
}

func (o *JobScheduleExistsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
