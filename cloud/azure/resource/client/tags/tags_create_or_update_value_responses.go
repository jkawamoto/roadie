package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/resource/models"
)

// TagsCreateOrUpdateValueReader is a Reader for the TagsCreateOrUpdateValue structure.
type TagsCreateOrUpdateValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TagsCreateOrUpdateValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTagsCreateOrUpdateValueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewTagsCreateOrUpdateValueCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTagsCreateOrUpdateValueOK creates a TagsCreateOrUpdateValueOK with default headers values
func NewTagsCreateOrUpdateValueOK() *TagsCreateOrUpdateValueOK {
	return &TagsCreateOrUpdateValueOK{}
}

/*TagsCreateOrUpdateValueOK handles this case with default header values.

OK - Returns information about the tag value.
*/
type TagsCreateOrUpdateValueOK struct {
	Payload *models.TagValue
}

func (o *TagsCreateOrUpdateValueOK) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}][%d] tagsCreateOrUpdateValueOK  %+v", 200, o.Payload)
}

func (o *TagsCreateOrUpdateValueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagsCreateOrUpdateValueCreated creates a TagsCreateOrUpdateValueCreated with default headers values
func NewTagsCreateOrUpdateValueCreated() *TagsCreateOrUpdateValueCreated {
	return &TagsCreateOrUpdateValueCreated{}
}

/*TagsCreateOrUpdateValueCreated handles this case with default header values.

Created - Returns information about the tag value.
*/
type TagsCreateOrUpdateValueCreated struct {
	Payload *models.TagValue
}

func (o *TagsCreateOrUpdateValueCreated) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}][%d] tagsCreateOrUpdateValueCreated  %+v", 201, o.Payload)
}

func (o *TagsCreateOrUpdateValueCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
