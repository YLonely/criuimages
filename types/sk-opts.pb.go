// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: sk-opts.proto

package types

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SkShutdown int32

const (
	SkShutdown_NONE  SkShutdown = 0
	SkShutdown_READ  SkShutdown = 1
	SkShutdown_WRITE SkShutdown = 2
	SkShutdown_BOTH  SkShutdown = 3
)

// Enum value maps for SkShutdown.
var (
	SkShutdown_name = map[int32]string{
		0: "NONE",
		1: "READ",
		2: "WRITE",
		3: "BOTH",
	}
	SkShutdown_value = map[string]int32{
		"NONE":  0,
		"READ":  1,
		"WRITE": 2,
		"BOTH":  3,
	}
)

func (x SkShutdown) Enum() *SkShutdown {
	p := new(SkShutdown)
	*p = x
	return p
}

func (x SkShutdown) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SkShutdown) Descriptor() protoreflect.EnumDescriptor {
	return file_sk_opts_proto_enumTypes[0].Descriptor()
}

func (SkShutdown) Type() protoreflect.EnumType {
	return &file_sk_opts_proto_enumTypes[0]
}

func (x SkShutdown) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SkShutdown) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SkShutdown(num)
	return nil
}

// Deprecated: Use SkShutdown.Descriptor instead.
func (SkShutdown) EnumDescriptor() ([]byte, []int) {
	return file_sk_opts_proto_rawDescGZIP(), []int{0}
}

type SkOptsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SoSndbuf     *uint32  `protobuf:"varint,1,req,name=so_sndbuf,json=soSndbuf" json:"so_sndbuf,omitempty"`
	SoRcvbuf     *uint32  `protobuf:"varint,2,req,name=so_rcvbuf,json=soRcvbuf" json:"so_rcvbuf,omitempty"`
	SoSndTmoSec  *uint64  `protobuf:"varint,3,req,name=so_snd_tmo_sec,json=soSndTmoSec" json:"so_snd_tmo_sec,omitempty"`
	SoSndTmoUsec *uint64  `protobuf:"varint,4,req,name=so_snd_tmo_usec,json=soSndTmoUsec" json:"so_snd_tmo_usec,omitempty"`
	SoRcvTmoSec  *uint64  `protobuf:"varint,5,req,name=so_rcv_tmo_sec,json=soRcvTmoSec" json:"so_rcv_tmo_sec,omitempty"`
	SoRcvTmoUsec *uint64  `protobuf:"varint,6,req,name=so_rcv_tmo_usec,json=soRcvTmoUsec" json:"so_rcv_tmo_usec,omitempty"`
	Reuseaddr    *bool    `protobuf:"varint,7,opt,name=reuseaddr" json:"reuseaddr,omitempty"`
	SoPriority   *uint32  `protobuf:"varint,8,opt,name=so_priority,json=soPriority" json:"so_priority,omitempty"`
	SoRcvlowat   *uint32  `protobuf:"varint,9,opt,name=so_rcvlowat,json=soRcvlowat" json:"so_rcvlowat,omitempty"`
	SoMark       *uint32  `protobuf:"varint,10,opt,name=so_mark,json=soMark" json:"so_mark,omitempty"`
	SoPasscred   *bool    `protobuf:"varint,11,opt,name=so_passcred,json=soPasscred" json:"so_passcred,omitempty"`
	SoPasssec    *bool    `protobuf:"varint,12,opt,name=so_passsec,json=soPasssec" json:"so_passsec,omitempty"`
	SoDontroute  *bool    `protobuf:"varint,13,opt,name=so_dontroute,json=soDontroute" json:"so_dontroute,omitempty"`
	SoNoCheck    *bool    `protobuf:"varint,14,opt,name=so_no_check,json=soNoCheck" json:"so_no_check,omitempty"`
	SoBoundDev   *string  `protobuf:"bytes,15,opt,name=so_bound_dev,json=soBoundDev" json:"so_bound_dev,omitempty"`
	SoFilter     []uint64 `protobuf:"fixed64,16,rep,name=so_filter,json=soFilter" json:"so_filter,omitempty"`
	SoReuseport  *bool    `protobuf:"varint,17,opt,name=so_reuseport,json=soReuseport" json:"so_reuseport,omitempty"`
	SoBroadcast  *bool    `protobuf:"varint,18,opt,name=so_broadcast,json=soBroadcast" json:"so_broadcast,omitempty"`
	SoKeepalive  *bool    `protobuf:"varint,19,opt,name=so_keepalive,json=soKeepalive" json:"so_keepalive,omitempty"`
	TcpKeepcnt   *uint32  `protobuf:"varint,20,opt,name=tcp_keepcnt,json=tcpKeepcnt" json:"tcp_keepcnt,omitempty"`
	TcpKeepidle  *uint32  `protobuf:"varint,21,opt,name=tcp_keepidle,json=tcpKeepidle" json:"tcp_keepidle,omitempty"`
	TcpKeepintvl *uint32  `protobuf:"varint,22,opt,name=tcp_keepintvl,json=tcpKeepintvl" json:"tcp_keepintvl,omitempty"`
}

