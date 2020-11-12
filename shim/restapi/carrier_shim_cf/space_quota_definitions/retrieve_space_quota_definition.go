// Code generated by go-swagger; DO NOT EDIT.

package space_quota_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RetrieveSpaceQuotaDefinitionHandlerFunc turns a function with the right signature into a retrieve space quota definition handler
type RetrieveSpaceQuotaDefinitionHandlerFunc func(RetrieveSpaceQuotaDefinitionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RetrieveSpaceQuotaDefinitionHandlerFunc) Handle(params RetrieveSpaceQuotaDefinitionParams) middleware.Responder {
	return fn(params)
}

// RetrieveSpaceQuotaDefinitionHandler interface for that can handle valid retrieve space quota definition params
type RetrieveSpaceQuotaDefinitionHandler interface {
	Handle(RetrieveSpaceQuotaDefinitionParams) middleware.Responder
}

// NewRetrieveSpaceQuotaDefinition creates a new http.Handler for the retrieve space quota definition operation
func NewRetrieveSpaceQuotaDefinition(ctx *middleware.Context, handler RetrieveSpaceQuotaDefinitionHandler) *RetrieveSpaceQuotaDefinition {
	return &RetrieveSpaceQuotaDefinition{Context: ctx, Handler: handler}
}

/*RetrieveSpaceQuotaDefinition swagger:route GET /space_quota_definitions/{guid} spaceQuotaDefinitions retrieveSpaceQuotaDefinition

Retrieve a Particular Space Quota Definition

curl --insecure -i %s/v2/space_quota_definitions/{guid} -X GET -H 'Authorization: %s'

*/
type RetrieveSpaceQuotaDefinition struct {
	Context *middleware.Context
	Handler RetrieveSpaceQuotaDefinitionHandler
}

func (o *RetrieveSpaceQuotaDefinition) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRetrieveSpaceQuotaDefinitionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}