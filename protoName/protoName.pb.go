// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: protoName/protoName.proto

package protoName

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Para leer/escribir en el log
type LogData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookName      string  `protobuf:"bytes,1,opt,name=bookName,proto3" json:"bookName,omitempty"`
	NumParts      int64   `protobuf:"varint,2,opt,name=numParts,proto3" json:"numParts,omitempty"`
	PartsLocation []*Part `protobuf:"bytes,3,rep,name=partsLocation,proto3" json:"partsLocation,omitempty"`
}

func (x *LogData) Reset() {
	*x = LogData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoName_protoName_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogData) ProtoMessage() {}

func (x *LogData) ProtoReflect() protoreflect.Message {
	mi := &file_protoName_protoName_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogData.ProtoReflect.Descriptor instead.
func (*LogData) Descriptor() ([]byte, []int) {
	return file_protoName_protoName_proto_rawDescGZIP(), []int{0}
}

func (x *LogData) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *LogData) GetNumParts() int64 {
	if x != nil {
		return x.NumParts
	}
	return 0
}

func (x *LogData) GetPartsLocation() []*Part {
	if x != nil {
		return x.PartsLocation
	}
	return nil
}

type Part struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index            int64  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	IpPuertoDatanode string `protobuf:"bytes,2,opt,name=ipPuertoDatanode,proto3" json:"ipPuertoDatanode,omitempty"`
}

func (x *Part) Reset() {
	*x = Part{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoName_protoName_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Part) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Part) ProtoMessage() {}

func (x *Part) ProtoReflect() protoreflect.Message {
	mi := &file_protoName_protoName_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Part.ProtoReflect.Descriptor instead.
func (*Part) Descriptor() ([]byte, []int) {
	return file_protoName_protoName_proto_rawDescGZIP(), []int{1}
}

func (x *Part) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Part) GetIpPuertoDatanode() string {
	if x != nil {
		return x.IpPuertoDatanode
	}
	return ""
}

// Quiero leer el log (se lo envio a todos)
type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Node      int64  `protobuf:"varint,2,opt,name=node,proto3" json:"node,omitempty"`
	Bookname  string `protobuf:"bytes,3,opt,name=bookname,proto3" json:"bookname,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoName_protoName_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoName_protoName_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_protoName_protoName_proto_rawDescGZIP(), []int{2}
}

func (x *ReadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReadRequest) GetNode() int64 {
	if x != nil {
		return x.Node
	}
	return 0
}

func (x *ReadRequest) GetBookname() string {
	if x != nil {
		return x.Bookname
	}
	return ""
}

func (x *ReadRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Quiero escribir el log (se lo envio a todos)
type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Node      int64  `protobuf:"varint,2,opt,name=node,proto3" json:"node,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoName_protoName_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoName_protoName_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_protoName_protoName_proto_rawDescGZIP(), []int{3}
}

func (x *WriteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WriteRequest) GetNode() int64 {
	if x != nil {
		return x.Node
	}
	return 0
}

func (x *WriteRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Enviamos una propuesta al nodo (al nodo que recibe el msg)
type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Node      int64  `protobuf:"varint,2,opt,name=node,proto3" json:"node,omitempty"`
	NumChunks int64  `protobuf:"varint,3,opt,name=numChunks,proto3" json:"numChunks,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoName_protoName_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

func (x *Proposal) ProtoReflect() protoreflect.Message {
	mi := &file_protoName_protoName_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_protoName_protoName_proto_rawDescGZIP(), []int{4}
}

func (x *Proposal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Proposal) GetNode() int64 {
	if x != nil {
		return x.Node
	}
	return 0
}

func (x *Proposal) GetNumChunks() int64 {
	if x != nil {
		return x.NumChunks
	}
	return 0
}

func (x *Proposal) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

//Nodes repart contiene el nro de nodos por cada datanode  EJ: [6,5,1] -> 6 irían al datanode 0
type ProposalToNameNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChunksNode1 []int64 `protobuf:"varint,2,rep,packed,name=chunksNode1,proto3" json:"chunksNode1,omitempty"`
	ChunksNode2 []int64 `protobuf:"varint,3,rep,packed,name=chunksNode2,proto3" json:"chunksNode2,omitempty"`
	ChunksNode3 []int64 `protobuf:"varint,4,rep,packed,name=chunksNode3,proto3" json:"chunksNode3,omitempty"`
	NumChunks   int64   `protobuf:"varint,5,opt,name=numChunks,proto3" json:"numChunks,omitempty"`
	Timestamp   int64   `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ProposalToNameNode) Reset() {
	*x = ProposalToNameNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoName_protoName_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalToNameNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalToNameNode) ProtoMessage() {}

