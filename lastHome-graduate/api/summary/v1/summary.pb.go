// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: summary.proto

package summary

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{0}
}

func (x *Summary) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Summary) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateSummaryRequest) Reset() {
	*x = CreateSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSummaryRequest) ProtoMessage() {}

func (x *CreateSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSummaryRequest.ProtoReflect.Descriptor instead.
func (*CreateSummaryRequest) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSummaryRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateSummaryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary *Summary `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *CreateSummaryReply) Reset() {
	*x = CreateSummaryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSummaryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSummaryReply) ProtoMessage() {}

func (x *CreateSummaryReply) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSummaryReply.ProtoReflect.Descriptor instead.
func (*CreateSummaryReply) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSummaryReply) GetSummary() *Summary {
	if x != nil {
		return x.Summary
	}
	return nil
}

type UpdateSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateSummaryRequest) Reset() {
	*x = UpdateSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSummaryRequest) ProtoMessage() {}

func (x *UpdateSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSummaryRequest.ProtoReflect.Descriptor instead.
func (*UpdateSummaryRequest) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSummaryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateSummaryRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateSummaryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateSummaryReply) Reset() {
	*x = UpdateSummaryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSummaryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSummaryReply) ProtoMessage() {}

func (x *UpdateSummaryReply) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSummaryReply.ProtoReflect.Descriptor instead.
func (*UpdateSummaryReply) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSummaryReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSummaryRequest) Reset() {
	*x = DeleteSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSummaryRequest) ProtoMessage() {}

func (x *DeleteSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSummaryRequest.ProtoReflect.Descriptor instead.
func (*DeleteSummaryRequest) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSummaryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteSummaryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSummaryReply) Reset() {
	*x = DeleteSummaryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSummaryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSummaryReply) ProtoMessage() {}

func (x *DeleteSummaryReply) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSummaryReply.ProtoReflect.Descriptor instead.
func (*DeleteSummaryReply) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSummaryReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSummaryRequest) Reset() {
	*x = GetSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummaryRequest) ProtoMessage() {}

func (x *GetSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummaryRequest.ProtoReflect.Descriptor instead.
func (*GetSummaryRequest) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{7}
}

func (x *GetSummaryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSummaryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary *Summary `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *GetSummaryReply) Reset() {
	*x = GetSummaryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummaryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummaryReply) ProtoMessage() {}

func (x *GetSummaryReply) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummaryReply.ProtoReflect.Descriptor instead.
func (*GetSummaryReply) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{8}
}

func (x *GetSummaryReply) GetSummary() *Summary {
	if x != nil {
		return x.Summary
	}
	return nil
}

type ListSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartId int64 `protobuf:"varint,1,opt,name=start_id,json=startId,proto3" json:"start_id,omitempty"`
	Limit   int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListSummaryRequest) Reset() {
	*x = ListSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSummaryRequest) ProtoMessage() {}

func (x *ListSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSummaryRequest.ProtoReflect.Descriptor instead.
func (*ListSummaryRequest) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{9}
}

func (x *ListSummaryRequest) GetStartId() int64 {
	if x != nil {
		return x.StartId
	}
	return 0
}

func (x *ListSummaryRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListSummaryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summaries []*Summary `protobuf:"bytes,1,rep,name=summaries,proto3" json:"summaries,omitempty"`
}

func (x *ListSummaryReply) Reset() {
	*x = ListSummaryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summary_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSummaryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSummaryReply) ProtoMessage() {}

func (x *ListSummaryReply) ProtoReflect() protoreflect.Message {
	mi := &file_summary_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSummaryReply.ProtoReflect.Descriptor instead.
func (*ListSummaryReply) Descriptor() ([]byte, []int) {
	return file_summary_proto_rawDescGZIP(), []int{10}
}

func (x *ListSummaryReply) GetSummaries() []*Summary {
	if x != nil {
		return x.Summaries
	}
	return nil
}

var File_summary_proto protoreflect.FileDescriptor

var file_summary_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x40, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x3c, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x45, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x42, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x09, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x32, 0x9d, 0x04, 0x0a, 0x0e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x3a, 0x01,
	0x2a, 0x12, 0x68, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x1d, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x1a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x5c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x1a, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x76, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x1b, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12,
	0x27, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x3f, 0x73, 0x3d, 0x7b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x26, 0x6c,
	0x3d, 0x7b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x7d, 0x42, 0x1d, 0x5a, 0x1b, 0x62, 0x6c, 0x6f, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_summary_proto_rawDescOnce sync.Once
	file_summary_proto_rawDescData = file_summary_proto_rawDesc
)

