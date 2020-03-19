// Code generated by go-swagger; DO NOT EDIT.

package iaas_service_broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new iaas service broker API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iaas service broker API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ServiceBrokerHealth gets current server health
*/
func (a *Client) ServiceBrokerHealth(params *ServiceBrokerHealthParams) (*ServiceBrokerHealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerHealthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.health",
		Method:             "GET",
		PathPattern:        "/broker/v1/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerHealthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerHealthOK), nil

}

/*
ServiceBrokerHealthHead gets current server health
*/
func (a *Client) ServiceBrokerHealthHead(params *ServiceBrokerHealthHeadParams) (*ServiceBrokerHealthHeadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerHealthHeadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.health.head",
		Method:             "HEAD",
		PathPattern:        "/broker/v1/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerHealthHeadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerHealthHeadOK), nil

}

/*
ServiceBrokerVersion gets current server version
*/
func (a *Client) ServiceBrokerVersion(params *ServiceBrokerVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBrokerVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.version",
		Method:             "GET",
		PathPattern:        "/broker/v1/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerVersionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}