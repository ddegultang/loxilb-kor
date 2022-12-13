// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigTunnelVxlanAllHandlerFunc turns a function with the right signature into a get config tunnel vxlan all handler
type GetConfigTunnelVxlanAllHandlerFunc func(GetConfigTunnelVxlanAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigTunnelVxlanAllHandlerFunc) Handle(params GetConfigTunnelVxlanAllParams) middleware.Responder {
	return fn(params)
}

// GetConfigTunnelVxlanAllHandler interface for that can handle valid get config tunnel vxlan all params
type GetConfigTunnelVxlanAllHandler interface {
	Handle(GetConfigTunnelVxlanAllParams) middleware.Responder
}

// NewGetConfigTunnelVxlanAll creates a new http.Handler for the get config tunnel vxlan all operation
func NewGetConfigTunnelVxlanAll(ctx *middleware.Context, handler GetConfigTunnelVxlanAllHandler) *GetConfigTunnelVxlanAll {
	return &GetConfigTunnelVxlanAll{Context: ctx, Handler: handler}
}

/*
	GetConfigTunnelVxlanAll swagger:route GET /config/tunnel/vxlan/all getConfigTunnelVxlanAll

# Get a list of vxlan configurations

Return a list of existing tunnels of a type. If there're no tunnels to return, empty list will be returned.
*/
type GetConfigTunnelVxlanAll struct {
	Context *middleware.Context
	Handler GetConfigTunnelVxlanAllHandler
}

func (o *GetConfigTunnelVxlanAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigTunnelVxlanAllParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetConfigTunnelVxlanAllOKBody get config tunnel vxlan all o k body
//
// swagger:model GetConfigTunnelVxlanAllOKBody
type GetConfigTunnelVxlanAllOKBody struct {

	// vxlan attr
	VxlanAttr []*models.VxlanEntry `json:"vxlanAttr"`
}

// Validate validates this get config tunnel vxlan all o k body
func (o *GetConfigTunnelVxlanAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVxlanAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigTunnelVxlanAllOKBody) validateVxlanAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.VxlanAttr) { // not required
		return nil
	}

	for i := 0; i < len(o.VxlanAttr); i++ {
		if swag.IsZero(o.VxlanAttr[i]) { // not required
			continue
		}

		if o.VxlanAttr[i] != nil {
			if err := o.VxlanAttr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigTunnelVxlanAllOK" + "." + "vxlanAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigTunnelVxlanAllOK" + "." + "vxlanAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get config tunnel vxlan all o k body based on the context it is used
func (o *GetConfigTunnelVxlanAllOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVxlanAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigTunnelVxlanAllOKBody) contextValidateVxlanAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.VxlanAttr); i++ {

		if o.VxlanAttr[i] != nil {
			if err := o.VxlanAttr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigTunnelVxlanAllOK" + "." + "vxlanAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigTunnelVxlanAllOK" + "." + "vxlanAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigTunnelVxlanAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigTunnelVxlanAllOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigTunnelVxlanAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
