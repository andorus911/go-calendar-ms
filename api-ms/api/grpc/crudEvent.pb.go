// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: crudEvent.proto

package gocalendar

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EventResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Owner       int64                `protobuf:"varint,4,opt,name=owner,proto3" json:"owner,omitempty"`
	StartTime   *timestamp.Timestamp `protobuf:"bytes,5,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime     *timestamp.Timestamp `protobuf:"bytes,6,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *EventResponseMessage) Reset() {
	*x = EventResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crudEvent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponseMessage) ProtoMessage() {}

func (x *EventResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_crudEvent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponseMessage.ProtoReflect.Descriptor instead.
func (*EventResponseMessage) Descriptor() ([]byte, []int) {
	return file_crudEvent_proto_rawDescGZIP(), []int{0}
}

func (x *EventResponseMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EventResponseMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EventResponseMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EventResponseMessage) GetOwner() int64 {
	if x != nil {
		return x.Owner
	}
	return 0
}

func (x *EventResponseMessage) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *EventResponseMessage) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type EventsResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*EventResponseMessage `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *EventsResponseMessage) Reset() {
	*x = EventsResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crudEvent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsResponseMessage) ProtoMessage() {}

func (x *EventsResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_crudEvent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsResponseMessage.ProtoReflect.Descriptor instead.
func (*EventsResponseMessage) Descriptor() ([]byte, []int) {
	return file_crudEvent_proto_rawDescGZIP(), []int{1}
}

func (x *EventsResponseMessage) GetEvents() []*EventResponseMessage {
	if x != nil {
		return x.Events
	}
	return nil
}

type CreateEventMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Owner       int64                `protobuf:"varint,3,opt,name=owner,proto3" json:"owner,omitempty"`
	StartTime   *timestamp.Timestamp `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime     *timestamp.Timestamp `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *CreateEventMessage) Reset() {
	*x = CreateEventMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crudEvent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventMessage) ProtoMessage() {}

func (x *CreateEventMessage) ProtoReflect() protoreflect.Message {
	mi := &file_crudEvent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventMessage.ProtoReflect.Descriptor instead.
func (*CreateEventMessage) Descriptor() ([]byte, []int) {
	return file_crudEvent_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEventMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateEventMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEventMessage) GetOwner() int64 {
	if x != nil {
		return x.Owner
	}
	return 0
}

func (x *CreateEventMessage) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *CreateEventMessage) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type ReadEventsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner int64 `protobuf:"varint,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *ReadEventsMessage) Reset() {
	*x = ReadEventsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crudEvent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEventsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEventsMessage) ProtoMessage() {}

func (x *ReadEventsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_crudEvent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEventsMessage.ProtoReflect.Descriptor instead.
func (*ReadEventsMessage) Descriptor() ([]byte, []int) {
	return file_crudEvent_proto_rawDescGZIP(), []int{3}
}

func (x *ReadEventsMessage) GetOwner() int64 {
	if x != nil {
		return x.Owner
	}
	return 0
}

type UpdateEventMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       *string `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// int64 owner = 4;
	StartTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=startTime,proto3,oneof" json:"startTime,omitempty"`
	EndTime   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=endTime,proto3,oneof" json:"endTime,omitempty"`
}

func (x *UpdateEventMessage) Reset() {
	*x = UpdateEventMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crudEvent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventMessage) ProtoMessage() {}

func (x *UpdateEventMessage) ProtoReflect() protoreflect.Message {
	mi := &file_crudEvent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventMessage.ProtoReflect.Descriptor instead.
func (*UpdateEventMessage) Descriptor() ([]byte, []int) {
	return file_crudEvent_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEventMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateEventMessage) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *UpdateEventMessage) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateEventMessage) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *UpdateEventMessage) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type DeleteEventMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEventMessage) Reset() {
	*x = DeleteEventMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crudEvent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventMessage) ProtoMessage() {}

