// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: location.proto

/*
Package location is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package location

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_LocationService_CreateLocation_0(ctx context.Context, marshaler runtime.Marshaler, client LocationServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq CreateLocationRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.CreateLocation(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_LocationService_CreateLocation_0(ctx context.Context, marshaler runtime.Marshaler, server LocationServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq CreateLocationRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.CreateLocation(ctx, &protoReq)
	return msg, metadata, err
}

func request_LocationService_GetLocations_0(ctx context.Context, marshaler runtime.Marshaler, client LocationServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.GetLocations(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_LocationService_GetLocations_0(ctx context.Context, marshaler runtime.Marshaler, server LocationServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.GetLocations(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterLocationServiceHandlerServer registers the http handlers for service LocationService to "mux".
// UnaryRPC     :call LocationServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterLocationServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterLocationServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server LocationServiceServer) error {
	mux.Handle(http.MethodPost, pattern_LocationService_CreateLocation_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/location.LocationService/CreateLocation", runtime.WithHTTPPathPattern("/v1/location"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_LocationService_CreateLocation_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_LocationService_CreateLocation_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_LocationService_GetLocations_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/location.LocationService/GetLocations", runtime.WithHTTPPathPattern("/v1/location"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_LocationService_GetLocations_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_LocationService_GetLocations_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterLocationServiceHandlerFromEndpoint is same as RegisterLocationServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterLocationServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterLocationServiceHandler(ctx, mux, conn)
}

// RegisterLocationServiceHandler registers the http handlers for service LocationService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterLocationServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterLocationServiceHandlerClient(ctx, mux, NewLocationServiceClient(conn))
}

// RegisterLocationServiceHandlerClient registers the http handlers for service LocationService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "LocationServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "LocationServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "LocationServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterLocationServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client LocationServiceClient) error {
	mux.Handle(http.MethodPost, pattern_LocationService_CreateLocation_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/location.LocationService/CreateLocation", runtime.WithHTTPPathPattern("/v1/location"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_LocationService_CreateLocation_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_LocationService_CreateLocation_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_LocationService_GetLocations_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/location.LocationService/GetLocations", runtime.WithHTTPPathPattern("/v1/location"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_LocationService_GetLocations_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_LocationService_GetLocations_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_LocationService_CreateLocation_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "location"}, ""))
	pattern_LocationService_GetLocations_0   = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "location"}, ""))
)

var (
	forward_LocationService_CreateLocation_0 = runtime.ForwardResponseMessage
	forward_LocationService_GetLocations_0   = runtime.ForwardResponseMessage
)
