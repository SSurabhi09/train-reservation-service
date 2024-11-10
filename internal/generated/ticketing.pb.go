// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: ticketing.proto

package ticketing

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

// Message for a Train Reservation Request
type CreateReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservation *Reservation `protobuf:"bytes,1,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *CreateReservationRequest) Reset() {
	*x = CreateReservationRequest{}
	mi := &file_ticketing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReservationRequest) ProtoMessage() {}

func (x *CreateReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReservationRequest.ProtoReflect.Descriptor instead.
func (*CreateReservationRequest) Descriptor() ([]byte, []int) {
	return file_ticketing_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReservationRequest) GetReservation() *Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

type GetReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationId string `protobuf:"bytes,1,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
}

func (x *GetReservationRequest) Reset() {
	*x = GetReservationRequest{}
	mi := &file_ticketing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReservationRequest) ProtoMessage() {}

func (x *GetReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReservationRequest.ProtoReflect.Descriptor instead.
func (*GetReservationRequest) Descriptor() ([]byte, []int) {
	return file_ticketing_proto_rawDescGZIP(), []int{1}
}

func (x *GetReservationRequest) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`     //User details
	Ticket *Ticket `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"` //Ticket details
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	mi := &file_ticketing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_ticketing_proto_rawDescGZIP(), []int{2}
}

func (x *Reservation) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Reservation) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

// Message for a User
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"` //[json_name = "firstName"]// without json_name field compiler anyways generates with camelCase
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`    // User's last name
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`                          // User's email
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_ticketing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_ticketing_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Message for Ticket
type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`     // Source
	To    string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`         // Destination
	Price float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"` // Price
	Seat  string  `protobuf:"bytes,4,opt,name=seat,proto3" json:"seat,omitempty"`     // Seat number under section A or B
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	mi := &file_ticketing_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_proto_msgTypes[4]
	if x != nil {
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
	return file_ticketing_proto_rawDescGZIP(), []int{4}
}

func (x *Ticket) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Ticket) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Ticket) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Ticket) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

// Message for a Reservation Response - receipt
type CreateReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success       bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                                 // Request success or failure
	Message       string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                                  // Message for success or failure cases
	ReservationId string `protobuf:"bytes,3,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"` // Reservation
}

func (x *CreateReservationResponse) Reset() {
	*x = CreateReservationResponse{}
	mi := &file_ticketing_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReservationResponse) ProtoMessage() {}

func (x *CreateReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReservationResponse.ProtoReflect.Descriptor instead.
func (*CreateReservationResponse) Descriptor() ([]byte, []int) {
	return file_ticketing_proto_rawDescGZIP(), []int{5}
}

func (x *CreateReservationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateReservationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateReservationResponse) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

type GetReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success            bool         `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                                                // Request success or failure
	Message            string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                                                 // Message for success or failure cases
	ReservationDetails *Reservation `protobuf:"bytes,3,opt,name=reservation_details,json=reservationDetails,proto3" json:"reservation_details,omitempty"` // Reservation
}

func (x *GetReservationResponse) Reset() {
	*x = GetReservationResponse{}
	mi := &file_ticketing_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReservationResponse) ProtoMessage() {}

func (x *GetReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReservationResponse.ProtoReflect.Descriptor instead.
func (*GetReservationResponse) Descriptor() ([]byte, []int) {
	return file_ticketing_proto_rawDescGZIP(), []int{6}
}

func (x *GetReservationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetReservationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetReservationResponse) GetReservationDetails() *Reservation {
	if x != nil {
		return x.ReservationDetails
	}
	return nil
}

var File_ticketing_proto protoreflect.FileDescriptor

var file_ticketing_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x18, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x5d, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22,
	0x58, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x56, 0x0a, 0x06, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x61,
	0x74, 0x22, 0x76, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x32, 0x92, 0x02, 0x0a, 0x10, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01,
	0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x16, 0x5a, 0x14, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x3b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticketing_proto_rawDescOnce sync.Once
	file_ticketing_proto_rawDescData = file_ticketing_proto_rawDesc
)

func file_ticketing_proto_rawDescGZIP() []byte {
	file_ticketing_proto_rawDescOnce.Do(func() {
		file_ticketing_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticketing_proto_rawDescData)
	})
	return file_ticketing_proto_rawDescData
}

var file_ticketing_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ticketing_proto_goTypes = []any{
	(*CreateReservationRequest)(nil),  // 0: ticketing.CreateReservationRequest
	(*GetReservationRequest)(nil),     // 1: ticketing.GetReservationRequest
	(*Reservation)(nil),               // 2: ticketing.Reservation
	(*User)(nil),                      // 3: ticketing.User
	(*Ticket)(nil),                    // 4: ticketing.Ticket
	(*CreateReservationResponse)(nil), // 5: ticketing.CreateReservationResponse
	(*GetReservationResponse)(nil),    // 6: ticketing.GetReservationResponse
}
var file_ticketing_proto_depIdxs = []int32{
	2, // 0: ticketing.CreateReservationRequest.reservation:type_name -> ticketing.Reservation
	3, // 1: ticketing.Reservation.user:type_name -> ticketing.User
	4, // 2: ticketing.Reservation.ticket:type_name -> ticketing.Ticket
	2, // 3: ticketing.GetReservationResponse.reservation_details:type_name -> ticketing.Reservation
	0, // 4: ticketing.TicketingService.CreateReservation:input_type -> ticketing.CreateReservationRequest
	1, // 5: ticketing.TicketingService.GetReservation:input_type -> ticketing.GetReservationRequest
	5, // 6: ticketing.TicketingService.CreateReservation:output_type -> ticketing.CreateReservationResponse
	6, // 7: ticketing.TicketingService.GetReservation:output_type -> ticketing.GetReservationResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ticketing_proto_init() }
func file_ticketing_proto_init() {
	if File_ticketing_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ticketing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticketing_proto_goTypes,
		DependencyIndexes: file_ticketing_proto_depIdxs,
		MessageInfos:      file_ticketing_proto_msgTypes,
	}.Build()
	File_ticketing_proto = out.File
	file_ticketing_proto_rawDesc = nil
	file_ticketing_proto_goTypes = nil
	file_ticketing_proto_depIdxs = nil
}
