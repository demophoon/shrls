// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: shrls.proto

package gen

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

type ShortURL_ShortURLType int32

const (
	ShortURL_LINK    ShortURL_ShortURLType = 0
	ShortURL_UPLOAD  ShortURL_ShortURLType = 1
	ShortURL_SNIPPET ShortURL_ShortURLType = 2
)

// Enum value maps for ShortURL_ShortURLType.
var (
	ShortURL_ShortURLType_name = map[int32]string{
		0: "LINK",
		1: "UPLOAD",
		2: "SNIPPET",
	}
	ShortURL_ShortURLType_value = map[string]int32{
		"LINK":    0,
		"UPLOAD":  1,
		"SNIPPET": 2,
	}
)

func (x ShortURL_ShortURLType) Enum() *ShortURL_ShortURLType {
	p := new(ShortURL_ShortURLType)
	*p = x
	return p
}

func (x ShortURL_ShortURLType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShortURL_ShortURLType) Descriptor() protoreflect.EnumDescriptor {
	return file_shrls_proto_enumTypes[0].Descriptor()
}

func (ShortURL_ShortURLType) Type() protoreflect.EnumType {
	return &file_shrls_proto_enumTypes[0]
}

func (x ShortURL_ShortURLType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShortURL_ShortURLType.Descriptor instead.
func (ShortURL_ShortURLType) EnumDescriptor() ([]byte, []int) {
	return file_shrls_proto_rawDescGZIP(), []int{3, 0}
}

// ========================================
// Server Methods
// ========================================
type GetShrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shrl *Ref_ShortURL `protobuf:"bytes,1,opt,name=shrl,proto3" json:"shrl,omitempty"`
}

func (x *GetShrlRequest) Reset() {
	*x = GetShrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shrls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShrlRequest) ProtoMessage() {}

func (x *GetShrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shrls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShrlRequest.ProtoReflect.Descriptor instead.
func (*GetShrlRequest) Descriptor() ([]byte, []int) {
	return file_shrls_proto_rawDescGZIP(), []int{0}
}

func (x *GetShrlRequest) GetShrl() *Ref_ShortURL {
	if x != nil {
		return x.Shrl
	}
	return nil
}

type GetShrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shrl *ShortURL `protobuf:"bytes,1,opt,name=shrl,proto3" json:"shrl,omitempty"`
}

func (x *GetShrlResponse) Reset() {
	*x = GetShrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shrls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShrlResponse) ProtoMessage() {}

func (x *GetShrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shrls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShrlResponse.ProtoReflect.Descriptor instead.
func (*GetShrlResponse) Descriptor() ([]byte, []int) {
	return file_shrls_proto_rawDescGZIP(), []int{1}
}

func (x *GetShrlResponse) GetShrl() *ShortURL {
	if x != nil {
		return x.Shrl
	}
	return nil
}

// ========================================
// References
// ========================================
type Ref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ref) Reset() {
	*x = Ref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shrls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ref) ProtoMessage() {}

func (x *Ref) ProtoReflect() protoreflect.Message {
	mi := &file_shrls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ref.ProtoReflect.Descriptor instead.
func (*Ref) Descriptor() ([]byte, []int) {
	return file_shrls_proto_rawDescGZIP(), []int{2}
}

// ========================================
// Objects
// ========================================
// Shortened Urls
type ShortURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                       // External ID for ShortURL
	Type    ShortURL_ShortURLType `protobuf:"varint,2,opt,name=type,proto3,enum=shrls.ShortURL_ShortURLType" json:"type,omitempty"` // ShortURL type
	Stub    string                `protobuf:"bytes,3,opt,name=stub,proto3" json:"stub,omitempty"`                                   // URL stub
	Content *ExpandedURL          `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`                             // Expanded destination
}

func (x *ShortURL) Reset() {
	*x = ShortURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shrls_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortURL) ProtoMessage() {}

func (x *ShortURL) ProtoReflect() protoreflect.Message {
	mi := &file_shrls_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortURL.ProtoReflect.Descriptor instead.
func (*ShortURL) Descriptor() ([]byte, []int) {
	return file_shrls_proto_rawDescGZIP(), []int{3}
}

func (x *ShortURL) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShortURL) GetType() ShortURL_ShortURLType {
	if x != nil {
		return x.Type
	}
	return ShortURL_LINK
}

func (x *ShortURL) GetStub() string {
	if x != nil {
		return x.Stub
	}
	return ""
}

func (x *ShortURL) GetContent() *ExpandedURL {
	if x != nil {
		return x.Content
	}
	return nil
}

// Text snippets
type Snippet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Snippet) Reset() {
	*x = Snippet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shrls_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snippet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snippet) ProtoMessage() {}

