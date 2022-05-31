// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: mafia/mafia.proto

package mafia

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url1 string `protobuf:"bytes,1,opt,name=url1,proto3" json:"url1,omitempty"`
	Url2 string `protobuf:"bytes,2,opt,name=url2,proto3" json:"url2,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_mafia_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_mafia_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_mafia_mafia_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetUrl1() string {
	if x != nil {
		return x.Url1
	}
	return ""
}

func (x *Request) GetUrl2() string {
	if x != nil {
		return x.Url2
	}
	return ""
}

type SResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *SResponse) Reset() {
	*x = SResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_mafia_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SResponse) ProtoMessage() {}

func (x *SResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_mafia_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SResponse.ProtoReflect.Descriptor instead.
func (*SResponse) Descriptor() ([]byte, []int) {
	return file_mafia_mafia_proto_rawDescGZIP(), []int{1}
}

func (x *SResponse) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type SRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *SRequest) Reset() {
	*x = SRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_mafia_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SRequest) ProtoMessage() {}

func (x *SRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_mafia_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SRequest.ProtoReflect.Descriptor instead.
func (*SRequest) Descriptor() ([]byte, []int) {
	return file_mafia_mafia_proto_rawDescGZIP(), []int{2}
}

func (x *SRequest) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url1  string `protobuf:"bytes,1,opt,name=url1,proto3" json:"url1,omitempty"`
	Url2  string `protobuf:"bytes,2,opt,name=url2,proto3" json:"url2,omitempty"`
	Count string `protobuf:"bytes,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mafia_mafia_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_mafia_mafia_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_mafia_mafia_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetUrl1() string {
	if x != nil {
		return x.Url1
	}
	return ""
}

func (x *Response) GetUrl2() string {
	if x != nil {
		return x.Url2
	}
	return ""
}

func (x *Response) GetCount() string {
	if x != nil {
		return x.Count
	}
	return ""
}

var File_mafia_mafia_proto protoreflect.FileDescriptor

var file_mafia_mafia_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2f, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x22, 0x31, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x32, 0x22, 0x1f, 0x0a,
	0x09, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x1e,
	0x0a, 0x08, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x48,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72,
	0x6c, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x31, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x72, 0x6c, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72,
	0x6c, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x8c, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x02, 0x44, 0x6f, 0x12, 0x0e, 0x2e, 0x6d, 0x61, 0x66,
	0x69, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x66,
	0x69, 0x61, 0x2e, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b,
	0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x04, 0x45,
	0x78, 0x69, 0x74, 0x12, 0x0e, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x2e, 0x53, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x67, 0x6f, 0x2f, 0x68, 0x77, 0x35, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x66, 0x69,
	0x61, 0x2f, 0x6d, 0x61, 0x66, 0x69, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mafia_mafia_proto_rawDescOnce sync.Once
	file_mafia_mafia_proto_rawDescData = file_mafia_mafia_proto_rawDesc
)

func file_mafia_mafia_proto_rawDescGZIP() []byte {
	file_mafia_mafia_proto_rawDescOnce.Do(func() {
		file_mafia_mafia_proto_rawDescData = protoimpl.X.CompressGZIP(file_mafia_mafia_proto_rawDescData)
	})
	return file_mafia_mafia_proto_rawDescData
}

var file_mafia_mafia_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mafia_mafia_proto_goTypes = []interface{}{
	(*Request)(nil),   // 0: mafia.Request
	(*SResponse)(nil), // 1: mafia.SResponse
	(*SRequest)(nil),  // 2: mafia.SRequest
	(*Response)(nil),  // 3: mafia.Response
}
var file_mafia_mafia_proto_depIdxs = []int32{
	0, // 0: mafia.Reverse.Do:input_type -> mafia.Request
	0, // 1: mafia.Reverse.Answer:input_type -> mafia.Request
	0, // 2: mafia.Reverse.Exit:input_type -> mafia.Request
	1, // 3: mafia.Reverse.Do:output_type -> mafia.SResponse
	3, // 4: mafia.Reverse.Answer:output_type -> mafia.Response
	1, // 5: mafia.Reverse.Exit:output_type -> mafia.SResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mafia_mafia_proto_init() }
func file_mafia_mafia_proto_init() {
	if File_mafia_mafia_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mafia_mafia_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_mafia_mafia_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SResponse); i {
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
		file_mafia_mafia_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SRequest); i {
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
		file_mafia_mafia_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_mafia_mafia_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mafia_mafia_proto_goTypes,
		DependencyIndexes: file_mafia_mafia_proto_depIdxs,
		MessageInfos:      file_mafia_mafia_proto_msgTypes,
	}.Build()
	File_mafia_mafia_proto = out.File
	file_mafia_mafia_proto_rawDesc = nil
	file_mafia_mafia_proto_goTypes = nil
	file_mafia_mafia_proto_depIdxs = nil
}