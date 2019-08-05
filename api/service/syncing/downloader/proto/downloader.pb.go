// Code generated by protoc-gen-go. DO NOT EDIT.
// source: downloader.proto

package downloader

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DownloaderRequest_RequestType int32

const (
	DownloaderRequest_HEADER          DownloaderRequest_RequestType = 0
	DownloaderRequest_BLOCK           DownloaderRequest_RequestType = 1
	DownloaderRequest_NEWBLOCK        DownloaderRequest_RequestType = 2
	DownloaderRequest_BLOCKHEIGHT     DownloaderRequest_RequestType = 3
	DownloaderRequest_REGISTER        DownloaderRequest_RequestType = 4
	DownloaderRequest_REGISTERTIMEOUT DownloaderRequest_RequestType = 5
	DownloaderRequest_UNKNOWN         DownloaderRequest_RequestType = 6
)

var DownloaderRequest_RequestType_name = map[int32]string{
	0: "HEADER",
	1: "BLOCK",
	2: "NEWBLOCK",
	3: "BLOCKHEIGHT",
	4: "REGISTER",
	5: "REGISTERTIMEOUT",
	6: "UNKNOWN",
}

var DownloaderRequest_RequestType_value = map[string]int32{
	"HEADER":          0,
	"BLOCK":           1,
	"NEWBLOCK":        2,
	"BLOCKHEIGHT":     3,
	"REGISTER":        4,
	"REGISTERTIMEOUT": 5,
	"UNKNOWN":         6,
}

func (x DownloaderRequest_RequestType) String() string {
	return proto.EnumName(DownloaderRequest_RequestType_name, int32(x))
}

func (DownloaderRequest_RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a99ec95c7ab1ff1, []int{0, 0}
}

type DownloaderResponse_RegisterResponseType int32

const (
	DownloaderResponse_SUCCESS DownloaderResponse_RegisterResponseType = 0
	DownloaderResponse_FAIL    DownloaderResponse_RegisterResponseType = 1
	DownloaderResponse_INSYNC  DownloaderResponse_RegisterResponseType = 2
)

var DownloaderResponse_RegisterResponseType_name = map[int32]string{
	0: "SUCCESS",
	1: "FAIL",
	2: "INSYNC",
}

var DownloaderResponse_RegisterResponseType_value = map[string]int32{
	"SUCCESS": 0,
	"FAIL":    1,
	"INSYNC":  2,
}

func (x DownloaderResponse_RegisterResponseType) String() string {
	return proto.EnumName(DownloaderResponse_RegisterResponseType_name, int32(x))
}

func (DownloaderResponse_RegisterResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a99ec95c7ab1ff1, []int{1, 0}
}

