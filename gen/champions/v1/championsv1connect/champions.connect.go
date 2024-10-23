// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: champions/v1/champions.proto

package championsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "maxischmaxi/jstreams-server/gen/champions/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ChampionsServiceName is the fully-qualified name of the ChampionsService service.
	ChampionsServiceName = "champions.ChampionsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ChampionsServiceGetChampionsProcedure is the fully-qualified name of the ChampionsService's
	// GetChampions RPC.
	ChampionsServiceGetChampionsProcedure = "/champions.ChampionsService/GetChampions"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	championsServiceServiceDescriptor            = v1.File_champions_v1_champions_proto.Services().ByName("ChampionsService")
	championsServiceGetChampionsMethodDescriptor = championsServiceServiceDescriptor.Methods().ByName("GetChampions")
)

// ChampionsServiceClient is a client for the champions.ChampionsService service.
type ChampionsServiceClient interface {
	GetChampions(context.Context, *connect.Request[v1.GetChampionsRequest]) (*connect.Response[v1.GetChampionsResponse], error)
}

// NewChampionsServiceClient constructs a client for the champions.ChampionsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChampionsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ChampionsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &championsServiceClient{
		getChampions: connect.NewClient[v1.GetChampionsRequest, v1.GetChampionsResponse](
			httpClient,
			baseURL+ChampionsServiceGetChampionsProcedure,
			connect.WithSchema(championsServiceGetChampionsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// championsServiceClient implements ChampionsServiceClient.
type championsServiceClient struct {
	getChampions *connect.Client[v1.GetChampionsRequest, v1.GetChampionsResponse]
}

// GetChampions calls champions.ChampionsService.GetChampions.
func (c *championsServiceClient) GetChampions(ctx context.Context, req *connect.Request[v1.GetChampionsRequest]) (*connect.Response[v1.GetChampionsResponse], error) {
	return c.getChampions.CallUnary(ctx, req)
}

// ChampionsServiceHandler is an implementation of the champions.ChampionsService service.
type ChampionsServiceHandler interface {
	GetChampions(context.Context, *connect.Request[v1.GetChampionsRequest]) (*connect.Response[v1.GetChampionsResponse], error)
}

// NewChampionsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChampionsServiceHandler(svc ChampionsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	championsServiceGetChampionsHandler := connect.NewUnaryHandler(
		ChampionsServiceGetChampionsProcedure,
		svc.GetChampions,
		connect.WithSchema(championsServiceGetChampionsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/champions.ChampionsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ChampionsServiceGetChampionsProcedure:
			championsServiceGetChampionsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedChampionsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChampionsServiceHandler struct{}

func (UnimplementedChampionsServiceHandler) GetChampions(context.Context, *connect.Request[v1.GetChampionsRequest]) (*connect.Response[v1.GetChampionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("champions.ChampionsService.GetChampions is not implemented"))
}
