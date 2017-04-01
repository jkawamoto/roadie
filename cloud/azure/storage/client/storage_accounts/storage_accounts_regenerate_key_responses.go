package storage_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/storage/models"
)

// StorageAccountsRegenerateKeyReader is a Reader for the StorageAccountsRegenerateKey structure.
type StorageAccountsRegenerateKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageAccountsRegenerateKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageAccountsRegenerateKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageAccountsRegenerateKeyOK creates a StorageAccountsRegenerateKeyOK with default headers values
func NewStorageAccountsRegenerateKeyOK() *StorageAccountsRegenerateKeyOK {
	return &StorageAccountsRegenerateKeyOK{}
}

/*StorageAccountsRegenerateKeyOK handles this case with default header values.

OK -- specified key regenerated successfully.
*/
type StorageAccountsRegenerateKeyOK struct {
	Payload *models.StorageAccountListKeysResult
}

func (o *StorageAccountsRegenerateKeyOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/regenerateKey][%d] storageAccountsRegenerateKeyOK  %+v", 200, o.Payload)
}

func (o *StorageAccountsRegenerateKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageAccountListKeysResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
