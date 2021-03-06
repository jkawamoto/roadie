package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ResourcesCheckExistenceByIDReader is a Reader for the ResourcesCheckExistenceByID structure.
type ResourcesCheckExistenceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourcesCheckExistenceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewResourcesCheckExistenceByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewResourcesCheckExistenceByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResourcesCheckExistenceByIDNoContent creates a ResourcesCheckExistenceByIDNoContent with default headers values
func NewResourcesCheckExistenceByIDNoContent() *ResourcesCheckExistenceByIDNoContent {
	return &ResourcesCheckExistenceByIDNoContent{}
}

/*ResourcesCheckExistenceByIDNoContent handles this case with default header values.

No Content
*/
type ResourcesCheckExistenceByIDNoContent struct {
}

func (o *ResourcesCheckExistenceByIDNoContent) Error() string {
	return fmt.Sprintf("[HEAD /{resourceId}][%d] resourcesCheckExistenceByIdNoContent ", 204)
}

func (o *ResourcesCheckExistenceByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResourcesCheckExistenceByIDNotFound creates a ResourcesCheckExistenceByIDNotFound with default headers values
func NewResourcesCheckExistenceByIDNotFound() *ResourcesCheckExistenceByIDNotFound {
	return &ResourcesCheckExistenceByIDNotFound{}
}

/*ResourcesCheckExistenceByIDNotFound handles this case with default header values.

Not Found
*/
type ResourcesCheckExistenceByIDNotFound struct {
}

func (o *ResourcesCheckExistenceByIDNotFound) Error() string {
	return fmt.Sprintf("[HEAD /{resourceId}][%d] resourcesCheckExistenceByIdNotFound ", 404)
}

func (o *ResourcesCheckExistenceByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
