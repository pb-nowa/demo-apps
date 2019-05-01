// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// PostFinishHandlerFunc turns a function with the right signature into a post finish handler
type PostFinishHandlerFunc func(PostFinishParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostFinishHandlerFunc) Handle(params PostFinishParams) middleware.Responder {
	return fn(params)
}

// PostFinishHandler interface for that can handle valid post finish params
type PostFinishHandler interface {
	Handle(PostFinishParams) middleware.Responder
}

// NewPostFinish creates a new http.Handler for the post finish operation
func NewPostFinish(ctx *middleware.Context, handler PostFinishHandler) *PostFinish {
	return &PostFinish{Context: ctx, Handler: handler}
}

/*PostFinish swagger:route POST /finish postFinish

FE calls this to post a level

*/
type PostFinish struct {
	Context *middleware.Context
	Handler PostFinishHandler
}

func (o *PostFinish) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostFinishParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostFinishOKBody post finish o k body
// swagger:model PostFinishOKBody
type PostFinishOKBody struct {

	// reward amount, in wei (divide by 10^18 to get HRX)
	Reward float64 `json:"reward,omitempty"`
}

// Validate validates this post finish o k body
func (o *PostFinishOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostFinishOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostFinishOKBody) UnmarshalBinary(b []byte) error {
	var res PostFinishOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
