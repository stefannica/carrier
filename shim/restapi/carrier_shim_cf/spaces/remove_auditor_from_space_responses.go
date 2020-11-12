// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RemoveAuditorFromSpaceCreatedCode is the HTTP code returned for type RemoveAuditorFromSpaceCreated
const RemoveAuditorFromSpaceCreatedCode int = 201

/*RemoveAuditorFromSpaceCreated successful response

swagger:response removeAuditorFromSpaceCreated
*/
type RemoveAuditorFromSpaceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RemoveAuditorFromSpaceResponse `json:"body,omitempty"`
}

// NewRemoveAuditorFromSpaceCreated creates RemoveAuditorFromSpaceCreated with default headers values
func NewRemoveAuditorFromSpaceCreated() *RemoveAuditorFromSpaceCreated {

	return &RemoveAuditorFromSpaceCreated{}
}

// WithPayload adds the payload to the remove auditor from space created response
func (o *RemoveAuditorFromSpaceCreated) WithPayload(payload *models.RemoveAuditorFromSpaceResponse) *RemoveAuditorFromSpaceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove auditor from space created response
func (o *RemoveAuditorFromSpaceCreated) SetPayload(payload *models.RemoveAuditorFromSpaceResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAuditorFromSpaceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}