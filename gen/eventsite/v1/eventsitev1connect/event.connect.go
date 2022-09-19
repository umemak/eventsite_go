// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: eventsite/v1/event.proto

package eventsitev1connect

import (
	context "context"
	errors "errors"
	v1 "eventsite/gen/eventsite/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// EventsiteServiceName is the fully-qualified name of the EventsiteService service.
	EventsiteServiceName = "eventsite.v1.EventsiteService"
)

// EventsiteServiceClient is a client for the eventsite.v1.EventsiteService service.
type EventsiteServiceClient interface {
	// Get all events.
	GetEvents(context.Context, *connect_go.Request[v1.GetEventsRequest]) (*connect_go.Response[v1.EventsMessage], error)
	// Create event.
	PostEvents(context.Context, *connect_go.Request[v1.CreateEventsRequest]) (*connect_go.Response[v1.EventsMessage], error)
}

// NewEventsiteServiceClient constructs a client for the eventsite.v1.EventsiteService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEventsiteServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EventsiteServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &eventsiteServiceClient{
		getEvents: connect_go.NewClient[v1.GetEventsRequest, v1.EventsMessage](
			httpClient,
			baseURL+"/eventsite.v1.EventsiteService/GetEvents",
			opts...,
		),
		postEvents: connect_go.NewClient[v1.CreateEventsRequest, v1.EventsMessage](
			httpClient,
			baseURL+"/eventsite.v1.EventsiteService/PostEvents",
			opts...,
		),
	}
}

// eventsiteServiceClient implements EventsiteServiceClient.
type eventsiteServiceClient struct {
	getEvents  *connect_go.Client[v1.GetEventsRequest, v1.EventsMessage]
	postEvents *connect_go.Client[v1.CreateEventsRequest, v1.EventsMessage]
}

// GetEvents calls eventsite.v1.EventsiteService.GetEvents.
func (c *eventsiteServiceClient) GetEvents(ctx context.Context, req *connect_go.Request[v1.GetEventsRequest]) (*connect_go.Response[v1.EventsMessage], error) {
	return c.getEvents.CallUnary(ctx, req)
}

// PostEvents calls eventsite.v1.EventsiteService.PostEvents.
func (c *eventsiteServiceClient) PostEvents(ctx context.Context, req *connect_go.Request[v1.CreateEventsRequest]) (*connect_go.Response[v1.EventsMessage], error) {
	return c.postEvents.CallUnary(ctx, req)
}

// EventsiteServiceHandler is an implementation of the eventsite.v1.EventsiteService service.
type EventsiteServiceHandler interface {
	// Get all events.
	GetEvents(context.Context, *connect_go.Request[v1.GetEventsRequest]) (*connect_go.Response[v1.EventsMessage], error)
	// Create event.
	PostEvents(context.Context, *connect_go.Request[v1.CreateEventsRequest]) (*connect_go.Response[v1.EventsMessage], error)
}

// NewEventsiteServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEventsiteServiceHandler(svc EventsiteServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/eventsite.v1.EventsiteService/GetEvents", connect_go.NewUnaryHandler(
		"/eventsite.v1.EventsiteService/GetEvents",
		svc.GetEvents,
		opts...,
	))
	mux.Handle("/eventsite.v1.EventsiteService/PostEvents", connect_go.NewUnaryHandler(
		"/eventsite.v1.EventsiteService/PostEvents",
		svc.PostEvents,
		opts...,
	))
	return "/eventsite.v1.EventsiteService/", mux
}

// UnimplementedEventsiteServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEventsiteServiceHandler struct{}

func (UnimplementedEventsiteServiceHandler) GetEvents(context.Context, *connect_go.Request[v1.GetEventsRequest]) (*connect_go.Response[v1.EventsMessage], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("eventsite.v1.EventsiteService.GetEvents is not implemented"))
}

func (UnimplementedEventsiteServiceHandler) PostEvents(context.Context, *connect_go.Request[v1.CreateEventsRequest]) (*connect_go.Response[v1.EventsMessage], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("eventsite.v1.EventsiteService.PostEvents is not implemented"))
}