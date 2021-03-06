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

// ResourceGroupsExportTemplateReader is a Reader for the ResourceGroupsExportTemplate structure.
type ResourceGroupsExportTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceGroupsExportTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResourceGroupsExportTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResourceGroupsExportTemplateOK creates a ResourceGroupsExportTemplateOK with default headers values
func NewResourceGroupsExportTemplateOK() *ResourceGroupsExportTemplateOK {
	return &ResourceGroupsExportTemplateOK{}
}

/*ResourceGroupsExportTemplateOK handles this case with default header values.

OK - Returns the result of the export.
*/
type ResourceGroupsExportTemplateOK struct {
	Payload *models.ResourceGroupExportResult
}

func (o *ResourceGroupsExportTemplateOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/exportTemplate][%d] resourceGroupsExportTemplateOK  %+v", 200, o.Payload)
}

func (o *ResourceGroupsExportTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceGroupExportResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
