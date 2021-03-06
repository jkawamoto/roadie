package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// TagsDeleteValueReader is a Reader for the TagsDeleteValue structure.
type TagsDeleteValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TagsDeleteValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTagsDeleteValueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewTagsDeleteValueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTagsDeleteValueOK creates a TagsDeleteValueOK with default headers values
func NewTagsDeleteValueOK() *TagsDeleteValueOK {
	return &TagsDeleteValueOK{}
}

/*TagsDeleteValueOK handles this case with default header values.

OK
*/
type TagsDeleteValueOK struct {
}

func (o *TagsDeleteValueOK) Error() string {
	return fmt.Sprintf("[DELETE /subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}][%d] tagsDeleteValueOK ", 200)
}

func (o *TagsDeleteValueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTagsDeleteValueNoContent creates a TagsDeleteValueNoContent with default headers values
func NewTagsDeleteValueNoContent() *TagsDeleteValueNoContent {
	return &TagsDeleteValueNoContent{}
}

/*TagsDeleteValueNoContent handles this case with default header values.

No Content
*/
type TagsDeleteValueNoContent struct {
}

func (o *TagsDeleteValueNoContent) Error() string {
	return fmt.Sprintf("[DELETE /subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}][%d] tagsDeleteValueNoContent ", 204)
}

func (o *TagsDeleteValueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
