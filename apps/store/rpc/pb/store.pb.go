// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.23.1
// source: store.proto

package pb

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

type GoodsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sku           string                 `protobuf:"bytes,1,opt,name=Sku,proto3" json:"Sku,omitempty"`
	Stock         uint64                 `protobuf:"varint,2,opt,name=Stock,proto3" json:"Stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GoodsInfo) Reset() {
	*x = GoodsInfo{}
	mi := &file_store_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfo) ProtoMessage() {}

func (x *GoodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfo.ProtoReflect.Descriptor instead.
func (*GoodsInfo) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsInfo) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *GoodsInfo) GetStock() uint64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type StoreInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	StoreId       uint64                 `protobuf:"varint,2,opt,name=StoreId,proto3" json:"StoreId,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Uid           uint64                 `protobuf:"varint,4,opt,name=Uid,proto3" json:"Uid,omitempty"`
	CreateDate    string                 `protobuf:"bytes,5,opt,name=CreateDate,proto3" json:"CreateDate,omitempty"`
	Status        uint64                 `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	GoodsInfos    []*GoodsInfo           `protobuf:"bytes,7,rep,name=GoodsInfos,proto3" json:"GoodsInfos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StoreInfo) Reset() {
	*x = StoreInfo{}
	mi := &file_store_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreInfo) ProtoMessage() {}

func (x *StoreInfo) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreInfo.ProtoReflect.Descriptor instead.
func (*StoreInfo) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{1}
}

func (x *StoreInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StoreInfo) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *StoreInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoreInfo) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *StoreInfo) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

func (x *StoreInfo) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StoreInfo) GetGoodsInfos() []*GoodsInfo {
	if x != nil {
		return x.GoodsInfos
	}
	return nil
}

type RegistryStoreReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StoreId       uint64                 `protobuf:"varint,1,opt,name=StoreId,proto3" json:"StoreId,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegistryStoreReq) Reset() {
	*x = RegistryStoreReq{}
	mi := &file_store_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistryStoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryStoreReq) ProtoMessage() {}

func (x *RegistryStoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryStoreReq.ProtoReflect.Descriptor instead.
func (*RegistryStoreReq) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{2}
}

func (x *RegistryStoreReq) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *RegistryStoreReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegistryStoreResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Info          *StoreInfo             `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegistryStoreResp) Reset() {
	*x = RegistryStoreResp{}
	mi := &file_store_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistryStoreResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryStoreResp) ProtoMessage() {}

func (x *RegistryStoreResp) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryStoreResp.ProtoReflect.Descriptor instead.
func (*RegistryStoreResp) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{3}
}

func (x *RegistryStoreResp) GetInfo() *StoreInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type StoreInfoReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StoreId       uint64                 `protobuf:"varint,1,opt,name=StoreId,proto3" json:"StoreId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StoreInfoReq) Reset() {
	*x = StoreInfoReq{}
	mi := &file_store_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreInfoReq) ProtoMessage() {}

func (x *StoreInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreInfoReq.ProtoReflect.Descriptor instead.
func (*StoreInfoReq) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{4}
}

func (x *StoreInfoReq) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

type StoreInfoResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Info          *StoreInfo             `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StoreInfoResp) Reset() {
	*x = StoreInfoResp{}
	mi := &file_store_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreInfoResp) ProtoMessage() {}

func (x *StoreInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreInfoResp.ProtoReflect.Descriptor instead.
func (*StoreInfoResp) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{5}
}

func (x *StoreInfoResp) GetInfo() *StoreInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type AddGoodsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StoreId       uint64                 `protobuf:"varint,1,opt,name=StoreId,proto3" json:"StoreId,omitempty"`
	Sku           string                 `protobuf:"bytes,2,opt,name=Sku,proto3" json:"Sku,omitempty"`
	Stock         uint64                 `protobuf:"varint,3,opt,name=Stock,proto3" json:"Stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddGoodsReq) Reset() {
	*x = AddGoodsReq{}
	mi := &file_store_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddGoodsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGoodsReq) ProtoMessage() {}

func (x *AddGoodsReq) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGoodsReq.ProtoReflect.Descriptor instead.
func (*AddGoodsReq) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{6}
}

func (x *AddGoodsReq) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *AddGoodsReq) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *AddGoodsReq) GetStock() uint64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type AddGoodsResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddGoodsResp) Reset() {
	*x = AddGoodsResp{}
	mi := &file_store_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddGoodsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGoodsResp) ProtoMessage() {}

func (x *AddGoodsResp) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGoodsResp.ProtoReflect.Descriptor instead.
func (*AddGoodsResp) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{7}
}

type GoodsStockReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StoreId       uint64                 `protobuf:"varint,1,opt,name=StoreId,proto3" json:"StoreId,omitempty"`
	Sku           string                 `protobuf:"bytes,2,opt,name=Sku,proto3" json:"Sku,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GoodsStockReq) Reset() {
	*x = GoodsStockReq{}
	mi := &file_store_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoodsStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsStockReq) ProtoMessage() {}