// DownloaderRequest is the generic download request.
type DownloaderRequest struct {
	// Request type.
	Type DownloaderRequest_RequestType `protobuf:"varint,1,opt,name=type,proto3,enum=downloader.DownloaderRequest_RequestType" json:"type,omitempty"`
	// The hashes of the blocks we want to download.
	Hashes               [][]byte `protobuf:"bytes,2,rep,name=hashes,proto3" json:"hashes,omitempty"`
	PeerHash             []byte   `protobuf:"bytes,3,opt,name=peerHash,proto3" json:"peerHash,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,4,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Ip                   string   `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 string   `protobuf:"bytes,6,opt,name=port,proto3" json:"port,omitempty"`
	Size                 uint32   `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloaderRequest) Reset()         { *m = DownloaderRequest{} }
func (m *DownloaderRequest) String() string { return proto.CompactTextString(m) }
func (*DownloaderRequest) ProtoMessage()    {}
func (*DownloaderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a99ec95c7ab1ff1, []int{0}
}

func (m *DownloaderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloaderRequest.Unmarshal(m, b)
}
func (m *DownloaderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloaderRequest.Marshal(b, m, deterministic)
}
func (m *DownloaderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloaderRequest.Merge(m, src)
}
func (m *DownloaderRequest) XXX_Size() int {
	return xxx_messageInfo_DownloaderRequest.Size(m)
}
func (m *DownloaderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloaderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloaderRequest proto.InternalMessageInfo

func (m *DownloaderRequest) GetType() DownloaderRequest_RequestType {
	if m != nil {
		return m.Type
	}
	return DownloaderRequest_HEADER
}

func (m *DownloaderRequest) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func (m *DownloaderRequest) GetPeerHash() []byte {
	if m != nil {
		return m.PeerHash
	}
	return nil
}

func (m *DownloaderRequest) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *DownloaderRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DownloaderRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *DownloaderRequest) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

// DownloaderResponse is the generic response of DownloaderRequest.
type DownloaderResponse struct {
	// payload of Block.
	Payload [][]byte `protobuf:"bytes,1,rep,name=payload,proto3" json:"payload,omitempty"`
	// response of registration request
	Type                 DownloaderResponse_RegisterResponseType `protobuf:"varint,2,opt,name=type,proto3,enum=downloader.DownloaderResponse_RegisterResponseType" json:"type,omitempty"`
	BlockHeight          uint64                                  `protobuf:"varint,3,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *DownloaderResponse) Reset()         { *m = DownloaderResponse{} }
func (m *DownloaderResponse) String() string { return proto.CompactTextString(m) }
func (*DownloaderResponse) ProtoMessage()    {}
func (*DownloaderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a99ec95c7ab1ff1, []int{1}
}

func (m *DownloaderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloaderResponse.Unmarshal(m, b)
}
func (m *DownloaderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloaderResponse.Marshal(b, m, deterministic)
}
func (m *DownloaderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloaderResponse.Merge(m, src)
}
func (m *DownloaderResponse) XXX_Size() int {
	return xxx_messageInfo_DownloaderResponse.Size(m)
}
func (m *DownloaderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloaderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloaderResponse proto.InternalMessageInfo

func (m *DownloaderResponse) GetPayload() [][]byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *DownloaderResponse) GetType() DownloaderResponse_RegisterResponseType {
	if m != nil {
		return m.Type
	}
	return DownloaderResponse_SUCCESS
}

func (m *DownloaderResponse) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterEnum("downloader.DownloaderRequest_RequestType", DownloaderRequest_RequestType_name, DownloaderRequest_RequestType_value)
	proto.RegisterEnum("downloader.DownloaderResponse_RegisterResponseType", DownloaderResponse_RegisterResponseType_name, DownloaderResponse_RegisterResponseType_value)
	proto.RegisterType((*DownloaderRequest)(nil), "downloader.DownloaderRequest")
	proto.RegisterType((*DownloaderResponse)(nil), "downloader.DownloaderResponse")
}

func init() { proto.RegisterFile("downloader.proto", fileDescriptor_6a99ec95c7ab1ff1) }

