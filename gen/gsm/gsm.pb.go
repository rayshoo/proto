// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.0
// source: gsm.proto

package gsm

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

	Url        string  `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	SyncMode   string  `protobuf:"bytes,2,opt,name=syncMode,proto3" json:"syncMode,omitempty"`
	SyncPolicy string  `protobuf:"bytes,3,opt,name=syncPolicy,proto3" json:"syncPolicy,omitempty"`
	CommitId   *string `protobuf:"bytes,4,opt,name=commitId,proto3,oneof" json:"commitId,omitempty"`
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

func (x *SyncRequest) GetCommitId() string {
	if x != nil && x.CommitId != nil {
		return *x.CommitId
	}
	return ""
}

type SyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *string `protobuf:"bytes,1,opt,name=message,proto3,oneof" json:"message,omitempty"`
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
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_gsm_proto protoreflect.FileDescriptor

var file_gsm_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x73, 0x6d,
	0x22, 0x89, 0x01, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1f,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x0c,
	0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x47, 0x0a, 0x14, 0x47, 0x6f, 0x50, 0x61, 0x73,
	0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x2f, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x10, 0x2e, 0x67, 0x73, 0x6d, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x73, 0x6d, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x09, 0x5a, 0x07, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x73, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_gsm_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gsm_proto_goTypes = []interface{}{
	(*SyncRequest)(nil),  // 0: gsm.SyncRequest
	(*SyncResponse)(nil), // 1: gsm.SyncResponse
}
var file_gsm_proto_depIdxs = []int32{
	0, // 0: gsm.GoPassSecretsManager.Sync:input_type -> gsm.SyncRequest
	1, // 1: gsm.GoPassSecretsManager.Sync:output_type -> gsm.SyncResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
	}
	file_gsm_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_gsm_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gsm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
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
