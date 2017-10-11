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

// TagsCreateOrUpdateReader is a Reader for the TagsCreateOrUpdate structure.
type TagsCreateOrUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TagsCreateOrUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTagsCreateOrUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewTagsCreateOrUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTagsCreateOrUpdateOK creates a TagsCreateOrUpdateOK with default headers values
func NewTagsCreateOrUpdateOK() *TagsCreateOrUpdateOK {
	return &TagsCreateOrUpdateOK{}
}

/*TagsCreateOrUpdateOK handles this case with default header values.

OK - Returns information about the tag.
*/
type TagsCreateOrUpdateOK struct {
	Payload *models.TagDetails
}

func (o *TagsCreateOrUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/tagNames/{tagName}][%d] tagsCreateOrUpdateOK  %+v", 200, o.Payload)
}

func (o *TagsCreateOrUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTagsCreateOrUpdateCreated creates a TagsCreateOrUpdateCreated with default headers values
func NewTagsCreateOrUpdateCreated() *TagsCreateOrUpdateCreated {
	return &TagsCreateOrUpdateCreated{}
}

/*TagsCreateOrUpdateCreated handles this case with default header values.

Created - Returns information about the tag.
*/
type TagsCreateOrUpdateCreated struct {
	Payload *models.TagDetails
}

func (o *TagsCreateOrUpdateCreated) Error() string {
	return fmt.Sprintf("[PUT /subscriptions/{subscriptionId}/tagNames/{tagName}][%d] tagsCreateOrUpdateCreated  %+v", 201, o.Payload)
}

func (o *TagsCreateOrUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}