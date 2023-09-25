// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/customer.proto

package Sample_crud

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

type DOB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int64 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int64 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Date  int64 `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *DOB) Reset() {
	*x = DOB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DOB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DOB) ProtoMessage() {}

func (x *DOB) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DOB.ProtoReflect.Descriptor instead.
func (*DOB) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{0}
}

func (x *DOB) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *DOB) GetMonth() int64 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *DOB) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	PostalCode    string `protobuf:"bytes,4,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	Country       string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type CustomerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName      string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName       string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	CustomerId     string   `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	SsnNumber      int64    `protobuf:"varint,4,opt,name=ssn_number,json=ssnNumber,proto3" json:"ssn_number,omitempty"`
	PhoneNumber    int64    `protobuf:"varint,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email          string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Dob            *DOB     `protobuf:"bytes,7,opt,name=dob,proto3" json:"dob,omitempty"`
	Address        *Address `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	JobTitle       string   `protobuf:"bytes,9,opt,name=job_title,json=jobTitle,proto3" json:"job_title,omitempty"`
	AnnualIncome   float32  `protobuf:"fixed32,10,opt,name=annual_income,json=annualIncome,proto3" json:"annual_income,omitempty"`
	InitialDeposit float32  `protobuf:"fixed32,11,opt,name=initial_deposit,json=initialDeposit,proto3" json:"initial_deposit,omitempty"`
}

func (x *CustomerInfo) Reset() {
	*x = CustomerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerInfo) ProtoMessage() {}

func (x *CustomerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerInfo.ProtoReflect.Descriptor instead.
func (*CustomerInfo) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CustomerInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CustomerInfo) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CustomerInfo) GetSsnNumber() int64 {
	if x != nil {
		return x.SsnNumber
	}
	return 0
}

func (x *CustomerInfo) GetPhoneNumber() int64 {
	if x != nil {
		return x.PhoneNumber
	}
	return 0
}

func (x *CustomerInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerInfo) GetDob() *DOB {
	if x != nil {
		return x.Dob
	}
	return nil
}

func (x *CustomerInfo) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *CustomerInfo) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

func (x *CustomerInfo) GetAnnualIncome() float32 {
	if x != nil {
		return x.AnnualIncome
	}
	return 0
}

func (x *CustomerInfo) GetInitialDeposit() float32 {
	if x != nil {
		return x.InitialDeposit
	}
	return 0
}

type DBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SsnNumber   int64  `protobuf:"varint,1,opt,name=ssn_number,json=ssnNumber,proto3" json:"ssn_number,omitempty"`
	PhoneNumber int64  `protobuf:"varint,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *DBResponse) Reset() {
	*x = DBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBResponse) ProtoMessage() {}

func (x *DBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBResponse.ProtoReflect.Descriptor instead.
func (*DBResponse) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{3}
}

func (x *DBResponse) GetSsnNumber() int64 {
	if x != nil {
		return x.SsnNumber
	}
	return 0
}

func (x *DBResponse) GetPhoneNumber() int64 {
	if x != nil {
		return x.PhoneNumber
	}
	return 0
}

func (x *DBResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_proto_customer_proto protoreflect.FileDescriptor

var file_proto_customer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x4e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x22, 0x43, 0x0a, 0x03, 0x44, 0x4f, 0x42, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x95, 0x01, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x22, 0x8a, 0x03, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x73, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x73, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x03, 0x64, 0x6f,
	0x62, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x4e, 0x65, 0x74, 0x78, 0x64, 0x5f,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x4f, 0x42, 0x52, 0x03, 0x64,
	0x6f, 0x62, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x4e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x61, 0x6e, 0x6e, 0x75,
	0x61, 0x6c, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x22, 0x64, 0x0a, 0x0a, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x73, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x73, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0x5f, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x4e,
	0x65, 0x74, 0x78, 0x64, 0x5f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1b, 0x2e, 0x4e, 0x65,
	0x74, 0x78, 0x64, 0x5f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x42,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x61, 0x73, 0x65, 0x6e, 0x74, 0x68, 0x44, 0x2f,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x63, 0x72, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_customer_proto_rawDescOnce sync.Once
	file_proto_customer_proto_rawDescData = file_proto_customer_proto_rawDesc
)

func file_proto_customer_proto_rawDescGZIP() []byte {
	file_proto_customer_proto_rawDescOnce.Do(func() {
		file_proto_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_customer_proto_rawDescData)
	})
	return file_proto_customer_proto_rawDescData
}

var file_proto_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_customer_proto_goTypes = []interface{}{
	(*DOB)(nil),          // 0: Netxd_Customers.DOB
	(*Address)(nil),      // 1: Netxd_Customers.Address
	(*CustomerInfo)(nil), // 2: Netxd_Customers.CustomerInfo
	(*DBResponse)(nil),   // 3: Netxd_Customers.DBResponse
}
var file_proto_customer_proto_depIdxs = []int32{
	0, // 0: Netxd_Customers.CustomerInfo.dob:type_name -> Netxd_Customers.DOB
	1, // 1: Netxd_Customers.CustomerInfo.address:type_name -> Netxd_Customers.Address
	2, // 2: Netxd_Customers.CustomerService.CreateCustomer:input_type -> Netxd_Customers.CustomerInfo
	3, // 3: Netxd_Customers.CustomerService.CreateCustomer:output_type -> Netxd_Customers.DBResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_customer_proto_init() }
func file_proto_customer_proto_init() {
	if File_proto_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DOB); i {
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
		file_proto_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_proto_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerInfo); i {
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
		file_proto_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBResponse); i {
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
			RawDescriptor: file_proto_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_customer_proto_goTypes,
		DependencyIndexes: file_proto_customer_proto_depIdxs,
		MessageInfos:      file_proto_customer_proto_msgTypes,
	}.Build()
	File_proto_customer_proto = out.File
	file_proto_customer_proto_rawDesc = nil
	file_proto_customer_proto_goTypes = nil
	file_proto_customer_proto_depIdxs = nil
}