var fileDescriptor_6a99ec95c7ab1ff1 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x8e, 0xe3, 0x24, 0x93, 0xd0, 0x2e, 0x03, 0x42, 0xab, 0x0a, 0x90, 0xe5, 0x93,
	0xb9, 0xf8, 0xd0, 0x9e, 0x38, 0x70, 0x28, 0xee, 0x12, 0x5b, 0x2d, 0x8e, 0x58, 0x3b, 0x54, 0x1c,
	0x5d, 0xba, 0xaa, 0x2d, 0xaa, 0x7a, 0xf1, 0xba, 0x42, 0xe6, 0x2d, 0x91, 0x78, 0x20, 0xe4, 0x4d,
	0x5a, 0x5b, 0x02, 0x72, 0xda, 0xf9, 0xff, 0xd9, 0x19, 0xcd, 0x7c, 0x1a, 0xa0, 0xd7, 0xd5, 0x8f,
	0xbb, 0xdb, 0x2a, 0xbf, 0x96, 0x75, 0xa0, 0xea, 0xaa, 0xa9, 0x10, 0x7a, 0xc7, 0xfb, 0x65, 0xc1,
	0xd3, 0xb3, 0x47, 0x29, 0xe4, 0xf7, 0x7b, 0xa9, 0x1b, 0x7c, 0x07, 0x76, 0xd3, 0x2a, 0xc9, 0x88,
	0x4b, 0xfc, 0x83, 0xe3, 0x37, 0xc1, 0xa0, 0xc5, 0x5f, 0x9f, 0x83, 0xdd, 0x9b, 0xb5, 0x4a, 0x0a,
	0x53, 0x86, 0x2f, 0xc0, 0x29, 0x72, 0x5d, 0x48, 0xcd, 0x2c, 0x77, 0xec, 0x2f, 0xc5, 0x4e, 0xe1,
	0x11, 0xcc, 0x94, 0x94, 0x75, 0x94, 0xeb, 0x82, 0x8d, 0x5d, 0xe2, 0x2f, 0xc5, 0xa3, 0xc6, 0x97,
	0x30, 0xbf, 0xba, 0xad, 0xbe, 0x7e, 0x33, 0x49, 0xdb, 0x24, 0x7b, 0x03, 0x0f, 0xc0, 0x2a, 0x15,
	0x9b, 0xb8, 0xc4, 0x9f, 0x0b, 0xab, 0x54, 0x88, 0x60, 0xab, 0xaa, 0x6e, 0x98, 0x63, 0x1c, 0x13,
	0x77, 0x9e, 0x2e, 0x7f, 0x4a, 0x36, 0x75, 0x89, 0xff, 0x44, 0x98, 0xd8, 0xd3, 0xb0, 0x18, 0x8c,
	0x87, 0x00, 0x4e, 0xc4, 0x4f, 0xcf, 0xb8, 0xa0, 0x23, 0x9c, 0xc3, 0xe4, 0xfd, 0xc5, 0x3a, 0x3c,
	0xa7, 0x04, 0x97, 0x30, 0x4b, 0xf8, 0xe5, 0x56, 0x59, 0x78, 0x08, 0x0b, 0x13, 0x46, 0x3c, 0x5e,
	0x45, 0x19, 0x1d, 0x77, 0x69, 0xc1, 0x57, 0x71, 0x9a, 0x71, 0x41, 0x6d, 0x7c, 0x06, 0x87, 0x0f,
	0x2a, 0x8b, 0x3f, 0xf2, 0xf5, 0x26, 0xa3, 0x13, 0x5c, 0xc0, 0x74, 0x93, 0x9c, 0x27, 0xeb, 0xcb,
	0x84, 0x3a, 0xde, 0x6f, 0x02, 0x38, 0xc4, 0xa4, 0x55, 0x75, 0xa7, 0x25, 0x32, 0x98, 0xaa, 0xbc,
	0xed, 0x4c, 0x46, 0x0c, 0x96, 0x07, 0x89, 0xab, 0x1d, 0x6e, 0xcb, 0xe0, 0x3e, 0xf9, 0x1f, 0xee,
	0x6d, 0x9f, 0x40, 0xc8, 0x9b, 0x52, 0x37, 0xbd, 0x31, 0x00, 0xef, 0xc2, 0x62, 0xcb, 0x4c, 0x96,
	0x37, 0x45, 0x63, 0x18, 0xdb, 0x62, 0x68, 0x79, 0x6f, 0xe1, 0xf9, 0xbf, 0xea, 0xbb, 0x05, 0xd2,
	0x4d, 0x18, 0xf2, 0x34, 0xa5, 0x23, 0x9c, 0x81, 0xfd, 0xe1, 0x34, 0xbe, 0xa0, 0xa4, 0x03, 0x16,
	0x27, 0xe9, 0x97, 0x24, 0xa4, 0xd6, 0xf1, 0x67, 0x80, 0x7e, 0x1a, 0x8c, 0x60, 0xf2, 0xe9, 0x5e,
	0xd6, 0x2d, 0xbe, 0xda, 0x7b, 0x1d, 0x47, 0xaf, 0xf7, 0x6f, 0xe3, 0x8d, 0xae, 0x1c, 0x73, 0x95,
	0x27, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0xc1, 0xb5, 0x25, 0xa9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DownloaderClient is the client API for Downloader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DownloaderClient interface {
	Query(ctx context.Context, in *DownloaderRequest, opts ...grpc.CallOption) (*DownloaderResponse, error)
}

type downloaderClient struct {
	cc *grpc.ClientConn
}

func NewDownloaderClient(cc *grpc.ClientConn) DownloaderClient {
	return &downloaderClient{cc}
}

func (c *downloaderClient) Query(ctx context.Context, in *DownloaderRequest, opts ...grpc.CallOption) (*DownloaderResponse, error) {
	out := new(DownloaderResponse)
	err := c.cc.Invoke(ctx, "/downloader.Downloader/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownloaderServer is the server API for Downloader service.
type DownloaderServer interface {
	Query(context.Context, *DownloaderRequest) (*DownloaderResponse, error)
}

func RegisterDownloaderServer(s *grpc.Server, srv DownloaderServer) {
	s.RegisterService(&_Downloader_serviceDesc, srv)
}

func _Downloader_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloaderServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/downloader.Downloader/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloaderServer).Query(ctx, req.(*DownloaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Downloader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "downloader.Downloader",
	HandlerType: (*DownloaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Downloader_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "downloader.proto",
}