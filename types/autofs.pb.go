// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: autofs.proto

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

type AutofsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fd       *int32 `protobuf:"varint,1,req,name=fd" json:"fd,omitempty"`
	Pgrp     *int32 `protobuf:"varint,2,req,name=pgrp" json:"pgrp,omitempty"`
	Timeout  *int32 `protobuf:"varint,3,req,name=timeout" json:"timeout,omitempty"`
	Minproto *int32 `protobuf:"varint,4,req,name=minproto" json:"minproto,omitempty"`
	Maxproto *int32 `protobuf:"varint,5,req,name=maxproto" json:"maxproto,omitempty"`
	Mode     *int32 `protobuf:"varint,6,req,name=mode" json:"mode,omitempty"`
	Uid      *int32 `protobuf:"varint,7,opt,name=uid" json:"uid,omitempty"`
	Gid      *int32 `protobuf:"varint,8,opt,name=gid" json:"gid,omitempty"`
	ReadFd   *int32 `protobuf:"varint,9,opt,name=read_fd,json=readFd" json:"read_fd,omitempty"`
}

func (x *AutofsEntry) Reset() {
	*x = AutofsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autofs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutofsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutofsEntry) ProtoMessage() {}

func (x *AutofsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_autofs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutofsEntry.ProtoReflect.Descriptor instead.
func (*AutofsEntry) Descriptor() ([]byte, []int) {
	return file_autofs_proto_rawDescGZIP(), []int{0}
}

func (x *AutofsEntry) GetFd() int32 {
	if x != nil && x.Fd != nil {
		return *x.Fd
	}
	return 0
}

func (x *AutofsEntry) GetPgrp() int32 {
	if x != nil && x.Pgrp != nil {
		return *x.Pgrp
	}
	return 0
}

func (x *AutofsEntry) GetTimeout() int32 {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return 0
}

func (x *AutofsEntry) GetMinproto() int32 {
	if x != nil && x.Minproto != nil {
		return *x.Minproto
	}
	return 0
}

func (x *AutofsEntry) GetMaxproto() int32 {
	if x != nil && x.Maxproto != nil {
		return *x.Maxproto
	}
	return 0
}

func (x *AutofsEntry) GetMode() int32 {
	if x != nil && x.Mode != nil {
		return *x.Mode
	}
	return 0
}

func (x *AutofsEntry) GetUid() int32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *AutofsEntry) GetGid() int32 {
	if x != nil && x.Gid != nil {
		return *x.Gid
	}
	return 0
}

func (x *AutofsEntry) GetReadFd() int32 {
	if x != nil && x.ReadFd != nil {
		return *x.ReadFd
	}
	return 0
}

var File_autofs_proto protoreflect.FileDescriptor

var file_autofs_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xd5, 0x01, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x66, 0x73,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x64, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x02, 0x66, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x67, 0x72, 0x70, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x04, 0x70, 0x67, 0x72, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x02, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x67, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x66, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x61, 0x64, 0x46, 0x64, 0x42, 0x25, 0x5a,
	0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x4c, 0x6f, 0x6e,
	0x65, 0x6c, 0x79, 0x2f, 0x63, 0x72, 0x69, 0x75, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73,
}

var (
	file_autofs_proto_rawDescOnce sync.Once
	file_autofs_proto_rawDescData = file_autofs_proto_rawDesc
)

func file_autofs_proto_rawDescGZIP() []byte {
	file_autofs_proto_rawDescOnce.Do(func() {
		file_autofs_proto_rawDescData = protoimpl.X.CompressGZIP(file_autofs_proto_rawDescData)
	})
	return file_autofs_proto_rawDescData
}

var file_autofs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_autofs_proto_goTypes = []interface{}{
	(*AutofsEntry)(nil), // 0: types.autofs_entry
}
var file_autofs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_autofs_proto_init() }
func file_autofs_proto_init() {
	if File_autofs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autofs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutofsEntry); i {
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
			RawDescriptor: file_autofs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_autofs_proto_goTypes,
		DependencyIndexes: file_autofs_proto_depIdxs,
		MessageInfos:      file_autofs_proto_msgTypes,
	}.Build()
	File_autofs_proto = out.File
	file_autofs_proto_rawDesc = nil
	file_autofs_proto_goTypes = nil
	file_autofs_proto_depIdxs = nil
}