func (x *GoodsStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsStockReq.ProtoReflect.Descriptor instead.
func (*GoodsStockReq) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{8}
}

func (x *GoodsStockReq) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *GoodsStockReq) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

type GoodsStockResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Stock         uint64                 `protobuf:"varint,1,opt,name=Stock,proto3" json:"Stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GoodsStockResp) Reset() {
	*x = GoodsStockResp{}
	mi := &file_store_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoodsStockResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsStockResp) ProtoMessage() {}

func (x *GoodsStockResp) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsStockResp.ProtoReflect.Descriptor instead.
func (*GoodsStockResp) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{9}
}

func (x *GoodsStockResp) GetStock() uint64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

var File_store_proto protoreflect.FileDescriptor

var file_store_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x22, 0x33, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x53, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0xc5, 0x01, 0x0a, 0x09, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x30, 0x0a, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x22, 0x40, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x28,
	0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x4f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x6b, 0x75, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x3b, 0x0a, 0x0d, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x53,
	0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x6b, 0x75, 0x22, 0x26, 0x0a,
	0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x32, 0xf9, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12,
	0x42, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33,
	0x0a, 0x08, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x12, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_store_proto_rawDescOnce sync.Once
	file_store_proto_rawDescData = file_store_proto_rawDesc
)

func file_store_proto_rawDescGZIP() []byte {
	file_store_proto_rawDescOnce.Do(func() {
		file_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_proto_rawDescData)
	})
	return file_store_proto_rawDescData
}

var file_store_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_store_proto_goTypes = []any{
	(*GoodsInfo)(nil),         // 0: store.GoodsInfo
	(*StoreInfo)(nil),         // 1: store.StoreInfo
	(*RegistryStoreReq)(nil),  // 2: store.RegistryStoreReq
	(*RegistryStoreResp)(nil), // 3: store.RegistryStoreResp
	(*StoreInfoReq)(nil),      // 4: store.StoreInfoReq
	(*StoreInfoResp)(nil),     // 5: store.StoreInfoResp
	(*AddGoodsReq)(nil),       // 6: store.AddGoodsReq
	(*AddGoodsResp)(nil),      // 7: store.AddGoodsResp
	(*GoodsStockReq)(nil),     // 8: store.GoodsStockReq
	(*GoodsStockResp)(nil),    // 9: store.GoodsStockResp
}
var file_store_proto_depIdxs = []int32{
	0, // 0: store.StoreInfo.GoodsInfos:type_name -> store.GoodsInfo
	1, // 1: store.RegistryStoreResp.Info:type_name -> store.StoreInfo
	1, // 2: store.StoreInfoResp.Info:type_name -> store.StoreInfo
	2, // 3: store.Store.RegistryStore:input_type -> store.RegistryStoreReq
	4, // 4: store.Store.GetStoreInfo:input_type -> store.StoreInfoReq
	6, // 5: store.Store.AddGoods:input_type -> store.AddGoodsReq
	8, // 6: store.Store.GetGoodsStock:input_type -> store.GoodsStockReq
	3, // 7: store.Store.RegistryStore:output_type -> store.RegistryStoreResp
	5, // 8: store.Store.GetStoreInfo:output_type -> store.StoreInfoResp
	7, // 9: store.Store.AddGoods:output_type -> store.AddGoodsResp
	9, // 10: store.Store.GetGoodsStock:output_type -> store.GoodsStockResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_store_proto_init() }
func file_store_proto_init() {
	if File_store_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_store_proto_goTypes,
		DependencyIndexes: file_store_proto_depIdxs,
		MessageInfos:      file_store_proto_msgTypes,
	}.Build()
	File_store_proto = out.File
	file_store_proto_rawDesc = nil
	file_store_proto_goTypes = nil
	file_store_proto_depIdxs = nil
}
