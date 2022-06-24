// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateGame struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Red     string `protobuf:"bytes,2,opt,name=red,proto3" json:"red,omitempty"`
	Black   string `protobuf:"bytes,3,opt,name=black,proto3" json:"black,omitempty"`
}

func (m *MsgCreateGame) Reset()         { *m = MsgCreateGame{} }
func (m *MsgCreateGame) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGame) ProtoMessage()    {}
func (*MsgCreateGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89f7ca8d0309536, []int{0}
}
func (m *MsgCreateGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGame.Merge(m, src)
}
func (m *MsgCreateGame) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGame) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGame.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGame proto.InternalMessageInfo

func (m *MsgCreateGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateGame) GetRed() string {
	if m != nil {
		return m.Red
	}
	return ""
}

func (m *MsgCreateGame) GetBlack() string {
	if m != nil {
		return m.Black
	}
	return ""
}

type MsgCreateGameResponse struct {
	IdValue string `protobuf:"bytes,1,opt,name=idValue,proto3" json:"idValue,omitempty"`
}

func (m *MsgCreateGameResponse) Reset()         { *m = MsgCreateGameResponse{} }
func (m *MsgCreateGameResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGameResponse) ProtoMessage()    {}
func (*MsgCreateGameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89f7ca8d0309536, []int{1}
}
func (m *MsgCreateGameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGameResponse.Merge(m, src)
}
func (m *MsgCreateGameResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGameResponse proto.InternalMessageInfo

func (m *MsgCreateGameResponse) GetIdValue() string {
	if m != nil {
		return m.IdValue
	}
	return ""
}

type MsgPlayMove struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	IdValue string `protobuf:"bytes,2,opt,name=idValue,proto3" json:"idValue,omitempty"`
	FromX   uint64 `protobuf:"varint,3,opt,name=fromX,proto3" json:"fromX,omitempty"`
	FromY   uint64 `protobuf:"varint,4,opt,name=fromY,proto3" json:"fromY,omitempty"`
	ToX     uint64 `protobuf:"varint,5,opt,name=toX,proto3" json:"toX,omitempty"`
	ToY     uint64 `protobuf:"varint,6,opt,name=toY,proto3" json:"toY,omitempty"`
}

func (m *MsgPlayMove) Reset()         { *m = MsgPlayMove{} }
func (m *MsgPlayMove) String() string { return proto.CompactTextString(m) }
func (*MsgPlayMove) ProtoMessage()    {}
func (*MsgPlayMove) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89f7ca8d0309536, []int{2}
}
func (m *MsgPlayMove) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlayMove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlayMove.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlayMove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlayMove.Merge(m, src)
}
func (m *MsgPlayMove) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlayMove) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlayMove.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlayMove proto.InternalMessageInfo

func (m *MsgPlayMove) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgPlayMove) GetIdValue() string {
	if m != nil {
		return m.IdValue
	}
	return ""
}

func (m *MsgPlayMove) GetFromX() uint64 {
	if m != nil {
		return m.FromX
	}
	return 0
}

func (m *MsgPlayMove) GetFromY() uint64 {
	if m != nil {
		return m.FromY
	}
	return 0
}

func (m *MsgPlayMove) GetToX() uint64 {
	if m != nil {
		return m.ToX
	}
	return 0
}

func (m *MsgPlayMove) GetToY() uint64 {
	if m != nil {
		return m.ToY
	}
	return 0
}

