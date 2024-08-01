// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: submodule-food-delivery/order/office.proto

package order

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

// Office represents a market's branch where orders can be delivered.
type Office struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address   string  `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Latitude  float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"` // Add more fields as needed (e.g., contact information, opening hours)
}

func (x *Office) Reset() {
	*x = Office{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_food_delivery_order_office_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Office) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Office) ProtoMessage() {}

func (x *Office) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_food_delivery_order_office_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Office.ProtoReflect.Descriptor instead.
func (*Office) Descriptor() ([]byte, []int) {
	return file_submodule_food_delivery_order_office_proto_rawDescGZIP(), []int{0}
}

func (x *Office) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Office) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Office) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Office) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Office) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

// OfficeRequest is used for various office related gRPC methods.
type OfficeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OfficeRequest) Reset() {
	*x = OfficeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_food_delivery_order_office_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeRequest) ProtoMessage() {}

func (x *OfficeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_food_delivery_order_office_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeRequest.ProtoReflect.Descriptor instead.
func (*OfficeRequest) Descriptor() ([]byte, []int) {
	return file_submodule_food_delivery_order_office_proto_rawDescGZIP(), []int{1}
}

func (x *OfficeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// CreateOfficeRequest is used for creating a new office.
type CreateOfficeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Office *Office `protobuf:"bytes,1,opt,name=office,proto3" json:"office,omitempty"`
}

func (x *CreateOfficeRequest) Reset() {
	*x = CreateOfficeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_food_delivery_order_office_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOfficeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOfficeRequest) ProtoMessage() {}

func (x *CreateOfficeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_food_delivery_order_office_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOfficeRequest.ProtoReflect.Descriptor instead.
func (*CreateOfficeRequest) Descriptor() ([]byte, []int) {
	return file_submodule_food_delivery_order_office_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOfficeRequest) GetOffice() *Office {
	if x != nil {
		return x.Office
	}
	return nil
}

// UpdateOfficeRequest is used for updating an existing office.
type UpdateOfficeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Office *Office `protobuf:"bytes,1,opt,name=office,proto3" json:"office,omitempty"`
}

func (x *UpdateOfficeRequest) Reset() {
	*x = UpdateOfficeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_food_delivery_order_office_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOfficeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOfficeRequest) ProtoMessage() {}

func (x *UpdateOfficeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_food_delivery_order_office_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOfficeRequest.ProtoReflect.Descriptor instead.
func (*UpdateOfficeRequest) Descriptor() ([]byte, []int) {
	return file_submodule_food_delivery_order_office_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOfficeRequest) GetOffice() *Office {
	if x != nil {
		return x.Office
	}
	return nil
}

// PatchOfficeRequest is used for partially updating an existing office.
type PatchOfficeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address   string  `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Latitude  float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *PatchOfficeRequest) Reset() {
	*x = PatchOfficeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_food_delivery_order_office_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchOfficeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchOfficeRequest) ProtoMessage() {}

func (x *PatchOfficeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_food_delivery_order_office_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchOfficeRequest.ProtoReflect.Descriptor instead.
func (*PatchOfficeRequest) Descriptor() ([]byte, []int) {
	return file_submodule_food_delivery_order_office_proto_rawDescGZIP(), []int{4}
}

func (x *PatchOfficeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PatchOfficeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PatchOfficeRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PatchOfficeRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *PatchOfficeRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

// GetOfficesRequest is used for retrieving a list of offices.
type GetOfficesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit   int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetOfficesRequest) Reset() {
	*x = GetOfficesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_food_delivery_order_office_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOfficesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOfficesRequest) ProtoMessage() {}

func (x *GetOfficesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_food_delivery_order_office_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOfficesRequest.ProtoReflect.Descriptor instead.
func (*GetOfficesRequest) Descriptor() ([]byte, []int) {
	return file_submodule_food_delivery_order_office_proto_rawDescGZIP(), []int{5}
}

func (x *GetOfficesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetOfficesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetOfficesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetOfficesRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// GetOfficesResponse is used for returning a list of offices.
type GetOfficesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offices []*Office `protobuf:"bytes,1,rep,name=offices,proto3" json:"offices,omitempty"`
}

func (x *GetOfficesResponse) Reset() {
	*x = GetOfficesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_food_delivery_order_office_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOfficesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOfficesResponse) ProtoMessage() {}

func (x *GetOfficesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_food_delivery_order_office_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOfficesResponse.ProtoReflect.Descriptor instead.
func (*GetOfficesResponse) Descriptor() ([]byte, []int) {
	return file_submodule_food_delivery_order_office_proto_rawDescGZIP(), []int{6}
}

func (x *GetOfficesResponse) GetOffices() []*Office {
	if x != nil {
		return x.Offices
	}
	return nil
}

// DeleteOfficeRequest is used for deleting an office.
type DeleteOfficeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteOfficeRequest) Reset() {
	*x = DeleteOfficeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_food_delivery_order_office_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOfficeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOfficeRequest) ProtoMessage() {}

func (x *DeleteOfficeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_food_delivery_order_office_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOfficeRequest.ProtoReflect.Descriptor instead.
func (*DeleteOfficeRequest) Descriptor() ([]byte, []int) {
	return file_submodule_food_delivery_order_office_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteOfficeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteOfficeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteOfficeRes) Reset() {
	*x = DeleteOfficeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_food_delivery_order_office_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOfficeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOfficeRes) ProtoMessage() {}

func (x *DeleteOfficeRes) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_food_delivery_order_office_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOfficeRes.ProtoReflect.Descriptor instead.
func (*DeleteOfficeRes) Descriptor() ([]byte, []int) {
	return file_submodule_food_delivery_order_office_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteOfficeRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_submodule_food_delivery_order_office_proto protoreflect.FileDescriptor

var file_submodule_food_delivery_order_office_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2d, 0x66, 0x6f, 0x6f, 0x64,
	0x2d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0x80, 0x01, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x22, 0x3c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x12, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x22, 0x6b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x3d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x73, 0x22, 0x25,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x83, 0x03, 0x0a, 0x0d, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x12, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_submodule_food_delivery_order_office_proto_rawDescOnce sync.Once
	file_submodule_food_delivery_order_office_proto_rawDescData = file_submodule_food_delivery_order_office_proto_rawDesc
)

func file_submodule_food_delivery_order_office_proto_rawDescGZIP() []byte {
	file_submodule_food_delivery_order_office_proto_rawDescOnce.Do(func() {
		file_submodule_food_delivery_order_office_proto_rawDescData = protoimpl.X.CompressGZIP(file_submodule_food_delivery_order_office_proto_rawDescData)
	})
	return file_submodule_food_delivery_order_office_proto_rawDescData
}

var file_submodule_food_delivery_order_office_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_submodule_food_delivery_order_office_proto_goTypes = []any{
	(*Office)(nil),              // 0: order.Office
	(*OfficeRequest)(nil),       // 1: order.OfficeRequest
	(*CreateOfficeRequest)(nil), // 2: order.CreateOfficeRequest
	(*UpdateOfficeRequest)(nil), // 3: order.UpdateOfficeRequest
	(*PatchOfficeRequest)(nil),  // 4: order.PatchOfficeRequest
	(*GetOfficesRequest)(nil),   // 5: order.GetOfficesRequest
	(*GetOfficesResponse)(nil),  // 6: order.GetOfficesResponse
	(*DeleteOfficeRequest)(nil), // 7: order.DeleteOfficeRequest
	(*DeleteOfficeRes)(nil),     // 8: order.DeleteOfficeRes
}
var file_submodule_food_delivery_order_office_proto_depIdxs = []int32{
	0, // 0: order.CreateOfficeRequest.office:type_name -> order.Office
	0, // 1: order.UpdateOfficeRequest.office:type_name -> order.Office
	0, // 2: order.GetOfficesResponse.offices:type_name -> order.Office
	2, // 3: order.OfficeService.CreateOffice:input_type -> order.CreateOfficeRequest
	1, // 4: order.OfficeService.GetOffice:input_type -> order.OfficeRequest
	3, // 5: order.OfficeService.UpdateOffice:input_type -> order.UpdateOfficeRequest
	4, // 6: order.OfficeService.PatchOffice:input_type -> order.PatchOfficeRequest
	7, // 7: order.OfficeService.DeleteOffice:input_type -> order.DeleteOfficeRequest
	5, // 8: order.OfficeService.GetOffices:input_type -> order.GetOfficesRequest
	0, // 9: order.OfficeService.CreateOffice:output_type -> order.Office
	0, // 10: order.OfficeService.GetOffice:output_type -> order.Office
	0, // 11: order.OfficeService.UpdateOffice:output_type -> order.Office
	0, // 12: order.OfficeService.PatchOffice:output_type -> order.Office
	8, // 13: order.OfficeService.DeleteOffice:output_type -> order.DeleteOfficeRes
	6, // 14: order.OfficeService.GetOffices:output_type -> order.GetOfficesResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_submodule_food_delivery_order_office_proto_init() }
func file_submodule_food_delivery_order_office_proto_init() {
	if File_submodule_food_delivery_order_office_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_submodule_food_delivery_order_office_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Office); i {
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
		file_submodule_food_delivery_order_office_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OfficeRequest); i {
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
		file_submodule_food_delivery_order_office_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOfficeRequest); i {
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
		file_submodule_food_delivery_order_office_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateOfficeRequest); i {
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
		file_submodule_food_delivery_order_office_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PatchOfficeRequest); i {
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
		file_submodule_food_delivery_order_office_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetOfficesRequest); i {
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
		file_submodule_food_delivery_order_office_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetOfficesResponse); i {
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
		file_submodule_food_delivery_order_office_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteOfficeRequest); i {
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
		file_submodule_food_delivery_order_office_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteOfficeRes); i {
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
			RawDescriptor: file_submodule_food_delivery_order_office_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_submodule_food_delivery_order_office_proto_goTypes,
		DependencyIndexes: file_submodule_food_delivery_order_office_proto_depIdxs,
		MessageInfos:      file_submodule_food_delivery_order_office_proto_msgTypes,
	}.Build()
	File_submodule_food_delivery_order_office_proto = out.File
	file_submodule_food_delivery_order_office_proto_rawDesc = nil
	file_submodule_food_delivery_order_office_proto_goTypes = nil
	file_submodule_food_delivery_order_office_proto_depIdxs = nil
}