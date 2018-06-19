// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/endpoint/endpoint.proto

package endpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _ "github.com/gogo/protobuf/gogoproto"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Upstream host identifier.
type Endpoint struct {
	// The upstream host address.
	//
	// .. attention::
	//
	//   The form of host address depends on the given cluster type. For STATIC,
	//   it is expected to be a direct IP address (or something resolvable by the
	//   specified :ref:`resolver <envoy_api_field_core.SocketAddress.resolver_name>`
	//   in the Address). For LOGICAL or STRICT DNS, it is expected to be hostname,
	//   and will be resolved via DNS.
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// [#not-implemented-hide:] The optional health check configuration is used as
	// configuration for the health checker to contact the health checked host.
	//
	// .. attention::
	//
	//   This takes into effect only for upstream clusters with
	//   :ref:`active health checking <arch_overview_health_checking>` enabled.
	HealthCheckConfig    *Endpoint_HealthCheckConfig `protobuf:"bytes,2,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoint_ac874e8927bdf21f, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (dst *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(dst, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetHealthCheckConfig() *Endpoint_HealthCheckConfig {
	if m != nil {
		return m.HealthCheckConfig
	}
	return nil
}

// [#not-implemented-hide:] The optional health check configuration.
type Endpoint_HealthCheckConfig struct {
	// Optional alternative health check port value.
	//
	// By default the health check address port of an upstream host is the same
	// as the host's serving address port. This provides an alternative health
	// check port. Setting this with a non-zero value allows an upstream host
	// to have different health check address port.
	PortValue            uint32   `protobuf:"varint,1,opt,name=port_value,json=portValue,proto3" json:"port_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint_HealthCheckConfig) Reset()         { *m = Endpoint_HealthCheckConfig{} }
func (m *Endpoint_HealthCheckConfig) String() string { return proto.CompactTextString(m) }
func (*Endpoint_HealthCheckConfig) ProtoMessage()    {}
func (*Endpoint_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoint_ac874e8927bdf21f, []int{0, 0}
}
func (m *Endpoint_HealthCheckConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Unmarshal(m, b)
}
func (m *Endpoint_HealthCheckConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Marshal(b, m, deterministic)
}
func (dst *Endpoint_HealthCheckConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint_HealthCheckConfig.Merge(dst, src)
}
func (m *Endpoint_HealthCheckConfig) XXX_Size() int {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Size(m)
}
func (m *Endpoint_HealthCheckConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint_HealthCheckConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint_HealthCheckConfig proto.InternalMessageInfo

func (m *Endpoint_HealthCheckConfig) GetPortValue() uint32 {
	if m != nil {
		return m.PortValue
	}
	return 0
}

// An Endpoint that Envoy can route traffic to.
type LbEndpoint struct {
	// Upstream host identifier
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Optional health status when known and supplied by EDS server.
	HealthStatus core.HealthStatus `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"health_status,omitempty"`
	// The endpoint metadata specifies values that may be used by the load
	// balancer to select endpoints in a cluster for a given request. The filter
	// name should be specified as *envoy.lb*. An example boolean key-value pair
	// is *canary*, providing the optional canary status of the upstream host.
	// This may be matched against in a route's ForwardAction metadata_match field
	// to subset the endpoints considered in cluster load balancing.
	Metadata *core.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The optional load balancing weight of the upstream host, in the range 1 -
	// 128. Envoy uses the load balancing weight in some of the built in load
	// balancers. The load balancing weight for an endpoint is divided by the sum
	// of the weights of all endpoints in the endpoint's locality to produce a
	// percentage of traffic for the endpoint. This percentage is then further
	// weighted by the endpoint's locality's load balancing weight from
	// LocalityLbEndpoints. If unspecified, each host is presumed to have equal
	// weight in a locality.
	//
	// .. attention::
	//
	//   The limit of 128 is somewhat arbitrary, but is applied due to performance
	//   concerns with the current implementation and can be removed when
	//   `this issue <https://github.com/envoyproxy/envoy/issues/1285>`_ is fixed.
	LoadBalancingWeight  *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LbEndpoint) Reset()         { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()    {}
func (*LbEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoint_ac874e8927bdf21f, []int{1}
}
func (m *LbEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LbEndpoint.Unmarshal(m, b)
}
func (m *LbEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LbEndpoint.Marshal(b, m, deterministic)
}
func (dst *LbEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LbEndpoint.Merge(dst, src)
}
func (m *LbEndpoint) XXX_Size() int {
	return xxx_messageInfo_LbEndpoint.Size(m)
}
func (m *LbEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_LbEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_LbEndpoint proto.InternalMessageInfo

func (m *LbEndpoint) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *LbEndpoint) GetHealthStatus() core.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

func (m *LbEndpoint) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// A group of endpoints belonging to a Locality.
// One can have multiple LocalityLbEndpoints for a locality, but this is
// generally only done if the different groups need to have different load
// balancing weights or different priorities.
type LocalityLbEndpoints struct {
	// Identifies location of where the upstream hosts run.
	Locality *core.Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// The group of endpoints belonging to the locality specified.
	LbEndpoints []*LbEndpoint `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints,proto3" json:"lb_endpoints,omitempty"`
	// Optional: Per priority/region/zone/sub_zone weight - range 1-128. The load
	// balancing weight for a locality is divided by the sum of the weights of all
	// localities  at the same priority level to produce the effective percentage
	// of traffic for the locality.
	//
	// Locality weights are only considered when :ref:`locality weighted load
	// balancing <arch_overview_load_balancing_locality_weighted_lb>` is
	// configured. These weights are ignored otherwise. If no weights are
	// specificed when locality weighted load balancing is enabled, the cluster is
	// assumed to have a weight of 1.
	//
	// .. attention::
	//
	//   The limit of 128 is somewhat arbitrary, but is applied due to performance
	//   concerns with the current implementation and can be removed when
	//   `this issue <https://github.com/envoyproxy/envoy/issues/1285>`_ is fixed.
	LoadBalancingWeight *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	// Optional: the priority for this LocalityLbEndpoints. If unspecified this will
	// default to the highest priority (0).
	//
	// Under usual circumstances, Envoy will only select endpoints for the highest
	// priority (0). In the event all endpoints for a particular priority are
	// unavailable/unhealthy, Envoy will fail over to selecting endpoints for the
	// next highest priority group.
	//
	// Priorities should range from 0 (highest) to N (lowest) without skipping.
	Priority             uint32   `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalityLbEndpoints) Reset()         { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()    {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoint_ac874e8927bdf21f, []int{2}
}
func (m *LocalityLbEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityLbEndpoints.Unmarshal(m, b)
}
func (m *LocalityLbEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityLbEndpoints.Marshal(b, m, deterministic)
}
func (dst *LocalityLbEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityLbEndpoints.Merge(dst, src)
}
func (m *LocalityLbEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityLbEndpoints.Size(m)
}
func (m *LocalityLbEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityLbEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityLbEndpoints proto.InternalMessageInfo

func (m *LocalityLbEndpoints) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLbEndpoints() []*LbEndpoint {
	if m != nil {
		return m.LbEndpoints
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

func (m *LocalityLbEndpoints) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "envoy.api.v2.endpoint.Endpoint")
	proto.RegisterType((*Endpoint_HealthCheckConfig)(nil), "envoy.api.v2.endpoint.Endpoint.HealthCheckConfig")
	proto.RegisterType((*LbEndpoint)(nil), "envoy.api.v2.endpoint.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "envoy.api.v2.endpoint.LocalityLbEndpoints")
}

func init() {
	proto.RegisterFile("envoy/api/v2/endpoint/endpoint.proto", fileDescriptor_endpoint_ac874e8927bdf21f)
}

var fileDescriptor_endpoint_ac874e8927bdf21f = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0x4b, 0xc3, 0x26, 0xad, 0xd4, 0x0d, 0x15, 0x91, 0x29, 0x0d, 0x44, 0x3d, 0x44,
	0x1c, 0xd6, 0xc2, 0x45, 0xe2, 0xc0, 0x89, 0x14, 0x24, 0x40, 0xe5, 0xb2, 0x08, 0x90, 0x38, 0x60,
	0x8d, 0xed, 0xad, 0xbd, 0x62, 0xf1, 0x5a, 0xf6, 0xc6, 0x55, 0x6f, 0x7c, 0x05, 0xdf, 0xd0, 0x6f,
	0xe0, 0xc4, 0x91, 0x3f, 0xe0, 0xc6, 0x81, 0x1b, 0x7f, 0x81, 0xbc, 0xde, 0x75, 0xd3, 0x26, 0x15,
	0x97, 0xde, 0x76, 0x67, 0xde, 0xbc, 0x79, 0xef, 0x69, 0xd0, 0x3e, 0xcb, 0x2a, 0x79, 0xea, 0x41,
	0xce, 0xbd, 0xca, 0xf7, 0x58, 0x16, 0xe7, 0x92, 0x67, 0xaa, 0x7d, 0x90, 0xbc, 0x90, 0x4a, 0xe2,
	0x1d, 0x8d, 0x22, 0x90, 0x73, 0x52, 0xf9, 0xc4, 0x36, 0xdd, 0xf1, 0x85, 0xe1, 0x48, 0x16, 0xcc,
	0x83, 0x38, 0x2e, 0x58, 0x59, 0x36, 0x73, 0xee, 0xee, 0x32, 0x20, 0x84, 0x92, 0x99, 0xee, 0xfe,
	0x72, 0x37, 0x65, 0x20, 0x54, 0x1a, 0x44, 0x29, 0x8b, 0x3e, 0x1b, 0xd4, 0x5e, 0x22, 0x65, 0x22,
	0x98, 0xa7, 0x7f, 0xe1, 0xfc, 0xd8, 0x3b, 0x29, 0x20, 0xcf, 0x59, 0x61, 0x77, 0xdc, 0xa9, 0x40,
	0xf0, 0x18, 0x14, 0xf3, 0xec, 0xc3, 0x34, 0x6e, 0x27, 0x32, 0x91, 0xfa, 0xe9, 0xd5, 0xaf, 0xa6,
	0x3a, 0xf9, 0xe5, 0xa0, 0xde, 0x0b, 0x63, 0x00, 0x3f, 0x46, 0x1b, 0x46, 0xf0, 0xc8, 0xb9, 0xef,
	0x4c, 0xfb, 0xbe, 0x4b, 0x2e, 0x38, 0xad, 0x35, 0x91, 0x67, 0x0d, 0x82, 0x5a, 0x28, 0x06, 0x34,
	0x5c, 0xd4, 0x19, 0x44, 0x32, 0x3b, 0xe6, 0xc9, 0xa8, 0xa3, 0x19, 0x1e, 0x91, 0x95, 0x59, 0x11,
	0xbb, 0x93, 0xbc, 0xd4, 0xa3, 0x87, 0xf5, 0xe4, 0xa1, 0x1e, 0xa4, 0xdb, 0xe9, 0xe5, 0x92, 0xeb,
	0xa3, 0xed, 0x25, 0x1c, 0xbe, 0x87, 0x50, 0x2e, 0x0b, 0x15, 0x54, 0x20, 0xe6, 0x4c, 0x0b, 0xde,
	0xa4, 0xb7, 0xea, 0xca, 0xfb, 0xba, 0x30, 0x39, 0xeb, 0x20, 0x74, 0x14, 0xb6, 0xde, 0x9e, 0xa2,
	0x9e, 0x5d, 0x6e, 0xcc, 0x8d, 0xff, 0x23, 0x8d, 0xb6, 0x03, 0xf8, 0x39, 0xda, 0x34, 0x16, 0x4b,
	0x05, 0x6a, 0x5e, 0x6a, 0x73, 0x5b, 0x97, 0x19, 0x74, 0x3c, 0x8d, 0xce, 0xb7, 0x1a, 0x46, 0x07,
	0xe9, 0xc2, 0x0f, 0x3f, 0x41, 0xbd, 0x2f, 0x4c, 0x41, 0x0c, 0x0a, 0x46, 0x5d, 0x2d, 0xe1, 0xee,
	0x0a, 0x82, 0x37, 0x06, 0x42, 0x5b, 0x30, 0xfe, 0x84, 0x76, 0x84, 0x84, 0x38, 0x08, 0x41, 0x40,
	0x16, 0xf1, 0x2c, 0x09, 0x4e, 0x18, 0x4f, 0x52, 0x35, 0x5a, 0xd3, 0x2c, 0xbb, 0xa4, 0xb9, 0x09,
	0x62, 0x6f, 0x82, 0xbc, 0x7b, 0x95, 0xa9, 0x03, 0x5f, 0xe7, 0x30, 0x1b, 0x7c, 0xff, 0xfb, 0xa3,
	0xbb, 0xf1, 0x70, 0x7d, 0xf4, 0xd5, 0x99, 0x3a, 0x74, 0x58, 0x13, 0xcd, 0x2c, 0xcf, 0x07, 0x4d,
	0x33, 0xf9, 0xd6, 0x41, 0xc3, 0x23, 0x19, 0x81, 0xe0, 0xea, 0xf4, 0x3c, 0x32, 0x2d, 0x58, 0x98,
	0xb2, 0xc9, 0x6c, 0x95, 0x60, 0x3b, 0x49, 0x5b, 0x30, 0x7e, 0x8d, 0x06, 0x22, 0x0c, 0x6c, 0x7c,
	0x75, 0x5c, 0xdd, 0x69, 0xdf, 0x7f, 0x70, 0x45, 0xe0, 0xe7, 0x2b, 0x67, 0x6b, 0x3f, 0x7f, 0x8f,
	0x6f, 0xd0, 0xbe, 0x58, 0x10, 0x71, 0xa5, 0xf9, 0xee, 0xb5, 0x98, 0xc7, 0x2e, 0xea, 0xe5, 0x05,
	0x97, 0x45, 0x6d, 0x72, 0x5d, 0x1f, 0x51, 0xfb, 0x9f, 0x6d, 0x9d, 0xfd, 0xd9, 0x73, 0x3e, 0xb6,
	0x77, 0x10, 0xde, 0xd4, 0x4b, 0x0e, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x58, 0x83, 0xd6,
	0x27, 0x04, 0x00, 0x00,
}
