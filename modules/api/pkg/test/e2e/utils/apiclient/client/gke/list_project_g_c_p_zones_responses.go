// Code generated by go-swagger; DO NOT EDIT.

package gke

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListProjectGCPZonesReader is a Reader for the ListProjectGCPZones structure.
type ListProjectGCPZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectGCPZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectGCPZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectGCPZonesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectGCPZonesOK creates a ListProjectGCPZonesOK with default headers values
func NewListProjectGCPZonesOK() *ListProjectGCPZonesOK {
	return &ListProjectGCPZonesOK{}
}

/*
ListProjectGCPZonesOK describes a response with status code 200, with default header values.

GCPZoneList
*/
type ListProjectGCPZonesOK struct {
	Payload models.GCPZoneList
}

// IsSuccess returns true when this list project g c p zones o k response has a 2xx status code
func (o *ListProjectGCPZonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project g c p zones o k response has a 3xx status code
func (o *ListProjectGCPZonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project g c p zones o k response has a 4xx status code
func (o *ListProjectGCPZonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project g c p zones o k response has a 5xx status code
func (o *ListProjectGCPZonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project g c p zones o k response a status code equal to that given
func (o *ListProjectGCPZonesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectGCPZonesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/zones][%d] listProjectGCPZonesOK  %+v", 200, o.Payload)
}

func (o *ListProjectGCPZonesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/zones][%d] listProjectGCPZonesOK  %+v", 200, o.Payload)
}

func (o *ListProjectGCPZonesOK) GetPayload() models.GCPZoneList {
	return o.Payload
}

func (o *ListProjectGCPZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectGCPZonesDefault creates a ListProjectGCPZonesDefault with default headers values
func NewListProjectGCPZonesDefault(code int) *ListProjectGCPZonesDefault {
	return &ListProjectGCPZonesDefault{
		_statusCode: code,
	}
}

/*
ListProjectGCPZonesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectGCPZonesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project g c p zones default response
func (o *ListProjectGCPZonesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project g c p zones default response has a 2xx status code
func (o *ListProjectGCPZonesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project g c p zones default response has a 3xx status code
func (o *ListProjectGCPZonesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project g c p zones default response has a 4xx status code
func (o *ListProjectGCPZonesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project g c p zones default response has a 5xx status code
func (o *ListProjectGCPZonesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project g c p zones default response a status code equal to that given
func (o *ListProjectGCPZonesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectGCPZonesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/zones][%d] listProjectGCPZones default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectGCPZonesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/gcp/zones][%d] listProjectGCPZones default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectGCPZonesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectGCPZonesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}