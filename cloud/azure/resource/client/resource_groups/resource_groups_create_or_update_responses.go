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

// ResourceGroupsCreateOrUpdateReader is a Reader for the ResourceGroupsCreateOrUpdate structure.
type ResourceGroupsCreateOrUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceGroupsCreateOrUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResourceGroupsCreateOrUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewResourceGroupsCreateOrUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResourceGroupsCreateOrUpdateOK creates a ResourceGroupsCreateOrUpdateOK with default headers values
func NewResourceGroupsCreateOrUpdateOK() *ResourceGroupsCreateOrUpdateOK {
	return &ResourceGroupsCreateOrUpdateOK{}
}

/*ResourceGroupsCreateOrUpdateOK handles this case with default header values.

OK - Returns information about the new resource group.
*/
type ResourceGroupsCreateOrUpdateOK struct {
	Payload *models.ResourceGroup
}

func (o *ResourceGroupsCreateOrUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}][%d] resourceGroupsCreateOrUpdateOK  %+v", 200, o.Payload)
}

func (o *ResourceGroupsCreateOrUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupsCreateOrUpdateCreated creates a ResourceGroupsCreateOrUpdateCreated with default headers values
func NewResourceGroupsCreateOrUpdateCreated() *ResourceGroupsCreateOrUpdateCreated {
	return &ResourceGroupsCreateOrUpdateCreated{}
}

/*ResourceGroupsCreateOrUpdateCreated handles this case with default header values.

Created - Returns information about the new resource group.
*/
type ResourceGroupsCreateOrUpdateCreated struct {
	Payload *models.ResourceGroup
}

func (o *ResourceGroupsCreateOrUpdateCreated) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}][%d] resourceGroupsCreateOrUpdateCreated  %+v", 201, o.Payload)
}

func (o *ResourceGroupsCreateOrUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}