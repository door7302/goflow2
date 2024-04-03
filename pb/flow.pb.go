// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: flow.proto

package flowpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FlowMessage_FlowType int32

const (
	FlowMessage_FLOWUNKNOWN FlowMessage_FlowType = 0
	FlowMessage_SFLOW_5     FlowMessage_FlowType = 1
	FlowMessage_NETFLOW_V5  FlowMessage_FlowType = 2
	FlowMessage_NETFLOW_V9  FlowMessage_FlowType = 3
	FlowMessage_IPFIX       FlowMessage_FlowType = 4
)

// Enum value maps for FlowMessage_FlowType.
var (
	FlowMessage_FlowType_name = map[int32]string{
		0: "FLOWUNKNOWN",
		1: "SFLOW_5",
		2: "NETFLOW_V5",
		3: "NETFLOW_V9",
		4: "IPFIX",
	}
	FlowMessage_FlowType_value = map[string]int32{
		"FLOWUNKNOWN": 0,
		"SFLOW_5":     1,
		"NETFLOW_V5":  2,
		"NETFLOW_V9":  3,
		"IPFIX":       4,
	}
)

func (x FlowMessage_FlowType) Enum() *FlowMessage_FlowType {
	p := new(FlowMessage_FlowType)
	*p = x
	return p
}

func (x FlowMessage_FlowType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlowMessage_FlowType) Descriptor() protoreflect.EnumDescriptor {
	return file_flow_proto_enumTypes[0].Descriptor()
}

func (FlowMessage_FlowType) Type() protoreflect.EnumType {
	return &file_flow_proto_enumTypes[0]
}

