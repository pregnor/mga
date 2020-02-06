// +build !ignore_autogenerated

// Copyright 2020 Acme Inc.
// All rights reserved.
//
// Licensed under "Only for testing purposes" license.

// Code generated by mga tool. DO NOT EDIT.

package pkgdriver

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
	"sagikazarmark.dev/mga/internal/generate/kit/endpoint/testdata/generator/unnamed_param"
)

// endpointError identifies an error that should be returned as an error endpoint.
type endpointError interface {
	EndpointError() bool
}

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateTodo endpoint.Endpoint
}

// MakeEndpoints returns a(n) Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service unnamed_param.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Combine(middleware...)

	return Endpoints{CreateTodo: kitxendpoint.OperationNameMiddleware("unnamed_param.CreateTodo")(mw(MakeCreateTodoEndpoint(service)))}
}

// TraceEndpoints returns a(n) Endpoints struct where each endpoint is wrapped with a tracing middleware.
func TraceEndpoints(endpoints Endpoints) Endpoints {
	return Endpoints{CreateTodo: kitoc.TraceEndpoint("unnamed_param.CreateTodo")(endpoints.CreateTodo)}
}

// CreateTodoRequest is a request struct for CreateTodo endpoint.
type CreateTodoRequest struct {
	P0 string
}

// CreateTodoResponse is a response struct for CreateTodo endpoint.
type CreateTodoResponse struct {
	R0  string
	Err error
}

func (r *CreateTodoResponse) Failed() error {
	return r.Err
}

// MakeCreateTodoEndpoint returns an endpoint for the matching method of the underlying service.
func MakeCreateTodoEndpoint(service unnamed_param.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*CreateTodoRequest)

		r0, err := service.CreateTodo(ctx, req.P0)

		if err != nil {
			if endpointErr := endpointError(nil); errors.As(err, &endpointErr) && endpointErr.EndpointError() {
				return &CreateTodoResponse{
					Err: err,
					R0:  r0,
				}, err
			}

			return &CreateTodoResponse{
				Err: err,
				R0:  r0,
			}, nil
		}

		return &CreateTodoResponse{R0: r0}, nil
	}
}
