// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ossf/scorecard-webapp/app/generated/models"
)

// GetResultOKCode is the HTTP code returned for type GetResultOK
const GetResultOKCode int = 200

/*GetResultOK A JSON object of the repository's ScorecardResult

swagger:response getResultOK
*/
type GetResultOK struct {

	/*
	  In: Body
	*/
	Payload *models.ScorecardResult `json:"body,omitempty"`
}

// NewGetResultOK creates GetResultOK with default headers values
func NewGetResultOK() *GetResultOK {

	return &GetResultOK{}
}

// WithPayload adds the payload to the get result o k response
func (o *GetResultOK) WithPayload(payload *models.ScorecardResult) *GetResultOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get result o k response
func (o *GetResultOK) SetPayload(payload *models.ScorecardResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResultOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetResultBadRequestCode is the HTTP code returned for type GetResultBadRequest
const GetResultBadRequestCode int = 400

/*GetResultBadRequest The request provided to the server was invalid

swagger:response getResultBadRequest
*/
type GetResultBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetResultBadRequest creates GetResultBadRequest with default headers values
func NewGetResultBadRequest() *GetResultBadRequest {

	return &GetResultBadRequest{}
}

// WithPayload adds the payload to the get result bad request response
func (o *GetResultBadRequest) WithPayload(payload *models.Error) *GetResultBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get result bad request response
func (o *GetResultBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResultBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetResultNotFoundCode is the HTTP code returned for type GetResultNotFound
const GetResultNotFoundCode int = 404

/*GetResultNotFound The content requested could not be found

swagger:response getResultNotFound
*/
type GetResultNotFound struct {
}

// NewGetResultNotFound creates GetResultNotFound with default headers values
func NewGetResultNotFound() *GetResultNotFound {

	return &GetResultNotFound{}
}

// WriteResponse to the client
func (o *GetResultNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*GetResultDefault There was an internal error in the server while processing the request

swagger:response getResultDefault
*/
type GetResultDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetResultDefault creates GetResultDefault with default headers values
func NewGetResultDefault(code int) *GetResultDefault {
	if code <= 0 {
		code = 500
	}

	return &GetResultDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get result default response
func (o *GetResultDefault) WithStatusCode(code int) *GetResultDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get result default response
func (o *GetResultDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get result default response
func (o *GetResultDefault) WithPayload(payload *models.Error) *GetResultDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get result default response
func (o *GetResultDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResultDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
