// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: ticketService.proto

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

type TicketBookingStatus int32

const (
	TicketBookingStatus_TICKET_BOOKING_STATUS_UNKNOWN         TicketBookingStatus = 0
	TicketBookingStatus_TICKET_BOOKING_STATUS_PENDING         TicketBookingStatus = 1
	TicketBookingStatus_TICKET_BOOKING_STATUS_RESERVING_SEATS TicketBookingStatus = 2
	TicketBookingStatus_TICKET_BOOKING_STATUS_SEATS_BOOKED    TicketBookingStatus = 3
	TicketBookingStatus_TICKET_BOOKING_STATUS_CONFIRMED       TicketBookingStatus = 4
	TicketBookingStatus_TICKET_BOOKING_STATUS_REJECTED        TicketBookingStatus = 5
)

// Enum value maps for TicketBookingStatus.
var (
	TicketBookingStatus_name = map[int32]string{
		0: "TICKET_BOOKING_STATUS_UNKNOWN",
		1: "TICKET_BOOKING_STATUS_PENDING",
		2: "TICKET_BOOKING_STATUS_RESERVING_SEATS",
		3: "TICKET_BOOKING_STATUS_SEATS_BOOKED",
		4: "TICKET_BOOKING_STATUS_CONFIRMED",
		5: "TICKET_BOOKING_STATUS_REJECTED",
	}
	TicketBookingStatus_value = map[string]int32{
		"TICKET_BOOKING_STATUS_UNKNOWN":         0,
		"TICKET_BOOKING_STATUS_PENDING":         1,
		"TICKET_BOOKING_STATUS_RESERVING_SEATS": 2,
		"TICKET_BOOKING_STATUS_SEATS_BOOKED":    3,
		"TICKET_BOOKING_STATUS_CONFIRMED":       4,
		"TICKET_BOOKING_STATUS_REJECTED":        5,
	}
)

func (x TicketBookingStatus) Enum() *TicketBookingStatus {
	p := new(TicketBookingStatus)
	*p = x
	return p
}

func (x TicketBookingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TicketBookingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_ticketService_proto_enumTypes[0].Descriptor()
}

func (TicketBookingStatus) Type() protoreflect.EnumType {
	return &file_ticketService_proto_enumTypes[0]
}

func (x TicketBookingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TicketBookingStatus.Descriptor instead.
func (TicketBookingStatus) EnumDescriptor() ([]byte, []int) {
	return file_ticketService_proto_rawDescGZIP(), []int{0}
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	HallId    int32  `protobuf:"varint,2,opt,name=hall_id,json=hallId,proto3" json:"hall_id,omitempty"`
	Seats     int32  `protobuf:"varint,3,opt,name=seats,proto3" json:"seats,omitempty"`
	EntryCode string `protobuf:"bytes,4,opt,name=entry_code,json=entryCode,proto3" json:"entry_code,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_ticketService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_ticketService_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ticket) GetHallId() int32 {
	if x != nil {
		return x.HallId
	}
	return 0
}

func (x *Ticket) GetSeats() int32 {
	if x != nil {
		return x.Seats
	}
	return 0
}

func (x *Ticket) GetEntryCode() string {
	if x != nil {
		return x.EntryCode
	}
	return ""
}

type CreateTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HallId int32 `protobuf:"varint,1,opt,name=hall_id,json=hallId,proto3" json:"hall_id,omitempty"`
	Seats  int32 `protobuf:"varint,2,opt,name=seats,proto3" json:"seats,omitempty"`
}

func (x *CreateTicketRequest) Reset() {
	*x = CreateTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketRequest) ProtoMessage() {}

func (x *CreateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketRequest.ProtoReflect.Descriptor instead.
func (*CreateTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticketService_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTicketRequest) GetHallId() int32 {
	if x != nil {
		return x.HallId
	}
	return 0
}

func (x *CreateTicketRequest) GetSeats() int32 {
	if x != nil {
		return x.Seats
	}
	return 0
}

type CreateTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket             `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Status TicketBookingStatus `protobuf:"varint,2,opt,name=status,proto3,enum=proto.TicketBookingStatus" json:"status,omitempty"`
}

func (x *CreateTicketResponse) Reset() {
	*x = CreateTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketResponse) ProtoMessage() {}

func (x *CreateTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticketService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketResponse.ProtoReflect.Descriptor instead.
func (*CreateTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticketService_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTicketResponse) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

func (x *CreateTicketResponse) GetStatus() TicketBookingStatus {
	if x != nil {
		return x.Status
	}
	return TicketBookingStatus_TICKET_BOOKING_STATUS_UNKNOWN
}

var File_ticketService_proto protoreflect.FileDescriptor

var file_ticketService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x06,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x61, 0x6c, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x44, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x68,
	0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x61,
	0x6c, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x22, 0x71, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0xf7, 0x01,
	0x0a, 0x13, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x5f,
	0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x54, 0x49, 0x43, 0x4b,
	0x45, 0x54, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x29, 0x0a, 0x25, 0x54,
	0x49, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x45, 0x41, 0x54, 0x53, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54,
	0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x53, 0x45, 0x41, 0x54, 0x53, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x45, 0x44, 0x10, 0x03, 0x12, 0x23,
	0x0a, 0x1f, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x42, 0x4f,
	0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x05, 0x32, 0x5c, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticketService_proto_rawDescOnce sync.Once
	file_ticketService_proto_rawDescData = file_ticketService_proto_rawDesc
)

func file_ticketService_proto_rawDescGZIP() []byte {
	file_ticketService_proto_rawDescOnce.Do(func() {
		file_ticketService_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticketService_proto_rawDescData)
	})
	return file_ticketService_proto_rawDescData
}

var file_ticketService_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ticketService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ticketService_proto_goTypes = []interface{}{
	(TicketBookingStatus)(0),     // 0: proto.TicketBookingStatus
	(*Ticket)(nil),               // 1: proto.Ticket
	(*CreateTicketRequest)(nil),  // 2: proto.CreateTicketRequest
	(*CreateTicketResponse)(nil), // 3: proto.CreateTicketResponse
}
var file_ticketService_proto_depIdxs = []int32{
	1, // 0: proto.CreateTicketResponse.ticket:type_name -> proto.Ticket
	0, // 1: proto.CreateTicketResponse.status:type_name -> proto.TicketBookingStatus
	2, // 2: proto.TicketService.CreateTicket:input_type -> proto.CreateTicketRequest
	3, // 3: proto.TicketService.CreateTicket:output_type -> proto.CreateTicketResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ticketService_proto_init() }
func file_ticketService_proto_init() {
	if File_ticketService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ticketService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
		file_ticketService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTicketRequest); i {
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
		file_ticketService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTicketResponse); i {
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
			RawDescriptor: file_ticketService_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticketService_proto_goTypes,
		DependencyIndexes: file_ticketService_proto_depIdxs,
		EnumInfos:         file_ticketService_proto_enumTypes,
		MessageInfos:      file_ticketService_proto_msgTypes,
	}.Build()
	File_ticketService_proto = out.File
	file_ticketService_proto_rawDesc = nil
	file_ticketService_proto_goTypes = nil
	file_ticketService_proto_depIdxs = nil
}