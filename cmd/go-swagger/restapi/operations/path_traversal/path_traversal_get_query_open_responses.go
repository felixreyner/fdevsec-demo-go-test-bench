// Code generated by go-swagger; DO NOT EDIT.

package path_traversal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PathTraversalGetQueryOpenOKCode is the HTTP code returned for type PathTraversalGetQueryOpenOK
const PathTraversalGetQueryOpenOKCode int = 200

/*PathTraversalGetQueryOpenOK returns the rendered response as a string

swagger:response pathTraversalGetQueryOpenOK
*/
type PathTraversalGetQueryOpenOK struct {

	/*The response when succesful query happens
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPathTraversalGetQueryOpenOK creates PathTraversalGetQueryOpenOK with default headers values
func NewPathTraversalGetQueryOpenOK() *PathTraversalGetQueryOpenOK {

	return &PathTraversalGetQueryOpenOK{}
}

// WithPayload adds the payload to the path traversal get query open o k response
func (o *PathTraversalGetQueryOpenOK) WithPayload(payload string) *PathTraversalGetQueryOpenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the path traversal get query open o k response
func (o *PathTraversalGetQueryOpenOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PathTraversalGetQueryOpenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PathTraversalGetQueryOpenDefault Error occured

swagger:response pathTraversalGetQueryOpenDefault
*/
type PathTraversalGetQueryOpenDefault struct {
	_statusCode int
}

// NewPathTraversalGetQueryOpenDefault creates PathTraversalGetQueryOpenDefault with default headers values
func NewPathTraversalGetQueryOpenDefault(code int) *PathTraversalGetQueryOpenDefault {
	if code <= 0 {
		code = 500
	}

	return &PathTraversalGetQueryOpenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the path traversal get query open default response
func (o *PathTraversalGetQueryOpenDefault) WithStatusCode(code int) *PathTraversalGetQueryOpenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the path traversal get query open default response
func (o *PathTraversalGetQueryOpenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *PathTraversalGetQueryOpenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
