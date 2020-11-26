// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: core-s390.proto

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

type UserS390RegsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PswMask    *uint64  `protobuf:"varint,1,req,name=psw_mask,json=pswMask" json:"psw_mask,omitempty"`
	PswAddr    *uint64  `protobuf:"varint,2,req,name=psw_addr,json=pswAddr" json:"psw_addr,omitempty"`
	Gprs       []uint64 `protobuf:"varint,3,rep,name=gprs" json:"gprs,omitempty"`
	Acrs       []uint32 `protobuf:"varint,4,rep,name=acrs" json:"acrs,omitempty"`
	OrigGpr2   *uint64  `protobuf:"varint,5,req,name=orig_gpr2,json=origGpr2" json:"orig_gpr2,omitempty"`
	SystemCall *uint32  `protobuf:"varint,6,req,name=system_call,json=systemCall" json:"system_call,omitempty"`
}

func (x *UserS390RegsEntry) Reset() {
	*x = UserS390RegsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_s390_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserS390RegsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserS390RegsEntry) ProtoMessage() {}

func (x *UserS390RegsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_s390_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserS390RegsEntry.ProtoReflect.Descriptor instead.
func (*UserS390RegsEntry) Descriptor() ([]byte, []int) {
	return file_core_s390_proto_rawDescGZIP(), []int{0}
}

func (x *UserS390RegsEntry) GetPswMask() uint64 {
	if x != nil && x.PswMask != nil {
		return *x.PswMask
	}
	return 0
}

func (x *UserS390RegsEntry) GetPswAddr() uint64 {
	if x != nil && x.PswAddr != nil {
		return *x.PswAddr
	}
	return 0
}

func (x *UserS390RegsEntry) GetGprs() []uint64 {
	if x != nil {
		return x.Gprs
	}
	return nil
}

func (x *UserS390RegsEntry) GetAcrs() []uint32 {
	if x != nil {
		return x.Acrs
	}
	return nil
}

func (x *UserS390RegsEntry) GetOrigGpr2() uint64 {
	if x != nil && x.OrigGpr2 != nil {
		return *x.OrigGpr2
	}
	return 0
}

func (x *UserS390RegsEntry) GetSystemCall() uint32 {
	if x != nil && x.SystemCall != nil {
		return *x.SystemCall
	}
	return 0
}

type UserS390VxrsLowEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regs []uint64 `protobuf:"varint,1,rep,name=regs" json:"regs,omitempty"`
}

func (x *UserS390VxrsLowEntry) Reset() {
	*x = UserS390VxrsLowEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_s390_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserS390VxrsLowEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserS390VxrsLowEntry) ProtoMessage() {}

func (x *UserS390VxrsLowEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_s390_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserS390VxrsLowEntry.ProtoReflect.Descriptor instead.
func (*UserS390VxrsLowEntry) Descriptor() ([]byte, []int) {
	return file_core_s390_proto_rawDescGZIP(), []int{1}
}

func (x *UserS390VxrsLowEntry) GetRegs() []uint64 {
	if x != nil {
		return x.Regs
	}
	return nil
}

//
// The vxrs_high registers have 128 bit:
//
//   vxrs_high_0 = regs[0] << 64 | regs[1];
//   vxrs_high_1 = regs[2] << 64 | regs[3];
type UserS390VxrsHighEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regs []uint64 `protobuf:"varint,1,rep,name=regs" json:"regs,omitempty"`
}

func (x *UserS390VxrsHighEntry) Reset() {
	*x = UserS390VxrsHighEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_s390_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserS390VxrsHighEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserS390VxrsHighEntry) ProtoMessage() {}

func (x *UserS390VxrsHighEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_s390_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserS390VxrsHighEntry.ProtoReflect.Descriptor instead.
func (*UserS390VxrsHighEntry) Descriptor() ([]byte, []int) {
	return file_core_s390_proto_rawDescGZIP(), []int{2}
}

func (x *UserS390VxrsHighEntry) GetRegs() []uint64 {
	if x != nil {
		return x.Regs
	}
	return nil
}

type UserS390FpregsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fpc  *uint32  `protobuf:"varint,1,req,name=fpc" json:"fpc,omitempty"`
	Fprs []uint64 `protobuf:"varint,2,rep,name=fprs" json:"fprs,omitempty"`
}

