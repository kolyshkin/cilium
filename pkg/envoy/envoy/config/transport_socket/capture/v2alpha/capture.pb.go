// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/transport_socket/capture/v2alpha/capture.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// File format.
type FileSink_Format int32

const (
	// Binary proto format as per :ref:`Trace
	// <envoy_api_msg_data.tap.v2alpha.Trace>`.
	FileSink_PROTO_BINARY FileSink_Format = 0
	// Text proto format as per :ref:`Trace
	// <envoy_api_msg_data.tap.v2alpha.Trace>`.
	FileSink_PROTO_TEXT FileSink_Format = 1
)

var FileSink_Format_name = map[int32]string{
	0: "PROTO_BINARY",
	1: "PROTO_TEXT",
}
var FileSink_Format_value = map[string]int32{
	"PROTO_BINARY": 0,
	"PROTO_TEXT":   1,
}

func (x FileSink_Format) String() string {
	return proto.EnumName(FileSink_Format_name, int32(x))
}
func (FileSink_Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_capture_95048b78d4910e85, []int{0, 0}
}

// File sink.
//
// .. warning::
//
//   The current file sink implementation buffers the entire trace in memory
//   prior to writing. This will OOM for long lived sockets and/or where there
//   is a large amount of traffic on the socket.
type FileSink struct {
	// Path prefix. The output file will be of the form <path_prefix>_<id>.pb, where <id> is an
	// identifier distinguishing the recorded trace for individual socket instances (the Envoy
	// connection ID).
	PathPrefix           string          `protobuf:"bytes,1,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	Format               FileSink_Format `protobuf:"varint,2,opt,name=format,proto3,enum=envoy.config.transport_socket.capture.v2alpha.FileSink_Format" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FileSink) Reset()         { *m = FileSink{} }
func (m *FileSink) String() string { return proto.CompactTextString(m) }
func (*FileSink) ProtoMessage()    {}
func (*FileSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_capture_95048b78d4910e85, []int{0}
}
func (m *FileSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSink.Unmarshal(m, b)
}
func (m *FileSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSink.Marshal(b, m, deterministic)
}
func (dst *FileSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSink.Merge(dst, src)
}
func (m *FileSink) XXX_Size() int {
	return xxx_messageInfo_FileSink.Size(m)
}
func (m *FileSink) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSink.DiscardUnknown(m)
}

var xxx_messageInfo_FileSink proto.InternalMessageInfo

func (m *FileSink) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *FileSink) GetFormat() FileSink_Format {
	if m != nil {
		return m.Format
	}
	return FileSink_PROTO_BINARY
}

