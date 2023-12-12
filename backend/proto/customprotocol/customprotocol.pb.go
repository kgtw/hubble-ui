// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: customprotocol/customprotocol.proto

package customprotocol

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

type Error_Kind int32

const (
	Error_Unknown Error_Kind = 0
	Error_Grpc    Error_Kind = 1
)

// Enum value maps for Error_Kind.
var (
	Error_Kind_name = map[int32]string{
		0: "Unknown",
		1: "Grpc",
	}
	Error_Kind_value = map[string]int32{
		"Unknown": 0,
		"Grpc":    1,
	}
)

func (x Error_Kind) Enum() *Error_Kind {
	p := new(Error_Kind)
	*p = x
	return p
}

func (x Error_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_customprotocol_customprotocol_proto_enumTypes[0].Descriptor()
}

func (Error_Kind) Type() protoreflect.EnumType {
	return &file_customprotocol_customprotocol_proto_enumTypes[0]
}

func (x Error_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_Kind.Descriptor instead.
func (Error_Kind) EnumDescriptor() ([]byte, []int) {
	return file_customprotocol_customprotocol_proto_rawDescGZIP(), []int{2, 0}
}

// NOTE: The same structure is used for both Request and Response
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Body *Body `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customprotocol_customprotocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_customprotocol_customprotocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_customprotocol_customprotocol_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Message) GetBody() *Body {
	if x != nil {
		return x.Body
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NOTE: This field can be used for debug or ux research purposes.
	TraceId string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// NOTE: This identifier can be used to mark logical channel or stream. For
	// example when we establish a stream, we also need to route user requests
	// to stored (cached) streams. The same is applied when used with chunks/etc.
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// NOTE: This is a name of route this message delivered to/from.
	RouteName string `protobuf:"bytes,3,opt,name=route_name,json=routeName,proto3" json:"route_name,omitempty"`
	// NOTE: Should be true if sender handler is terminating
	IsTerminated bool `protobuf:"varint,4,opt,name=is_terminated,json=isTerminated,proto3" json:"is_terminated,omitempty"`
	// NOTE: Indicates that outgoing message is not ready yet. Polling request
	// must be repeated after poll_delay_ms (see below).
	IsNotReady bool `protobuf:"varint,5,opt,name=is_not_ready,json=isNotReady,proto3" json:"is_not_ready,omitempty"`
	// NOTE: Error (handler or internal) occured, its details serialized to body
	IsError bool     `protobuf:"varint,6,opt,name=is_error,json=isError,proto3" json:"is_error,omitempty"`
	Errors  []*Error `protobuf:"bytes,7,rep,name=errors,proto3" json:"errors,omitempty"`
	// NOTE: Server will use this field to notify the client on what time it
	// should wait before doing next poll request.
	PollDelayMs uint64 `protobuf:"varint,8,opt,name=poll_delay_ms,json=pollDelayMs,proto3" json:"poll_delay_ms,omitempty"`
	// NOTE: When sent from backend, this flag means that there are no messages
	// in outgoing channel. Combined with `is_terminated` flag this one can be
	// used to prevent further polling/closing requests.
	IsEmpty bool `protobuf:"varint,9,opt,name=is_empty,json=isEmpty,proto3" json:"is_empty,omitempty"`
	// NOTE: This fields are useful for organizing page navigation or data
	// transfer in similar manner
	LastDatumId string `protobuf:"bytes,100,opt,name=last_datum_id,json=lastDatumId,proto3" json:"last_datum_id,omitempty"`
	NextDatumId string `protobuf:"bytes,101,opt,name=next_datum_id,json=nextDatumId,proto3" json:"next_datum_id,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customprotocol_customprotocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_customprotocol_customprotocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_customprotocol_customprotocol_proto_rawDescGZIP(), []int{1}
}

func (x *Meta) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *Meta) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Meta) GetRouteName() string {
	if x != nil {
		return x.RouteName
	}
	return ""
}

func (x *Meta) GetIsTerminated() bool {
	if x != nil {
		return x.IsTerminated
	}
	return false
}

func (x *Meta) GetIsNotReady() bool {
	if x != nil {
		return x.IsNotReady
	}
	return false
}

