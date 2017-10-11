package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/batch/models"
)

// FileListFromTaskReader is a Reader for the FileListFromTask structure.
type FileListFromTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileListFromTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFileListFromTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFileListFromTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileListFromTaskOK creates a FileListFromTaskOK with default headers values
func NewFileListFromTaskOK() *FileListFromTaskOK {
	return &FileListFromTaskOK{}
}

/*FileListFromTaskOK handles this case with default header values.

A response containing the list of files.
*/
type FileListFromTaskOK struct {
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

	Payload *models.NodeFileListResult
}

func (o *FileListFromTaskOK) Error() string {
	return fmt.Sprintf("[GET /jobs/{jobId}/tasks/{taskId}/files][%d] fileListFromTaskOK  %+v", 200, o.Payload)
}

func (o *FileListFromTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response header client-request-id
	o.ClientRequestID = response.GetHeader("client-request-id")

	// response header request-id
	o.RequestID = response.GetHeader("request-id")

	o.Payload = new(models.NodeFileListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileListFromTaskDefault creates a FileListFromTaskDefault with default headers values
func NewFileListFromTaskDefault(code int) *FileListFromTaskDefault {
	return &FileListFromTaskDefault{
		_statusCode: code,
	}
}

/*FileListFromTaskDefault handles this case with default header values.

The error from the Batch service.
*/
type FileListFromTaskDefault struct {
	_statusCode int

	Payload *models.BatchError
}

// Code gets the status code for the file list from task default response
func (o *FileListFromTaskDefault) Code() int {
	return o._statusCode
}

func (o *FileListFromTaskDefault) Error() string {
	return fmt.Sprintf("[GET /jobs/{jobId}/tasks/{taskId}/files][%d] File_ListFromTask default  %+v", o._statusCode, o.Payload)
}

func (o *FileListFromTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}