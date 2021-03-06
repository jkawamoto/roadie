package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/batch/models"
)

// FileGetNodeFilePropertiesFromTaskReader is a Reader for the FileGetNodeFilePropertiesFromTask structure.
type FileGetNodeFilePropertiesFromTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileGetNodeFilePropertiesFromTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFileGetNodeFilePropertiesFromTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFileGetNodeFilePropertiesFromTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileGetNodeFilePropertiesFromTaskOK creates a FileGetNodeFilePropertiesFromTaskOK with default headers values
func NewFileGetNodeFilePropertiesFromTaskOK() *FileGetNodeFilePropertiesFromTaskOK {
	return &FileGetNodeFilePropertiesFromTaskOK{}
}

/*FileGetNodeFilePropertiesFromTaskOK handles this case with default header values.

A response containing the file properties.
*/
type FileGetNodeFilePropertiesFromTaskOK struct {
	/*The length of the file.
	 */
	ContentLength int64
	/*The content type of the file.
	 */
	ContentType string
	/*The ETag HTTP response header. This is an opaque string. You can use it to detect whether the resource has changed between requests. In particular, you can pass the ETag to one of the If-Modified-Since, If-Unmodified-Since, If-Match or If-None-Match headers.
	 */
	ETag string
	/*The time at which the resource was last modified.
	 */
	LastModified string
	/*The client-request-id provided by the client during the request. This will be returned only if the return-client-request-id parameter was set to true.
	 */
	ClientRequestID string
	/*Whether the object represents a directory.
	 */
	OcpBatchFileIsdirectory bool
	/*The file mode attribute in octal format.
	 */
	OcpBatchFileMode string
	/*The URL of the file.
	 */
	OcpBatchFileURL string
	/*The file creation time.
	 */
	OcpCreationTime string
	/*This header uniquely identifies the request that was made and can be used for troubleshooting the request. If a request is consistently failing and you have verified that the request is properly formulated, you may use this value to report the error to Microsoft. In your report, include the value of this header, the approximate time that the request was made, the Batch account against which the request was made, and the region that account resides in.
	 */
	RequestID string
}

func (o *FileGetNodeFilePropertiesFromTaskOK) Error() string {
	return fmt.Sprintf("[HEAD /jobs/{jobId}/tasks/{taskId}/files/{fileName}][%d] fileGetNodeFilePropertiesFromTaskOK ", 200)
}

func (o *FileGetNodeFilePropertiesFromTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	// response header Content-Type
	o.ContentType = response.GetHeader("Content-Type")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response header client-request-id
	o.ClientRequestID = response.GetHeader("client-request-id")

	// response header ocp-batch-file-isdirectory
	ocpBatchFileIsdirectory, err := swag.ConvertBool(response.GetHeader("ocp-batch-file-isdirectory"))
	if err != nil {
		return errors.InvalidType("ocp-batch-file-isdirectory", "header", "bool", response.GetHeader("ocp-batch-file-isdirectory"))
	}
	o.OcpBatchFileIsdirectory = ocpBatchFileIsdirectory

	// response header ocp-batch-file-mode
	o.OcpBatchFileMode = response.GetHeader("ocp-batch-file-mode")

	// response header ocp-batch-file-url
	o.OcpBatchFileURL = response.GetHeader("ocp-batch-file-url")

	// response header ocp-creation-time
	o.OcpCreationTime = response.GetHeader("ocp-creation-time")

	// response header request-id
	o.RequestID = response.GetHeader("request-id")

	return nil
}

// NewFileGetNodeFilePropertiesFromTaskDefault creates a FileGetNodeFilePropertiesFromTaskDefault with default headers values
func NewFileGetNodeFilePropertiesFromTaskDefault(code int) *FileGetNodeFilePropertiesFromTaskDefault {
	return &FileGetNodeFilePropertiesFromTaskDefault{
		_statusCode: code,
	}
}

/*FileGetNodeFilePropertiesFromTaskDefault handles this case with default header values.

The error from the Batch service.
*/
type FileGetNodeFilePropertiesFromTaskDefault struct {
	_statusCode int

	Payload *models.BatchError
}

// Code gets the status code for the file get node file properties from task default response
func (o *FileGetNodeFilePropertiesFromTaskDefault) Code() int {
	return o._statusCode
}

func (o *FileGetNodeFilePropertiesFromTaskDefault) Error() string {
	return fmt.Sprintf("[HEAD /jobs/{jobId}/tasks/{taskId}/files/{fileName}][%d] File_GetNodeFilePropertiesFromTask default  %+v", o._statusCode, o.Payload)
}

func (o *FileGetNodeFilePropertiesFromTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
