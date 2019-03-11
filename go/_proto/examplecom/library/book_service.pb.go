// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examplecom/library/book_service.proto

/*
Package examplecom_library is a generated protocol buffer package.

It is generated from these files:
	examplecom/library/book_service.proto

It has these top-level messages:
	Book
	GetBookRequest
	QueryBooksRequest
*/
package examplecom_library

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Book struct {
	Isbn   int64  `protobuf:"varint,1,opt,name=isbn" json:"isbn,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Book) GetIsbn() int64 {
	if m != nil {
		return m.Isbn
	}
	return 0
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type GetBookRequest struct {
	Isbn int64 `protobuf:"varint,1,opt,name=isbn" json:"isbn,omitempty"`
}

func (m *GetBookRequest) Reset()                    { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()               {}
func (*GetBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetBookRequest) GetIsbn() int64 {
	if m != nil {
		return m.Isbn
	}
	return 0
}

type QueryBooksRequest struct {
	AuthorPrefix string `protobuf:"bytes,1,opt,name=author_prefix,json=authorPrefix" json:"author_prefix,omitempty"`
}

func (m *QueryBooksRequest) Reset()                    { *m = QueryBooksRequest{} }
func (m *QueryBooksRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryBooksRequest) ProtoMessage()               {}
func (*QueryBooksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryBooksRequest) GetAuthorPrefix() string {
	if m != nil {
		return m.AuthorPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*Book)(nil), "examplecom.library.Book")
	proto.RegisterType((*GetBookRequest)(nil), "examplecom.library.GetBookRequest")
	proto.RegisterType((*QueryBooksRequest)(nil), "examplecom.library.QueryBooksRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BookService service

type BookServiceClient interface {
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpc.CallOption) (BookService_QueryBooksClient, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/examplecom.library.BookService/GetBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpc.CallOption) (BookService_QueryBooksClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BookService_serviceDesc.Streams[0], c.cc, "/examplecom.library.BookService/QueryBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceQueryBooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_QueryBooksClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type bookServiceQueryBooksClient struct {
	grpc.ClientStream
}

func (x *bookServiceQueryBooksClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BookService service

type BookServiceServer interface {
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	QueryBooks(*QueryBooksRequest, BookService_QueryBooksServer) error
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examplecom.library.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_QueryBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryBooksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).QueryBooks(m, &bookServiceQueryBooksServer{stream})
}

type BookService_QueryBooksServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type bookServiceQueryBooksServer struct {
	grpc.ServerStream
}

func (x *bookServiceQueryBooksServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "examplecom.library.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryBooks",
			Handler:       _BookService_QueryBooks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "examplecom/library/book_service.proto",
}

func init() { proto.RegisterFile("examplecom/library/book_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x4d, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0x4c, 0x2a, 0x4a, 0x2c, 0xaa, 0xd4, 0x4f,
	0xca, 0xcf, 0xcf, 0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x42, 0x28, 0xd3, 0x83, 0x2a, 0x53, 0xf2, 0xe0, 0x62, 0x71, 0xca, 0xcf, 0xcf,
	0x16, 0x12, 0xe2, 0x62, 0xc9, 0x2c, 0x4e, 0xca, 0x93, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x02,
	0xb3, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38,
	0x83, 0x20, 0x1c, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xd2, 0x92, 0x8c, 0xfc, 0x22, 0x09, 0x66, 0xb0,
	0x30, 0x94, 0xa7, 0xa4, 0xc2, 0xc5, 0xe7, 0x9e, 0x5a, 0x02, 0x32, 0x2c, 0x28, 0xb5, 0xb0, 0x34,
	0xb5, 0xb8, 0x04, 0x9b, 0x99, 0x4a, 0x16, 0x5c, 0x82, 0x81, 0xa5, 0xa9, 0x45, 0x95, 0x20, 0x75,
	0xc5, 0x30, 0x85, 0xca, 0x5c, 0xbc, 0x10, 0x43, 0xe2, 0x0b, 0x8a, 0x52, 0xd3, 0x32, 0x2b, 0xc0,
	0x3a, 0x38, 0x83, 0x78, 0x20, 0x82, 0x01, 0x60, 0x31, 0xa3, 0xd5, 0x8c, 0x5c, 0xdc, 0x20, 0x5d,
	0xc1, 0x10, 0x3f, 0x09, 0x79, 0x72, 0xb1, 0x43, 0xed, 0x13, 0x52, 0xd2, 0xc3, 0xf4, 0x99, 0x1e,
	0xaa, 0x63, 0xa4, 0x24, 0xb0, 0xa9, 0x01, 0x29, 0x50, 0x62, 0x10, 0x0a, 0xe4, 0xe2, 0x42, 0x38,
	0x4a, 0x48, 0x15, 0x9b, 0x4a, 0x0c, 0x47, 0xe3, 0x33, 0xd0, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0xe4,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x45, 0x57, 0xc6, 0x9b, 0x01, 0x00, 0x00,
}