func (x FlowMessage_FlowType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlowMessage_FlowType.Descriptor instead.
func (FlowMessage_FlowType) EnumDescriptor() ([]byte, []int) {
	return file_flow_proto_rawDescGZIP(), []int{0, 0}
}

type FlowMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           FlowMessage_FlowType `protobuf:"varint,1,opt,name=type,proto3,enum=flowpb.FlowMessage_FlowType" json:"type,omitempty"`
	TimeReceivedNs uint64               `protobuf:"varint,110,opt,name=time_received_ns,json=timeReceivedNs,proto3" json:"time_received_ns,omitempty"`
	SequenceNum    uint32               `protobuf:"varint,4,opt,name=sequence_num,json=sequenceNum,proto3" json:"sequence_num,omitempty"`
	SamplingRate   uint64               `protobuf:"varint,3,opt,name=sampling_rate,json=samplingRate,proto3" json:"sampling_rate,omitempty"`
	// Sampler information
	SamplerAddress []byte `protobuf:"bytes,11,opt,name=sampler_address,json=samplerAddress,proto3" json:"sampler_address,omitempty"`
	// Found inside packet
	TimeFlowStartNs uint64 `protobuf:"varint,111,opt,name=time_flow_start_ns,json=timeFlowStartNs,proto3" json:"time_flow_start_ns,omitempty"`
	TimeFlowEndNs   uint64 `protobuf:"varint,112,opt,name=time_flow_end_ns,json=timeFlowEndNs,proto3" json:"time_flow_end_ns,omitempty"`
	// Size of the sampled packet
	Bytes   uint64 `protobuf:"varint,9,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Packets uint64 `protobuf:"varint,10,opt,name=packets,proto3" json:"packets,omitempty"`
	// Source/destination addresses
	SrcAddr []byte `protobuf:"bytes,6,opt,name=src_addr,json=srcAddr,proto3" json:"src_addr,omitempty"`
	DstAddr []byte `protobuf:"bytes,7,opt,name=dst_addr,json=dstAddr,proto3" json:"dst_addr,omitempty"`
	// Layer 3 protocol (IPv4/IPv6/ARP/MPLS...)
	Etype uint32 `protobuf:"varint,30,opt,name=etype,proto3" json:"etype,omitempty"`
	// Layer 4 protocol
	Proto uint32 `protobuf:"varint,20,opt,name=proto,proto3" json:"proto,omitempty"`
	// Ports for UDP and TCP
	SrcPort uint32 `protobuf:"varint,21,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	DstPort uint32 `protobuf:"varint,22,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
	// Interfaces
	InIf  uint32 `protobuf:"varint,18,opt,name=in_if,json=inIf,proto3" json:"in_if,omitempty"`
	OutIf uint32 `protobuf:"varint,19,opt,name=out_if,json=outIf,proto3" json:"out_if,omitempty"`
	// Ethernet information
	SrcMac uint64 `protobuf:"varint,27,opt,name=src_mac,json=srcMac,proto3" json:"src_mac,omitempty"`
	DstMac uint64 `protobuf:"varint,28,opt,name=dst_mac,json=dstMac,proto3" json:"dst_mac,omitempty"`
	// Vlan
	SrcVlan uint32 `protobuf:"varint,33,opt,name=src_vlan,json=srcVlan,proto3" json:"src_vlan,omitempty"`
	DstVlan uint32 `protobuf:"varint,34,opt,name=dst_vlan,json=dstVlan,proto3" json:"dst_vlan,omitempty"`
	// 802.1q VLAN in sampled packet
	VlanId uint32 `protobuf:"varint,29,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	// IP and TCP special flags
	IpTos            uint32 `protobuf:"varint,23,opt,name=ip_tos,json=ipTos,proto3" json:"ip_tos,omitempty"`
	ForwardingStatus uint32 `protobuf:"varint,24,opt,name=forwarding_status,json=forwardingStatus,proto3" json:"forwarding_status,omitempty"`
	IpTtl            uint32 `protobuf:"varint,25,opt,name=ip_ttl,json=ipTtl,proto3" json:"ip_ttl,omitempty"`
	IpFlags          uint32 `protobuf:"varint,38,opt,name=ip_flags,json=ipFlags,proto3" json:"ip_flags,omitempty"`
	TcpFlags         uint32 `protobuf:"varint,26,opt,name=tcp_flags,json=tcpFlags,proto3" json:"tcp_flags,omitempty"`
	IcmpType         uint32 `protobuf:"varint,31,opt,name=icmp_type,json=icmpType,proto3" json:"icmp_type,omitempty"`
	IcmpCode         uint32 `protobuf:"varint,32,opt,name=icmp_code,json=icmpCode,proto3" json:"icmp_code,omitempty"`
	Ipv6FlowLabel    uint32 `protobuf:"varint,37,opt,name=ipv6_flow_label,json=ipv6FlowLabel,proto3" json:"ipv6_flow_label,omitempty"`
	// Fragments (IPv4/IPv6)
	FragmentId     uint32 `protobuf:"varint,35,opt,name=fragment_id,json=fragmentId,proto3" json:"fragment_id,omitempty"`
	FragmentOffset uint32 `protobuf:"varint,36,opt,name=fragment_offset,json=fragmentOffset,proto3" json:"fragment_offset,omitempty"`
	// Autonomous system information
	SrcAs     uint32 `protobuf:"varint,14,opt,name=src_as,json=srcAs,proto3" json:"src_as,omitempty"`
	DstAs     uint32 `protobuf:"varint,15,opt,name=dst_as,json=dstAs,proto3" json:"dst_as,omitempty"`
	NextHop   []byte `protobuf:"bytes,12,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	NextHopAs uint32 `protobuf:"varint,13,opt,name=next_hop_as,json=nextHopAs,proto3" json:"next_hop_as,omitempty"`
	// Prefix size
	SrcNet uint32 `protobuf:"varint,16,opt,name=src_net,json=srcNet,proto3" json:"src_net,omitempty"`
	DstNet uint32 `protobuf:"varint,17,opt,name=dst_net,json=dstNet,proto3" json:"dst_net,omitempty"`
	// BGP information
	BgpNextHop     []byte   `protobuf:"bytes,100,opt,name=bgp_next_hop,json=bgpNextHop,proto3" json:"bgp_next_hop,omitempty"`
	BgpCommunities []uint32 `protobuf:"varint,101,rep,packed,name=bgp_communities,json=bgpCommunities,proto3" json:"bgp_communities,omitempty"`
	AsPath         []uint32 `protobuf:"varint,102,rep,packed,name=as_path,json=asPath,proto3" json:"as_path,omitempty"`
	// MPLS information
	MplsTtl             []uint32 `protobuf:"varint,80,rep,packed,name=mpls_ttl,json=mplsTtl,proto3" json:"mpls_ttl,omitempty"`
	MplsLabel           []uint32 `protobuf:"varint,81,rep,packed,name=mpls_label,json=mplsLabel,proto3" json:"mpls_label,omitempty"`
	MplsIp              [][]byte `protobuf:"bytes,82,rep,name=mpls_ip,json=mplsIp,proto3" json:"mpls_ip,omitempty"`
	ObservationDomainId uint32   `protobuf:"varint,70,opt,name=observation_domain_id,json=observationDomainId,proto3" json:"observation_domain_id,omitempty"`
	ObservationPointId  uint32   `protobuf:"varint,71,opt,name=observation_point_id,json=observationPointId,proto3" json:"observation_point_id,omitempty"`
	// SRH SRv6 Header - fields decoding
	SrhSegmentsIPv6Left     uint32   `protobuf:"varint,150,opt,name=srhSegmentsIPv6Left,proto3" json:"srhSegmentsIPv6Left,omitempty"`
	SrhLastEntryIPv6        uint32   `protobuf:"varint,151,opt,name=srhLastEntryIPv6,proto3" json:"srhLastEntryIPv6,omitempty"`
	SrhFlagsIPv6            uint32   `protobuf:"varint,152,opt,name=srhFlagsIPv6,proto3" json:"srhFlagsIPv6,omitempty"`
	SrhTagIPv6              uint32   `protobuf:"varint,153,opt,name=srhTagIPv6,proto3" json:"srhTagIPv6,omitempty"`
	SrhSegmentIPv6BasicList [][]byte `protobuf:"bytes,154,rep,name=srhSegmentIPv6BasicList,proto3" json:"srhSegmentIPv6BasicList,omitempty"`
	// Tunnel IP - inner frame fields decoding
	InnerFrameSrcAddr        []byte `protobuf:"bytes,160,opt,name=innerFrame_src_addr,json=innerFrameSrcAddr,proto3" json:"innerFrame_src_addr,omitempty"`
	InnerFrameDstAddr        []byte `protobuf:"bytes,161,opt,name=innerFrame_dst_addr,json=innerFrameDstAddr,proto3" json:"innerFrame_dst_addr,omitempty"`
	InnerFrameProto          uint32 `protobuf:"varint,162,opt,name=innerFrame_proto,json=innerFrameProto,proto3" json:"innerFrame_proto,omitempty"`
	InnerFrameSrcPort        uint32 `protobuf:"varint,163,opt,name=innerFrame_src_port,json=innerFrameSrcPort,proto3" json:"innerFrame_src_port,omitempty"`
	InnerFrameDstPort        uint32 `protobuf:"varint,164,opt,name=innerFrame_dst_port,json=innerFrameDstPort,proto3" json:"innerFrame_dst_port,omitempty"`
	InnerFrameIpv6FlowLabel  uint32 `protobuf:"varint,165,opt,name=innerFrame_ipv6_flow_label,json=innerFrameIpv6FlowLabel,proto3" json:"innerFrame_ipv6_flow_label,omitempty"`
	InnerFrameIpTos          uint32 `protobuf:"varint,166,opt,name=innerFrame_ip_tos,json=innerFrameIpTos,proto3" json:"innerFrame_ip_tos,omitempty"`
	InnerFrameIpTtl          uint32 `protobuf:"varint,167,opt,name=innerFrame_ip_ttl,json=innerFrameIpTtl,proto3" json:"innerFrame_ip_ttl,omitempty"`
	InnerFrameIpFlags        uint32 `protobuf:"varint,168,opt,name=innerFrame_ip_flags,json=innerFrameIpFlags,proto3" json:"innerFrame_ip_flags,omitempty"`
	InnerFrameTcpFlags       uint32 `protobuf:"varint,169,opt,name=innerFrame_tcp_flags,json=innerFrameTcpFlags,proto3" json:"innerFrame_tcp_flags,omitempty"`
	InnerFrameFragmentId     uint32 `protobuf:"varint,170,opt,name=innerFrame_fragment_id,json=innerFrameFragmentId,proto3" json:"innerFrame_fragment_id,omitempty"`
	InnerFrameFragmentOffset uint32 `protobuf:"varint,171,opt,name=innerFrame_fragment_offset,json=innerFrameFragmentOffset,proto3" json:"innerFrame_fragment_offset,omitempty"`
	InnerFrameIcmpType       uint32 `protobuf:"varint,172,opt,name=innerFrame_icmp_type,json=innerFrameIcmpType,proto3" json:"innerFrame_icmp_type,omitempty"`
	InnerFrameIcmpCode       uint32 `protobuf:"varint,173,opt,name=innerFrame_icmp_code,json=innerFrameIcmpCode,proto3" json:"innerFrame_icmp_code,omitempty"`
	InnerFramePayloadLen     uint32 `protobuf:"varint,174,opt,name=innerFrame_payload_len,json=innerFramePayloadLen,proto3" json:"innerFrame_payload_len,omitempty"`
}