func (x *DeleteEventMessage) ProtoReflect() protoreflect.Message {
	mi := &file_crudEvent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventMessage.ProtoReflect.Descriptor instead.
func (*DeleteEventMessage) Descriptor() ([]byte, []int) {
	return file_crudEvent_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteEventMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_crudEvent_proto protoreflect.FileDescriptor

var file_crudEvent_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x72, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x14,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x51, 0x0a, 0x15, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x52, 0x65,
	0x61, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x94, 0x02, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3d,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x03, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xc9, 0x02, 0x0a, 0x0e, 0x47, 0x6f, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x43, 0x52, 0x55, 0x44, 0x12, 0x4f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x20, 0x2e, 0x67, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e, 0x67, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x20, 0x2e, 0x67, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0e,
	0x5a, 0x0c, 0x2e, 0x3b, 0x67, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crudEvent_proto_rawDescOnce sync.Once
	file_crudEvent_proto_rawDescData = file_crudEvent_proto_rawDesc
)

func file_crudEvent_proto_rawDescGZIP() []byte {
	file_crudEvent_proto_rawDescOnce.Do(func() {
		file_crudEvent_proto_rawDescData = protoimpl.X.CompressGZIP(file_crudEvent_proto_rawDescData)
	})
	return file_crudEvent_proto_rawDescData
}

var file_crudEvent_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_crudEvent_proto_goTypes = []interface{}{
	(*EventResponseMessage)(nil),  // 0: gocalendar.EventResponseMessage
	(*EventsResponseMessage)(nil), // 1: gocalendar.EventsResponseMessage
	(*CreateEventMessage)(nil),    // 2: gocalendar.CreateEventMessage
	(*ReadEventsMessage)(nil),     // 3: gocalendar.ReadEventsMessage
	(*UpdateEventMessage)(nil),    // 4: gocalendar.UpdateEventMessage
	(*DeleteEventMessage)(nil),    // 5: gocalendar.DeleteEventMessage
	(*timestamp.Timestamp)(nil),   // 6: google.protobuf.Timestamp
	(*empty.Empty)(nil),           // 7: google.protobuf.Empty
}
var file_crudEvent_proto_depIdxs = []int32{
	6,  // 0: gocalendar.EventResponseMessage.startTime:type_name -> google.protobuf.Timestamp
	6,  // 1: gocalendar.EventResponseMessage.endTime:type_name -> google.protobuf.Timestamp
	0,  // 2: gocalendar.EventsResponseMessage.events:type_name -> gocalendar.EventResponseMessage
	6,  // 3: gocalendar.CreateEventMessage.startTime:type_name -> google.protobuf.Timestamp
	6,  // 4: gocalendar.CreateEventMessage.endTime:type_name -> google.protobuf.Timestamp
	6,  // 5: gocalendar.UpdateEventMessage.startTime:type_name -> google.protobuf.Timestamp
	6,  // 6: gocalendar.UpdateEventMessage.endTime:type_name -> google.protobuf.Timestamp
	2,  // 7: gocalendar.GoCalendarCRUD.createEvent:input_type -> gocalendar.CreateEventMessage
	3,  // 8: gocalendar.GoCalendarCRUD.readEvents:input_type -> gocalendar.ReadEventsMessage
	4,  // 9: gocalendar.GoCalendarCRUD.updateEvent:input_type -> gocalendar.UpdateEventMessage
	5,  // 10: gocalendar.GoCalendarCRUD.deleteEvent:input_type -> gocalendar.DeleteEventMessage
	0,  // 11: gocalendar.GoCalendarCRUD.createEvent:output_type -> gocalendar.EventResponseMessage
	1,  // 12: gocalendar.GoCalendarCRUD.readEvents:output_type -> gocalendar.EventsResponseMessage
	0,  // 13: gocalendar.GoCalendarCRUD.updateEvent:output_type -> gocalendar.EventResponseMessage
	7,  // 14: gocalendar.GoCalendarCRUD.deleteEvent:output_type -> google.protobuf.Empty
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_crudEvent_proto_init() }
func file_crudEvent_proto_init() {
	if File_crudEvent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crudEvent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponseMessage); i {
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
		file_crudEvent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsResponseMessage); i {
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
		file_crudEvent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventMessage); i {
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
		file_crudEvent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEventsMessage); i {
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
		file_crudEvent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventMessage); i {
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
		file_crudEvent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventMessage); i {
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
	file_crudEvent_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crudEvent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crudEvent_proto_goTypes,
		DependencyIndexes: file_crudEvent_proto_depIdxs,
		MessageInfos:      file_crudEvent_proto_msgTypes,
	}.Build()
	File_crudEvent_proto = out.File
	file_crudEvent_proto_rawDesc = nil
	file_crudEvent_proto_goTypes = nil
	file_crudEvent_proto_depIdxs = nil
}
