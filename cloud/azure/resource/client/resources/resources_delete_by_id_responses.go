package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ResourcesDeleteByIDReader is a Reader for the ResourcesDeleteByID structure.
type ResourcesDeleteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourcesDeleteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResourcesDeleteByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewResourcesDeleteByIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewResourcesDeleteByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResourcesDeleteByIDOK creates a ResourcesDeleteByIDOK with default headers values
func NewResourcesDeleteByIDOK() *ResourcesDeleteByIDOK {
	return &ResourcesDeleteByIDOK{}
}

/*ResourcesDeleteByIDOK handles this case with default header values.

OK
*/
type ResourcesDeleteByIDOK struct {
}

func (o *ResourcesDeleteByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /{resourceId}][%d] resourcesDeleteByIdOK ", 200)
}

func (o *ResourcesDeleteByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResourcesDeleteByIDAccepted creates a ResourcesDeleteByIDAccepted with default headers values
func NewResourcesDeleteByIDAccepted() *ResourcesDeleteByIDAccepted {
	return &ResourcesDeleteByIDAccepted{}
}

/*ResourcesDeleteByIDAccepted handles this case with default header values.

Accepted
*/
type ResourcesDeleteByIDAccepted struct {
}

func (o *ResourcesDeleteByIDAccepted) Error() string {
	return fmt.Sprintf("[DELETE /{resourceId}][%d] resourcesDeleteByIdAccepted ", 202)
}

func (o *ResourcesDeleteByIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResourcesDeleteByIDNoContent creates a ResourcesDeleteByIDNoContent with default headers values
func NewResourcesDeleteByIDNoContent() *ResourcesDeleteByIDNoContent {
	return &ResourcesDeleteByIDNoContent{}
}

/*ResourcesDeleteByIDNoContent handles this case with default header values.

No Content
*/
type ResourcesDeleteByIDNoContent struct {
}

func (o *ResourcesDeleteByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{resourceId}][%d] resourcesDeleteByIdNoContent ", 204)
}

func (o *ResourcesDeleteByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
