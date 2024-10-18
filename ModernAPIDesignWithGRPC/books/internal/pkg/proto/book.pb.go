// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.6
// source: books/internal/pkg/proto/book.proto

package proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[0]
	if x != nil {
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
	return file_books_internal_pkg_proto_book_proto_rawDescGZIP(), []int{0}
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isbn      int32  `protobuf:"varint,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Publisher string `protobuf:"bytes,3,opt,name=publisher,proto3" json:"publisher,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_books_internal_pkg_proto_book_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetIsbn() int32 {
	if x != nil {
		return x.Isbn
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

type AddBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AddBookResponse) Reset() {
	*x = AddBookResponse{}
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookResponse) ProtoMessage() {}

func (x *AddBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookResponse.ProtoReflect.Descriptor instead.
func (*AddBookResponse) Descriptor() ([]byte, []int) {
	return file_books_internal_pkg_proto_book_proto_rawDescGZIP(), []int{2}
}

func (x *AddBookResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ListBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *ListBooksResponse) Reset() {
	*x = ListBooksResponse{}
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksResponse) ProtoMessage() {}

func (x *ListBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksResponse.ProtoReflect.Descriptor instead.
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return file_books_internal_pkg_proto_book_proto_rawDescGZIP(), []int{3}
}

func (x *ListBooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type GetBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isbn int32 `protobuf:"varint,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_books_internal_pkg_proto_book_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookRequest) GetIsbn() int32 {
	if x != nil {
		return x.Isbn
	}
	return 0
}

type RemoveBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isbn int32 `protobuf:"varint,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
}

func (x *RemoveBookRequest) Reset() {
	*x = RemoveBookRequest{}
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBookRequest) ProtoMessage() {}

func (x *RemoveBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBookRequest.ProtoReflect.Descriptor instead.
func (*RemoveBookRequest) Descriptor() ([]byte, []int) {
	return file_books_internal_pkg_proto_book_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveBookRequest) GetIsbn() int32 {
	if x != nil {
		return x.Isbn
	}
	return 0
}

type RemoveBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RemoveBookResponse) Reset() {
	*x = RemoveBookResponse{}
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBookResponse) ProtoMessage() {}

func (x *RemoveBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBookResponse.ProtoReflect.Descriptor instead.
func (*RemoveBookResponse) Descriptor() ([]byte, []int) {
	return file_books_internal_pkg_proto_book_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveBookResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateBookResponse) Reset() {
	*x = UpdateBookResponse{}
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookResponse) ProtoMessage() {}

func (x *UpdateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_internal_pkg_proto_book_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookResponse.ProtoReflect.Descriptor instead.
func (*UpdateBookResponse) Descriptor() ([]byte, []int) {
	return file_books_internal_pkg_proto_book_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBookResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_books_internal_pkg_proto_book_proto protoreflect.FileDescriptor

var file_books_internal_pkg_proto_book_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x72, 0x6f, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x4c, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x73, 0x62, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x22, 0x29, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x22, 0x27, 0x0a, 0x11, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x69,
	0x73, 0x62, 0x6e, 0x22, 0x2c, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x2c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0x90, 0x02, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x3f, 0x0a,
	0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_internal_pkg_proto_book_proto_rawDescOnce sync.Once
	file_books_internal_pkg_proto_book_proto_rawDescData = file_books_internal_pkg_proto_book_proto_rawDesc
)

func file_books_internal_pkg_proto_book_proto_rawDescGZIP() []byte {
	file_books_internal_pkg_proto_book_proto_rawDescOnce.Do(func() {
		file_books_internal_pkg_proto_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_internal_pkg_proto_book_proto_rawDescData)
	})
	return file_books_internal_pkg_proto_book_proto_rawDescData
}

var file_books_internal_pkg_proto_book_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_books_internal_pkg_proto_book_proto_goTypes = []any{
	(*Empty)(nil),              // 0: prot.Empty
	(*Book)(nil),               // 1: prot.Book
	(*AddBookResponse)(nil),    // 2: prot.AddBookResponse
	(*ListBooksResponse)(nil),  // 3: prot.ListBooksResponse
	(*GetBookRequest)(nil),     // 4: prot.GetBookRequest
	(*RemoveBookRequest)(nil),  // 5: prot.RemoveBookRequest
	(*RemoveBookResponse)(nil), // 6: prot.RemoveBookResponse
	(*UpdateBookResponse)(nil), // 7: prot.UpdateBookResponse
}
var file_books_internal_pkg_proto_book_proto_depIdxs = []int32{
	1, // 0: prot.ListBooksResponse.books:type_name -> prot.Book
	1, // 1: prot.BookService.AddBook:input_type -> prot.Book
	0, // 2: prot.BookService.ListBooks:input_type -> prot.Empty
	4, // 3: prot.BookService.GetBook:input_type -> prot.GetBookRequest
	5, // 4: prot.BookService.RemoveBook:input_type -> prot.RemoveBookRequest
	1, // 5: prot.BookService.UpdateBook:input_type -> prot.Book
	2, // 6: prot.BookService.AddBook:output_type -> prot.AddBookResponse
	3, // 7: prot.BookService.ListBooks:output_type -> prot.ListBooksResponse
	1, // 8: prot.BookService.GetBook:output_type -> prot.Book
	6, // 9: prot.BookService.RemoveBook:output_type -> prot.RemoveBookResponse
	7, // 10: prot.BookService.UpdateBook:output_type -> prot.UpdateBookResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_books_internal_pkg_proto_book_proto_init() }
func file_books_internal_pkg_proto_book_proto_init() {
	if File_books_internal_pkg_proto_book_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_books_internal_pkg_proto_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_books_internal_pkg_proto_book_proto_goTypes,
		DependencyIndexes: file_books_internal_pkg_proto_book_proto_depIdxs,
		MessageInfos:      file_books_internal_pkg_proto_book_proto_msgTypes,
	}.Build()
	File_books_internal_pkg_proto_book_proto = out.File
	file_books_internal_pkg_proto_book_proto_rawDesc = nil
	file_books_internal_pkg_proto_book_proto_goTypes = nil
	file_books_internal_pkg_proto_book_proto_depIdxs = nil
}