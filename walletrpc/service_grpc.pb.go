// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: service.proto

package walletrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CompactTxStreamerClient is the client API for CompactTxStreamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompactTxStreamerClient interface {
	GetLiteWalletBlockGroup(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*BlockID, error)
	// Return the height of the tip of the best chain
	GetLatestBlock(ctx context.Context, in *ChainSpec, opts ...grpc.CallOption) (*BlockID, error)
	// Return the compact block corresponding to the given block identifier
	GetBlock(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*CompactBlock, error)
	// Return a list of consecutive compact blocks
	GetBlockRange(ctx context.Context, in *BlockRange, opts ...grpc.CallOption) (CompactTxStreamer_GetBlockRangeClient, error)
	// Get the historical and current prices
	GetARRRPrice(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*PriceResponse, error)
	GetCurrentARRRPrice(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PriceResponse, error)
	// Return the requested full (not compact) transaction (as from pirated)
	GetTransaction(ctx context.Context, in *TxFilter, opts ...grpc.CallOption) (*RawTransaction, error)
	// Submit the given transaction to the Zcash network
	SendTransaction(ctx context.Context, in *RawTransaction, opts ...grpc.CallOption) (*SendResponse, error)
	// Return the txids corresponding to the given t-address within the given block range
	GetTaddressTxids(ctx context.Context, in *TransparentAddressBlockFilter, opts ...grpc.CallOption) (CompactTxStreamer_GetTaddressTxidsClient, error)
	GetTaddressBalance(ctx context.Context, in *AddressList, opts ...grpc.CallOption) (*Balance, error)
	GetTaddressBalanceStream(ctx context.Context, opts ...grpc.CallOption) (CompactTxStreamer_GetTaddressBalanceStreamClient, error)
	// Return the compact transactions currently in the mempool; the results
	// can be a few seconds out of date. If the Exclude list is empty, return
	// all transactions; otherwise return all *except* those in the Exclude list
	// (if any); this allows the client to avoid receiving transactions that it
	// already has (from an earlier call to this rpc). The transaction IDs in the
	// Exclude list can be shortened to any number of bytes to make the request
	// more bandwidth-efficient; if two or more transactions in the mempool
	// match a shortened txid, they are all sent (none is excluded). Transactions
	// in the exclude list that don't exist in the mempool are ignored.
	GetMempoolTx(ctx context.Context, in *Exclude, opts ...grpc.CallOption) (CompactTxStreamer_GetMempoolTxClient, error)
	// Return a stream of current Mempool transactions. This will keep the output stream open while
	// there are mempool transactions. It will close the returned stream when a new block is mined.
	GetMempoolStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CompactTxStreamer_GetMempoolStreamClient, error)
	// GetTreeState returns the note commitment tree state corresponding to the given block.
	// See section 3.7 of the Zcash protocol specification. It returns several other useful
	// values also (even though they can be obtained using GetBlock).
	// The block can be specified by either height or hash.
	GetTreeState(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*TreeState, error)
	GetAddressUtxos(ctx context.Context, in *GetAddressUtxosArg, opts ...grpc.CallOption) (*GetAddressUtxosReplyList, error)
	GetAddressUtxosStream(ctx context.Context, in *GetAddressUtxosArg, opts ...grpc.CallOption) (CompactTxStreamer_GetAddressUtxosStreamClient, error)
	// Return information about this lightwalletd instance and the blockchain
	GetLightdInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LightdInfo, error)
	// Testing-only, requires lightwalletd --ping-very-insecure (do not enable in production)
	Ping(ctx context.Context, in *Duration, opts ...grpc.CallOption) (*PingResponse, error)
}

type compactTxStreamerClient struct {
	cc grpc.ClientConnInterface
}

func NewCompactTxStreamerClient(cc grpc.ClientConnInterface) CompactTxStreamerClient {
	return &compactTxStreamerClient{cc}
}

