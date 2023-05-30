// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: proto/fetcher.proto

package fetcher

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

type FetchStatus int32

const (
	FetchStatus_DONE        FetchStatus = 0
	FetchStatus_IN_PROGRESS FetchStatus = 1
	FetchStatus_FAILED      FetchStatus = 2
)

// Enum value maps for FetchStatus.
var (
	FetchStatus_name = map[int32]string{
		0: "DONE",
		1: "IN_PROGRESS",
		2: "FAILED",
	}
	FetchStatus_value = map[string]int32{
		"DONE":        0,
		"IN_PROGRESS": 1,
		"FAILED":      2,
	}
)

func (x FetchStatus) Enum() *FetchStatus {
	p := new(FetchStatus)
	*p = x
	return p
}

func (x FetchStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FetchStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_fetcher_proto_enumTypes[0].Descriptor()
}

func (FetchStatus) Type() protoreflect.EnumType {
	return &file_proto_fetcher_proto_enumTypes[0]
}

func (x FetchStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FetchStatus.Descriptor instead.
func (FetchStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_fetcher_proto_rawDescGZIP(), []int{0}
}

type SortField int32

const (
	SortField_NAME  SortField = 0
	SortField_PRICE SortField = 1
)

// Enum value maps for SortField.
var (
	SortField_name = map[int32]string{
		0: "NAME",
		1: "PRICE",
	}
	SortField_value = map[string]int32{
		"NAME":  0,
		"PRICE": 1,
	}
)

func (x SortField) Enum() *SortField {
	p := new(SortField)
	*p = x
	return p
}

func (x SortField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortField) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_fetcher_proto_enumTypes[1].Descriptor()
}

func (SortField) Type() protoreflect.EnumType {
	return &file_proto_fetcher_proto_enumTypes[1]
}

func (x SortField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortField.Descriptor instead.
func (SortField) EnumDescriptor() ([]byte, []int) {
	return file_proto_fetcher_proto_rawDescGZIP(), []int{1}
}

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  FetchStatus `protobuf:"varint,1,opt,name=status,proto3,enum=fetcher.FetchStatus" json:"status,omitempty"`
	Content string      `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_proto_rawDescGZIP(), []int{1}
}

func (x *FetchResponse) GetStatus() FetchStatus {
	if x != nil {
		return x.Status
	}
	return FetchStatus_DONE
}

func (x *FetchResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging  *PagingParams `protobuf:"bytes,1,opt,name=paging,proto3" json:"paging,omitempty"`
	Sorting []*SortParams `protobuf:"bytes,2,rep,name=sorting,proto3" json:"sorting,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetPaging() *PagingParams {
	if x != nil {
		return x.Paging
	}
	return nil
}

func (x *GetRequest) GetSorting() []*SortParams {
	if x != nil {
		return x.Sorting
	}
	return nil
}

type PagingParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *PagingParams) Reset() {
	*x = PagingParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingParams) ProtoMessage() {}

func (x *PagingParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingParams.ProtoReflect.Descriptor instead.
func (*PagingParams) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_proto_rawDescGZIP(), []int{3}
}

func (x *PagingParams) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PagingParams) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type SortParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field SortField `protobuf:"varint,1,opt,name=field,proto3,enum=fetcher.SortField" json:"field,omitempty"`
	Asc   bool      `protobuf:"varint,2,opt,name=asc,proto3" json:"asc,omitempty"`
}

func (x *SortParams) Reset() {
	*x = SortParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortParams) ProtoMessage() {}

func (x *SortParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortParams.ProtoReflect.Descriptor instead.
func (*SortParams) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_proto_rawDescGZIP(), []int{4}
}

func (x *SortParams) GetField() SortField {
	if x != nil {
		return x.Field
	}
	return SortField_NAME
}

