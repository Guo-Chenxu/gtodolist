// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: desc/task.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 定义消息类型
type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     int64  `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	Status  int32  `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CreateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Data    string `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *CreateResp) Reset() {
	*x = CreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResp) ProtoMessage() {}

func (x *CreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResp.ProtoReflect.Descriptor instead.
func (*CreateResp) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *CreateResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	View      int64  `protobuf:"varint,4,opt,name=View,proto3" json:"View,omitempty"`
	Status    int32  `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	CreateAt  int64  `protobuf:"varint,6,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	StartTime int64  `protobuf:"varint,7,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime   int64  `protobuf:"varint,8,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Task) GetView() int64 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *Task) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Task) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Task) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Task) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task  []*Task `protobuf:"bytes,1,rep,name=Task,proto3" json:"Task,omitempty"`
	Total int64   `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetTask() []*Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *Data) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Start int64 `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{4}
}

func (x *ListReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ListReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListReq) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

type ListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Data    *Data  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *ListResp) Reset() {
	*x = ListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResp) ProtoMessage() {}

func (x *ListResp) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResp.ProtoReflect.Descriptor instead.
func (*ListResp) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{5}
}

func (x *ListResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ListResp) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ShowReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Uid int64 `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (x *ShowReq) Reset() {
	*x = ShowReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowReq) ProtoMessage() {}

func (x *ShowReq) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowReq.ProtoReflect.Descriptor instead.
func (*ShowReq) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{6}
}

func (x *ShowReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShowReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ShowResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Task    *Task  `protobuf:"bytes,2,opt,name=Task,proto3" json:"Task,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *ShowResp) Reset() {
	*x = ShowResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowResp) ProtoMessage() {}

func (x *ShowResp) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowResp.ProtoReflect.Descriptor instead.
func (*ShowResp) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{7}
}

func (x *ShowResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ShowResp) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *ShowResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ShowResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Uid     int64  `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Content string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	Status  int32  `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UpdateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Data    string `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *UpdateResp) Reset() {
	*x = UpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResp) ProtoMessage() {}

func (x *UpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResp.ProtoReflect.Descriptor instead.
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *UpdateResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  int64  `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Info string `protobuf:"bytes,2,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{10}
}

func (x *SearchReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SearchReq) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type SearchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Data    *Data  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *SearchResp) Reset() {
	*x = SearchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResp) ProtoMessage() {}

func (x *SearchResp) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResp.ProtoReflect.Descriptor instead.
func (*SearchResp) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{11}
}

func (x *SearchResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SearchResp) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SearchResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SearchResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64  `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DeleteReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Data    string `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	Error   string `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *DeleteResp) Reset() {
	*x = DeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desc_task_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResp) ProtoMessage() {}

func (x *DeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_desc_task_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResp.ProtoReflect.Descriptor instead.
func (*DeleteResp) Descriptor() ([]byte, []int) {
	return file_desc_task_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *DeleteResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_desc_task_proto protoreflect.FileDescriptor

var file_desc_task_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x65, 0x73, 0x63, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x65, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x68, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0xc6, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69,
	0x65, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x47, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x55, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x22,
	0x70, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x2b, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x22, 0x70,
	0x0a, 0x08, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x75, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x68, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x31, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x72, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2d, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0x8b, 0x02, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x72, 0x70, 0x63, 0x12, 0x2b, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x25, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_desc_task_proto_rawDescOnce sync.Once
	file_desc_task_proto_rawDescData = file_desc_task_proto_rawDesc
)

func file_desc_task_proto_rawDescGZIP() []byte {
	file_desc_task_proto_rawDescOnce.Do(func() {
		file_desc_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_desc_task_proto_rawDescData)
	})
	return file_desc_task_proto_rawDescData
}

var file_desc_task_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_desc_task_proto_goTypes = []interface{}{
	(*CreateReq)(nil),  // 0: pb.CreateReq
	(*CreateResp)(nil), // 1: pb.CreateResp
	(*Task)(nil),       // 2: pb.Task
	(*Data)(nil),       // 3: pb.Data
	(*ListReq)(nil),    // 4: pb.ListReq
	(*ListResp)(nil),   // 5: pb.ListResp
	(*ShowReq)(nil),    // 6: pb.ShowReq
	(*ShowResp)(nil),   // 7: pb.ShowResp
	(*UpdateReq)(nil),  // 8: pb.UpdateReq
	(*UpdateResp)(nil), // 9: pb.UpdateResp
	(*SearchReq)(nil),  // 10: pb.SearchReq
	(*SearchResp)(nil), // 11: pb.SearchResp
	(*DeleteReq)(nil),  // 12: pb.DeleteReq
	(*DeleteResp)(nil), // 13: pb.DeleteResp
}
var file_desc_task_proto_depIdxs = []int32{
	2,  // 0: pb.Data.Task:type_name -> pb.Task
	3,  // 1: pb.ListResp.Data:type_name -> pb.Data
	2,  // 2: pb.ShowResp.Task:type_name -> pb.Task
	3,  // 3: pb.SearchResp.Data:type_name -> pb.Data
	0,  // 4: pb.taskrpc.CreateTask:input_type -> pb.CreateReq
	4,  // 5: pb.taskrpc.ListTask:input_type -> pb.ListReq
	6,  // 6: pb.taskrpc.ShowTask:input_type -> pb.ShowReq
	8,  // 7: pb.taskrpc.UpdateTask:input_type -> pb.UpdateReq
	10, // 8: pb.taskrpc.SearchTask:input_type -> pb.SearchReq
	12, // 9: pb.taskrpc.DeleteTask:input_type -> pb.DeleteReq
	1,  // 10: pb.taskrpc.CreateTask:output_type -> pb.CreateResp
	5,  // 11: pb.taskrpc.ListTask:output_type -> pb.ListResp
	7,  // 12: pb.taskrpc.ShowTask:output_type -> pb.ShowResp
	9,  // 13: pb.taskrpc.UpdateTask:output_type -> pb.UpdateResp
	11, // 14: pb.taskrpc.SearchTask:output_type -> pb.SearchResp
	13, // 15: pb.taskrpc.DeleteTask:output_type -> pb.DeleteResp
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_desc_task_proto_init() }
func file_desc_task_proto_init() {
	if File_desc_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_desc_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_desc_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResp); i {
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
		file_desc_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_desc_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_desc_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_desc_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResp); i {
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
		file_desc_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowReq); i {
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
		file_desc_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowResp); i {
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
		file_desc_task_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_desc_task_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResp); i {
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
		file_desc_task_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
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
		file_desc_task_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResp); i {
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
		file_desc_task_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_desc_task_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResp); i {
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
			RawDescriptor: file_desc_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_desc_task_proto_goTypes,
		DependencyIndexes: file_desc_task_proto_depIdxs,
		MessageInfos:      file_desc_task_proto_msgTypes,
	}.Build()
	File_desc_task_proto = out.File
	file_desc_task_proto_rawDesc = nil
	file_desc_task_proto_goTypes = nil
	file_desc_task_proto_depIdxs = nil
}
