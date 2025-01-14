// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ChangePassword(params *ChangePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChangePasswordOK, error)

	Invite(params *InviteParams, opts ...ClientOption) (*InviteCreated, error)

	Login(params *LoginParams, opts ...ClientOption) (*LoginOK, error)

	RegenerateToken(params *RegenerateTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegenerateTokenOK, error)

	Register(params *RegisterParams, opts ...ClientOption) (*RegisterOK, error)

	ResetPassword(params *ResetPasswordParams, opts ...ClientOption) (*ResetPasswordOK, error)

	ResetPasswordRequest(params *ResetPasswordRequestParams, opts ...ClientOption) (*ResetPasswordRequestCreated, error)

	Verify(params *VerifyParams, opts ...ClientOption) (*VerifyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ChangePassword change password API
*/
func (a *Client) ChangePassword(params *ChangePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChangePasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "changePassword",
		Method:             "POST",
		PathPattern:        "/changePassword",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChangePasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangePasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for changePassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Invite invite API
*/
func (a *Client) Invite(params *InviteParams, opts ...ClientOption) (*InviteCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInviteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "invite",
		Method:             "POST",
		PathPattern:        "/invite",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InviteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InviteCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for invite: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Login login API
*/
func (a *Client) Login(params *LoginParams, opts ...ClientOption) (*LoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "login",
		Method:             "POST",
		PathPattern:        "/login",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for login: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RegenerateToken regenerate token API
*/
func (a *Client) RegenerateToken(params *RegenerateTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegenerateTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegenerateTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "regenerateToken",
		Method:             "POST",
		PathPattern:        "/regenerateToken",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegenerateTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegenerateTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for regenerateToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Register register API
*/
func (a *Client) Register(params *RegisterParams, opts ...ClientOption) (*RegisterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "register",
		Method:             "POST",
		PathPattern:        "/register",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegisterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for register: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ResetPassword reset password API
*/
func (a *Client) ResetPassword(params *ResetPasswordParams, opts ...ClientOption) (*ResetPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetPassword",
		Method:             "POST",
		PathPattern:        "/resetPassword",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResetPasswordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResetPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resetPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ResetPasswordRequest reset password request API
*/
func (a *Client) ResetPasswordRequest(params *ResetPasswordRequestParams, opts ...ClientOption) (*ResetPasswordRequestCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetPasswordRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetPasswordRequest",
		Method:             "POST",
		PathPattern:        "/resetPasswordRequest",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResetPasswordRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResetPasswordRequestCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resetPasswordRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Verify verify API
*/
func (a *Client) Verify(params *VerifyParams, opts ...ClientOption) (*VerifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "verify",
		Method:             "POST",
		PathPattern:        "/verify",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VerifyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VerifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for verify: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}