func (x *Snippet) ProtoReflect() protoreflect.Message {
	mi := &file_shrls_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snippet.ProtoReflect.Descriptor instead.
func (*Snippet) Descriptor() ([]byte, []int) {
	return file_shrls_proto_rawDescGZIP(), []int{4}
}

func (x *Snippet) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Snippet) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// Url Redirects
type Redirect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Favicon []byte `protobuf:"bytes,2,opt,name=favicon,proto3" json:"favicon,omitempty"`
}

func (x *Redirect) Reset() {
	*x = Redirect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shrls_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redirect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redirect) ProtoMessage() {}

func (x *Redirect) ProtoReflect() protoreflect.Message {
	mi := &file_shrls_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redirect.ProtoReflect.Descriptor instead.
func (*Redirect) Descriptor() ([]byte, []int) {
	return file_shrls_proto_rawDescGZIP(), []int{5}
}

func (x *Redirect) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Redirect) GetFavicon() []byte {
	if x != nil {
		return x.Favicon
	}
	return nil
}

type ExpandedURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*ExpandedURL_Url
	//	*ExpandedURL_File
	//	*ExpandedURL_Snippet
	Content isExpandedURL_Content `protobuf_oneof:"content"`
}

func (x *ExpandedURL) Reset() {
	*x = ExpandedURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shrls_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandedURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandedURL) ProtoMessage() {}

func (x *ExpandedURL) ProtoReflect() protoreflect.Message {
	mi := &file_shrls_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandedURL.ProtoReflect.Descriptor instead.
func (*ExpandedURL) Descriptor() ([]byte, []int) {
	return file_shrls_proto_rawDescGZIP(), []int{6}
}

func (m *ExpandedURL) GetContent() isExpandedURL_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *ExpandedURL) GetUrl() *Redirect {
	if x, ok := x.GetContent().(*ExpandedURL_Url); ok {
		return x.Url
	}
	return nil
}

func (x *ExpandedURL) GetFile() []byte {
	if x, ok := x.GetContent().(*ExpandedURL_File); ok {
		return x.File
	}
	return nil
}

func (x *ExpandedURL) GetSnippet() *Snippet {
	if x, ok := x.GetContent().(*ExpandedURL_Snippet); ok {
		return x.Snippet
	}
	return nil
}

type isExpandedURL_Content interface {
	isExpandedURL_Content()
}

type ExpandedURL_Url struct {
	Url *Redirect `protobuf:"bytes,1,opt,name=url,proto3,oneof"`
}

type ExpandedURL_File struct {
	File []byte `protobuf:"bytes,2,opt,name=file,proto3,oneof"`
}

type ExpandedURL_Snippet struct {
	Snippet *Snippet `protobuf:"bytes,3,opt,name=snippet,proto3,oneof"`
}

func (*ExpandedURL_Url) isExpandedURL_Content() {}

func (*ExpandedURL_File) isExpandedURL_Content() {}

func (*ExpandedURL_Snippet) isExpandedURL_Content() {}

// Reference ShortURL
type Ref_ShortURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Ref:
	//
	//	*Ref_ShortURL_Id
	//	*Ref_ShortURL_Alias
	Ref isRef_ShortURL_Ref `protobuf_oneof:"ref"`
}

func (x *Ref_ShortURL) Reset() {
	*x = Ref_ShortURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shrls_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ref_ShortURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ref_ShortURL) ProtoMessage() {}