func (x *SkOptsEntry) Reset() {
	*x = SkOptsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sk_opts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkOptsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkOptsEntry) ProtoMessage() {}

func (x *SkOptsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sk_opts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkOptsEntry.ProtoReflect.Descriptor instead.
func (*SkOptsEntry) Descriptor() ([]byte, []int) {
	return file_sk_opts_proto_rawDescGZIP(), []int{0}
}

func (x *SkOptsEntry) GetSoSndbuf() uint32 {
	if x != nil && x.SoSndbuf != nil {
		return *x.SoSndbuf
	}
	return 0
}

func (x *SkOptsEntry) GetSoRcvbuf() uint32 {
	if x != nil && x.SoRcvbuf != nil {
		return *x.SoRcvbuf
	}
	return 0
}

func (x *SkOptsEntry) GetSoSndTmoSec() uint64 {
	if x != nil && x.SoSndTmoSec != nil {
		return *x.SoSndTmoSec
	}
	return 0
}

func (x *SkOptsEntry) GetSoSndTmoUsec() uint64 {
	if x != nil && x.SoSndTmoUsec != nil {
		return *x.SoSndTmoUsec
	}
	return 0
}

func (x *SkOptsEntry) GetSoRcvTmoSec() uint64 {
	if x != nil && x.SoRcvTmoSec != nil {
		return *x.SoRcvTmoSec
	}
	return 0
}

func (x *SkOptsEntry) GetSoRcvTmoUsec() uint64 {
	if x != nil && x.SoRcvTmoUsec != nil {
		return *x.SoRcvTmoUsec
	}
	return 0
}

func (x *SkOptsEntry) GetReuseaddr() bool {
	if x != nil && x.Reuseaddr != nil {
		return *x.Reuseaddr
	}
	return false
}

func (x *SkOptsEntry) GetSoPriority() uint32 {
	if x != nil && x.SoPriority != nil {
		return *x.SoPriority
	}
	return 0
}

func (x *SkOptsEntry) GetSoRcvlowat() uint32 {
	if x != nil && x.SoRcvlowat != nil {
		return *x.SoRcvlowat
	}
	return 0
}

func (x *SkOptsEntry) GetSoMark() uint32 {
	if x != nil && x.SoMark != nil {
		return *x.SoMark
	}
	return 0
}

func (x *SkOptsEntry) GetSoPasscred() bool {
	if x != nil && x.SoPasscred != nil {
		return *x.SoPasscred
	}
	return false
}

func (x *SkOptsEntry) GetSoPasssec() bool {
	if x != nil && x.SoPasssec != nil {
		return *x.SoPasssec
	}
	return false
}

func (x *SkOptsEntry) GetSoDontroute() bool {
	if x != nil && x.SoDontroute != nil {
		return *x.SoDontroute
	}
	return false
}

func (x *SkOptsEntry) GetSoNoCheck() bool {
	if x != nil && x.SoNoCheck != nil {
		return *x.SoNoCheck
	}
	return false
}

func (x *SkOptsEntry) GetSoBoundDev() string {
	if x != nil && x.SoBoundDev != nil {
		return *x.SoBoundDev
	}
	return ""
}

func (x *SkOptsEntry) GetSoFilter() []uint64 {
	if x != nil {
		return x.SoFilter
	}
	return nil
}

func (x *SkOptsEntry) GetSoReuseport() bool {
	if x != nil && x.SoReuseport != nil {
		return *x.SoReuseport
	}
	return false
}

func (x *SkOptsEntry) GetSoBroadcast() bool {
	if x != nil && x.SoBroadcast != nil {
		return *x.SoBroadcast
	}
	return false
}

func (x *SkOptsEntry) GetSoKeepalive() bool {
	if x != nil && x.SoKeepalive != nil {
		return *x.SoKeepalive
	}
	return false
}

func (x *SkOptsEntry) GetTcpKeepcnt() uint32 {
	if x != nil && x.TcpKeepcnt != nil {
		return *x.TcpKeepcnt
	}
	return 0
}

func (x *SkOptsEntry) GetTcpKeepidle() uint32 {
	if x != nil && x.TcpKeepidle != nil {
		return *x.TcpKeepidle
	}
	return 0
}

func (x *SkOptsEntry) GetTcpKeepintvl() uint32 {
	if x != nil && x.TcpKeepintvl != nil {
		return *x.TcpKeepintvl
	}
	return 0
}

var File_sk_opts_proto protoreflect.FileDescriptor

var file_sk_opts_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x6b, 0x2d, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xee, 0x05, 0x0a, 0x0d, 0x73, 0x6b, 0x5f, 0x6f, 0x70,
	0x74, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x5f, 0x73,
	0x6e, 0x64, 0x62, 0x75, 0x66, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x6f, 0x53,
	0x6e, 0x64, 0x62, 0x75, 0x66, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x5f, 0x72, 0x63, 0x76, 0x62,
	0x75, 0x66, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x6f, 0x52, 0x63, 0x76, 0x62,
	0x75, 0x66, 0x12, 0x23, 0x0a, 0x0e, 0x73, 0x6f, 0x5f, 0x73, 0x6e, 0x64, 0x5f, 0x74, 0x6d, 0x6f,
	0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x6f, 0x53, 0x6e,
	0x64, 0x54, 0x6d, 0x6f, 0x53, 0x65, 0x63, 0x12, 0x25, 0x0a, 0x0f, 0x73, 0x6f, 0x5f, 0x73, 0x6e,
	0x64, 0x5f, 0x74, 0x6d, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x18, 0x04, 0x20, 0x02, 0x28, 0x04,
	0x52, 0x0c, 0x73, 0x6f, 0x53, 0x6e, 0x64, 0x54, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x63, 0x12, 0x23,
	0x0a, 0x0e, 0x73, 0x6f, 0x5f, 0x72, 0x63, 0x76, 0x5f, 0x74, 0x6d, 0x6f, 0x5f, 0x73, 0x65, 0x63,
	0x18, 0x05, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x6f, 0x52, 0x63, 0x76, 0x54, 0x6d, 0x6f,
	0x53, 0x65, 0x63, 0x12, 0x25, 0x0a, 0x0f, 0x73, 0x6f, 0x5f, 0x72, 0x63, 0x76, 0x5f, 0x74, 0x6d,
	0x6f, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x18, 0x06, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x6f,
	0x52, 0x63, 0x76, 0x54, 0x6d, 0x6f, 0x55, 0x73, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x75, 0x73, 0x65, 0x61, 0x64, 0x64, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72,
	0x65, 0x75, 0x73, 0x65, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x5f, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73,
	0x6f, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x5f,
	0x72, 0x63, 0x76, 0x6c, 0x6f, 0x77, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x6f, 0x52, 0x63, 0x76, 0x6c, 0x6f, 0x77, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f,
	0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x6f, 0x4d,
	0x61, 0x72, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x63, 0x72,
	0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x6f, 0x50, 0x61, 0x73, 0x73,
	0x63, 0x72, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x73,
	0x65, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x6f, 0x50, 0x61, 0x73, 0x73,
	0x73, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x5f, 0x64, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6f, 0x44, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x6f, 0x5f, 0x6e, 0x6f, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x6f, 0x4e,
	0x6f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x6f, 0x5f, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x10, 0x20, 0x03, 0x28, 0x06, 0x52, 0x08, 0x73, 0x6f, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x5f, 0x72, 0x65, 0x75, 0x73,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6f, 0x52,
	0x65, 0x75, 0x73, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x5f, 0x62,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x73, 0x6f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x6f, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x73, 0x6f, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x63, 0x70, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x63, 0x6e, 0x74, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x63, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x63, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x63, 0x70, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x64, 0x6c, 0x65, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x63, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x69, 0x64,
	0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x63, 0x70, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e,
	0x74, 0x76, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x63, 0x70, 0x4b, 0x65,
	0x65, 0x70, 0x69, 0x6e, 0x74, 0x76, 0x6c, 0x2a, 0x36, 0x0a, 0x0b, 0x73, 0x6b, 0x5f, 0x73, 0x68,
	0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52,
	0x49, 0x54, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x54, 0x48, 0x10, 0x03, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x4c,
	0x6f, 0x6e, 0x65, 0x6c, 0x79, 0x2f, 0x63, 0x72, 0x69, 0x75, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
}

