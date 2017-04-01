package virtual_network_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VirtualNetworkGatewaysGeneratevpnclientpackageReader is a Reader for the VirtualNetworkGatewaysGeneratevpnclientpackage structure.
type VirtualNetworkGatewaysGeneratevpnclientpackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualNetworkGatewaysGeneratevpnclientpackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewVirtualNetworkGatewaysGeneratevpnclientpackageAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualNetworkGatewaysGeneratevpnclientpackageAccepted creates a VirtualNetworkGatewaysGeneratevpnclientpackageAccepted with default headers values
func NewVirtualNetworkGatewaysGeneratevpnclientpackageAccepted() *VirtualNetworkGatewaysGeneratevpnclientpackageAccepted {
	return &VirtualNetworkGatewaysGeneratevpnclientpackageAccepted{}
}

/*VirtualNetworkGatewaysGeneratevpnclientpackageAccepted handles this case with default header values.

Vpn client package url
*/
type VirtualNetworkGatewaysGeneratevpnclientpackageAccepted struct {
	Payload string
}

func (o *VirtualNetworkGatewaysGeneratevpnclientpackageAccepted) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/generatevpnclientpackage][%d] virtualNetworkGatewaysGeneratevpnclientpackageAccepted  %+v", 202, o.Payload)
}

func (o *VirtualNetworkGatewaysGeneratevpnclientpackageAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