func (x *FlowMessage) Reset() {
	*x = FlowMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowMessage) ProtoMessage() {}

func (x *FlowMessage) ProtoReflect() protoreflect.Message {
	mi := &file_flow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowMessage.ProtoReflect.Descriptor instead.
func (*FlowMessage) Descriptor() ([]byte, []int) {
	return file_flow_proto_rawDescGZIP(), []int{0}
}

func (x *FlowMessage) GetType() FlowMessage_FlowType {
	if x != nil {
		return x.Type
	}
	return FlowMessage_FLOWUNKNOWN
}

func (x *FlowMessage) GetTimeReceivedNs() uint64 {
	if x != nil {
		return x.TimeReceivedNs
	}
	return 0
}

func (x *FlowMessage) GetSequenceNum() uint32 {
	if x != nil {
		return x.SequenceNum
	}
	return 0
}

func (x *FlowMessage) GetSamplingRate() uint64 {
	if x != nil {
		return x.SamplingRate
	}
	return 0
}

func (x *FlowMessage) GetSamplerAddress() []byte {
	if x != nil {
		return x.SamplerAddress
	}
	return nil
}

func (x *FlowMessage) GetTimeFlowStartNs() uint64 {
	if x != nil {
		return x.TimeFlowStartNs
	}
	return 0
}

func (x *FlowMessage) GetTimeFlowEndNs() uint64 {
	if x != nil {
		return x.TimeFlowEndNs
	}
	return 0
}

func (x *FlowMessage) GetBytes() uint64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *FlowMessage) GetPackets() uint64 {
	if x != nil {
		return x.Packets
	}
	return 0
}

