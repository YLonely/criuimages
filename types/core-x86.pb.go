// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: core-x86.proto

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

type UserX86RegsMode int32

const (
	UserX86RegsMode_NATIVE UserX86RegsMode = 1
	UserX86RegsMode_COMPAT UserX86RegsMode = 2
)

// Enum value maps for UserX86RegsMode.
var (
	UserX86RegsMode_name = map[int32]string{
		1: "NATIVE",
		2: "COMPAT",
	}
	UserX86RegsMode_value = map[string]int32{
		"NATIVE": 1,
		"COMPAT": 2,
	}
)

func (x UserX86RegsMode) Enum() *UserX86RegsMode {
	p := new(UserX86RegsMode)
	*p = x
	return p
}

func (x UserX86RegsMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserX86RegsMode) Descriptor() protoreflect.EnumDescriptor {
	return file_core_x86_proto_enumTypes[0].Descriptor()
}

func (UserX86RegsMode) Type() protoreflect.EnumType {
	return &file_core_x86_proto_enumTypes[0]
}

func (x UserX86RegsMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *UserX86RegsMode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = UserX86RegsMode(num)
	return nil
}

// Deprecated: Use UserX86RegsMode.Descriptor instead.
func (UserX86RegsMode) EnumDescriptor() ([]byte, []int) {
	return file_core_x86_proto_rawDescGZIP(), []int{0}
}

// Reusing entry for both 64 and 32 bits register sets
type UserX86RegsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R15    *uint64          `protobuf:"varint,1,req,name=r15" json:"r15,omitempty"`
	R14    *uint64          `protobuf:"varint,2,req,name=r14" json:"r14,omitempty"`
	R13    *uint64          `protobuf:"varint,3,req,name=r13" json:"r13,omitempty"`
	R12    *uint64          `protobuf:"varint,4,req,name=r12" json:"r12,omitempty"`
	Bp     *uint64          `protobuf:"varint,5,req,name=bp" json:"bp,omitempty"`
	Bx     *uint64          `protobuf:"varint,6,req,name=bx" json:"bx,omitempty"`
	R11    *uint64          `protobuf:"varint,7,req,name=r11" json:"r11,omitempty"`
	R10    *uint64          `protobuf:"varint,8,req,name=r10" json:"r10,omitempty"`
	R9     *uint64          `protobuf:"varint,9,req,name=r9" json:"r9,omitempty"`
	R8     *uint64          `protobuf:"varint,10,req,name=r8" json:"r8,omitempty"`
	Ax     *uint64          `protobuf:"varint,11,req,name=ax" json:"ax,omitempty"`
	Cx     *uint64          `protobuf:"varint,12,req,name=cx" json:"cx,omitempty"`
	Dx     *uint64          `protobuf:"varint,13,req,name=dx" json:"dx,omitempty"`
	Si     *uint64          `protobuf:"varint,14,req,name=si" json:"si,omitempty"`
	Di     *uint64          `protobuf:"varint,15,req,name=di" json:"di,omitempty"`
	OrigAx *uint64          `protobuf:"varint,16,req,name=orig_ax,json=origAx" json:"orig_ax,omitempty"`
	Ip     *uint64          `protobuf:"varint,17,req,name=ip" json:"ip,omitempty"`
	Cs     *uint64          `protobuf:"varint,18,req,name=cs" json:"cs,omitempty"`
	Flags  *uint64          `protobuf:"varint,19,req,name=flags" json:"flags,omitempty"`
	Sp     *uint64          `protobuf:"varint,20,req,name=sp" json:"sp,omitempty"`
	Ss     *uint64          `protobuf:"varint,21,req,name=ss" json:"ss,omitempty"`
	FsBase *uint64          `protobuf:"varint,22,req,name=fs_base,json=fsBase" json:"fs_base,omitempty"`
	GsBase *uint64          `protobuf:"varint,23,req,name=gs_base,json=gsBase" json:"gs_base,omitempty"`
	Ds     *uint64          `protobuf:"varint,24,req,name=ds" json:"ds,omitempty"`
	Es     *uint64          `protobuf:"varint,25,req,name=es" json:"es,omitempty"`
	Fs     *uint64          `protobuf:"varint,26,req,name=fs" json:"fs,omitempty"`
	Gs     *uint64          `protobuf:"varint,27,req,name=gs" json:"gs,omitempty"`
	Mode   *UserX86RegsMode `protobuf:"varint,28,opt,name=mode,enum=types.UserX86RegsMode,def=1" json:"mode,omitempty"`
}

