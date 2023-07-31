// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: replitlmpb/replitlm.proto

package replitlmpb

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

// The request message.
type CodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prompt string `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
	Lang   string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *CodeRequest) Reset() {
	*x = CodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replitlmpb_replitlm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeRequest) ProtoMessage() {}

func (x *CodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_replitlmpb_replitlm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeRequest.ProtoReflect.Descriptor instead.
func (*CodeRequest) Descriptor() ([]byte, []int) {
	return file_replitlmpb_replitlm_proto_rawDescGZIP(), []int{0}
}

func (x *CodeRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *CodeRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

// The response message.
type CodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret  *Result `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Code int32   `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string  `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CodeResponse) Reset() {
	*x = CodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replitlmpb_replitlm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeResponse) ProtoMessage() {}

func (x *CodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_replitlmpb_replitlm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeResponse.ProtoReflect.Descriptor instead.
func (*CodeResponse) Descriptor() ([]byte, []int) {
	return file_replitlmpb_replitlm_proto_rawDescGZIP(), []int{1}
}

func (x *CodeResponse) GetRet() *Result {
	if x != nil {
		return x.Ret
	}
	return nil
}

func (x *CodeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CodeResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeList           []string `protobuf:"bytes,1,rep,name=code_list,json=codeList,proto3" json:"code_list,omitempty"`
	CompletionTokenNum int32    `protobuf:"varint,2,opt,name=completion_token_num,json=completionTokenNum,proto3" json:"completion_token_num,omitempty"`
	PromptTokenNum     int32    `protobuf:"varint,3,opt,name=prompt_token_num,json=promptTokenNum,proto3" json:"prompt_token_num,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replitlmpb_replitlm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_replitlmpb_replitlm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_replitlmpb_replitlm_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetCodeList() []string {
	if x != nil {
		return x.CodeList
	}
	return nil
}

func (x *Result) GetCompletionTokenNum() int32 {
	if x != nil {
		return x.CompletionTokenNum
	}
	return 0
}

func (x *Result) GetPromptTokenNum() int32 {
	if x != nil {
		return x.PromptTokenNum
	}
	return 0
}

var File_replitlmpb_replitlm_proto protoreflect.FileDescriptor

var file_replitlmpb_replitlm_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x74, 0x6c, 0x6d, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x74, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x74, 0x6c, 0x6d, 0x22, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67,
	0x22, 0x58, 0x0a, 0x0c, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x74, 0x6c, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x03, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x81, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x32, 0x4e,
	0x0a, 0x0d, 0x43, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x3d, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x15, 0x2e,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x74, 0x6c, 0x6d, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x74, 0x6c, 0x6d, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14,
	0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x74,
	0x6c, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_replitlmpb_replitlm_proto_rawDescOnce sync.Once
	file_replitlmpb_replitlm_proto_rawDescData = file_replitlmpb_replitlm_proto_rawDesc
)

func file_replitlmpb_replitlm_proto_rawDescGZIP() []byte {
	file_replitlmpb_replitlm_proto_rawDescOnce.Do(func() {
		file_replitlmpb_replitlm_proto_rawDescData = protoimpl.X.CompressGZIP(file_replitlmpb_replitlm_proto_rawDescData)
	})
	return file_replitlmpb_replitlm_proto_rawDescData
}

var file_replitlmpb_replitlm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_replitlmpb_replitlm_proto_goTypes = []interface{}{
	(*CodeRequest)(nil),  // 0: replitlm.CodeRequest
	(*CodeResponse)(nil), // 1: replitlm.CodeResponse
	(*Result)(nil),       // 2: replitlm.Result
}
var file_replitlmpb_replitlm_proto_depIdxs = []int32{
	2, // 0: replitlm.CodeResponse.ret:type_name -> replitlm.Result
	0, // 1: replitlm.CodeGenerator.SendPrompt:input_type -> replitlm.CodeRequest
	1, // 2: replitlm.CodeGenerator.SendPrompt:output_type -> replitlm.CodeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_replitlmpb_replitlm_proto_init() }
func file_replitlmpb_replitlm_proto_init() {
	if File_replitlmpb_replitlm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_replitlmpb_replitlm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeRequest); i {
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
		file_replitlmpb_replitlm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeResponse); i {
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
		file_replitlmpb_replitlm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_replitlmpb_replitlm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_replitlmpb_replitlm_proto_goTypes,
		DependencyIndexes: file_replitlmpb_replitlm_proto_depIdxs,
		MessageInfos:      file_replitlmpb_replitlm_proto_msgTypes,
	}.Build()
	File_replitlmpb_replitlm_proto = out.File
	file_replitlmpb_replitlm_proto_rawDesc = nil
	file_replitlmpb_replitlm_proto_goTypes = nil
	file_replitlmpb_replitlm_proto_depIdxs = nil
}