func (x *FlowMessage) GetSrcAddr() []byte {
	if x != nil {
		return x.SrcAddr
	}
	return nil
}

func (x *FlowMessage) GetDstAddr() []byte {
	if x != nil {
		return x.DstAddr
	}
	return nil
}

func (x *FlowMessage) GetEtype() uint32 {
	if x != nil {
		return x.Etype
	}
	return 0
}

func (x *FlowMessage) GetProto() uint32 {
	if x != nil {
		return x.Proto
	}
	return 0
}

func (x *FlowMessage) GetSrcPort() uint32 {
	if x != nil {
		return x.SrcPort
	}
	return 0
}

func (x *FlowMessage) GetDstPort() uint32 {
	if x != nil {
		return x.DstPort
	}
	return 0
}

func (x *FlowMessage) GetInIf() uint32 {
	if x != nil {
		return x.InIf
	}
	return 0
}

func (x *FlowMessage) GetOutIf() uint32 {
	if x != nil {
		return x.OutIf
	}
	return 0
}

func (x *FlowMessage) GetSrcMac() uint64 {
	if x != nil {
		return x.SrcMac
	}
	return 0
}

func (x *FlowMessage) GetDstMac() uint64 {
	if x != nil {
		return x.DstMac
	}
	return 0
}

func (x *FlowMessage) GetSrcVlan() uint32 {
	if x != nil {
		return x.SrcVlan
	}
	return 0
}

func (x *FlowMessage) GetDstVlan() uint32 {
	if x != nil {
		return x.DstVlan
	}
	return 0
}

func (x *FlowMessage) GetVlanId() uint32 {
	if x != nil {
		return x.VlanId
	}
	return 0
}

func (x *FlowMessage) GetIpTos() uint32 {
	if x != nil {
		return x.IpTos
	}
	return 0
}

func (x *FlowMessage) GetForwardingStatus() uint32 {
	if x != nil {
		return x.ForwardingStatus
	}
	return 0
}

func (x *FlowMessage) GetIpTtl() uint32 {
	if x != nil {
		return x.IpTtl
	}
	return 0
}

