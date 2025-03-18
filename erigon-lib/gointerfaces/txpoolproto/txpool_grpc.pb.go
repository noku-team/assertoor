// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: txpool/txpool.proto

package txpoolproto

import (
	context "context"
	typesproto "github.com/erigontech/erigon-lib/gointerfaces/typesproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Txpool_Version_FullMethodName      = "/txpool.Txpool/Version"
	Txpool_FindUnknown_FullMethodName  = "/txpool.Txpool/FindUnknown"
	Txpool_Add_FullMethodName          = "/txpool.Txpool/Add"
	Txpool_Transactions_FullMethodName = "/txpool.Txpool/Transactions"
	Txpool_All_FullMethodName          = "/txpool.Txpool/All"
	Txpool_Pending_FullMethodName      = "/txpool.Txpool/Pending"
	Txpool_OnAdd_FullMethodName        = "/txpool.Txpool/OnAdd"
	Txpool_Status_FullMethodName       = "/txpool.Txpool/Status"
	Txpool_Nonce_FullMethodName        = "/txpool.Txpool/Nonce"
	Txpool_GetBlobs_FullMethodName     = "/txpool.Txpool/GetBlobs"
)

// TxpoolClient is the client API for Txpool service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TxpoolClient interface {
	// Version returns the service version number
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*typesproto.VersionReply, error)
	// preserves incoming order, changes amount, unknown hashes will be omitted
	FindUnknown(ctx context.Context, in *TxHashes, opts ...grpc.CallOption) (*TxHashes, error)
	// Expecting signed transactions. Preserves incoming order and amount
	// Adding txs as local (use P2P to add remote txs)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	// preserves incoming order and amount, if some transaction doesn't exists in pool - returns nil in this slot
	Transactions(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (*TransactionsReply, error)
	// returns all transactions from tx pool
	All(ctx context.Context, in *AllRequest, opts ...grpc.CallOption) (*AllReply, error)
	// Returns all pending (processable) transactions, in ready-for-mining order
	Pending(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PendingReply, error)
	// subscribe to new transactions add event
	OnAdd(ctx context.Context, in *OnAddRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OnAddReply], error)
	// returns high level status
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	// returns nonce for given account
	Nonce(ctx context.Context, in *NonceRequest, opts ...grpc.CallOption) (*NonceReply, error)
	// returns the list of blobs and proofs for a given list of blob hashes
	GetBlobs(ctx context.Context, in *GetBlobsRequest, opts ...grpc.CallOption) (*GetBlobsReply, error)
}

type txpoolClient struct {
	cc grpc.ClientConnInterface
}

func NewTxpoolClient(cc grpc.ClientConnInterface) TxpoolClient {
	return &txpoolClient{cc}
}