func (c *compactTxStreamerClient) GetLiteWalletBlockGroup(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*BlockID, error) {
	out := new(BlockID)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetLiteWalletBlockGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetLatestBlock(ctx context.Context, in *ChainSpec, opts ...grpc.CallOption) (*BlockID, error) {
	out := new(BlockID)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetBlock(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*CompactBlock, error) {
	out := new(CompactBlock)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetBlockRange(ctx context.Context, in *BlockRange, opts ...grpc.CallOption) (CompactTxStreamer_GetBlockRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompactTxStreamer_ServiceDesc.Streams[0], "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetBlockRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetBlockRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompactTxStreamer_GetBlockRangeClient interface {
	Recv() (*CompactBlock, error)
	grpc.ClientStream
}

type compactTxStreamerGetBlockRangeClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetBlockRangeClient) Recv() (*CompactBlock, error) {
	m := new(CompactBlock)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetARRRPrice(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*PriceResponse, error) {
	out := new(PriceResponse)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetARRRPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetCurrentARRRPrice(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PriceResponse, error) {
	out := new(PriceResponse)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetCurrentARRRPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetTransaction(ctx context.Context, in *TxFilter, opts ...grpc.CallOption) (*RawTransaction, error) {
	out := new(RawTransaction)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) SendTransaction(ctx context.Context, in *RawTransaction, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/SendTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetTaddressTxids(ctx context.Context, in *TransparentAddressBlockFilter, opts ...grpc.CallOption) (CompactTxStreamer_GetTaddressTxidsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompactTxStreamer_ServiceDesc.Streams[1], "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetTaddressTxids", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetTaddressTxidsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompactTxStreamer_GetTaddressTxidsClient interface {
	Recv() (*RawTransaction, error)
	grpc.ClientStream
}

type compactTxStreamerGetTaddressTxidsClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetTaddressTxidsClient) Recv() (*RawTransaction, error) {
	m := new(RawTransaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetTaddressBalance(ctx context.Context, in *AddressList, opts ...grpc.CallOption) (*Balance, error) {
	out := new(Balance)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetTaddressBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetTaddressBalanceStream(ctx context.Context, opts ...grpc.CallOption) (CompactTxStreamer_GetTaddressBalanceStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompactTxStreamer_ServiceDesc.Streams[2], "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetTaddressBalanceStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetTaddressBalanceStreamClient{stream}
	return x, nil
}

type CompactTxStreamer_GetTaddressBalanceStreamClient interface {
	Send(*Address) error
	CloseAndRecv() (*Balance, error)
	grpc.ClientStream
}

type compactTxStreamerGetTaddressBalanceStreamClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetTaddressBalanceStreamClient) Send(m *Address) error {
	return x.ClientStream.SendMsg(m)
}

func (x *compactTxStreamerGetTaddressBalanceStreamClient) CloseAndRecv() (*Balance, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Balance)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetMempoolTx(ctx context.Context, in *Exclude, opts ...grpc.CallOption) (CompactTxStreamer_GetMempoolTxClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompactTxStreamer_ServiceDesc.Streams[3], "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetMempoolTx", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetMempoolTxClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompactTxStreamer_GetMempoolTxClient interface {
	Recv() (*CompactTx, error)
	grpc.ClientStream
}

type compactTxStreamerGetMempoolTxClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetMempoolTxClient) Recv() (*CompactTx, error) {
	m := new(CompactTx)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetMempoolStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CompactTxStreamer_GetMempoolStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompactTxStreamer_ServiceDesc.Streams[4], "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetMempoolStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetMempoolStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompactTxStreamer_GetMempoolStreamClient interface {
	Recv() (*RawTransaction, error)
	grpc.ClientStream
}

type compactTxStreamerGetMempoolStreamClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetMempoolStreamClient) Recv() (*RawTransaction, error) {
	m := new(RawTransaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetTreeState(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*TreeState, error) {
	out := new(TreeState)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetTreeState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetAddressUtxos(ctx context.Context, in *GetAddressUtxosArg, opts ...grpc.CallOption) (*GetAddressUtxosReplyList, error) {
	out := new(GetAddressUtxosReplyList)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetAddressUtxos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetAddressUtxosStream(ctx context.Context, in *GetAddressUtxosArg, opts ...grpc.CallOption) (CompactTxStreamer_GetAddressUtxosStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompactTxStreamer_ServiceDesc.Streams[5], "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetAddressUtxosStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetAddressUtxosStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompactTxStreamer_GetAddressUtxosStreamClient interface {
	Recv() (*GetAddressUtxosReply, error)
	grpc.ClientStream
}

type compactTxStreamerGetAddressUtxosStreamClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetAddressUtxosStreamClient) Recv() (*GetAddressUtxosReply, error) {
	m := new(GetAddressUtxosReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetLightdInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LightdInfo, error) {
	out := new(LightdInfo)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetLightdInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) Ping(ctx context.Context, in *Duration, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/pirate.wallet.sdk.rpc.CompactTxStreamer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompactTxStreamerServer is the server API for CompactTxStreamer service.
// All implementations must embed UnimplementedCompactTxStreamerServer
// for forward compatibility
type CompactTxStreamerServer interface {
	GetLiteWalletBlockGroup(context.Context, *BlockID) (*BlockID, error)
	// Return the height of the tip of the best chain
	GetLatestBlock(context.Context, *ChainSpec) (*BlockID, error)
	// Return the compact block corresponding to the given block identifier
	GetBlock(context.Context, *BlockID) (*CompactBlock, error)
	// Return a list of consecutive compact blocks
	GetBlockRange(*BlockRange, CompactTxStreamer_GetBlockRangeServer) error
	// Get the historical and current prices
	GetARRRPrice(context.Context, *PriceRequest) (*PriceResponse, error)
	GetCurrentARRRPrice(context.Context, *Empty) (*PriceResponse, error)
	// Return the requested full (not compact) transaction (as from pirated)
	GetTransaction(context.Context, *TxFilter) (*RawTransaction, error)
	// Submit the given transaction to the Zcash network
	SendTransaction(context.Context, *RawTransaction) (*SendResponse, error)
	// Return the txids corresponding to the given t-address within the given block range
	GetTaddressTxids(*TransparentAddressBlockFilter, CompactTxStreamer_GetTaddressTxidsServer) error
	GetTaddressBalance(context.Context, *AddressList) (*Balance, error)
	GetTaddressBalanceStream(CompactTxStreamer_GetTaddressBalanceStreamServer) error
	// Return the compact transactions currently in the mempool; the results
	// can be a few seconds out of date. If the Exclude list is empty, return
	// all transactions; otherwise return all *except* those in the Exclude list
	// (if any); this allows the client to avoid receiving transactions that it
	// already has (from an earlier call to this rpc). The transaction IDs in the
	// Exclude list can be shortened to any number of bytes to make the request
	// more bandwidth-efficient; if two or more transactions in the mempool
	// match a shortened txid, they are all sent (none is excluded). Transactions
	// in the exclude list that don't exist in the mempool are ignored.
	GetMempoolTx(*Exclude, CompactTxStreamer_GetMempoolTxServer) error
	// Return a stream of current Mempool transactions. This will keep the output stream open while
	// there are mempool transactions. It will close the returned stream when a new block is mined.
	GetMempoolStream(*Empty, CompactTxStreamer_GetMempoolStreamServer) error
	// GetTreeState returns the note commitment tree state corresponding to the given block.
	// See section 3.7 of the Zcash protocol specification. It returns several other useful
	// values also (even though they can be obtained using GetBlock).
	// The block can be specified by either height or hash.
	GetTreeState(context.Context, *BlockID) (*TreeState, error)
	GetAddressUtxos(context.Context, *GetAddressUtxosArg) (*GetAddressUtxosReplyList, error)
	GetAddressUtxosStream(*GetAddressUtxosArg, CompactTxStreamer_GetAddressUtxosStreamServer) error
	// Return information about this lightwalletd instance and the blockchain
	GetLightdInfo(context.Context, *Empty) (*LightdInfo, error)
	// Testing-only, requires lightwalletd --ping-very-insecure (do not enable in production)
	Ping(context.Context, *Duration) (*PingResponse, error)
	mustEmbedUnimplementedCompactTxStreamerServer()
}

// UnimplementedCompactTxStreamerServer must be embedded to have forward compatible implementations.
type UnimplementedCompactTxStreamerServer struct {
}

func (UnimplementedCompactTxStreamerServer) GetLiteWalletBlockGroup(context.Context, *BlockID) (*BlockID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiteWalletBlockGroup not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetLatestBlock(context.Context, *ChainSpec) (*BlockID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlock not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetBlock(context.Context, *BlockID) (*CompactBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetBlockRange(*BlockRange, CompactTxStreamer_GetBlockRangeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlockRange not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetARRRPrice(context.Context, *PriceRequest) (*PriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetARRRPrice not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetCurrentARRRPrice(context.Context, *Empty) (*PriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentARRRPrice not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetTransaction(context.Context, *TxFilter) (*RawTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedCompactTxStreamerServer) SendTransaction(context.Context, *RawTransaction) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTransaction not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetTaddressTxids(*TransparentAddressBlockFilter, CompactTxStreamer_GetTaddressTxidsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTaddressTxids not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetTaddressBalance(context.Context, *AddressList) (*Balance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaddressBalance not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetTaddressBalanceStream(CompactTxStreamer_GetTaddressBalanceStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTaddressBalanceStream not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetMempoolTx(*Exclude, CompactTxStreamer_GetMempoolTxServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMempoolTx not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetMempoolStream(*Empty, CompactTxStreamer_GetMempoolStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMempoolStream not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetTreeState(context.Context, *BlockID) (*TreeState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTreeState not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetAddressUtxos(context.Context, *GetAddressUtxosArg) (*GetAddressUtxosReplyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressUtxos not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetAddressUtxosStream(*GetAddressUtxosArg, CompactTxStreamer_GetAddressUtxosStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAddressUtxosStream not implemented")
}
func (UnimplementedCompactTxStreamerServer) GetLightdInfo(context.Context, *Empty) (*LightdInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLightdInfo not implemented")
}
func (UnimplementedCompactTxStreamerServer) Ping(context.Context, *Duration) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCompactTxStreamerServer) mustEmbedUnimplementedCompactTxStreamerServer() {}

// UnsafeCompactTxStreamerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompactTxStreamerServer will
// result in compilation errors.
type UnsafeCompactTxStreamerServer interface {
	mustEmbedUnimplementedCompactTxStreamerServer()
}

func RegisterCompactTxStreamerServer(s grpc.ServiceRegistrar, srv CompactTxStreamerServer) {
	s.RegisterService(&CompactTxStreamer_ServiceDesc, srv)
}

func _CompactTxStreamer_GetLiteWalletBlockGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetLiteWalletBlockGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetLiteWalletBlockGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetLiteWalletBlockGroup(ctx, req.(*BlockID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetLatestBlock(ctx, req.(*ChainSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetBlock(ctx, req.(*BlockID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetBlockRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockRange)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompactTxStreamerServer).GetBlockRange(m, &compactTxStreamerGetBlockRangeServer{stream})
}

type CompactTxStreamer_GetBlockRangeServer interface {
	Send(*CompactBlock) error
	grpc.ServerStream
}

type compactTxStreamerGetBlockRangeServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetBlockRangeServer) Send(m *CompactBlock) error {
	return x.ServerStream.SendMsg(m)
}

func _CompactTxStreamer_GetARRRPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetARRRPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetARRRPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetARRRPrice(ctx, req.(*PriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetCurrentARRRPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetCurrentARRRPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetCurrentARRRPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetCurrentARRRPrice(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetTransaction(ctx, req.(*TxFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).SendTransaction(ctx, req.(*RawTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetTaddressTxids_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransparentAddressBlockFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompactTxStreamerServer).GetTaddressTxids(m, &compactTxStreamerGetTaddressTxidsServer{stream})
}

type CompactTxStreamer_GetTaddressTxidsServer interface {
	Send(*RawTransaction) error
	grpc.ServerStream
}

type compactTxStreamerGetTaddressTxidsServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetTaddressTxidsServer) Send(m *RawTransaction) error {
	return x.ServerStream.SendMsg(m)
}

func _CompactTxStreamer_GetTaddressBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetTaddressBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetTaddressBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetTaddressBalance(ctx, req.(*AddressList))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetTaddressBalanceStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CompactTxStreamerServer).GetTaddressBalanceStream(&compactTxStreamerGetTaddressBalanceStreamServer{stream})
}

type CompactTxStreamer_GetTaddressBalanceStreamServer interface {
	SendAndClose(*Balance) error
	Recv() (*Address, error)
	grpc.ServerStream
}

type compactTxStreamerGetTaddressBalanceStreamServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetTaddressBalanceStreamServer) SendAndClose(m *Balance) error {
	return x.ServerStream.SendMsg(m)
}

func (x *compactTxStreamerGetTaddressBalanceStreamServer) Recv() (*Address, error) {
	m := new(Address)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CompactTxStreamer_GetMempoolTx_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Exclude)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompactTxStreamerServer).GetMempoolTx(m, &compactTxStreamerGetMempoolTxServer{stream})
}

type CompactTxStreamer_GetMempoolTxServer interface {
	Send(*CompactTx) error
	grpc.ServerStream
}

type compactTxStreamerGetMempoolTxServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetMempoolTxServer) Send(m *CompactTx) error {
	return x.ServerStream.SendMsg(m)
}

func _CompactTxStreamer_GetMempoolStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompactTxStreamerServer).GetMempoolStream(m, &compactTxStreamerGetMempoolStreamServer{stream})
}

type CompactTxStreamer_GetMempoolStreamServer interface {
	Send(*RawTransaction) error
	grpc.ServerStream
}

type compactTxStreamerGetMempoolStreamServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetMempoolStreamServer) Send(m *RawTransaction) error {
	return x.ServerStream.SendMsg(m)
}

func _CompactTxStreamer_GetTreeState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetTreeState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetTreeState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetTreeState(ctx, req.(*BlockID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetAddressUtxos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressUtxosArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetAddressUtxos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetAddressUtxos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetAddressUtxos(ctx, req.(*GetAddressUtxosArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetAddressUtxosStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAddressUtxosArg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompactTxStreamerServer).GetAddressUtxosStream(m, &compactTxStreamerGetAddressUtxosStreamServer{stream})
}

type CompactTxStreamer_GetAddressUtxosStreamServer interface {
	Send(*GetAddressUtxosReply) error
	grpc.ServerStream
}

type compactTxStreamerGetAddressUtxosStreamServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetAddressUtxosStreamServer) Send(m *GetAddressUtxosReply) error {
	return x.ServerStream.SendMsg(m)
}

func _CompactTxStreamer_GetLightdInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetLightdInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/GetLightdInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetLightdInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Duration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pirate.wallet.sdk.rpc.CompactTxStreamer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).Ping(ctx, req.(*Duration))
	}
	return interceptor(ctx, in, info, handler)
}

// CompactTxStreamer_ServiceDesc is the grpc.ServiceDesc for CompactTxStreamer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompactTxStreamer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pirate.wallet.sdk.rpc.CompactTxStreamer",
	HandlerType: (*CompactTxStreamerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLiteWalletBlockGroup",
			Handler:    _CompactTxStreamer_GetLiteWalletBlockGroup_Handler,
		},
		{
			MethodName: "GetLatestBlock",
			Handler:    _CompactTxStreamer_GetLatestBlock_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _CompactTxStreamer_GetBlock_Handler,
		},
		{
			MethodName: "GetARRRPrice",
			Handler:    _CompactTxStreamer_GetARRRPrice_Handler,
		},
		{
			MethodName: "GetCurrentARRRPrice",
			Handler:    _CompactTxStreamer_GetCurrentARRRPrice_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _CompactTxStreamer_GetTransaction_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _CompactTxStreamer_SendTransaction_Handler,
		},
		{
			MethodName: "GetTaddressBalance",
			Handler:    _CompactTxStreamer_GetTaddressBalance_Handler,
		},
		{
			MethodName: "GetTreeState",
			Handler:    _CompactTxStreamer_GetTreeState_Handler,
		},
		{
			MethodName: "GetAddressUtxos",
			Handler:    _CompactTxStreamer_GetAddressUtxos_Handler,
		},
		{
			MethodName: "GetLightdInfo",
			Handler:    _CompactTxStreamer_GetLightdInfo_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _CompactTxStreamer_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlockRange",
			Handler:       _CompactTxStreamer_GetBlockRange_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTaddressTxids",
			Handler:       _CompactTxStreamer_GetTaddressTxids_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTaddressBalanceStream",
			Handler:       _CompactTxStreamer_GetTaddressBalanceStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetMempoolTx",
			Handler:       _CompactTxStreamer_GetMempoolTx_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMempoolStream",
			Handler:       _CompactTxStreamer_GetMempoolStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAddressUtxosStream",
			Handler:       _CompactTxStreamer_GetAddressUtxosStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