func (x *ProposalToNameNode) ProtoReflect() protoreflect.Message {
	mi := &file_protoName_protoName_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalToNameNode.ProtoReflect.Descriptor instead.
func (*ProposalToNameNode) Descriptor() ([]byte, []int) {
	return file_protoName_protoName_proto_rawDescGZIP(), []int{5}
}

func (x *ProposalToNameNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProposalToNameNode) GetChunksNode1() []int64 {
	if x != nil {
		return x.ChunksNode1
	}
	return nil
}

func (x *ProposalToNameNode) GetChunksNode2() []int64 {
	if x != nil {
		return x.ChunksNode2
	}
	return nil
}

func (x *ProposalToNameNode) GetChunksNode3() []int64 {
	if x != nil {
		return x.ChunksNode3
	}
	return nil
}

func (x *ProposalToNameNode) GetNumChunks() int64 {
	if x != nil {
		return x.NumChunks
	}
	return 0
}

func (x *ProposalToNameNode) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// 1 = apruebo  0 = rechazo
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Response  bool   `protobuf:"varint,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoName_protoName_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_protoName_protoName_proto_msgTypes[6]
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
	return file_protoName_protoName_proto_rawDescGZIP(), []int{6}
}

func (x *Response) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Response) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Response) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type EveryBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []string `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *EveryBook) Reset() {
	*x = EveryBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoName_protoName_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EveryBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EveryBook) ProtoMessage() {}

func (x *EveryBook) ProtoReflect() protoreflect.Message {
	mi := &file_protoName_protoName_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EveryBook.ProtoReflect.Descriptor instead.
func (*EveryBook) Descriptor() ([]byte, []int) {
	return file_protoName_protoName_proto_rawDescGZIP(), []int{7}
}

func (x *EveryBook) GetBooks() []string {
	if x != nil {
		return x.Books
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoName_protoName_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protoName_protoName_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protoName_protoName_proto_rawDescGZIP(), []int{8}
}

var File_protoName_protoName_proto protoreflect.FileDescriptor

var file_protoName_protoName_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x78, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x0d, 0x70, 0x61, 0x72,
	0x74, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x48, 0x0a, 0x04, 0x50, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2a,
	0x0a, 0x10, 0x69, 0x70, 0x50, 0x75, 0x65, 0x72, 0x74, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x70, 0x50, 0x75, 0x65, 0x72,
	0x74, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x6b, 0x0a, 0x0b, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x50, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x6a, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75, 0x6d,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x75,
	0x6d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xc6, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x31, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x31, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x32, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x32,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x33, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x4e, 0x6f, 0x64,
	0x65, 0x33, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x54,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0x87, 0x02, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x54, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x08, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x45,
	0x76, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x53, 0x6f, 0x6c, 0x61, 0x72, 0x56,
	0x2f, 0x74, 0x61, 0x72, 0x65, 0x61, 0x32, 0x2d, 0x73, 0x64, 0x2d, 0x77, 0x69, 0x6e, 0x64, 0x75,
	0x63, 0x6c, 0x6f, 0x76, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoName_protoName_proto_rawDescOnce sync.Once
	file_protoName_protoName_proto_rawDescData = file_protoName_protoName_proto_rawDesc
)

func file_protoName_protoName_proto_rawDescGZIP() []byte {
	file_protoName_protoName_proto_rawDescOnce.Do(func() {
		file_protoName_protoName_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoName_protoName_proto_rawDescData)
	})
	return file_protoName_protoName_proto_rawDescData
}

var file_protoName_protoName_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protoName_protoName_proto_goTypes = []interface{}{
	(*LogData)(nil),            // 0: protoName.LogData
	(*Part)(nil),               // 1: protoName.Part
	(*ReadRequest)(nil),        // 2: protoName.ReadRequest
	(*WriteRequest)(nil),       // 3: protoName.WriteRequest
	(*Proposal)(nil),           // 4: protoName.Proposal
	(*ProposalToNameNode)(nil), // 5: protoName.ProposalToNameNode
	(*Response)(nil),           // 6: protoName.Response
	(*EveryBook)(nil),          // 7: protoName.EveryBook
	(*Empty)(nil),              // 8: protoName.Empty
}
var file_protoName_protoName_proto_depIdxs = []int32{
	1, // 0: protoName.LogData.partsLocation:type_name -> protoName.Part
	5, // 1: protoName.ProtoNameService.DistributeProposal:input_type -> protoName.ProposalToNameNode
	0, // 2: protoName.ProtoNameService.WriteLog:input_type -> protoName.LogData
	2, // 3: protoName.ProtoNameService.ClientRequest:input_type -> protoName.ReadRequest
	8, // 4: protoName.ProtoNameService.GetBooks:input_type -> protoName.Empty
	6, // 5: protoName.ProtoNameService.DistributeProposal:output_type -> protoName.Response
	8, // 6: protoName.ProtoNameService.WriteLog:output_type -> protoName.Empty
	0, // 7: protoName.ProtoNameService.ClientRequest:output_type -> protoName.LogData
	7, // 8: protoName.ProtoNameService.GetBooks:output_type -> protoName.EveryBook
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protoName_protoName_proto_init() }
func file_protoName_protoName_proto_init() {
	if File_protoName_protoName_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoName_protoName_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogData); i {
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
		file_protoName_protoName_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Part); i {
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
		file_protoName_protoName_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_protoName_protoName_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_protoName_protoName_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
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
		file_protoName_protoName_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalToNameNode); i {
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
		file_protoName_protoName_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_protoName_protoName_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EveryBook); i {
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
		file_protoName_protoName_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_protoName_protoName_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoName_protoName_proto_goTypes,
		DependencyIndexes: file_protoName_protoName_proto_depIdxs,
		MessageInfos:      file_protoName_protoName_proto_msgTypes,
	}.Build()
	File_protoName_protoName_proto = out.File
	file_protoName_protoName_proto_rawDesc = nil
	file_protoName_protoName_proto_goTypes = nil
	file_protoName_protoName_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProtoNameServiceClient is the client API for ProtoNameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProtoNameServiceClient interface {
	// Metodos NameNode-Datanode
	// Coordinación entre DataNodes y entre NameNode con DataNode para verificar la propuesta
	DistributeProposal(ctx context.Context, in *ProposalToNameNode, opts ...grpc.CallOption) (*Response, error)
	//Función para escribir el Log
	WriteLog(ctx context.Context, in *LogData, opts ...grpc.CallOption) (*Empty, error)
	// Metodos cliente-NameNode
	// El cliente pide leer el Log para saber ubicación de los Chunks
	ClientRequest(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*LogData, error)
	GetBooks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EveryBook, error)
}

type protoNameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProtoNameServiceClient(cc grpc.ClientConnInterface) ProtoNameServiceClient {
	return &protoNameServiceClient{cc}
}

func (c *protoNameServiceClient) DistributeProposal(ctx context.Context, in *ProposalToNameNode, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protoName.ProtoNameService/DistributeProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoNameServiceClient) WriteLog(ctx context.Context, in *LogData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protoName.ProtoNameService/WriteLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoNameServiceClient) ClientRequest(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*LogData, error) {
	out := new(LogData)
	err := c.cc.Invoke(ctx, "/protoName.ProtoNameService/ClientRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoNameServiceClient) GetBooks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EveryBook, error) {
	out := new(EveryBook)
	err := c.cc.Invoke(ctx, "/protoName.ProtoNameService/GetBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtoNameServiceServer is the server API for ProtoNameService service.
type ProtoNameServiceServer interface {
	// Metodos NameNode-Datanode
	// Coordinación entre DataNodes y entre NameNode con DataNode para verificar la propuesta
	DistributeProposal(context.Context, *ProposalToNameNode) (*Response, error)
	//Función para escribir el Log
	WriteLog(context.Context, *LogData) (*Empty, error)
	// Metodos cliente-NameNode
	// El cliente pide leer el Log para saber ubicación de los Chunks
	ClientRequest(context.Context, *ReadRequest) (*LogData, error)
	GetBooks(context.Context, *Empty) (*EveryBook, error)
}

// UnimplementedProtoNameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProtoNameServiceServer struct {
}

func (*UnimplementedProtoNameServiceServer) DistributeProposal(context.Context, *ProposalToNameNode) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeProposal not implemented")
}
func (*UnimplementedProtoNameServiceServer) WriteLog(context.Context, *LogData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteLog not implemented")
}
func (*UnimplementedProtoNameServiceServer) ClientRequest(context.Context, *ReadRequest) (*LogData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientRequest not implemented")
}
func (*UnimplementedProtoNameServiceServer) GetBooks(context.Context, *Empty) (*EveryBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}

func RegisterProtoNameServiceServer(s *grpc.Server, srv ProtoNameServiceServer) {
	s.RegisterService(&_ProtoNameService_serviceDesc, srv)
}

func _ProtoNameService_DistributeProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposalToNameNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoNameServiceServer).DistributeProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoName.ProtoNameService/DistributeProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoNameServiceServer).DistributeProposal(ctx, req.(*ProposalToNameNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoNameService_WriteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoNameServiceServer).WriteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoName.ProtoNameService/WriteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoNameServiceServer).WriteLog(ctx, req.(*LogData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoNameService_ClientRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoNameServiceServer).ClientRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoName.ProtoNameService/ClientRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoNameServiceServer).ClientRequest(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoNameService_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoNameServiceServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoName.ProtoNameService/GetBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoNameServiceServer).GetBooks(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProtoNameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoName.ProtoNameService",
	HandlerType: (*ProtoNameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DistributeProposal",
			Handler:    _ProtoNameService_DistributeProposal_Handler,
		},
		{
			MethodName: "WriteLog",
			Handler:    _ProtoNameService_WriteLog_Handler,
		},
		{
			MethodName: "ClientRequest",
			Handler:    _ProtoNameService_ClientRequest_Handler,
		},
		{
			MethodName: "GetBooks",
			Handler:    _ProtoNameService_GetBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoName/protoName.proto",
}