func (x *SortParams) GetAsc() bool {
	if x != nil {
		return x.Asc
	}
	return false
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items   []*Item       `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Paging  *PagingParams `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
	Sorting []*SortParams `protobuf:"bytes,3,rep,name=sorting,proto3" json:"sorting,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_proto_rawDescGZIP(), []int{5}
}

func (x *GetResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GetResponse) GetPaging() *PagingParams {
	if x != nil {
		return x.Paging
	}
	return nil
}

func (x *GetResponse) GetSorting() []*SortParams {
	if x != nil {
		return x.Sorting
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fetcher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fetcher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_proto_fetcher_proto_rawDescGZIP(), []int{6}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_proto_fetcher_proto protoreflect.FileDescriptor

var file_proto_fetcher_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x20,
	0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x57, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x6a, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x73, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x3e, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x48, 0x0a, 0x0a, 0x53, 0x6f, 0x72, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x73, 0x63, 0x22,
	0x90, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53,
	0x6f, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x30, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x2a, 0x34, 0x0a, 0x0b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x20, 0x0a, 0x09, 0x53, 0x6f,
	0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x01, 0x32, 0x7f, 0x0a, 0x0e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38,
	0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x13, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fetcher_proto_rawDescOnce sync.Once
	file_proto_fetcher_proto_rawDescData = file_proto_fetcher_proto_rawDesc
)

func file_proto_fetcher_proto_rawDescGZIP() []byte {
	file_proto_fetcher_proto_rawDescOnce.Do(func() {
		file_proto_fetcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fetcher_proto_rawDescData)
	})
	return file_proto_fetcher_proto_rawDescData
}

var file_proto_fetcher_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_fetcher_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_fetcher_proto_goTypes = []interface{}{
	(FetchStatus)(0),      // 0: fetcher.FetchStatus
	(SortField)(0),        // 1: fetcher.SortField
	(*FetchRequest)(nil),  // 2: fetcher.FetchRequest
	(*FetchResponse)(nil), // 3: fetcher.FetchResponse
	(*GetRequest)(nil),    // 4: fetcher.GetRequest
	(*PagingParams)(nil),  // 5: fetcher.PagingParams
	(*SortParams)(nil),    // 6: fetcher.SortParams
	(*GetResponse)(nil),   // 7: fetcher.GetResponse
	(*Item)(nil),          // 8: fetcher.Item
}
var file_proto_fetcher_proto_depIdxs = []int32{
	0, // 0: fetcher.FetchResponse.status:type_name -> fetcher.FetchStatus
	5, // 1: fetcher.GetRequest.paging:type_name -> fetcher.PagingParams
	6, // 2: fetcher.GetRequest.sorting:type_name -> fetcher.SortParams
	1, // 3: fetcher.SortParams.field:type_name -> fetcher.SortField
	8, // 4: fetcher.GetResponse.items:type_name -> fetcher.Item
	5, // 5: fetcher.GetResponse.paging:type_name -> fetcher.PagingParams
	6, // 6: fetcher.GetResponse.sorting:type_name -> fetcher.SortParams
	2, // 7: fetcher.FetcherService.Fetch:input_type -> fetcher.FetchRequest
	4, // 8: fetcher.FetcherService.List:input_type -> fetcher.GetRequest
	3, // 9: fetcher.FetcherService.Fetch:output_type -> fetcher.FetchResponse
	7, // 10: fetcher.FetcherService.List:output_type -> fetcher.GetResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_fetcher_proto_init() }
func file_proto_fetcher_proto_init() {
	if File_proto_fetcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_fetcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_proto_fetcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_proto_fetcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_proto_fetcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingParams); i {
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
		file_proto_fetcher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortParams); i {
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
		file_proto_fetcher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_proto_fetcher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_proto_fetcher_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fetcher_proto_goTypes,
		DependencyIndexes: file_proto_fetcher_proto_depIdxs,
		EnumInfos:         file_proto_fetcher_proto_enumTypes,
		MessageInfos:      file_proto_fetcher_proto_msgTypes,
	}.Build()
	File_proto_fetcher_proto = out.File
	file_proto_fetcher_proto_rawDesc = nil
	file_proto_fetcher_proto_goTypes = nil
	file_proto_fetcher_proto_depIdxs = nil
}