func (x *Ref_ShortURL) ProtoReflect() protoreflect.Message {
	mi := &file_shrls_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ref_ShortURL.ProtoReflect.Descriptor instead.
func (*Ref_ShortURL) Descriptor() ([]byte, []int) {
	return file_shrls_proto_rawDescGZIP(), []int{2, 0}
}

func (m *Ref_ShortURL) GetRef() isRef_ShortURL_Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (x *Ref_ShortURL) GetId() string {
	if x, ok := x.GetRef().(*Ref_ShortURL_Id); ok {
		return x.Id
	}
	return ""
}

func (x *Ref_ShortURL) GetAlias() string {
	if x, ok := x.GetRef().(*Ref_ShortURL_Alias); ok {
		return x.Alias
	}
	return ""
}

type isRef_ShortURL_Ref interface {
	isRef_ShortURL_Ref()
}

type Ref_ShortURL_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type Ref_ShortURL_Alias struct {
	Alias string `protobuf:"bytes,2,opt,name=alias,proto3,oneof"`
}

func (*Ref_ShortURL_Id) isRef_ShortURL_Ref() {}

func (*Ref_ShortURL_Alias) isRef_ShortURL_Ref() {}

var File_shrls_proto protoreflect.FileDescriptor

var file_shrls_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x72, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x68, 0x72, 0x6c, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x72, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x73, 0x68, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x68, 0x72, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x2e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x04, 0x73, 0x68, 0x72, 0x6c, 0x22, 0x36, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x04, 0x73, 0x68, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x68, 0x72, 0x6c, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52,
	0x04, 0x73, 0x68, 0x72, 0x6c, 0x22, 0x42, 0x0a, 0x03, 0x52, 0x65, 0x66, 0x1a, 0x3b, 0x0a, 0x08,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x22, 0xc1, 0x01, 0x0a, 0x08, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x68, 0x72, 0x6c, 0x73, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x75, 0x62,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74, 0x75, 0x62, 0x12, 0x2c, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x68, 0x72, 0x6c, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x55, 0x52,
	0x4c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x0c, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49,
	0x4e, 0x4b, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4e, 0x49, 0x50, 0x50, 0x45, 0x54, 0x10, 0x02, 0x22, 0x33, 0x0a,
	0x07, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x0b, 0x45, 0x78,
	0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x12, 0x23, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x68, 0x72, 0x6c, 0x73, 0x2e, 0x52,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x68, 0x72, 0x6c, 0x73, 0x2e, 0x53, 0x6e,
	0x69, 0x70, 0x70, 0x65, 0x74, 0x48, 0x00, 0x52, 0x07, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74,
	0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x53, 0x0a, 0x05, 0x53,
	0x68, 0x72, 0x6c, 0x73, 0x12, 0x4a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x68, 0x72, 0x6c, 0x12,
	0x15, 0x2e, 0x73, 0x68, 0x72, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x72, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x68, 0x72, 0x6c, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x72, 0x6c,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x61, 0x73, 0x63, 0x61,
	0x64, 0x69, 0x61, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x68, 0x6f, 0x6f, 0x6e, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x68, 0x6f, 0x6f, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x68, 0x72, 0x6c, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shrls_proto_rawDescOnce sync.Once
	file_shrls_proto_rawDescData = file_shrls_proto_rawDesc
)

func file_shrls_proto_rawDescGZIP() []byte {
	file_shrls_proto_rawDescOnce.Do(func() {
		file_shrls_proto_rawDescData = protoimpl.X.CompressGZIP(file_shrls_proto_rawDescData)
	})
	return file_shrls_proto_rawDescData
}

var file_shrls_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shrls_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_shrls_proto_goTypes = []interface{}{
	(ShortURL_ShortURLType)(0), // 0: shrls.ShortURL.ShortURLType
	(*GetShrlRequest)(nil),     // 1: shrls.GetShrlRequest
	(*GetShrlResponse)(nil),    // 2: shrls.GetShrlResponse
	(*Ref)(nil),                // 3: shrls.Ref
	(*ShortURL)(nil),           // 4: shrls.ShortURL
	(*Snippet)(nil),            // 5: shrls.Snippet
	(*Redirect)(nil),           // 6: shrls.Redirect
	(*ExpandedURL)(nil),        // 7: shrls.ExpandedURL
	(*Ref_ShortURL)(nil),       // 8: shrls.Ref.ShortURL
}
var file_shrls_proto_depIdxs = []int32{
	8, // 0: shrls.GetShrlRequest.shrl:type_name -> shrls.Ref.ShortURL
	4, // 1: shrls.GetShrlResponse.shrl:type_name -> shrls.ShortURL
	0, // 2: shrls.ShortURL.type:type_name -> shrls.ShortURL.ShortURLType
	7, // 3: shrls.ShortURL.content:type_name -> shrls.ExpandedURL
	6, // 4: shrls.ExpandedURL.url:type_name -> shrls.Redirect
	5, // 5: shrls.ExpandedURL.snippet:type_name -> shrls.Snippet
	1, // 6: shrls.Shrls.GetShrl:input_type -> shrls.GetShrlRequest
	2, // 7: shrls.Shrls.GetShrl:output_type -> shrls.GetShrlResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_shrls_proto_init() }
func file_shrls_proto_init() {
	if File_shrls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shrls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShrlRequest); i {
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
		file_shrls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShrlResponse); i {
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
		file_shrls_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ref); i {
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
		file_shrls_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortURL); i {
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
		file_shrls_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snippet); i {
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
		file_shrls_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redirect); i {
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
		file_shrls_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandedURL); i {
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
		file_shrls_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ref_ShortURL); i {
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
	file_shrls_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ExpandedURL_Url)(nil),
		(*ExpandedURL_File)(nil),
		(*ExpandedURL_Snippet)(nil),
	}
	file_shrls_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*Ref_ShortURL_Id)(nil),
		(*Ref_ShortURL_Alias)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shrls_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shrls_proto_goTypes,
		DependencyIndexes: file_shrls_proto_depIdxs,
		EnumInfos:         file_shrls_proto_enumTypes,
		MessageInfos:      file_shrls_proto_msgTypes,
	}.Build()
	File_shrls_proto = out.File
	file_shrls_proto_rawDesc = nil
	file_shrls_proto_goTypes = nil
	file_shrls_proto_depIdxs = nil
}