func (x *FlowMessage) GetIpFlags() uint32 {
	if x != nil {
		return x.IpFlags
	}
	return 0
}

func (x *FlowMessage) GetTcpFlags() uint32 {
	if x != nil {
		return x.TcpFlags
	}
	return 0
}

func (x *FlowMessage) GetIcmpType() uint32 {
	if x != nil {
		return x.IcmpType
	}
	return 0
}

func (x *FlowMessage) GetIcmpCode() uint32 {
	if x != nil {
		return x.IcmpCode
	}
	return 0
}

func (x *FlowMessage) GetIpv6FlowLabel() uint32 {
	if x != nil {
		return x.Ipv6FlowLabel
	}
	return 0
}

func (x *FlowMessage) GetFragmentId() uint32 {
	if x != nil {
		return x.FragmentId
	}
	return 0
}

func (x *FlowMessage) GetFragmentOffset() uint32 {
	if x != nil {
		return x.FragmentOffset
	}
	return 0
}

func (x *FlowMessage) GetSrcAs() uint32 {
	if x != nil {
		return x.SrcAs
	}
	return 0
}

func (x *FlowMessage) GetDstAs() uint32 {
	if x != nil {
		return x.DstAs
	}
	return 0
}

func (x *FlowMessage) GetNextHop() []byte {
	if x != nil {
		return x.NextHop
	}
	return nil
}

func (x *FlowMessage) GetNextHopAs() uint32 {
	if x != nil {
		return x.NextHopAs
	}
	return 0
}

func (x *FlowMessage) GetSrcNet() uint32 {
	if x != nil {
		return x.SrcNet
	}
	return 0
}

func (x *FlowMessage) GetDstNet() uint32 {
	if x != nil {
		return x.DstNet
	}
	return 0
}

func (x *FlowMessage) GetBgpNextHop() []byte {
	if x != nil {
		return x.BgpNextHop
	}
	return nil
}

func (x *FlowMessage) GetBgpCommunities() []uint32 {
	if x != nil {
		return x.BgpCommunities
	}
	return nil
}

func (x *FlowMessage) GetAsPath() []uint32 {
	if x != nil {
		return x.AsPath
	}
	return nil
}

func (x *FlowMessage) GetMplsTtl() []uint32 {
	if x != nil {
		return x.MplsTtl
	}
	return nil
}

func (x *FlowMessage) GetMplsLabel() []uint32 {
	if x != nil {
		return x.MplsLabel
	}
	return nil
}

func (x *FlowMessage) GetMplsIp() [][]byte {
	if x != nil {
		return x.MplsIp
	}
	return nil
}

func (x *FlowMessage) GetObservationDomainId() uint32 {
	if x != nil {
		return x.ObservationDomainId
	}
	return 0
}

func (x *FlowMessage) GetObservationPointId() uint32 {
	if x != nil {
		return x.ObservationPointId
	}
	return 0
}

func (x *FlowMessage) GetSrhSegmentsIPv6Left() uint32 {
	if x != nil {
		return x.SrhSegmentsIPv6Left
	}
	return 0
}

func (x *FlowMessage) GetSrhLastEntryIPv6() uint32 {
	if x != nil {
		return x.SrhLastEntryIPv6
	}
	return 0
}

func (x *FlowMessage) GetSrhFlagsIPv6() uint32 {
	if x != nil {
		return x.SrhFlagsIPv6
	}
	return 0
}

func (x *FlowMessage) GetSrhTagIPv6() uint32 {
	if x != nil {
		return x.SrhTagIPv6
	}
	return 0
}

func (x *FlowMessage) GetSrhSegmentIPv6BasicList() [][]byte {
	if x != nil {
		return x.SrhSegmentIPv6BasicList
	}
	return nil
}

func (x *FlowMessage) GetInnerFrameSrcAddr() []byte {
	if x != nil {
		return x.InnerFrameSrcAddr
	}
	return nil
}

func (x *FlowMessage) GetInnerFrameDstAddr() []byte {
	if x != nil {
		return x.InnerFrameDstAddr
	}
	return nil
}

