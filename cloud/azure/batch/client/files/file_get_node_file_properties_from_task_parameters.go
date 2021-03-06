package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFileGetNodeFilePropertiesFromTaskParams creates a new FileGetNodeFilePropertiesFromTaskParams object
// with the default values initialized.
func NewFileGetNodeFilePropertiesFromTaskParams() *FileGetNodeFilePropertiesFromTaskParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &FileGetNodeFilePropertiesFromTaskParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewFileGetNodeFilePropertiesFromTaskParamsWithTimeout creates a new FileGetNodeFilePropertiesFromTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFileGetNodeFilePropertiesFromTaskParamsWithTimeout(timeout time.Duration) *FileGetNodeFilePropertiesFromTaskParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &FileGetNodeFilePropertiesFromTaskParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewFileGetNodeFilePropertiesFromTaskParamsWithContext creates a new FileGetNodeFilePropertiesFromTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewFileGetNodeFilePropertiesFromTaskParamsWithContext(ctx context.Context) *FileGetNodeFilePropertiesFromTaskParams {
	var (
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &FileGetNodeFilePropertiesFromTaskParams{
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*FileGetNodeFilePropertiesFromTaskParams contains all the parameters to send to the API endpoint
for the file get node file properties from task operation typically these are written to a http.Request
*/
type FileGetNodeFilePropertiesFromTaskParams struct {

	/*IfModifiedSince
	  Specify this header to perform the operation only if the resource has been modified since the specified date/time.

	*/
	IfModifiedSince *string
	/*IfUnmodifiedSince
	  Specify this header to perform the operation only if the resource has not been modified since the specified date/time.

	*/
	IfUnmodifiedSince *string
	/*APIVersion
	  Client API Version.

	*/
	APIVersion string
	/*ClientRequestID
	  The caller-generated request identity, in the form of a GUID with no decoration such as curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.

	*/
	ClientRequestID *string
	/*FileName
	  The path to the task file that you want to get the properties of.

	*/
	FileName string
	/*JobID
	  The ID of the job that contains the task.

	*/
	JobID string
	/*OcpDate
	  The time the request was issued. If not specified, this header will be automatically populated with the current system clock time.

	*/
	OcpDate *string
	/*ReturnClientRequestID
	  Whether the server should return the client-request-id in the response.

	*/
	ReturnClientRequestID *bool
	/*TaskID
	  The ID of the task whose file you want to get the properties of.

	*/
	TaskID string
	/*Timeout
	  The maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.

	*/
	Timeout *int32

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithRequestTimeout(timeout time.Duration) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithContext(ctx context.Context) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIfModifiedSince adds the ifModifiedSince to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithIfModifiedSince(ifModifiedSince *string) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetIfModifiedSince(ifModifiedSince)
	return o
}

// SetIfModifiedSince adds the ifModifiedSince to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetIfModifiedSince(ifModifiedSince *string) {
	o.IfModifiedSince = ifModifiedSince
}

// WithIfUnmodifiedSince adds the ifUnmodifiedSince to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithIfUnmodifiedSince(ifUnmodifiedSince *string) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetIfUnmodifiedSince(ifUnmodifiedSince)
	return o
}

// SetIfUnmodifiedSince adds the ifUnmodifiedSince to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetIfUnmodifiedSince(ifUnmodifiedSince *string) {
	o.IfUnmodifiedSince = ifUnmodifiedSince
}

// WithAPIVersion adds the aPIVersion to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithAPIVersion(aPIVersion string) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithClientRequestID(clientRequestID *string) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithFileName adds the fileName to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithFileName(fileName string) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WithJobID adds the jobID to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithJobID(jobID string) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WithOcpDate adds the ocpDate to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithOcpDate(ocpDate *string) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithReturnClientRequestID adds the returnClientRequestID to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithReturnClientRequestID(returnClientRequestID *bool) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTaskID adds the taskID to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithTaskID(taskID string) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WithTimeout adds the timeout to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) WithTimeout(timeout *int32) *FileGetNodeFilePropertiesFromTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file get node file properties from task params
func (o *FileGetNodeFilePropertiesFromTaskParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *FileGetNodeFilePropertiesFromTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.requestTimeout)
	var res []error

	if o.IfModifiedSince != nil {

		// header param If-Modified-Since
		if err := r.SetHeaderParam("If-Modified-Since", *o.IfModifiedSince); err != nil {
			return err
		}

	}

	if o.IfUnmodifiedSince != nil {

		// header param If-Unmodified-Since
		if err := r.SetHeaderParam("If-Unmodified-Since", *o.IfUnmodifiedSince); err != nil {
			return err
		}

	}

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	if o.ClientRequestID != nil {

		// header param client-request-id
		if err := r.SetHeaderParam("client-request-id", *o.ClientRequestID); err != nil {
			return err
		}

	}

	// path param fileName
	if err := r.SetPathParam("fileName", o.FileName); err != nil {
		return err
	}

	// path param jobId
	if err := r.SetPathParam("jobId", o.JobID); err != nil {
		return err
	}

	if o.OcpDate != nil {

		// header param ocp-date
		if err := r.SetHeaderParam("ocp-date", *o.OcpDate); err != nil {
			return err
		}

	}

	if o.ReturnClientRequestID != nil {

		// header param return-client-request-id
		if err := r.SetHeaderParam("return-client-request-id", swag.FormatBool(*o.ReturnClientRequestID)); err != nil {
			return err
		}

	}

	// path param taskId
	if err := r.SetPathParam("taskId", o.TaskID); err != nil {
		return err
	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int32
		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt32(qrTimeout)
		if qTimeout != "" {
			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
