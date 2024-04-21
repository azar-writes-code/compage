// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: api/v1/project.proto

package project

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

type GenerateCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompageCoreVersion  string `protobuf:"bytes,1,opt,name=compageCoreVersion,proto3" json:"compageCoreVersion,omitempty"`
	ProjectName         string `protobuf:"bytes,2,opt,name=projectName,proto3" json:"projectName,omitempty"`
	ProjectJSON         string `protobuf:"bytes,3,opt,name=projectJSON,proto3" json:"projectJSON,omitempty"`
	ProjectMetadata     string `protobuf:"bytes,4,opt,name=projectMetadata,proto3" json:"projectMetadata,omitempty"`
	GitRepositoryName   string `protobuf:"bytes,5,opt,name=gitRepositoryName,proto3" json:"gitRepositoryName,omitempty"`
	GitPlatformName     string `protobuf:"bytes,6,opt,name=gitPlatformName,proto3" json:"gitPlatformName,omitempty"`
	GitPlatformURL      string `protobuf:"bytes,7,opt,name=gitPlatformURL,proto3" json:"gitPlatformURL,omitempty"`
	GitPlatformUserName string `protobuf:"bytes,8,opt,name=gitPlatformUserName,proto3" json:"gitPlatformUserName,omitempty"`
}

func (x *GenerateCodeRequest) Reset() {
	*x = GenerateCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCodeRequest) ProtoMessage() {}

func (x *GenerateCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateCodeRequest.ProtoReflect.Descriptor instead.
func (*GenerateCodeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_project_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateCodeRequest) GetCompageCoreVersion() string {
	if x != nil {
		return x.CompageCoreVersion
	}
	return ""
}

func (x *GenerateCodeRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *GenerateCodeRequest) GetProjectJSON() string {
	if x != nil {
		return x.ProjectJSON
	}
	return ""
}

func (x *GenerateCodeRequest) GetProjectMetadata() string {
	if x != nil {
		return x.ProjectMetadata
	}
	return ""
}

func (x *GenerateCodeRequest) GetGitRepositoryName() string {
	if x != nil {
		return x.GitRepositoryName
	}
	return ""
}

func (x *GenerateCodeRequest) GetGitPlatformName() string {
	if x != nil {
		return x.GitPlatformName
	}
	return ""
}

func (x *GenerateCodeRequest) GetGitPlatformURL() string {
	if x != nil {
		return x.GitPlatformURL
	}
	return ""
}

func (x *GenerateCodeRequest) GetGitPlatformUserName() string {
	if x != nil {
		return x.GitPlatformUserName
	}
	return ""
}

type GenerateCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunk []byte `protobuf:"bytes,1,opt,name=fileChunk,proto3" json:"fileChunk,omitempty"`
}

func (x *GenerateCodeResponse) Reset() {
	*x = GenerateCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCodeResponse) ProtoMessage() {}

func (x *GenerateCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateCodeResponse.ProtoReflect.Descriptor instead.
func (*GenerateCodeResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_project_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateCodeResponse) GetFileChunk() []byte {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

var File_api_v1_project_proto protoreflect.FileDescriptor

var file_api_v1_project_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0xe5,
	0x02, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x11, 0x67, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x67, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x67, 0x69, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x69, 0x74,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x67, 0x69, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x55, 0x52, 0x4c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x69, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x55, 0x52, 0x4c, 0x12, 0x30, 0x0a, 0x13, 0x67, 0x69, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x67, 0x69, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0xb0, 0x01, 0x0a,
	0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4f,
	0x0a, 0x0e, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x6c, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_project_proto_rawDescOnce sync.Once
	file_api_v1_project_proto_rawDescData = file_api_v1_project_proto_rawDesc
)

func file_api_v1_project_proto_rawDescGZIP() []byte {
	file_api_v1_project_proto_rawDescOnce.Do(func() {
		file_api_v1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_project_proto_rawDescData)
	})
	return file_api_v1_project_proto_rawDescData
}

var file_api_v1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_project_proto_goTypes = []interface{}{
	(*GenerateCodeRequest)(nil),  // 0: api.v1.GenerateCodeRequest
	(*GenerateCodeResponse)(nil), // 1: api.v1.GenerateCodeResponse
}
var file_api_v1_project_proto_depIdxs = []int32{
	0, // 0: api.v1.ProjectService.GenerateCode:input_type -> api.v1.GenerateCodeRequest
	0, // 1: api.v1.ProjectService.RegenerateCode:input_type -> api.v1.GenerateCodeRequest
	1, // 2: api.v1.ProjectService.GenerateCode:output_type -> api.v1.GenerateCodeResponse
	1, // 3: api.v1.ProjectService.RegenerateCode:output_type -> api.v1.GenerateCodeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_project_proto_init() }
func file_api_v1_project_proto_init() {
	if File_api_v1_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateCodeRequest); i {
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
		file_api_v1_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateCodeResponse); i {
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
			RawDescriptor: file_api_v1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_project_proto_goTypes,
		DependencyIndexes: file_api_v1_project_proto_depIdxs,
		MessageInfos:      file_api_v1_project_proto_msgTypes,
	}.Build()
	File_api_v1_project_proto = out.File
	file_api_v1_project_proto_rawDesc = nil
	file_api_v1_project_proto_goTypes = nil
	file_api_v1_project_proto_depIdxs = nil
}