func (c *txpoolClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*typesproto.VersionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(typesproto.VersionReply)
	err := c.cc.Invoke(ctx, Txpool_Version_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txpoolClient) FindUnknown(ctx context.Context, in *TxHashes, opts ...grpc.CallOption) (*TxHashes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TxHashes)
	err := c.cc.Invoke(ctx, Txpool_FindUnknown_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txpoolClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddReply)
	err := c.cc.Invoke(ctx, Txpool_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txpoolClient) Transactions(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (*TransactionsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionsReply)
	err := c.cc.Invoke(ctx, Txpool_Transactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txpoolClient) All(ctx context.Context, in *AllRequest, opts ...grpc.CallOption) (*AllReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllReply)
	err := c.cc.Invoke(ctx, Txpool_All_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txpoolClient) Pending(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PendingReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PendingReply)
	err := c.cc.Invoke(ctx, Txpool_Pending_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txpoolClient) OnAdd(ctx context.Context, in *OnAddRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OnAddReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Txpool_ServiceDesc.Streams[0], Txpool_OnAdd_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[OnAddRequest, OnAddReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Txpool_OnAddClient = grpc.ServerStreamingClient[OnAddReply]

func (c *txpoolClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, Txpool_Status_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txpoolClient) Nonce(ctx context.Context, in *NonceRequest, opts ...grpc.CallOption) (*NonceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NonceReply)
	err := c.cc.Invoke(ctx, Txpool_Nonce_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txpoolClient) GetBlobs(ctx context.Context, in *GetBlobsRequest, opts ...grpc.CallOption) (*GetBlobsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlobsReply)
	err := c.cc.Invoke(ctx, Txpool_GetBlobs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TxpoolServer is the server API for Txpool service.
// All implementations must embed UnimplementedTxpoolServer
// for forward compatibility.
type TxpoolServer interface {
	// Version returns the service version number
	Version(context.Context, *emptypb.Empty) (*typesproto.VersionReply, error)
	// preserves incoming order, changes amount, unknown hashes will be omitted
	FindUnknown(context.Context, *TxHashes) (*TxHashes, error)
	// Expecting signed transactions. Preserves incoming order and amount
	// Adding txs as local (use P2P to add remote txs)
	Add(context.Context, *AddRequest) (*AddReply, error)
	// preserves incoming order and amount, if some transaction doesn't exists in pool - returns nil in this slot
	Transactions(context.Context, *TransactionsRequest) (*TransactionsReply, error)
	// returns all transactions from tx pool
	All(context.Context, *AllRequest) (*AllReply, error)
	// Returns all pending (processable) transactions, in ready-for-mining order
	Pending(context.Context, *emptypb.Empty) (*PendingReply, error)
	// subscribe to new transactions add event
	OnAdd(*OnAddRequest, grpc.ServerStreamingServer[OnAddReply]) error
	// returns high level status
	Status(context.Context, *StatusRequest) (*StatusReply, error)
	// returns nonce for given account
	Nonce(context.Context, *NonceRequest) (*NonceReply, error)
	// returns the list of blobs and proofs for a given list of blob hashes
	GetBlobs(context.Context, *GetBlobsRequest) (*GetBlobsReply, error)
	mustEmbedUnimplementedTxpoolServer()
}

// UnimplementedTxpoolServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTxpoolServer struct{}

func (UnimplementedTxpoolServer) Version(context.Context, *emptypb.Empty) (*typesproto.VersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedTxpoolServer) FindUnknown(context.Context, *TxHashes) (*TxHashes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUnknown not implemented")
}
func (UnimplementedTxpoolServer) Add(context.Context, *AddRequest) (*AddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedTxpoolServer) Transactions(context.Context, *TransactionsRequest) (*TransactionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transactions not implemented")
}
func (UnimplementedTxpoolServer) All(context.Context, *AllRequest) (*AllReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedTxpoolServer) Pending(context.Context, *emptypb.Empty) (*PendingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pending not implemented")
}
func (UnimplementedTxpoolServer) OnAdd(*OnAddRequest, grpc.ServerStreamingServer[OnAddReply]) error {
	return status.Errorf(codes.Unimplemented, "method OnAdd not implemented")
}
func (UnimplementedTxpoolServer) Status(context.Context, *StatusRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedTxpoolServer) Nonce(context.Context, *NonceRequest) (*NonceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nonce not implemented")
}
func (UnimplementedTxpoolServer) GetBlobs(context.Context, *GetBlobsRequest) (*GetBlobsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlobs not implemented")
}
func (UnimplementedTxpoolServer) mustEmbedUnimplementedTxpoolServer() {}
func (UnimplementedTxpoolServer) testEmbeddedByValue()                {}

// UnsafeTxpoolServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TxpoolServer will
// result in compilation errors.
type UnsafeTxpoolServer interface {
	mustEmbedUnimplementedTxpoolServer()
}

func RegisterTxpoolServer(s grpc.ServiceRegistrar, srv TxpoolServer) {
	// If the following call pancis, it indicates UnimplementedTxpoolServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Txpool_ServiceDesc, srv)
}

func _Txpool_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxpoolServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Txpool_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxpoolServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Txpool_FindUnknown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxHashes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxpoolServer).FindUnknown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Txpool_FindUnknown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxpoolServer).FindUnknown(ctx, req.(*TxHashes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Txpool_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxpoolServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Txpool_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxpoolServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Txpool_Transactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxpoolServer).Transactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Txpool_Transactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxpoolServer).Transactions(ctx, req.(*TransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Txpool_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxpoolServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Txpool_All_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxpoolServer).All(ctx, req.(*AllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Txpool_Pending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxpoolServer).Pending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Txpool_Pending_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxpoolServer).Pending(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Txpool_OnAdd_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OnAddRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TxpoolServer).OnAdd(m, &grpc.GenericServerStream[OnAddRequest, OnAddReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Txpool_OnAddServer = grpc.ServerStreamingServer[OnAddReply]

func _Txpool_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxpoolServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Txpool_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxpoolServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Txpool_Nonce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxpoolServer).Nonce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Txpool_Nonce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxpoolServer).Nonce(ctx, req.(*NonceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Txpool_GetBlobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxpoolServer).GetBlobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Txpool_GetBlobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxpoolServer).GetBlobs(ctx, req.(*GetBlobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Txpool_ServiceDesc is the grpc.ServiceDesc for Txpool service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Txpool_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "txpool.Txpool",
	HandlerType: (*TxpoolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Txpool_Version_Handler,
		},
		{
			MethodName: "FindUnknown",
			Handler:    _Txpool_FindUnknown_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Txpool_Add_Handler,
		},
		{
			MethodName: "Transactions",
			Handler:    _Txpool_Transactions_Handler,
		},
		{
			MethodName: "All",
			Handler:    _Txpool_All_Handler,
		},
		{
			MethodName: "Pending",
			Handler:    _Txpool_Pending_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Txpool_Status_Handler,
		},
		{
			MethodName: "Nonce",
			Handler:    _Txpool_Nonce_Handler,
		},
		{
			MethodName: "GetBlobs",
			Handler:    _Txpool_GetBlobs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnAdd",
			Handler:       _Txpool_OnAdd_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "txpool/txpool.proto",
}