func (x *UserS390FpregsEntry) Reset() {
	*x = UserS390FpregsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_s390_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserS390FpregsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserS390FpregsEntry) ProtoMessage() {}

func (x *UserS390FpregsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_s390_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserS390FpregsEntry.ProtoReflect.Descriptor instead.
func (*UserS390FpregsEntry) Descriptor() ([]byte, []int) {
	return file_core_s390_proto_rawDescGZIP(), []int{3}
}

func (x *UserS390FpregsEntry) GetFpc() uint32 {
	if x != nil && x.Fpc != nil {
		return *x.Fpc
	}
	return 0
}

func (x *UserS390FpregsEntry) GetFprs() []uint64 {
	if x != nil {
		return x.Fprs
	}
	return nil
}

type UserS390GsCbEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regs []uint64 `protobuf:"varint,1,rep,name=regs" json:"regs,omitempty"`
}

func (x *UserS390GsCbEntry) Reset() {
	*x = UserS390GsCbEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_s390_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserS390GsCbEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserS390GsCbEntry) ProtoMessage() {}

func (x *UserS390GsCbEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_s390_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserS390GsCbEntry.ProtoReflect.Descriptor instead.
func (*UserS390GsCbEntry) Descriptor() ([]byte, []int) {
	return file_core_s390_proto_rawDescGZIP(), []int{4}
}

func (x *UserS390GsCbEntry) GetRegs() []uint64 {
	if x != nil {
		return x.Regs
	}
	return nil
}

type UserS390RiEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RiOn *uint32  `protobuf:"varint,1,req,name=ri_on,json=riOn" json:"ri_on,omitempty"`
	Regs []uint64 `protobuf:"varint,2,rep,name=regs" json:"regs,omitempty"`
}

func (x *UserS390RiEntry) Reset() {
	*x = UserS390RiEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_s390_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserS390RiEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserS390RiEntry) ProtoMessage() {}

func (x *UserS390RiEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_s390_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserS390RiEntry.ProtoReflect.Descriptor instead.
func (*UserS390RiEntry) Descriptor() ([]byte, []int) {
	return file_core_s390_proto_rawDescGZIP(), []int{5}
}

func (x *UserS390RiEntry) GetRiOn() uint32 {
	if x != nil && x.RiOn != nil {
		return *x.RiOn
	}
	return 0
}

func (x *UserS390RiEntry) GetRegs() []uint64 {
	if x != nil {
		return x.Regs
	}
	return nil
}

type ThreadInfoS390 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClearTidAddr *uint64                `protobuf:"varint,1,req,name=clear_tid_addr,json=clearTidAddr" json:"clear_tid_addr,omitempty"`
	Gpregs       *UserS390RegsEntry     `protobuf:"bytes,2,req,name=gpregs" json:"gpregs,omitempty"`
	Fpregs       *UserS390FpregsEntry   `protobuf:"bytes,3,req,name=fpregs" json:"fpregs,omitempty"`
	VxrsLow      *UserS390VxrsLowEntry  `protobuf:"bytes,4,opt,name=vxrs_low,json=vxrsLow" json:"vxrs_low,omitempty"`
	VxrsHigh     *UserS390VxrsHighEntry `protobuf:"bytes,5,opt,name=vxrs_high,json=vxrsHigh" json:"vxrs_high,omitempty"`
	GsCb         *UserS390GsCbEntry     `protobuf:"bytes,6,opt,name=gs_cb,json=gsCb" json:"gs_cb,omitempty"`
	GsBc         *UserS390GsCbEntry     `protobuf:"bytes,7,opt,name=gs_bc,json=gsBc" json:"gs_bc,omitempty"`
	RiCb         *UserS390RiEntry       `protobuf:"bytes,8,opt,name=ri_cb,json=riCb" json:"ri_cb,omitempty"`
}

func (x *ThreadInfoS390) Reset() {
	*x = ThreadInfoS390{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_s390_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadInfoS390) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadInfoS390) ProtoMessage() {}