func (x *FlowMessage) GetInnerFrameProto() uint32 {
	if x != nil {
		return x.InnerFrameProto
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameSrcPort() uint32 {
	if x != nil {
		return x.InnerFrameSrcPort
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameDstPort() uint32 {
	if x != nil {
		return x.InnerFrameDstPort
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameIpv6FlowLabel() uint32 {
	if x != nil {
		return x.InnerFrameIpv6FlowLabel
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameIpTos() uint32 {
	if x != nil {
		return x.InnerFrameIpTos
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameIpTtl() uint32 {
	if x != nil {
		return x.InnerFrameIpTtl
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameIpFlags() uint32 {
	if x != nil {
		return x.InnerFrameIpFlags
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameTcpFlags() uint32 {
	if x != nil {
		return x.InnerFrameTcpFlags
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameFragmentId() uint32 {
	if x != nil {
		return x.InnerFrameFragmentId
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameFragmentOffset() uint32 {
	if x != nil {
		return x.InnerFrameFragmentOffset
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameIcmpType() uint32 {
	if x != nil {
		return x.InnerFrameIcmpType
	}
	return 0
}

func (x *FlowMessage) GetInnerFrameIcmpCode() uint32 {
	if x != nil {
		return x.InnerFrameIcmpCode
	}
	return 0
}

func (x *FlowMessage) GetInnerFramePayloadLen() uint32 {
	if x != nil {
		return x.InnerFramePayloadLen
	}
	return 0
}

var File_flow_proto protoreflect.FileDescriptor

var file_flow_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x6c,
	0x6f, 0x77, 0x70, 0x62, 0x22, 0xd1, 0x13, 0x0a, 0x0b, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x70, 0x62, 0x2e, 0x46, 0x6c, 0x6f, 0x77,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x6e, 0x73, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x4e, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x4e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2b, 0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x6e, 0x73, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x74,
	0x69, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x73, 0x12, 0x27,
	0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x6e, 0x73, 0x18, 0x70, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x6c,
	0x6f, 0x77, 0x45, 0x6e, 0x64, 0x4e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x72, 0x63, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x72, 0x63,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x13, 0x0a, 0x05, 0x69, 0x6e, 0x5f, 0x69, 0x66, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x69, 0x6e, 0x49, 0x66, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x66, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6f, 0x75, 0x74, 0x49, 0x66, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x72, 0x63, 0x5f, 0x6d, 0x61, 0x63, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x72,
	0x63, 0x4d, 0x61, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x63, 0x18,
	0x1c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x64, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x72, 0x63, 0x5f, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x73, 0x72, 0x63, 0x56, 0x6c, 0x61, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f,
	0x76, 0x6c, 0x61, 0x6e, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x73, 0x74, 0x56,
	0x6c, 0x61, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x1d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x76, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x69, 0x70, 0x5f, 0x74, 0x6f, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x70,
	0x54, 0x6f, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x15, 0x0a, 0x06, 0x69, 0x70, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x69, 0x70, 0x54, 0x74, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x70, 0x5f, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0x26, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x70, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x63, 0x70, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x63, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x63, 0x6d, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x69, 0x63, 0x6d, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x63, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x69, 0x63, 0x6d, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x70, 0x76,
	0x36, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x25, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x69, 0x70, 0x76, 0x36, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x23, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x24, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x72, 0x61,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73,
	0x72, 0x63, 0x5f, 0x61, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x72, 0x63,
	0x41, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x73, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x64, 0x73, 0x74, 0x41, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x68, 0x6f, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x65, 0x78,
	0x74, 0x48, 0x6f, 0x70, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70,
	0x5f, 0x61, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x48,
	0x6f, 0x70, 0x41, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x72, 0x63, 0x5f, 0x6e, 0x65, 0x74, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x72, 0x63, 0x4e, 0x65, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x64, 0x73, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x64, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x67, 0x70, 0x5f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x67,
	0x70, 0x4e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x67, 0x70, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0e, 0x62, 0x67, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x66, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x06, 0x61, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x70,
	0x6c, 0x73, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x50, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x70,
	0x6c, 0x73, 0x54, 0x74, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x70, 0x6c, 0x73, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x51, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x70, 0x6c, 0x73, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x70, 0x6c, 0x73, 0x5f, 0x69, 0x70, 0x18,
	0x52, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x6d, 0x70, 0x6c, 0x73, 0x49, 0x70, 0x12, 0x32, 0x0a,
	0x15, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x14, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x47, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x13, 0x73, 0x72, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x49, 0x50, 0x76, 0x36, 0x4c, 0x65, 0x66, 0x74, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x13, 0x73, 0x72, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x50,
	0x76, 0x36, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x2b, 0x0a, 0x10, 0x73, 0x72, 0x68, 0x4c, 0x61, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x50, 0x76, 0x36, 0x18, 0x97, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x73, 0x72, 0x68, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x50, 0x76, 0x36, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x72, 0x68, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x49,
	0x50, 0x76, 0x36, 0x18, 0x98, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x72, 0x68, 0x46,
	0x6c, 0x61, 0x67, 0x73, 0x49, 0x50, 0x76, 0x36, 0x12, 0x1f, 0x0a, 0x0a, 0x73, 0x72, 0x68, 0x54,
	0x61, 0x67, 0x49, 0x50, 0x76, 0x36, 0x18, 0x99, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73,
	0x72, 0x68, 0x54, 0x61, 0x67, 0x49, 0x50, 0x76, 0x36, 0x12, 0x39, 0x0a, 0x17, 0x73, 0x72, 0x68,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x76, 0x36, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x9a, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x17, 0x73, 0x72, 0x68,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x76, 0x36, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x5f, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0xa0, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x11, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x53, 0x72,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x5f, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0xa1, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x11, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x44,
	0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0xa2, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x5f, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0xa3, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x11, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x53, 0x72, 0x63, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x5f, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0xa4, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x11, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x44, 0x73, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x3c, 0x0a, 0x1a, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0xa5, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x70, 0x76, 0x36, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x70, 0x5f, 0x74, 0x6f, 0x73, 0x18, 0xa6, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x70, 0x54, 0x6f, 0x73, 0x12,
	0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x70,
	0x5f, 0x74, 0x74, 0x6c, 0x18, 0xa7, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x70, 0x54, 0x74, 0x6c, 0x12, 0x2f, 0x0a, 0x13,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x70, 0x5f, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0xa8, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x31, 0x0a,
	0x14, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x63, 0x70, 0x5f,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0xa9, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x63, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x73,
	0x12, 0x35, 0x0a, 0x16, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x66,
	0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0xaa, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x14, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x46, 0x72, 0x61,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x1a, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0xab, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x14, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x63, 0x6d, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xac,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x49, 0x63, 0x6d, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x63, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0xad, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x49, 0x63, 0x6d, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x16,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0xae, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x4c, 0x65, 0x6e, 0x22, 0x53, 0x0a, 0x08, 0x46, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x46, 0x4c, 0x4f, 0x57, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x35, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x45, 0x54, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x56, 0x35, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x45, 0x54, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x56, 0x39, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x49, 0x50, 0x46, 0x49, 0x58, 0x10, 0x04, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x74, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x72, 0x2f, 0x67, 0x6f, 0x66, 0x6c, 0x6f, 0x77, 0x32, 0x2f, 0x70, 0x62, 0x3b, 0x66, 0x6c, 0x6f,
	0x77, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flow_proto_rawDescOnce sync.Once
	file_flow_proto_rawDescData = file_flow_proto_rawDesc
)

func file_flow_proto_rawDescGZIP() []byte {
	file_flow_proto_rawDescOnce.Do(func() {
		file_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_flow_proto_rawDescData)
	})
	return file_flow_proto_rawDescData
}

var file_flow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flow_proto_goTypes = []interface{}{
	(FlowMessage_FlowType)(0), // 0: flowpb.FlowMessage.FlowType
	(*FlowMessage)(nil),       // 1: flowpb.FlowMessage
}
var file_flow_proto_depIdxs = []int32{
	0, // 0: flowpb.FlowMessage.type:type_name -> flowpb.FlowMessage.FlowType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flow_proto_init() }
func file_flow_proto_init() {
	if File_flow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flow_proto_goTypes,
		DependencyIndexes: file_flow_proto_depIdxs,
		EnumInfos:         file_flow_proto_enumTypes,
		MessageInfos:      file_flow_proto_msgTypes,
	}.Build()
	File_flow_proto = out.File
	file_flow_proto_rawDesc = nil
	file_flow_proto_goTypes = nil
	file_flow_proto_depIdxs = nil
}