// Default values for UserX86RegsEntry fields.
const (
	Default_UserX86RegsEntry_Mode = UserX86RegsMode_NATIVE
)

func (x *UserX86RegsEntry) Reset() {
	*x = UserX86RegsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_x86_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserX86RegsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserX86RegsEntry) ProtoMessage() {}

func (x *UserX86RegsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_x86_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserX86RegsEntry.ProtoReflect.Descriptor instead.
func (*UserX86RegsEntry) Descriptor() ([]byte, []int) {
	return file_core_x86_proto_rawDescGZIP(), []int{0}
}

func (x *UserX86RegsEntry) GetR15() uint64 {
	if x != nil && x.R15 != nil {
		return *x.R15
	}
	return 0
}

func (x *UserX86RegsEntry) GetR14() uint64 {
	if x != nil && x.R14 != nil {
		return *x.R14
	}
	return 0
}

func (x *UserX86RegsEntry) GetR13() uint64 {
	if x != nil && x.R13 != nil {
		return *x.R13
	}
	return 0
}

func (x *UserX86RegsEntry) GetR12() uint64 {
	if x != nil && x.R12 != nil {
		return *x.R12
	}
	return 0
}

func (x *UserX86RegsEntry) GetBp() uint64 {
	if x != nil && x.Bp != nil {
		return *x.Bp
	}
	return 0
}

func (x *UserX86RegsEntry) GetBx() uint64 {
	if x != nil && x.Bx != nil {
		return *x.Bx
	}
	return 0
}

func (x *UserX86RegsEntry) GetR11() uint64 {
	if x != nil && x.R11 != nil {
		return *x.R11
	}
	return 0
}

func (x *UserX86RegsEntry) GetR10() uint64 {
	if x != nil && x.R10 != nil {
		return *x.R10
	}
	return 0
}

func (x *UserX86RegsEntry) GetR9() uint64 {
	if x != nil && x.R9 != nil {
		return *x.R9
	}
	return 0
}

func (x *UserX86RegsEntry) GetR8() uint64 {
	if x != nil && x.R8 != nil {
		return *x.R8
	}
	return 0
}

func (x *UserX86RegsEntry) GetAx() uint64 {
	if x != nil && x.Ax != nil {
		return *x.Ax
	}
	return 0
}

func (x *UserX86RegsEntry) GetCx() uint64 {
	if x != nil && x.Cx != nil {
		return *x.Cx
	}
	return 0
}

func (x *UserX86RegsEntry) GetDx() uint64 {
	if x != nil && x.Dx != nil {
		return *x.Dx
	}
	return 0
}

func (x *UserX86RegsEntry) GetSi() uint64 {
	if x != nil && x.Si != nil {
		return *x.Si
	}
	return 0
}

func (x *UserX86RegsEntry) GetDi() uint64 {
	if x != nil && x.Di != nil {
		return *x.Di
	}
	return 0
}

func (x *UserX86RegsEntry) GetOrigAx() uint64 {
	if x != nil && x.OrigAx != nil {
		return *x.OrigAx
	}
	return 0
}

func (x *UserX86RegsEntry) GetIp() uint64 {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return 0
}

func (x *UserX86RegsEntry) GetCs() uint64 {
	if x != nil && x.Cs != nil {
		return *x.Cs
	}
	return 0
}

func (x *UserX86RegsEntry) GetFlags() uint64 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *UserX86RegsEntry) GetSp() uint64 {
	if x != nil && x.Sp != nil {
		return *x.Sp
	}
	return 0
}

