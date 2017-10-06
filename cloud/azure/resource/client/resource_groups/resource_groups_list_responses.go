package resource_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/resource/models"
)

// ResourceGroupsListReader is a Reader for the ResourceGroupsList structure.
type ResourceGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResourceGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResourceGroupsListOK creates a ResourceGroupsListOK with default headers values
func NewResourceGroupsListOK() *ResourceGroupsListOK {
	return &ResourceGroupsListOK{}
}

/*ResourceGroupsListOK handles this case with default header values.

OK - Returns an array of resource groups.
*/
type ResourceGroupsListOK struct {
	Payload *models.ResourceGroupListResult
}

func (o *ResourceGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscriptionId}/resourcegroups][%d] resourceGroupsListOK  %+v", 200, o.Payload)
}

func (o *ResourceGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceGroupListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
