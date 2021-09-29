// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities endpoints
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package sdkutilities

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "sdk-utilities" service endpoints.
type Endpoints struct {
	Supply goa.Endpoint
}

// NewEndpoints wraps the methods of the "sdk-utilities" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Supply: NewSupplyEndpoint(s),
	}
}

// Use applies the given middleware to all the "sdk-utilities" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Supply = m(e.Supply)
}

// NewSupplyEndpoint returns an endpoint function that calls the method
// "supply" of service "sdk-utilities".
func NewSupplyEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SupplyPayload)
		return s.Supply(ctx, p)
	}
}