func (x *UserX86RegsEntry) GetSs() uint64 {
	if x != nil && x.Ss != nil {
		return *x.Ss
	}
	return 0
}

func (x *UserX86RegsEntry) GetFsBase() uint64 {
	if x != nil && x.FsBase != nil {
		return *x.FsBase
	}
	return 0
}

func (x *UserX86RegsEntry) GetGsBase() uint64 {
	if x != nil && x.GsBase != nil {
		return *x.GsBase
	}
	return 0
}

func (x *UserX86RegsEntry) GetDs() uint64 {
	if x != nil && x.Ds != nil {
		return *x.Ds
	}
	return 0
}

func (x *UserX86RegsEntry) GetEs() uint64 {
	if x != nil && x.Es != nil {
		return *x.Es
	}
	return 0
}

func (x *UserX86RegsEntry) GetFs() uint64 {
	if x != nil && x.Fs != nil {
		return *x.Fs
	}
	return 0
}

func (x *UserX86RegsEntry) GetGs() uint64 {
	if x != nil && x.Gs != nil {
		return *x.Gs
	}
	return 0
}

func (x *UserX86RegsEntry) GetMode() UserX86RegsMode {
	if x != nil && x.Mode != nil {
		return *x.Mode
	}
	return Default_UserX86RegsEntry_Mode
}

type UserX86XsaveEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// standart xsave features
	XstateBv *uint64 `protobuf:"varint,1,req,name=xstate_bv,json=xstateBv" json:"xstate_bv,omitempty"`
	// AVX components: 16x 256-bit ymm registers, hi 128 bits
	YmmhSpace []uint32 `protobuf:"varint,2,rep,name=ymmh_space,json=ymmhSpace" json:"ymmh_space,omitempty"`
	// MPX components
	BndregState []uint64 `protobuf:"varint,3,rep,name=bndreg_state,json=bndregState" json:"bndreg_state,omitempty"`
	BndcsrState []uint64 `protobuf:"varint,4,rep,name=bndcsr_state,json=bndcsrState" json:"bndcsr_state,omitempty"`
	// AVX512 components: k0-k7, ZMM_Hi256, Hi16_ZMM
	OpmaskReg []uint64 `protobuf:"varint,5,rep,name=opmask_reg,json=opmaskReg" json:"opmask_reg,omitempty"`
	ZmmUpper  []uint64 `protobuf:"varint,6,rep,name=zmm_upper,json=zmmUpper" json:"zmm_upper,omitempty"`
	Hi16Zmm   []uint64 `protobuf:"varint,7,rep,name=hi16_zmm,json=hi16Zmm" json:"hi16_zmm,omitempty"`
	// Protected keys
	Pkru []uint32 `protobuf:"varint,8,rep,name=pkru" json:"pkru,omitempty"`
}

func (x *UserX86XsaveEntry) Reset() {
	*x = UserX86XsaveEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_x86_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserX86XsaveEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserX86XsaveEntry) ProtoMessage() {}

