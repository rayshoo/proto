// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.0
// source: gsm.proto

package gsm

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url        string                  `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	SyncMode   string                  `protobuf:"bytes,2,opt,name=syncMode,proto3" json:"syncMode,omitempty"`
	SyncPolicy string                  `protobuf:"bytes,3,opt,name=syncPolicy,proto3" json:"syncPolicy,omitempty"`
	DryRun     bool                    `protobuf:"varint,4,opt,name=dryRun,proto3" json:"dryRun,omitempty"`
	CommitId   *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=commitId,proto3" json:"commitId,omitempty"`
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gsm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_gsm_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SyncRequest) GetSyncMode() string {
	if x != nil {
		return x.SyncMode
	}
	return ""
}

func (x *SyncRequest) GetSyncPolicy() string {
	if x != nil {
		return x.SyncPolicy
	}
	return ""
}

func (x *SyncRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *SyncRequest) GetCommitId() *wrapperspb.StringValue {
	if x != nil {
		return x.CommitId
	}
	return nil
}

type SyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message       string                      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	DebugResponse *SyncResponse_DebugResponse `protobuf:"bytes,2,opt,name=debugResponse,proto3" json:"debugResponse,omitempty"`
}

func (x *SyncResponse) Reset() {
	*x = SyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResponse) ProtoMessage() {}

func (x *SyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gsm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResponse.ProtoReflect.Descriptor instead.
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return file_gsm_proto_rawDescGZIP(), []int{1}
}

func (x *SyncResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SyncResponse) GetDebugResponse() *SyncResponse_DebugResponse {
	if x != nil {
		return x.DebugResponse
	}
	return nil
}

type PostHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string                    `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Secrets   map[string]*emptypb.Empty `protobuf:"bytes,2,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DryRun    bool                      `protobuf:"varint,3,opt,name=dryRun,proto3" json:"dryRun,omitempty"`
}

func (x *PostHookRequest) Reset() {
	*x = PostHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostHookRequest) ProtoMessage() {}

func (x *PostHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gsm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostHookRequest.ProtoReflect.Descriptor instead.
func (*PostHookRequest) Descriptor() ([]byte, []int) {
	return file_gsm_proto_rawDescGZIP(), []int{2}
}

func (x *PostHookRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *PostHookRequest) GetSecrets() map[string]*emptypb.Empty {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *PostHookRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

type PostHookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PostHookResponse) Reset() {
	*x = PostHookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostHookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostHookResponse) ProtoMessage() {}

func (x *PostHookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gsm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostHookResponse.ProtoReflect.Descriptor instead.
func (*PostHookResponse) Descriptor() ([]byte, []int) {
	return file_gsm_proto_rawDescGZIP(), []int{3}
}

func (x *PostHookResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SyncResponse_DebugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SyncResponse_DebugResponse) Reset() {
	*x = SyncResponse_DebugResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResponse_DebugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResponse_DebugResponse) ProtoMessage() {}

func (x *SyncResponse_DebugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gsm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResponse_DebugResponse.ProtoReflect.Descriptor instead.
func (*SyncResponse_DebugResponse) Descriptor() ([]byte, []int) {
	return file_gsm_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SyncResponse_DebugResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *SyncResponse_DebugResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_gsm_proto protoreflect.FileDescriptor

var file_gsm_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x73, 0x6d,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0b,
	0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x79, 0x6e, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x79, 0x6e, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x6e,
	0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x79, 0x6e, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x79,
	0x52, 0x75, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75,
	0x6e, 0x12, 0x38, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x0c,
	0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x64, 0x65, 0x62, 0x75, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x67, 0x73, 0x6d, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x47, 0x0a,
	0x0d, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x48,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x73, 0x6d, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x1a, 0x52, 0x0a,
	0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x2c, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x47, 0x0a, 0x14, 0x47, 0x6f, 0x50, 0x61, 0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12,
	0x10, 0x2e, 0x67, 0x73, 0x6d, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x67, 0x73, 0x6d, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x32, 0x6f, 0x0a, 0x18, 0x47, 0x6f, 0x50, 0x61,
	0x73, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x53, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x6f, 0x6f, 0x6b,
	0x12, 0x14, 0x2e, 0x67, 0x73, 0x6d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x73, 0x6d, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x48, 0x6f, 0x6f, 0x6b, 0x3a, 0x01, 0x2a, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x73, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gsm_proto_rawDescOnce sync.Once
	file_gsm_proto_rawDescData = file_gsm_proto_rawDesc
)

func file_gsm_proto_rawDescGZIP() []byte {
	file_gsm_proto_rawDescOnce.Do(func() {
		file_gsm_proto_rawDescData = protoimpl.X.CompressGZIP(file_gsm_proto_rawDescData)
	})
	return file_gsm_proto_rawDescData
}

var file_gsm_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gsm_proto_goTypes = []interface{}{
	(*SyncRequest)(nil),                // 0: gsm.SyncRequest
	(*SyncResponse)(nil),               // 1: gsm.SyncResponse
	(*PostHookRequest)(nil),            // 2: gsm.PostHookRequest
	(*PostHookResponse)(nil),           // 3: gsm.PostHookResponse
	(*SyncResponse_DebugResponse)(nil), // 4: gsm.SyncResponse.DebugResponse
	nil,                                // 5: gsm.PostHookRequest.SecretsEntry
	(*wrapperspb.StringValue)(nil),     // 6: google.protobuf.StringValue
	(*emptypb.Empty)(nil),              // 7: google.protobuf.Empty
}
var file_gsm_proto_depIdxs = []int32{
	6, // 0: gsm.SyncRequest.commitId:type_name -> google.protobuf.StringValue
	4, // 1: gsm.SyncResponse.debugResponse:type_name -> gsm.SyncResponse.DebugResponse
	5, // 2: gsm.PostHookRequest.secrets:type_name -> gsm.PostHookRequest.SecretsEntry
	7, // 3: gsm.PostHookRequest.SecretsEntry.value:type_name -> google.protobuf.Empty
	0, // 4: gsm.GoPassSecretsManager.Sync:input_type -> gsm.SyncRequest
	2, // 5: gsm.GoPassSecretsManagerHook.PostHook:input_type -> gsm.PostHookRequest
	1, // 6: gsm.GoPassSecretsManager.Sync:output_type -> gsm.SyncResponse
	3, // 7: gsm.GoPassSecretsManagerHook.PostHook:output_type -> gsm.PostHookResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gsm_proto_init() }
func file_gsm_proto_init() {
	if File_gsm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gsm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_gsm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResponse); i {
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
		file_gsm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostHookRequest); i {
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
		file_gsm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostHookResponse); i {
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
		file_gsm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResponse_DebugResponse); i {
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
			RawDescriptor: file_gsm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_gsm_proto_goTypes,
		DependencyIndexes: file_gsm_proto_depIdxs,
		MessageInfos:      file_gsm_proto_msgTypes,
	}.Build()
	File_gsm_proto = out.File
	file_gsm_proto_rawDesc = nil
	file_gsm_proto_goTypes = nil
	file_gsm_proto_depIdxs = nil
}
