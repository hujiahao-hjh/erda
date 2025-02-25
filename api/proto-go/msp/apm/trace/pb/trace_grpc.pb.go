// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: trace.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TraceServiceClient interface {
	GetSpans(ctx context.Context, in *GetSpansRequest, opts ...grpc.CallOption) (*GetSpansResponse, error)
	GetSpanDashboards(ctx context.Context, in *GetSpanDashboardsRequest, opts ...grpc.CallOption) (*GetSpanDashboardsResponse, error)
	GetTraces(ctx context.Context, in *GetTracesRequest, opts ...grpc.CallOption) (*GetTracesResponse, error)
	GetTraceQueryConditions(ctx context.Context, in *GetTraceQueryConditionsRequest, opts ...grpc.CallOption) (*GetTraceQueryConditionsResponse, error)
	GetTraceDebugHistories(ctx context.Context, in *GetTraceDebugHistoriesRequest, opts ...grpc.CallOption) (*GetTraceDebugHistoriesResponse, error)
	GetTraceDebugByRequestID(ctx context.Context, in *GetTraceDebugRequest, opts ...grpc.CallOption) (*GetTraceDebugResponse, error)
	CreateTraceDebug(ctx context.Context, in *CreateTraceDebugRequest, opts ...grpc.CallOption) (*CreateTraceDebugResponse, error)
	StopTraceDebug(ctx context.Context, in *StopTraceDebugRequest, opts ...grpc.CallOption) (*StopTraceDebugResponse, error)
	GetTraceDebugHistoryStatusByRequestID(ctx context.Context, in *GetTraceDebugStatusByRequestIDRequest, opts ...grpc.CallOption) (*GetTraceDebugStatusByRequestIDResponse, error)
}

type traceServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewTraceServiceClient(cc grpc1.ClientConnInterface) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) GetSpans(ctx context.Context, in *GetSpansRequest, opts ...grpc.CallOption) (*GetSpansResponse, error) {
	out := new(GetSpansResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/GetSpans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) GetSpanDashboards(ctx context.Context, in *GetSpanDashboardsRequest, opts ...grpc.CallOption) (*GetSpanDashboardsResponse, error) {
	out := new(GetSpanDashboardsResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/GetSpanDashboards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) GetTraces(ctx context.Context, in *GetTracesRequest, opts ...grpc.CallOption) (*GetTracesResponse, error) {
	out := new(GetTracesResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/GetTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) GetTraceQueryConditions(ctx context.Context, in *GetTraceQueryConditionsRequest, opts ...grpc.CallOption) (*GetTraceQueryConditionsResponse, error) {
	out := new(GetTraceQueryConditionsResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/GetTraceQueryConditions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) GetTraceDebugHistories(ctx context.Context, in *GetTraceDebugHistoriesRequest, opts ...grpc.CallOption) (*GetTraceDebugHistoriesResponse, error) {
	out := new(GetTraceDebugHistoriesResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/GetTraceDebugHistories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) GetTraceDebugByRequestID(ctx context.Context, in *GetTraceDebugRequest, opts ...grpc.CallOption) (*GetTraceDebugResponse, error) {
	out := new(GetTraceDebugResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/GetTraceDebugByRequestID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) CreateTraceDebug(ctx context.Context, in *CreateTraceDebugRequest, opts ...grpc.CallOption) (*CreateTraceDebugResponse, error) {
	out := new(CreateTraceDebugResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/CreateTraceDebug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) StopTraceDebug(ctx context.Context, in *StopTraceDebugRequest, opts ...grpc.CallOption) (*StopTraceDebugResponse, error) {
	out := new(StopTraceDebugResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/StopTraceDebug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) GetTraceDebugHistoryStatusByRequestID(ctx context.Context, in *GetTraceDebugStatusByRequestIDRequest, opts ...grpc.CallOption) (*GetTraceDebugStatusByRequestIDResponse, error) {
	out := new(GetTraceDebugStatusByRequestIDResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/GetTraceDebugHistoryStatusByRequestID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceServiceServer is the server API for TraceService service.
// All implementations should embed UnimplementedTraceServiceServer
// for forward compatibility
type TraceServiceServer interface {
	GetSpans(context.Context, *GetSpansRequest) (*GetSpansResponse, error)
	GetSpanDashboards(context.Context, *GetSpanDashboardsRequest) (*GetSpanDashboardsResponse, error)
	GetTraces(context.Context, *GetTracesRequest) (*GetTracesResponse, error)
	GetTraceQueryConditions(context.Context, *GetTraceQueryConditionsRequest) (*GetTraceQueryConditionsResponse, error)
	GetTraceDebugHistories(context.Context, *GetTraceDebugHistoriesRequest) (*GetTraceDebugHistoriesResponse, error)
	GetTraceDebugByRequestID(context.Context, *GetTraceDebugRequest) (*GetTraceDebugResponse, error)
	CreateTraceDebug(context.Context, *CreateTraceDebugRequest) (*CreateTraceDebugResponse, error)
	StopTraceDebug(context.Context, *StopTraceDebugRequest) (*StopTraceDebugResponse, error)
	GetTraceDebugHistoryStatusByRequestID(context.Context, *GetTraceDebugStatusByRequestIDRequest) (*GetTraceDebugStatusByRequestIDResponse, error)
}

// UnimplementedTraceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (*UnimplementedTraceServiceServer) GetSpans(context.Context, *GetSpansRequest) (*GetSpansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpans not implemented")
}
func (*UnimplementedTraceServiceServer) GetSpanDashboards(context.Context, *GetSpanDashboardsRequest) (*GetSpanDashboardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpanDashboards not implemented")
}
func (*UnimplementedTraceServiceServer) GetTraces(context.Context, *GetTracesRequest) (*GetTracesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraces not implemented")
}
func (*UnimplementedTraceServiceServer) GetTraceQueryConditions(context.Context, *GetTraceQueryConditionsRequest) (*GetTraceQueryConditionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraceQueryConditions not implemented")
}
func (*UnimplementedTraceServiceServer) GetTraceDebugHistories(context.Context, *GetTraceDebugHistoriesRequest) (*GetTraceDebugHistoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraceDebugHistories not implemented")
}
func (*UnimplementedTraceServiceServer) GetTraceDebugByRequestID(context.Context, *GetTraceDebugRequest) (*GetTraceDebugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraceDebugByRequestID not implemented")
}
func (*UnimplementedTraceServiceServer) CreateTraceDebug(context.Context, *CreateTraceDebugRequest) (*CreateTraceDebugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTraceDebug not implemented")
}
func (*UnimplementedTraceServiceServer) StopTraceDebug(context.Context, *StopTraceDebugRequest) (*StopTraceDebugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopTraceDebug not implemented")
}
func (*UnimplementedTraceServiceServer) GetTraceDebugHistoryStatusByRequestID(context.Context, *GetTraceDebugStatusByRequestIDRequest) (*GetTraceDebugStatusByRequestIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraceDebugHistoryStatusByRequestID not implemented")
}

func RegisterTraceServiceServer(s grpc1.ServiceRegistrar, srv TraceServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_TraceService_serviceDesc(srv, opts...), srv)
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.msp.apm.trace.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "trace.proto",
}

func _get_TraceService_serviceDesc(srv TraceServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_TraceService_GetSpans_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetSpans(ctx, req.(*GetSpansRequest))
	}
	var _TraceService_GetSpans_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_GetSpans_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetSpans", srv)
		_TraceService_GetSpans_Handler = h.Interceptor(_TraceService_GetSpans_Handler)
	}

	_TraceService_GetSpanDashboards_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetSpanDashboards(ctx, req.(*GetSpanDashboardsRequest))
	}
	var _TraceService_GetSpanDashboards_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_GetSpanDashboards_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetSpanDashboards", srv)
		_TraceService_GetSpanDashboards_Handler = h.Interceptor(_TraceService_GetSpanDashboards_Handler)
	}

	_TraceService_GetTraces_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTraces(ctx, req.(*GetTracesRequest))
	}
	var _TraceService_GetTraces_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_GetTraces_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraces", srv)
		_TraceService_GetTraces_Handler = h.Interceptor(_TraceService_GetTraces_Handler)
	}

	_TraceService_GetTraceQueryConditions_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTraceQueryConditions(ctx, req.(*GetTraceQueryConditionsRequest))
	}
	var _TraceService_GetTraceQueryConditions_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_GetTraceQueryConditions_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraceQueryConditions", srv)
		_TraceService_GetTraceQueryConditions_Handler = h.Interceptor(_TraceService_GetTraceQueryConditions_Handler)
	}

	_TraceService_GetTraceDebugHistories_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTraceDebugHistories(ctx, req.(*GetTraceDebugHistoriesRequest))
	}
	var _TraceService_GetTraceDebugHistories_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_GetTraceDebugHistories_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraceDebugHistories", srv)
		_TraceService_GetTraceDebugHistories_Handler = h.Interceptor(_TraceService_GetTraceDebugHistories_Handler)
	}

	_TraceService_GetTraceDebugByRequestID_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTraceDebugByRequestID(ctx, req.(*GetTraceDebugRequest))
	}
	var _TraceService_GetTraceDebugByRequestID_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_GetTraceDebugByRequestID_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraceDebugByRequestID", srv)
		_TraceService_GetTraceDebugByRequestID_Handler = h.Interceptor(_TraceService_GetTraceDebugByRequestID_Handler)
	}

	_TraceService_CreateTraceDebug_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateTraceDebug(ctx, req.(*CreateTraceDebugRequest))
	}
	var _TraceService_CreateTraceDebug_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_CreateTraceDebug_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "CreateTraceDebug", srv)
		_TraceService_CreateTraceDebug_Handler = h.Interceptor(_TraceService_CreateTraceDebug_Handler)
	}

	_TraceService_StopTraceDebug_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.StopTraceDebug(ctx, req.(*StopTraceDebugRequest))
	}
	var _TraceService_StopTraceDebug_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_StopTraceDebug_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "StopTraceDebug", srv)
		_TraceService_StopTraceDebug_Handler = h.Interceptor(_TraceService_StopTraceDebug_Handler)
	}

	_TraceService_GetTraceDebugHistoryStatusByRequestID_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTraceDebugHistoryStatusByRequestID(ctx, req.(*GetTraceDebugStatusByRequestIDRequest))
	}
	var _TraceService_GetTraceDebugHistoryStatusByRequestID_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_GetTraceDebugHistoryStatusByRequestID_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraceDebugHistoryStatusByRequestID", srv)
		_TraceService_GetTraceDebugHistoryStatusByRequestID_Handler = h.Interceptor(_TraceService_GetTraceDebugHistoryStatusByRequestID_Handler)
	}

	var serviceDesc = _TraceService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "GetSpans",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetSpansRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).GetSpans(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_GetSpans_info)
				}
				if interceptor == nil {
					return _TraceService_GetSpans_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/GetSpans",
				}
				return interceptor(ctx, in, info, _TraceService_GetSpans_Handler)
			},
		},
		{
			MethodName: "GetSpanDashboards",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetSpanDashboardsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).GetSpanDashboards(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_GetSpanDashboards_info)
				}
				if interceptor == nil {
					return _TraceService_GetSpanDashboards_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/GetSpanDashboards",
				}
				return interceptor(ctx, in, info, _TraceService_GetSpanDashboards_Handler)
			},
		},
		{
			MethodName: "GetTraces",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetTracesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).GetTraces(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_GetTraces_info)
				}
				if interceptor == nil {
					return _TraceService_GetTraces_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/GetTraces",
				}
				return interceptor(ctx, in, info, _TraceService_GetTraces_Handler)
			},
		},
		{
			MethodName: "GetTraceQueryConditions",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetTraceQueryConditionsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).GetTraceQueryConditions(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_GetTraceQueryConditions_info)
				}
				if interceptor == nil {
					return _TraceService_GetTraceQueryConditions_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/GetTraceQueryConditions",
				}
				return interceptor(ctx, in, info, _TraceService_GetTraceQueryConditions_Handler)
			},
		},
		{
			MethodName: "GetTraceDebugHistories",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetTraceDebugHistoriesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).GetTraceDebugHistories(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_GetTraceDebugHistories_info)
				}
				if interceptor == nil {
					return _TraceService_GetTraceDebugHistories_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/GetTraceDebugHistories",
				}
				return interceptor(ctx, in, info, _TraceService_GetTraceDebugHistories_Handler)
			},
		},
		{
			MethodName: "GetTraceDebugByRequestID",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetTraceDebugRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).GetTraceDebugByRequestID(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_GetTraceDebugByRequestID_info)
				}
				if interceptor == nil {
					return _TraceService_GetTraceDebugByRequestID_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/GetTraceDebugByRequestID",
				}
				return interceptor(ctx, in, info, _TraceService_GetTraceDebugByRequestID_Handler)
			},
		},
		{
			MethodName: "CreateTraceDebug",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateTraceDebugRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).CreateTraceDebug(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_CreateTraceDebug_info)
				}
				if interceptor == nil {
					return _TraceService_CreateTraceDebug_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/CreateTraceDebug",
				}
				return interceptor(ctx, in, info, _TraceService_CreateTraceDebug_Handler)
			},
		},
		{
			MethodName: "StopTraceDebug",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(StopTraceDebugRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).StopTraceDebug(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_StopTraceDebug_info)
				}
				if interceptor == nil {
					return _TraceService_StopTraceDebug_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/StopTraceDebug",
				}
				return interceptor(ctx, in, info, _TraceService_StopTraceDebug_Handler)
			},
		},
		{
			MethodName: "GetTraceDebugHistoryStatusByRequestID",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetTraceDebugStatusByRequestIDRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).GetTraceDebugHistoryStatusByRequestID(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_GetTraceDebugHistoryStatusByRequestID_info)
				}
				if interceptor == nil {
					return _TraceService_GetTraceDebugHistoryStatusByRequestID_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/GetTraceDebugHistoryStatusByRequestID",
				}
				return interceptor(ctx, in, info, _TraceService_GetTraceDebugHistoryStatusByRequestID_Handler)
			},
		},
	}
	return &serviceDesc
}