func (x *UserX86XsaveEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_x86_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserX86XsaveEntry.ProtoReflect.Descriptor instead.
func (*UserX86XsaveEntry) Descriptor() ([]byte, []int) {
	return file_core_x86_proto_rawDescGZIP(), []int{1}
}

func (x *UserX86XsaveEntry) GetXstateBv() uint64 {
	if x != nil && x.XstateBv != nil {
		return *x.XstateBv
	}
	return 0
}

func (x *UserX86XsaveEntry) GetYmmhSpace() []uint32 {
	if x != nil {
		return x.YmmhSpace
	}
	return nil
}

func (x *UserX86XsaveEntry) GetBndregState() []uint64 {
	if x != nil {
		return x.BndregState
	}
	return nil
}

func (x *UserX86XsaveEntry) GetBndcsrState() []uint64 {
	if x != nil {
		return x.BndcsrState
	}
	return nil
}

func (x *UserX86XsaveEntry) GetOpmaskReg() []uint64 {
	if x != nil {
		return x.OpmaskReg
	}
	return nil
}

func (x *UserX86XsaveEntry) GetZmmUpper() []uint64 {
	if x != nil {
		return x.ZmmUpper
	}
	return nil
}

func (x *UserX86XsaveEntry) GetHi16Zmm() []uint64 {
	if x != nil {
		return x.Hi16Zmm
	}
	return nil
}

func (x *UserX86XsaveEntry) GetPkru() []uint32 {
	if x != nil {
		return x.Pkru
	}
	return nil
}

type UserX86FpregsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fxsave data
	Cwd       *uint32  `protobuf:"varint,1,req,name=cwd" json:"cwd,omitempty"`
	Swd       *uint32  `protobuf:"varint,2,req,name=swd" json:"swd,omitempty"`
	Twd       *uint32  `protobuf:"varint,3,req,name=twd" json:"twd,omitempty"`
	Fop       *uint32  `protobuf:"varint,4,req,name=fop" json:"fop,omitempty"`
	Rip       *uint64  `protobuf:"varint,5,req,name=rip" json:"rip,omitempty"`
	Rdp       *uint64  `protobuf:"varint,6,req,name=rdp" json:"rdp,omitempty"`
	Mxcsr     *uint32  `protobuf:"varint,7,req,name=mxcsr" json:"mxcsr,omitempty"`
	MxcsrMask *uint32  `protobuf:"varint,8,req,name=mxcsr_mask,json=mxcsrMask" json:"mxcsr_mask,omitempty"`
	StSpace   []uint32 `protobuf:"varint,9,rep,name=st_space,json=stSpace" json:"st_space,omitempty"`
	XmmSpace  []uint32 `protobuf:"varint,10,rep,name=xmm_space,json=xmmSpace" json:"xmm_space,omitempty"`
	// Unused, but present for backward compatibility
	Padding []uint32 `protobuf:"varint,11,rep,name=padding" json:"padding,omitempty"`
	// xsave extension
	Xsave *UserX86XsaveEntry `protobuf:"bytes,13,opt,name=xsave" json:"xsave,omitempty"`
}

func (x *UserX86FpregsEntry) Reset() {
	*x = UserX86FpregsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_x86_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserX86FpregsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserX86FpregsEntry) ProtoMessage() {}