var (
	file_sk_opts_proto_rawDescOnce sync.Once
	file_sk_opts_proto_rawDescData = file_sk_opts_proto_rawDesc
)

func file_sk_opts_proto_rawDescGZIP() []byte {
	file_sk_opts_proto_rawDescOnce.Do(func() {
		file_sk_opts_proto_rawDescData = protoimpl.X.CompressGZIP(file_sk_opts_proto_rawDescData)
	})
	return file_sk_opts_proto_rawDescData
}

var file_sk_opts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sk_opts_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sk_opts_proto_goTypes = []interface{}{
	(SkShutdown)(0),     // 0: types.sk_shutdown
	(*SkOptsEntry)(nil), // 1: types.sk_opts_entry
}
var file_sk_opts_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sk_opts_proto_init() }
func file_sk_opts_proto_init() {
	if File_sk_opts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sk_opts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkOptsEntry); i {
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
			RawDescriptor: file_sk_opts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sk_opts_proto_goTypes,
		DependencyIndexes: file_sk_opts_proto_depIdxs,
		EnumInfos:         file_sk_opts_proto_enumTypes,
		MessageInfos:      file_sk_opts_proto_msgTypes,
	}.Build()
	File_sk_opts_proto = out.File
	file_sk_opts_proto_rawDesc = nil
	file_sk_opts_proto_goTypes = nil
	file_sk_opts_proto_depIdxs = nil
}
