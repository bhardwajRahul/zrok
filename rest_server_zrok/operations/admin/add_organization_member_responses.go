// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddOrganizationMemberCreatedCode is the HTTP code returned for type AddOrganizationMemberCreated
const AddOrganizationMemberCreatedCode int = 201

/*
AddOrganizationMemberCreated member added

swagger:response addOrganizationMemberCreated
*/
type AddOrganizationMemberCreated struct {
}

// NewAddOrganizationMemberCreated creates AddOrganizationMemberCreated with default headers values
func NewAddOrganizationMemberCreated() *AddOrganizationMemberCreated {

	return &AddOrganizationMemberCreated{}
}

// WriteResponse to the client
func (o *AddOrganizationMemberCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// AddOrganizationMemberUnauthorizedCode is the HTTP code returned for type AddOrganizationMemberUnauthorized
const AddOrganizationMemberUnauthorizedCode int = 401

/*
AddOrganizationMemberUnauthorized unauthorized

swagger:response addOrganizationMemberUnauthorized
*/
type AddOrganizationMemberUnauthorized struct {
}

// NewAddOrganizationMemberUnauthorized creates AddOrganizationMemberUnauthorized with default headers values
func NewAddOrganizationMemberUnauthorized() *AddOrganizationMemberUnauthorized {

	return &AddOrganizationMemberUnauthorized{}
}

// WriteResponse to the client
func (o *AddOrganizationMemberUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// AddOrganizationMemberNotFoundCode is the HTTP code returned for type AddOrganizationMemberNotFound
const AddOrganizationMemberNotFoundCode int = 404

/*
AddOrganizationMemberNotFound not found

swagger:response addOrganizationMemberNotFound
*/
type AddOrganizationMemberNotFound struct {
}

// NewAddOrganizationMemberNotFound creates AddOrganizationMemberNotFound with default headers values
func NewAddOrganizationMemberNotFound() *AddOrganizationMemberNotFound {

	return &AddOrganizationMemberNotFound{}
}

// WriteResponse to the client
func (o *AddOrganizationMemberNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// AddOrganizationMemberInternalServerErrorCode is the HTTP code returned for type AddOrganizationMemberInternalServerError
const AddOrganizationMemberInternalServerErrorCode int = 500

/*
AddOrganizationMemberInternalServerError internal server error

swagger:response addOrganizationMemberInternalServerError
*/
type AddOrganizationMemberInternalServerError struct {
}

// NewAddOrganizationMemberInternalServerError creates AddOrganizationMemberInternalServerError with default headers values
func NewAddOrganizationMemberInternalServerError() *AddOrganizationMemberInternalServerError {

	return &AddOrganizationMemberInternalServerError{}
}

// WriteResponse to the client
func (o *AddOrganizationMemberInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}