func (x *UserX86FpregsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_x86_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserX86FpregsEntry.ProtoReflect.Descriptor instead.
func (*UserX86FpregsEntry) Descriptor() ([]byte, []int) {
	return file_core_x86_proto_rawDescGZIP(), []int{2}
}

func (x *UserX86FpregsEntry) GetCwd() uint32 {
	if x != nil && x.Cwd != nil {
		return *x.Cwd
	}
	return 0
}

func (x *UserX86FpregsEntry) GetSwd() uint32 {
	if x != nil && x.Swd != nil {
		return *x.Swd
	}
	return 0
}

func (x *UserX86FpregsEntry) GetTwd() uint32 {
	if x != nil && x.Twd != nil {
		return *x.Twd
	}
	return 0
}

func (x *UserX86FpregsEntry) GetFop() uint32 {
	if x != nil && x.Fop != nil {
		return *x.Fop
	}
	return 0
}

func (x *UserX86FpregsEntry) GetRip() uint64 {
	if x != nil && x.Rip != nil {
		return *x.Rip
	}
	return 0
}

func (x *UserX86FpregsEntry) GetRdp() uint64 {
	if x != nil && x.Rdp != nil {
		return *x.Rdp
	}
	return 0
}

func (x *UserX86FpregsEntry) GetMxcsr() uint32 {
	if x != nil && x.Mxcsr != nil {
		return *x.Mxcsr
	}
	return 0
}

func (x *UserX86FpregsEntry) GetMxcsrMask() uint32 {
	if x != nil && x.MxcsrMask != nil {
		return *x.MxcsrMask
	}
	return 0
}

func (x *UserX86FpregsEntry) GetStSpace() []uint32 {
	if x != nil {
		return x.StSpace
	}
	return nil
}

func (x *UserX86FpregsEntry) GetXmmSpace() []uint32 {
	if x != nil {
		return x.XmmSpace
	}
	return nil
}

func (x *UserX86FpregsEntry) GetPadding() []uint32 {
	if x != nil {
		return x.Padding
	}
	return nil
}

func (x *UserX86FpregsEntry) GetXsave() *UserX86XsaveEntry {
	if x != nil {
		return x.Xsave
	}
	return nil
}

type UserDescT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryNumber *uint32 `protobuf:"varint,1,req,name=entry_number,json=entryNumber" json:"entry_number,omitempty"`
	// this is for GDT, not for MSRs - 32-bit base
	BaseAddr      *uint32 `protobuf:"varint,2,req,name=base_addr,json=baseAddr" json:"base_addr,omitempty"`
	Limit         *uint32 `protobuf:"varint,3,req,name=limit" json:"limit,omitempty"`
	Seg_32Bit     *bool   `protobuf:"varint,4,req,name=seg_32bit,json=seg32bit" json:"seg_32bit,omitempty"`
	ContentsH     *bool   `protobuf:"varint,5,req,name=contents_h,json=contentsH" json:"contents_h,omitempty"`
	ContentsL     *bool   `protobuf:"varint,6,req,name=contents_l,json=contentsL" json:"contents_l,omitempty"`
	ReadExecOnly  *bool   `protobuf:"varint,7,req,name=read_exec_only,json=readExecOnly,def=1" json:"read_exec_only,omitempty"`
	LimitInPages  *bool   `protobuf:"varint,8,req,name=limit_in_pages,json=limitInPages" json:"limit_in_pages,omitempty"`
	SegNotPresent *bool   `protobuf:"varint,9,req,name=seg_not_present,json=segNotPresent,def=1" json:"seg_not_present,omitempty"`
	Useable       *bool   `protobuf:"varint,10,req,name=useable" json:"useable,omitempty"`
}

// Default values for UserDescT fields.
const (
	Default_UserDescT_ReadExecOnly  = bool(true)
	Default_UserDescT_SegNotPresent = bool(true)
)

func (x *UserDescT) Reset() {
	*x = UserDescT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_x86_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDescT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDescT) ProtoMessage() {}

func (x *UserDescT) ProtoReflect() protoreflect.Message {
	mi := &file_core_x86_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDescT.ProtoReflect.Descriptor instead.
func (*UserDescT) Descriptor() ([]byte, []int) {
	return file_core_x86_proto_rawDescGZIP(), []int{3}
}

func (x *UserDescT) GetEntryNumber() uint32 {
	if x != nil && x.EntryNumber != nil {
		return *x.EntryNumber
	}
	return 0
}

func (x *UserDescT) GetBaseAddr() uint32 {
	if x != nil && x.BaseAddr != nil {
		return *x.BaseAddr
	}
	return 0
}

func (x *UserDescT) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *UserDescT) GetSeg_32Bit() bool {
	if x != nil && x.Seg_32Bit != nil {
		return *x.Seg_32Bit
	}
	return false
}

func (x *UserDescT) GetContentsH() bool {
	if x != nil && x.ContentsH != nil {
		return *x.ContentsH
	}
	return false
}

func (x *UserDescT) GetContentsL() bool {
	if x != nil && x.ContentsL != nil {
		return *x.ContentsL
	}
	return false
}

func (x *UserDescT) GetReadExecOnly() bool {
	if x != nil && x.ReadExecOnly != nil {
		return *x.ReadExecOnly
	}
	return Default_UserDescT_ReadExecOnly
}