func file_summary_proto_rawDescGZIP() []byte {
	file_summary_proto_rawDescOnce.Do(func() {
		file_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_summary_proto_rawDescData)
	})
	return file_summary_proto_rawDescData
}

var file_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_summary_proto_goTypes = []interface{}{
	(*Summary)(nil),              // 0: summary.Summary
	(*CreateSummaryRequest)(nil), // 1: summary.CreateSummaryRequest
	(*CreateSummaryReply)(nil),   // 2: summary.CreateSummaryReply
	(*UpdateSummaryRequest)(nil), // 3: summary.UpdateSummaryRequest
	(*UpdateSummaryReply)(nil),   // 4: summary.UpdateSummaryReply
	(*DeleteSummaryRequest)(nil), // 5: summary.DeleteSummaryRequest
	(*DeleteSummaryReply)(nil),   // 6: summary.DeleteSummaryReply
	(*GetSummaryRequest)(nil),    // 7: summary.GetSummaryRequest
	(*GetSummaryReply)(nil),      // 8: summary.GetSummaryReply
	(*ListSummaryRequest)(nil),   // 9: summary.ListSummaryRequest
	(*ListSummaryReply)(nil),     // 10: summary.ListSummaryReply
}
var file_summary_proto_depIdxs = []int32{
	0,  // 0: summary.CreateSummaryReply.summary:type_name -> summary.Summary
	0,  // 1: summary.GetSummaryReply.summary:type_name -> summary.Summary
	0,  // 2: summary.ListSummaryReply.summaries:type_name -> summary.Summary
	1,  // 3: summary.SummaryService.CreateSummary:input_type -> summary.CreateSummaryRequest
	3,  // 4: summary.SummaryService.UpdateSummary:input_type -> summary.UpdateSummaryRequest
	5,  // 5: summary.SummaryService.DeleteSummary:input_type -> summary.DeleteSummaryRequest
	7,  // 6: summary.SummaryService.GetSummary:input_type -> summary.GetSummaryRequest
	9,  // 7: summary.SummaryService.ListSummary:input_type -> summary.ListSummaryRequest
	2,  // 8: summary.SummaryService.CreateSummary:output_type -> summary.CreateSummaryReply
	4,  // 9: summary.SummaryService.UpdateSummary:output_type -> summary.UpdateSummaryReply
	6,  // 10: summary.SummaryService.DeleteSummary:output_type -> summary.DeleteSummaryReply
	8,  // 11: summary.SummaryService.GetSummary:output_type -> summary.GetSummaryReply
	10, // 12: summary.SummaryService.ListSummary:output_type -> summary.ListSummaryReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_summary_proto_init() }
func file_summary_proto_init() {
	if File_summary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_summary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary); i {
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
		file_summary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSummaryRequest); i {
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
		file_summary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSummaryReply); i {
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
		file_summary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSummaryRequest); i {
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
		file_summary_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSummaryReply); i {
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
		file_summary_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSummaryRequest); i {
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
		file_summary_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSummaryReply); i {
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
		file_summary_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummaryRequest); i {
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
		file_summary_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummaryReply); i {
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
		file_summary_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSummaryRequest); i {
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
		file_summary_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSummaryReply); i {
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
			RawDescriptor: file_summary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_summary_proto_goTypes,
		DependencyIndexes: file_summary_proto_depIdxs,
		MessageInfos:      file_summary_proto_msgTypes,
	}.Build()
	File_summary_proto = out.File
	file_summary_proto_rawDesc = nil
	file_summary_proto_goTypes = nil
	file_summary_proto_depIdxs = nil
}