func (x *Meta) GetIsError() bool {
	if x != nil {
		return x.IsError
	}
	return false
}

func (x *Meta) GetErrors() []*Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *Meta) GetPollDelayMs() uint64 {
	if x != nil {
		return x.PollDelayMs
	}
	return 0
}

func (x *Meta) GetIsEmpty() bool {
	if x != nil {
		return x.IsEmpty
	}
	return false
}

func (x *Meta) GetLastDatumId() string {
	if x != nil {
		return x.LastDatumId
	}
	return ""
}

func (x *Meta) GetNextDatumId() string {
	if x != nil {
		return x.NextDatumId
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind    Error_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=customprotocol.Error_Kind" json:"kind,omitempty"`
	Code    int32      `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customprotocol_customprotocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_customprotocol_customprotocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_customprotocol_customprotocol_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetKind() Error_Kind {
	if x != nil {
		return x.Kind
	}
	return Error_Unknown
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Body) Reset() {
	*x = Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customprotocol_customprotocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Body) ProtoMessage() {}

func (x *Body) ProtoReflect() protoreflect.Message {
	mi := &file_customprotocol_customprotocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Body.ProtoReflect.Descriptor instead.
func (*Body) Descriptor() ([]byte, []int) {
	return file_customprotocol_customprotocol_proto_rawDescGZIP(), []int{3}
}

func (x *Body) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_customprotocol_customprotocol_proto protoreflect.FileDescriptor

var file_customprotocol_customprotocol_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x5d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0xf7, 0x02, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x73, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x6f, 0x6c, 0x6c, 0x5f,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x70, 0x6f, 0x6c, 0x6c, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x84,
	0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x4b, 0x69,
	0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1d, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47,
	0x72, 0x70, 0x63, 0x10, 0x01, 0x22, 0x20, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customprotocol_customprotocol_proto_rawDescOnce sync.Once
	file_customprotocol_customprotocol_proto_rawDescData = file_customprotocol_customprotocol_proto_rawDesc
)

func file_customprotocol_customprotocol_proto_rawDescGZIP() []byte {
	file_customprotocol_customprotocol_proto_rawDescOnce.Do(func() {
		file_customprotocol_customprotocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_customprotocol_customprotocol_proto_rawDescData)
	})
	return file_customprotocol_customprotocol_proto_rawDescData
}

var file_customprotocol_customprotocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_customprotocol_customprotocol_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_customprotocol_customprotocol_proto_goTypes = []interface{}{
	(Error_Kind)(0), // 0: customprotocol.Error.Kind
	(*Message)(nil), // 1: customprotocol.Message
	(*Meta)(nil),    // 2: customprotocol.Meta
	(*Error)(nil),   // 3: customprotocol.Error
	(*Body)(nil),    // 4: customprotocol.Body
}
var file_customprotocol_customprotocol_proto_depIdxs = []int32{
	2, // 0: customprotocol.Message.meta:type_name -> customprotocol.Meta
	4, // 1: customprotocol.Message.body:type_name -> customprotocol.Body
	3, // 2: customprotocol.Meta.errors:type_name -> customprotocol.Error
	0, // 3: customprotocol.Error.kind:type_name -> customprotocol.Error.Kind
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_customprotocol_customprotocol_proto_init() }
func file_customprotocol_customprotocol_proto_init() {
	if File_customprotocol_customprotocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customprotocol_customprotocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_customprotocol_customprotocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_customprotocol_customprotocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_customprotocol_customprotocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Body); i {
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
			RawDescriptor: file_customprotocol_customprotocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customprotocol_customprotocol_proto_goTypes,
		DependencyIndexes: file_customprotocol_customprotocol_proto_depIdxs,
		EnumInfos:         file_customprotocol_customprotocol_proto_enumTypes,
		MessageInfos:      file_customprotocol_customprotocol_proto_msgTypes,
	}.Build()
	File_customprotocol_customprotocol_proto = out.File
	file_customprotocol_customprotocol_proto_rawDesc = nil
	file_customprotocol_customprotocol_proto_goTypes = nil
	file_customprotocol_customprotocol_proto_depIdxs = nil
}