func (x *UserDescT) GetLimitInPages() bool {
	if x != nil && x.LimitInPages != nil {
		return *x.LimitInPages
	}
	return false
}

func (x *UserDescT) GetSegNotPresent() bool {
	if x != nil && x.SegNotPresent != nil {
		return *x.SegNotPresent
	}
	return Default_UserDescT_SegNotPresent
}

func (x *UserDescT) GetUseable() bool {
	if x != nil && x.Useable != nil {
		return *x.Useable
	}
	return false
}

type ThreadInfoX86 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClearTidAddr *uint64             `protobuf:"varint,1,req,name=clear_tid_addr,json=clearTidAddr" json:"clear_tid_addr,omitempty"`
	Gpregs       *UserX86RegsEntry   `protobuf:"bytes,2,req,name=gpregs" json:"gpregs,omitempty"`
	Fpregs       *UserX86FpregsEntry `protobuf:"bytes,3,req,name=fpregs" json:"fpregs,omitempty"`
	Tls          []*UserDescT        `protobuf:"bytes,4,rep,name=tls" json:"tls,omitempty"`
}

func (x *ThreadInfoX86) Reset() {
	*x = ThreadInfoX86{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_x86_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadInfoX86) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadInfoX86) ProtoMessage() {}

func (x *ThreadInfoX86) ProtoReflect() protoreflect.Message {
	mi := &file_core_x86_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadInfoX86.ProtoReflect.Descriptor instead.
func (*ThreadInfoX86) Descriptor() ([]byte, []int) {
	return file_core_x86_proto_rawDescGZIP(), []int{4}
}

func (x *ThreadInfoX86) GetClearTidAddr() uint64 {
	if x != nil && x.ClearTidAddr != nil {
		return *x.ClearTidAddr
	}
	return 0
}

func (x *ThreadInfoX86) GetGpregs() *UserX86RegsEntry {
	if x != nil {
		return x.Gpregs
	}
	return nil
}

func (x *ThreadInfoX86) GetFpregs() *UserX86FpregsEntry {
	if x != nil {
		return x.Fpregs
	}
	return nil
}

func (x *ThreadInfoX86) GetTls() []*UserDescT {
	if x != nil {
		return x.Tls
	}
	return nil
}

var File_core_x86_proto protoreflect.FileDescriptor

var file_core_x86_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x78, 0x38, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x04, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x78, 0x38, 0x36,
	0x5f, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x31, 0x35, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x35, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x31, 0x34, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x34, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x31, 0x33, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x03, 0x72, 0x31,
	0x33, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x32, 0x18, 0x04, 0x20, 0x02, 0x28, 0x04, 0x52, 0x03,
	0x72, 0x31, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x70, 0x18, 0x05, 0x20, 0x02, 0x28, 0x04, 0x52,
	0x02, 0x62, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x78, 0x18, 0x06, 0x20, 0x02, 0x28, 0x04, 0x52,
	0x02, 0x62, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x31, 0x18, 0x07, 0x20, 0x02, 0x28, 0x04,
	0x52, 0x03, 0x72, 0x31, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x31, 0x30, 0x18, 0x08, 0x20, 0x02,
	0x28, 0x04, 0x52, 0x03, 0x72, 0x31, 0x30, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x39, 0x18, 0x09, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x02, 0x72, 0x39, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x38, 0x18, 0x0a, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x02, 0x72, 0x38, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x78, 0x18, 0x0b, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x02, 0x61, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x78, 0x18, 0x0c, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x02, 0x63, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x78, 0x18, 0x0d, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x02, 0x64, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x69, 0x18, 0x0e, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x02, 0x73, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x69, 0x18, 0x0f, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x02, 0x64, 0x69, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x5f,
	0x61, 0x78, 0x18, 0x10, 0x20, 0x02, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x41, 0x78,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x11, 0x20, 0x02, 0x28, 0x04, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x63, 0x73, 0x18, 0x12, 0x20, 0x02, 0x28, 0x04, 0x52, 0x02, 0x63, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x13, 0x20, 0x02, 0x28, 0x04, 0x52,
	0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x70, 0x18, 0x14, 0x20, 0x02,
	0x28, 0x04, 0x52, 0x02, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x73, 0x18, 0x15, 0x20, 0x02,
	0x28, 0x04, 0x52, 0x02, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x73, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x16, 0x20, 0x02, 0x28, 0x04, 0x52, 0x06, 0x66, 0x73, 0x42, 0x61, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x67, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18, 0x17, 0x20, 0x02, 0x28, 0x04,
	0x52, 0x06, 0x67, 0x73, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x73, 0x18, 0x18,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x02, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x73, 0x18, 0x19,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x02, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x73, 0x18, 0x1a,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x02, 0x66, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x67, 0x73, 0x18, 0x1b,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x02, 0x67, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x1c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x78, 0x38, 0x36, 0x5f, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x3a, 0x06, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22,
	0x83, 0x02, 0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x78, 0x38, 0x36, 0x5f, 0x78, 0x73, 0x61,
	0x76, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x78, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x62, 0x76, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x08, 0x78, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x76, 0x12, 0x1d, 0x0a, 0x0a, 0x79, 0x6d, 0x6d, 0x68, 0x5f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x79, 0x6d, 0x6d, 0x68, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6e, 0x64, 0x72, 0x65, 0x67, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6e, 0x64, 0x72,
	0x65, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6e, 0x64, 0x63, 0x73,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x62,
	0x6e, 0x64, 0x63, 0x73, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x70,
	0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09,
	0x6f, 0x70, 0x6d, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x7a, 0x6d, 0x6d,
	0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x7a, 0x6d,
	0x6d, 0x55, 0x70, 0x70, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x69, 0x31, 0x36, 0x5f, 0x7a,
	0x6d, 0x6d, 0x18, 0x07, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x68, 0x69, 0x31, 0x36, 0x5a, 0x6d,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6b, 0x72, 0x75, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x04, 0x70, 0x6b, 0x72, 0x75, 0x22, 0xbd, 0x02, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x78,
	0x38, 0x36, 0x5f, 0x66, 0x70, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x77, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x77,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x77, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03,
	0x73, 0x77, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x77, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x03, 0x74, 0x77, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x03, 0x66, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x70, 0x18, 0x05,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x03, 0x72, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x64, 0x70,
	0x18, 0x06, 0x20, 0x02, 0x28, 0x04, 0x52, 0x03, 0x72, 0x64, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x78, 0x63, 0x73, 0x72, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x78, 0x63, 0x73,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x78, 0x63, 0x73, 0x72, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x08, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x78, 0x63, 0x73, 0x72, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x78,
	0x6d, 0x6d, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08,
	0x78, 0x6d, 0x6d, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x61, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x31, 0x0a, 0x05, 0x78, 0x73, 0x61, 0x76, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x78,
	0x38, 0x36, 0x5f, 0x78, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x78, 0x73, 0x61, 0x76, 0x65, 0x22, 0xd8, 0x02, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x5f, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x65, 0x67, 0x5f, 0x33, 0x32, 0x62, 0x69, 0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x08, 0x52, 0x08,
	0x73, 0x65, 0x67, 0x33, 0x32, 0x62, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x68, 0x18, 0x05, 0x20, 0x02, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x5f, 0x6c, 0x18, 0x06, 0x20, 0x02, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x12, 0x2a, 0x0a, 0x0e, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x02, 0x28, 0x08, 0x3a, 0x04,
	0x74, 0x72, 0x75, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x45, 0x78, 0x65, 0x63, 0x4f, 0x6e,
	0x6c, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x08, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0c, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x49, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x67, 0x5f,
	0x6e, 0x6f, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x02, 0x28,
	0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x0d, 0x73, 0x65, 0x67, 0x4e, 0x6f, 0x74, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x75, 0x73, 0x65, 0x61, 0x62, 0x6c, 0x65,
	0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x78, 0x38, 0x36, 0x12, 0x2b, 0x0a, 0x0e, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x74, 0x69,
	0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x42, 0x05, 0xd2, 0x3f,
	0x02, 0x08, 0x01, 0x52, 0x0c, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x69, 0x64, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x39, 0x0a, 0x06, 0x67, 0x70, 0x72, 0x65, 0x67, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x78,
	0x38, 0x36, 0x5f, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x05, 0xd2,
	0x3f, 0x02, 0x08, 0x01, 0x52, 0x06, 0x67, 0x70, 0x72, 0x65, 0x67, 0x73, 0x12, 0x34, 0x0a, 0x06,
	0x66, 0x70, 0x72, 0x65, 0x67, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x78, 0x38, 0x36, 0x5f, 0x66, 0x70,
	0x72, 0x65, 0x67, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x70, 0x72, 0x65,
	0x67, 0x73, 0x12, 0x24, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x5f, 0x74, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x2a, 0x2c, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x78, 0x38, 0x36, 0x5f, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f,
	0x4d, 0x50, 0x41, 0x54, 0x10, 0x02, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x4c, 0x6f, 0x6e, 0x65, 0x6c, 0x79, 0x2f, 0x63, 0x72, 0x69,
	0x75, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
}

