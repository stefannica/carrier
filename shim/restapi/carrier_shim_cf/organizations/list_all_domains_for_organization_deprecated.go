// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllDomainsForOrganizationDeprecatedHandlerFunc turns a function with the right signature into a list all domains for organization deprecated handler
type ListAllDomainsForOrganizationDeprecatedHandlerFunc func(ListAllDomainsForOrganizationDeprecatedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllDomainsForOrganizationDeprecatedHandlerFunc) Handle(params ListAllDomainsForOrganizationDeprecatedParams) middleware.Responder {
	return fn(params)
}

// ListAllDomainsForOrganizationDeprecatedHandler interface for that can handle valid list all domains for organization deprecated params
type ListAllDomainsForOrganizationDeprecatedHandler interface {
	Handle(ListAllDomainsForOrganizationDeprecatedParams) middleware.Responder
}

// NewListAllDomainsForOrganizationDeprecated creates a new http.Handler for the list all domains for organization deprecated operation
func NewListAllDomainsForOrganizationDeprecated(ctx *middleware.Context, handler ListAllDomainsForOrganizationDeprecatedHandler) *ListAllDomainsForOrganizationDeprecated {
	return &ListAllDomainsForOrganizationDeprecated{Context: ctx, Handler: handler}
}

/*ListAllDomainsForOrganizationDeprecated swagger:route GET /organizations/{guid}/domains organizations listAllDomainsForOrganizationDeprecated

List all Domains for the Organization (deprecated)

curl --insecure -i %s/v2/organizations/{guid}/domains -X GET -H 'Authorization: %s'

*/
type ListAllDomainsForOrganizationDeprecated struct {
	Context *middleware.Context
	Handler ListAllDomainsForOrganizationDeprecatedHandler
}

func (o *ListAllDomainsForOrganizationDeprecated) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllDomainsForOrganizationDeprecatedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}