package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// ModifyUserReader is a Reader for the ModifyUser structure.
type ModifyUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewModifyUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewModifyUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewModifyUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModifyUserOK creates a ModifyUserOK with default headers values
func NewModifyUserOK() *ModifyUserOK {
	return &ModifyUserOK{}
}

/*ModifyUserOK handles this case with default header values.

Successfully modified the specified user.
*/
type ModifyUserOK struct {
	Payload ModifyUserOKBody
}

func (o *ModifyUserOK) Error() string {
	return fmt.Sprintf("[PATCH /users/{name}][%d] modifyUserOK  %+v", 200, o.Payload)
}

func (o *ModifyUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyUserUnauthorized creates a ModifyUserUnauthorized with default headers values
func NewModifyUserUnauthorized() *ModifyUserUnauthorized {
	return &ModifyUserUnauthorized{}
}

/*ModifyUserUnauthorized handles this case with default header values.

Unauthorized
*/
type ModifyUserUnauthorized struct {
	Payload ModifyUserUnauthorizedBody
}

func (o *ModifyUserUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /users/{name}][%d] modifyUserUnauthorized  %+v", 401, o.Payload)
}

func (o *ModifyUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyUserForbidden creates a ModifyUserForbidden with default headers values
func NewModifyUserForbidden() *ModifyUserForbidden {
	return &ModifyUserForbidden{}
}

/*ModifyUserForbidden handles this case with default header values.

Forbidden
*/
type ModifyUserForbidden struct {
	Payload ModifyUserForbiddenBody
}

func (o *ModifyUserForbidden) Error() string {
	return fmt.Sprintf("[PATCH /users/{name}][%d] modifyUserForbidden  %+v", 403, o.Payload)
}

func (o *ModifyUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyUserDefault creates a ModifyUserDefault with default headers values
func NewModifyUserDefault(code int) *ModifyUserDefault {
	return &ModifyUserDefault{
		_statusCode: code,
	}
}

/*ModifyUserDefault handles this case with default header values.

Unexpected error
*/
type ModifyUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the modify user default response
func (o *ModifyUserDefault) Code() int {
	return o._statusCode
}

func (o *ModifyUserDefault) Error() string {
	return fmt.Sprintf("[PATCH /users/{name}][%d] modifyUser default  %+v", o._statusCode, o.Payload)
}

func (o *ModifyUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ModifyUserForbiddenBody modify user forbidden body
swagger:model ModifyUserForbiddenBody
*/
type ModifyUserForbiddenBody interface{}

/*ModifyUserOKBody modify user o k body
swagger:model ModifyUserOKBody
*/
type ModifyUserOKBody interface{}

/*ModifyUserUnauthorizedBody modify user unauthorized body
swagger:model ModifyUserUnauthorizedBody
*/
type ModifyUserUnauthorizedBody interface{}