type MsgPlayMoveResponse struct {
	IdValue   string `protobuf:"bytes,1,opt,name=idValue,proto3" json:"idValue,omitempty"`
	CapturedX int64  `protobuf:"varint,2,opt,name=capturedX,proto3" json:"capturedX,omitempty"`
	CapturedY int64  `protobuf:"varint,3,opt,name=capturedY,proto3" json:"capturedY,omitempty"`
	Winner    string `protobuf:"bytes,4,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (m *MsgPlayMoveResponse) Reset()         { *m = MsgPlayMoveResponse{} }
func (m *MsgPlayMoveResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPlayMoveResponse) ProtoMessage()    {}
func (*MsgPlayMoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89f7ca8d0309536, []int{3}
}
func (m *MsgPlayMoveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlayMoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlayMoveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlayMoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlayMoveResponse.Merge(m, src)
}
func (m *MsgPlayMoveResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlayMoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlayMoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlayMoveResponse proto.InternalMessageInfo

func (m *MsgPlayMoveResponse) GetIdValue() string {
	if m != nil {
		return m.IdValue
	}
	return ""
}

func (m *MsgPlayMoveResponse) GetCapturedX() int64 {
	if m != nil {
		return m.CapturedX
	}
	return 0
}

func (m *MsgPlayMoveResponse) GetCapturedY() int64 {
	if m != nil {
		return m.CapturedY
	}
	return 0
}

func (m *MsgPlayMoveResponse) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgCreateGame)(nil), "dvirtayeb.checkers.checkers.MsgCreateGame")
	proto.RegisterType((*MsgCreateGameResponse)(nil), "dvirtayeb.checkers.checkers.MsgCreateGameResponse")
	proto.RegisterType((*MsgPlayMove)(nil), "dvirtayeb.checkers.checkers.MsgPlayMove")
	proto.RegisterType((*MsgPlayMoveResponse)(nil), "dvirtayeb.checkers.checkers.MsgPlayMoveResponse")
}

func init() { proto.RegisterFile("checkers/tx.proto", fileDescriptor_b89f7ca8d0309536) }

var fileDescriptor_b89f7ca8d0309536 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x65, 0x28, 0xf0, 0x7d, 0xdc, 0x2f, 0x5f, 0xa2, 0xf5, 0x27, 0x0d, 0x9a, 0xc6, 0x74, 0x45,
	0x5c, 0x14, 0xc5, 0x37, 0xd0, 0x85, 0x6e, 0x9a, 0x68, 0x17, 0x86, 0xba, 0x9b, 0xb6, 0x43, 0x69,
	0x28, 0x4c, 0x33, 0x33, 0x20, 0xac, 0x7d, 0x01, 0x7d, 0x2b, 0x97, 0x2c, 0x4d, 0xdc, 0x18, 0x78,
	0x11, 0xd3, 0x81, 0x29, 0x65, 0x03, 0xec, 0xce, 0x39, 0xf7, 0xf6, 0xf4, 0xdc, 0xb9, 0x17, 0x0e,
	0x83, 0x1e, 0x09, 0xfa, 0x84, 0xf1, 0x96, 0x98, 0xd8, 0x29, 0xa3, 0x82, 0xea, 0x67, 0xe1, 0x38,
	0x66, 0x02, 0x4f, 0x89, 0x6f, 0xab, 0x62, 0x0e, 0xac, 0x27, 0xf8, 0xef, 0xf0, 0xe8, 0x8e, 0x11,
	0x2c, 0xc8, 0x3d, 0x1e, 0x10, 0xdd, 0x80, 0x3f, 0x41, 0xc6, 0x28, 0x33, 0xd0, 0x05, 0x6a, 0xd6,
	0x5d, 0x45, 0xf5, 0x03, 0xd0, 0x18, 0x09, 0x8d, 0xb2, 0x54, 0x33, 0xa8, 0x1f, 0x43, 0xd5, 0x4f,
	0x70, 0xd0, 0x37, 0x34, 0xa9, 0x2d, 0x89, 0x75, 0x0d, 0x27, 0x1b, 0x96, 0x2e, 0xe1, 0x29, 0x1d,
	0x72, 0x69, 0x1d, 0x87, 0xcf, 0x38, 0x19, 0x11, 0x65, 0xbd, 0xa2, 0xd6, 0x07, 0x82, 0x7f, 0x0e,
	0x8f, 0x1e, 0x13, 0x3c, 0x75, 0xe8, 0x78, 0x5b, 0x88, 0x82, 0x47, 0x79, 0xc3, 0x23, 0x0b, 0xd3,
	0x65, 0x74, 0xd0, 0x91, 0x61, 0x2a, 0xee, 0x92, 0x28, 0xd5, 0x33, 0x2a, 0x6b, 0xd5, 0xcb, 0x46,
	0x11, 0xb4, 0x63, 0x54, 0xa5, 0x96, 0xc1, 0xa5, 0xe2, 0x19, 0x35, 0xa5, 0x78, 0xd6, 0x1b, 0x82,
	0xa3, 0x42, 0xa6, 0xdd, 0x53, 0xe8, 0xe7, 0x50, 0x0f, 0x70, 0x2a, 0x46, 0x8c, 0x84, 0x1d, 0x99,
	0x4e, 0x73, 0xd7, 0x42, 0xb1, 0xea, 0xc9, 0x8c, 0x85, 0xaa, 0xa7, 0x9f, 0x42, 0xed, 0x35, 0x1e,
	0x0e, 0x09, 0x93, 0x41, 0xeb, 0xee, 0x8a, 0xb5, 0xbf, 0x11, 0x68, 0x0e, 0x8f, 0xf4, 0x04, 0xa0,
	0xb0, 0xa4, 0x4b, 0x7b, 0xcb, 0x4e, 0xed, 0x8d, 0xd7, 0x6f, 0xb4, 0xf7, 0xef, 0xcd, 0x67, 0xec,
	0xc2, 0xdf, 0x7c, 0x17, 0xcd, 0x5d, 0xdf, 0xab, 0xce, 0xc6, 0xd5, 0xbe, 0x9d, 0xea, 0x3f, 0xb7,
	0x0f, 0x9f, 0x73, 0x13, 0xcd, 0xe6, 0x26, 0xfa, 0x99, 0x9b, 0xe8, 0x7d, 0x61, 0x96, 0x66, 0x0b,
	0xb3, 0xf4, 0xb5, 0x30, 0x4b, 0x2f, 0x76, 0x14, 0x8b, 0xde, 0xc8, 0xb7, 0x03, 0x3a, 0x68, 0xe5,
	0xae, 0xad, 0xfc, 0xb8, 0x27, 0x6b, 0x28, 0xa6, 0x29, 0xe1, 0x7e, 0x4d, 0xde, 0xfa, 0xcd, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0xe3, 0xbc, 0x2d, 0x00, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateGame(ctx context.Context, in *MsgCreateGame, opts ...grpc.CallOption) (*MsgCreateGameResponse, error)
	PlayMove(ctx context.Context, in *MsgPlayMove, opts ...grpc.CallOption) (*MsgPlayMoveResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateGame(ctx context.Context, in *MsgCreateGame, opts ...grpc.CallOption) (*MsgCreateGameResponse, error) {
	out := new(MsgCreateGameResponse)
	err := c.cc.Invoke(ctx, "/dvirtayeb.checkers.checkers.Msg/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PlayMove(ctx context.Context, in *MsgPlayMove, opts ...grpc.CallOption) (*MsgPlayMoveResponse, error) {
	out := new(MsgPlayMoveResponse)
	err := c.cc.Invoke(ctx, "/dvirtayeb.checkers.checkers.Msg/PlayMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateGame(context.Context, *MsgCreateGame) (*MsgCreateGameResponse, error)
	PlayMove(context.Context, *MsgPlayMove) (*MsgPlayMoveResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateGame(ctx context.Context, req *MsgCreateGame) (*MsgCreateGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (*UnimplementedMsgServer) PlayMove(ctx context.Context, req *MsgPlayMove) (*MsgPlayMoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayMove not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dvirtayeb.checkers.checkers.Msg/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateGame(ctx, req.(*MsgCreateGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PlayMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPlayMove)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PlayMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dvirtayeb.checkers.checkers.Msg/PlayMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PlayMove(ctx, req.(*MsgPlayMove))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dvirtayeb.checkers.checkers.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGame",
			Handler:    _Msg_CreateGame_Handler,
		},
		{
			MethodName: "PlayMove",
			Handler:    _Msg_PlayMove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkers/tx.proto",
}

func (m *MsgCreateGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Black) > 0 {
		i -= len(m.Black)
		copy(dAtA[i:], m.Black)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Black)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Red) > 0 {
		i -= len(m.Red)
		copy(dAtA[i:], m.Red)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Red)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateGameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IdValue) > 0 {
		i -= len(m.IdValue)
		copy(dAtA[i:], m.IdValue)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IdValue)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPlayMove) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlayMove) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlayMove) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ToY != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ToY))
		i--
		dAtA[i] = 0x30
	}
	if m.ToX != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ToX))
		i--
		dAtA[i] = 0x28
	}
	if m.FromY != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.FromY))
		i--
		dAtA[i] = 0x20
	}
	if m.FromX != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.FromX))
		i--
		dAtA[i] = 0x18
	}
	if len(m.IdValue) > 0 {
		i -= len(m.IdValue)
		copy(dAtA[i:], m.IdValue)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IdValue)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPlayMoveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlayMoveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlayMoveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0x22
	}
	if m.CapturedY != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CapturedY))
		i--
		dAtA[i] = 0x18
	}
	if m.CapturedX != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CapturedX))
		i--
		dAtA[i] = 0x10
	}
	if len(m.IdValue) > 0 {
		i -= len(m.IdValue)
		copy(dAtA[i:], m.IdValue)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IdValue)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Red)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Black)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateGameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IdValue)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgPlayMove) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.IdValue)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.FromX != 0 {
		n += 1 + sovTx(uint64(m.FromX))
	}
	if m.FromY != 0 {
		n += 1 + sovTx(uint64(m.FromY))
	}
	if m.ToX != 0 {
		n += 1 + sovTx(uint64(m.ToX))
	}
	if m.ToY != 0 {
		n += 1 + sovTx(uint64(m.ToY))
	}
	return n
}

func (m *MsgPlayMoveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IdValue)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CapturedX != 0 {
		n += 1 + sovTx(uint64(m.CapturedX))
	}
	if m.CapturedY != 0 {
		n += 1 + sovTx(uint64(m.CapturedY))
	}
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Red", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Red = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Black", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Black = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateGameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateGameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgPlayMove) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPlayMove: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlayMove: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromX", wireType)
			}
			m.FromX = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromX |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromY", wireType)
			}
			m.FromY = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromY |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToX", wireType)
			}
			m.ToX = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ToX |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToY", wireType)
			}
			m.ToY = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ToY |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgPlayMoveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgPlayMoveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlayMoveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapturedX", wireType)
			}
			m.CapturedX = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CapturedX |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapturedY", wireType)
			}
			m.CapturedY = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CapturedY |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)