func (x *ThreadInfoS390) ProtoReflect() protoreflect.Message {
	mi := &file_core_s390_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadInfoS390.ProtoReflect.Descriptor instead.
func (*ThreadInfoS390) Descriptor() ([]byte, []int) {
	return file_core_s390_proto_rawDescGZIP(), []int{6}
}

func (x *ThreadInfoS390) GetClearTidAddr() uint64 {
	if x != nil && x.ClearTidAddr != nil {
		return *x.ClearTidAddr
	}
	return 0
}

func (x *ThreadInfoS390) GetGpregs() *UserS390RegsEntry {
	if x != nil {
		return x.Gpregs
	}
	return nil
}

func (x *ThreadInfoS390) GetFpregs() *UserS390FpregsEntry {
	if x != nil {
		return x.Fpregs
	}
	return nil
}

func (x *ThreadInfoS390) GetVxrsLow() *UserS390VxrsLowEntry {
	if x != nil {
		return x.VxrsLow
	}
	return nil
}

func (x *ThreadInfoS390) GetVxrsHigh() *UserS390VxrsHighEntry {
	if x != nil {
		return x.VxrsHigh
	}
	return nil
}

func (x *ThreadInfoS390) GetGsCb() *UserS390GsCbEntry {
	if x != nil {
		return x.GsCb
	}
	return nil
}

func (x *ThreadInfoS390) GetGsBc() *UserS390GsCbEntry {
	if x != nil {
		return x.GsBc
	}
	return nil
}

func (x *ThreadInfoS390) GetRiCb() *UserS390RiEntry {
	if x != nil {
		return x.RiCb
	}
	return nil
}

var File_core_s390_proto protoreflect.FileDescriptor

var file_core_s390_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x73, 0x33, 0x39, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x33,
	0x39, 0x30, 0x5f, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x70, 0x73, 0x77, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52,
	0x07, 0x70, 0x73, 0x77, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x73, 0x77, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x07, 0x70, 0x73, 0x77, 0x41,
	0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x70, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x04, 0x67, 0x70, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x63, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x04, 0x61, 0x63, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x72, 0x69, 0x67, 0x5f, 0x67, 0x70, 0x72, 0x32, 0x18, 0x05, 0x20, 0x02, 0x28, 0x04, 0x52, 0x08,
	0x6f, 0x72, 0x69, 0x67, 0x47, 0x70, 0x72, 0x32, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x22, 0x2e, 0x0a, 0x18, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x5f, 0x76, 0x78, 0x72, 0x73, 0x5f, 0x6c, 0x6f, 0x77, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x04, 0x72, 0x65, 0x67, 0x73, 0x22, 0x2f, 0x0a, 0x19, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x5f, 0x76, 0x78, 0x72, 0x73, 0x5f, 0x68, 0x69, 0x67, 0x68,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x72, 0x65, 0x67, 0x73, 0x22, 0x3e, 0x0a, 0x16, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x5f, 0x66, 0x70, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x70, 0x63, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x03, 0x66, 0x70, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x70, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x66, 0x70, 0x72, 0x73, 0x22, 0x2b, 0x0a, 0x15, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x5f, 0x67, 0x73, 0x5f, 0x63, 0x62, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x04, 0x72, 0x65, 0x67, 0x73, 0x22, 0x3d, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x33, 0x39, 0x30, 0x5f, 0x72, 0x69, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x13, 0x0a,
	0x05, 0x72, 0x69, 0x5f, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x69,
	0x4f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x04, 0x72, 0x65, 0x67, 0x73, 0x22, 0xed, 0x03, 0x0a, 0x10, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x12, 0x2b, 0x0a, 0x0e, 0x63,
	0x6c, 0x65, 0x61, 0x72, 0x5f, 0x74, 0x69, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x04, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x63, 0x6c, 0x65, 0x61,
	0x72, 0x54, 0x69, 0x64, 0x41, 0x64, 0x64, 0x72, 0x12, 0x3a, 0x0a, 0x06, 0x67, 0x70, 0x72, 0x65,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x5f, 0x72, 0x65, 0x67, 0x73, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x06, 0x67, 0x70,
	0x72, 0x65, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x06, 0x66, 0x70, 0x72, 0x65, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x5f, 0x66, 0x70, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x06, 0x66, 0x70, 0x72, 0x65,
	0x67, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x76, 0x78, 0x72, 0x73, 0x5f, 0x6c, 0x6f, 0x77, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x5f, 0x76, 0x78, 0x72, 0x73, 0x5f, 0x6c, 0x6f, 0x77, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x07, 0x76, 0x78,
	0x72, 0x73, 0x4c, 0x6f, 0x77, 0x12, 0x44, 0x0a, 0x09, 0x76, 0x78, 0x72, 0x73, 0x5f, 0x68, 0x69,
	0x67, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x5f, 0x76, 0x78, 0x72, 0x73, 0x5f,
	0x68, 0x69, 0x67, 0x68, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08,
	0x01, 0x52, 0x08, 0x76, 0x78, 0x72, 0x73, 0x48, 0x69, 0x67, 0x68, 0x12, 0x38, 0x0a, 0x05, 0x67,
	0x73, 0x5f, 0x63, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x5f, 0x67, 0x73, 0x5f,
	0x63, 0x62, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52,
	0x04, 0x67, 0x73, 0x43, 0x62, 0x12, 0x38, 0x0a, 0x05, 0x67, 0x73, 0x5f, 0x62, 0x63, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x33, 0x39, 0x30, 0x5f, 0x67, 0x73, 0x5f, 0x63, 0x62, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x04, 0x67, 0x73, 0x42, 0x63, 0x12,
	0x35, 0x0a, 0x05, 0x72, 0x69, 0x5f, 0x63, 0x62, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x33, 0x39, 0x30,
	0x5f, 0x72, 0x69, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01,
	0x52, 0x04, 0x72, 0x69, 0x43, 0x62, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x4c, 0x6f, 0x6e, 0x65, 0x6c, 0x79, 0x2f, 0x63, 0x72, 0x69,
	0x75, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
}