// Configuration for capture transport socket. This wraps another transport socket, providing the
// ability to interpose and record in plain text any traffic that is surfaced to Envoy.
type Capture struct {
	// Types that are valid to be assigned to SinkSelector:
	//	*Capture_FileSink
	SinkSelector isCapture_SinkSelector `protobuf_oneof:"sink_selector"`
	// The underlying transport socket being wrapped.
	TransportSocket      *core.TransportSocket `protobuf:"bytes,2,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Capture) Reset()         { *m = Capture{} }
func (m *Capture) String() string { return proto.CompactTextString(m) }
func (*Capture) ProtoMessage()    {}
func (*Capture) Descriptor() ([]byte, []int) {
	return fileDescriptor_capture_95048b78d4910e85, []int{1}
}
func (m *Capture) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capture.Unmarshal(m, b)
}
func (m *Capture) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capture.Marshal(b, m, deterministic)
}
func (dst *Capture) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capture.Merge(dst, src)
}
func (m *Capture) XXX_Size() int {
	return xxx_messageInfo_Capture.Size(m)
}
func (m *Capture) XXX_DiscardUnknown() {
	xxx_messageInfo_Capture.DiscardUnknown(m)
}

var xxx_messageInfo_Capture proto.InternalMessageInfo

type isCapture_SinkSelector interface {
	isCapture_SinkSelector()
}

type Capture_FileSink struct {
	FileSink *FileSink `protobuf:"bytes,1,opt,name=file_sink,json=fileSink,proto3,oneof"`
}

func (*Capture_FileSink) isCapture_SinkSelector() {}

func (m *Capture) GetSinkSelector() isCapture_SinkSelector {
	if m != nil {
		return m.SinkSelector
	}
	return nil
}

func (m *Capture) GetFileSink() *FileSink {
	if x, ok := m.GetSinkSelector().(*Capture_FileSink); ok {
		return x.FileSink
	}
	return nil
}

func (m *Capture) GetTransportSocket() *core.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Capture) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Capture_OneofMarshaler, _Capture_OneofUnmarshaler, _Capture_OneofSizer, []interface{}{
		(*Capture_FileSink)(nil),
	}
}

func _Capture_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Capture)
	// sink_selector
	switch x := m.SinkSelector.(type) {
	case *Capture_FileSink:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FileSink); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Capture.SinkSelector has unexpected type %T", x)
	}
	return nil
}

func _Capture_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Capture)
	switch tag {
	case 1: // sink_selector.file_sink
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FileSink)
		err := b.DecodeMessage(msg)
		m.SinkSelector = &Capture_FileSink{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Capture_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Capture)
	// sink_selector
	switch x := m.SinkSelector.(type) {
	case *Capture_FileSink:
		s := proto.Size(x.FileSink)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FileSink)(nil), "envoy.config.transport_socket.capture.v2alpha.FileSink")
	proto.RegisterType((*Capture)(nil), "envoy.config.transport_socket.capture.v2alpha.Capture")
	proto.RegisterEnum("envoy.config.transport_socket.capture.v2alpha.FileSink_Format", FileSink_Format_name, FileSink_Format_value)
}

func init() {
	proto.RegisterFile("envoy/config/transport_socket/capture/v2alpha/capture.proto", fileDescriptor_capture_95048b78d4910e85)
}

var fileDescriptor_capture_95048b78d4910e85 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xd7, 0x21, 0x73, 0x7b, 0xd3, 0xad, 0xe4, 0x34, 0x44, 0x70, 0xf4, 0x34, 0x04, 0x13,
	0x88, 0x07, 0x0f, 0x82, 0xe0, 0xc4, 0xa1, 0x07, 0xdd, 0xc8, 0xca, 0x50, 0x2f, 0x25, 0x2b, 0xa9,
	0x0b, 0xab, 0x4d, 0x48, 0x63, 0xd1, 0xbf, 0xca, 0xbf, 0xc1, 0xff, 0x4c, 0x9a, 0xb6, 0x07, 0x77,
	0xd3, 0x63, 0x3e, 0xbe, 0xef, 0xf7, 0xbe, 0xbc, 0x07, 0x97, 0x22, 0x2b, 0xd4, 0x27, 0x89, 0x55,
	0x96, 0xc8, 0x57, 0x62, 0x0d, 0xcf, 0x72, 0xad, 0x8c, 0x8d, 0x72, 0x15, 0x6f, 0x85, 0x25, 0x31,
	0xd7, 0xf6, 0xdd, 0x08, 0x52, 0x50, 0x9e, 0xea, 0x0d, 0x6f, 0xde, 0x58, 0x1b, 0x65, 0x15, 0x3a,
	0x73, 0x61, 0x5c, 0x85, 0xf1, 0x6e, 0x18, 0x37, 0xe6, 0x3a, 0x7c, 0x74, 0x5c, 0xcd, 0xe2, 0x5a,
	0x92, 0x82, 0x92, 0x58, 0x19, 0x41, 0xd6, 0x3c, 0xaf, 0x61, 0xc1, 0x97, 0x07, 0xdd, 0x99, 0x4c,
	0xc5, 0x52, 0x66, 0x5b, 0x74, 0x02, 0x7d, 0xcd, 0xed, 0x26, 0xd2, 0x46, 0x24, 0xf2, 0x63, 0xe4,
	0x8d, 0xbd, 0x49, 0x8f, 0x41, 0x29, 0x2d, 0x9c, 0x82, 0x56, 0xd0, 0x49, 0x94, 0x79, 0xe3, 0x76,
	0xd4, 0x1e, 0x7b, 0x93, 0x01, 0xbd, 0xc2, 0x7f, 0xea, 0x82, 0x9b, 0x49, 0x78, 0xe6, 0x28, 0xac,
	0xa6, 0x05, 0xa7, 0xd0, 0xa9, 0x14, 0xe4, 0xc3, 0xc1, 0x82, 0xcd, 0xc3, 0x79, 0x34, 0xbd, 0x7f,
	0xbc, 0x66, 0xcf, 0x7e, 0x0b, 0x0d, 0x00, 0x2a, 0x25, 0xbc, 0x7d, 0x0a, 0x7d, 0x2f, 0xf8, 0xf6,
	0x60, 0xff, 0xa6, 0xe2, 0xa2, 0x15, 0xf4, 0x12, 0x99, 0x8a, 0x28, 0x97, 0xd9, 0xd6, 0xd5, 0xed,
	0xd3, 0x8b, 0x7f, 0x56, 0xba, 0x6b, 0xb1, 0x6e, 0xd2, 0x2c, 0xe2, 0x01, 0xfc, 0xdd, 0xa0, 0xfb,
	0x71, 0x9f, 0x06, 0x35, 0x9e, 0x6b, 0x89, 0x0b, 0x8a, 0xcb, 0x75, 0xe2, 0xb0, 0xb1, 0x2e, 0x9d,
	0x93, 0x0d, 0xed, 0x6f, 0x61, 0x3a, 0x84, 0xc3, 0xb2, 0x61, 0x94, 0x8b, 0x54, 0xc4, 0x56, 0x99,
	0xe9, 0xde, 0x4b, 0xbb, 0xa0, 0xeb, 0x8e, 0x3b, 0xc1, 0xf9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xce, 0x8a, 0x4e, 0xe0, 0x0e, 0x02, 0x00, 0x00,
}