var (
	file_core_x86_proto_rawDescOnce sync.Once
	file_core_x86_proto_rawDescData = file_core_x86_proto_rawDesc
)

func file_core_x86_proto_rawDescGZIP() []byte {
	file_core_x86_proto_rawDescOnce.Do(func() {
		file_core_x86_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_x86_proto_rawDescData)
	})
	return file_core_x86_proto_rawDescData
}

var file_core_x86_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_x86_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_core_x86_proto_goTypes = []interface{}{
	(UserX86RegsMode)(0),       // 0: types.user_x86_regs_mode
	(*UserX86RegsEntry)(nil),   // 1: types.user_x86_regs_entry
	(*UserX86XsaveEntry)(nil),  // 2: types.user_x86_xsave_entry
	(*UserX86FpregsEntry)(nil), // 3: types.user_x86_fpregs_entry
	(*UserDescT)(nil),          // 4: types.user_desc_t
	(*ThreadInfoX86)(nil),      // 5: types.thread_info_x86
}
var file_core_x86_proto_depIdxs = []int32{
	0, // 0: types.user_x86_regs_entry.mode:type_name -> types.user_x86_regs_mode
	2, // 1: types.user_x86_fpregs_entry.xsave:type_name -> types.user_x86_xsave_entry
	1, // 2: types.thread_info_x86.gpregs:type_name -> types.user_x86_regs_entry
	3, // 3: types.thread_info_x86.fpregs:type_name -> types.user_x86_fpregs_entry
	4, // 4: types.thread_info_x86.tls:type_name -> types.user_desc_t
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_core_x86_proto_init() }
func file_core_x86_proto_init() {
	if File_core_x86_proto != nil {
		return
	}
	file_opts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_x86_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserX86RegsEntry); i {
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
		file_core_x86_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserX86XsaveEntry); i {
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
		file_core_x86_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserX86FpregsEntry); i {
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
		file_core_x86_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDescT); i {
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
		file_core_x86_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreadInfoX86); i {
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
			RawDescriptor: file_core_x86_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_x86_proto_goTypes,
		DependencyIndexes: file_core_x86_proto_depIdxs,
		EnumInfos:         file_core_x86_proto_enumTypes,
		MessageInfos:      file_core_x86_proto_msgTypes,
	}.Build()
	File_core_x86_proto = out.File
	file_core_x86_proto_rawDesc = nil
	file_core_x86_proto_goTypes = nil
	file_core_x86_proto_depIdxs = nil
}