var (
	file_core_s390_proto_rawDescOnce sync.Once
	file_core_s390_proto_rawDescData = file_core_s390_proto_rawDesc
)

func file_core_s390_proto_rawDescGZIP() []byte {
	file_core_s390_proto_rawDescOnce.Do(func() {
		file_core_s390_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_s390_proto_rawDescData)
	})
	return file_core_s390_proto_rawDescData
}

var file_core_s390_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_core_s390_proto_goTypes = []interface{}{
	(*UserS390RegsEntry)(nil),     // 0: types.user_s390_regs_entry
	(*UserS390VxrsLowEntry)(nil),  // 1: types.user_s390_vxrs_low_entry
	(*UserS390VxrsHighEntry)(nil), // 2: types.user_s390_vxrs_high_entry
	(*UserS390FpregsEntry)(nil),   // 3: types.user_s390_fpregs_entry
	(*UserS390GsCbEntry)(nil),     // 4: types.user_s390_gs_cb_entry
	(*UserS390RiEntry)(nil),       // 5: types.user_s390_ri_entry
	(*ThreadInfoS390)(nil),        // 6: types.thread_info_s390
}
var file_core_s390_proto_depIdxs = []int32{
	0, // 0: types.thread_info_s390.gpregs:type_name -> types.user_s390_regs_entry
	3, // 1: types.thread_info_s390.fpregs:type_name -> types.user_s390_fpregs_entry
	1, // 2: types.thread_info_s390.vxrs_low:type_name -> types.user_s390_vxrs_low_entry
	2, // 3: types.thread_info_s390.vxrs_high:type_name -> types.user_s390_vxrs_high_entry
	4, // 4: types.thread_info_s390.gs_cb:type_name -> types.user_s390_gs_cb_entry
	4, // 5: types.thread_info_s390.gs_bc:type_name -> types.user_s390_gs_cb_entry
	5, // 6: types.thread_info_s390.ri_cb:type_name -> types.user_s390_ri_entry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_core_s390_proto_init() }
func file_core_s390_proto_init() {
	if File_core_s390_proto != nil {
		return
	}
	file_opts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_s390_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserS390RegsEntry); i {
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
		file_core_s390_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserS390VxrsLowEntry); i {
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
		file_core_s390_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserS390VxrsHighEntry); i {
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
		file_core_s390_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserS390FpregsEntry); i {
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
		file_core_s390_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserS390GsCbEntry); i {
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
		file_core_s390_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserS390RiEntry); i {
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
		file_core_s390_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreadInfoS390); i {
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
			RawDescriptor: file_core_s390_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_s390_proto_goTypes,
		DependencyIndexes: file_core_s390_proto_depIdxs,
		MessageInfos:      file_core_s390_proto_msgTypes,
	}.Build()
	File_core_s390_proto = out.File
	file_core_s390_proto_rawDesc = nil
	file_core_s390_proto_goTypes = nil
	file_core_s390_proto_depIdxs = nil
}
