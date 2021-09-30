// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: testing.proto

package testdata

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestEnum int32

const (
	TestEnum_TestEnum0 TestEnum = 0
	TestEnum_TestEnum1 TestEnum = 1
)

// Enum value maps for TestEnum.
var (
	TestEnum_name = map[int32]string{
		0: "TestEnum0",
		1: "TestEnum1",
	}
	TestEnum_value = map[string]int32{
		"TestEnum0": 0,
		"TestEnum1": 1,
	}
)

func (x TestEnum) Enum() *TestEnum {
	p := new(TestEnum)
	*p = x
	return p
}

func (x TestEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_testing_proto_enumTypes[0].Descriptor()
}

func (TestEnum) Type() protoreflect.EnumType {
	return &file_testing_proto_enumTypes[0]
}

func (x TestEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TestEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TestEnum(num)
	return nil
}

// Deprecated: Use TestEnum.Descriptor instead.
func (TestEnum) EnumDescriptor() ([]byte, []int) {
	return file_testing_proto_rawDescGZIP(), []int{0}
}

type AllSupportedTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32Value   *int32    `protobuf:"varint,1,opt,name=int32_value,json=int32Value" json:"int32_value,omitempty"`
	Int64Value   *int64    `protobuf:"varint,2,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	Uint32Value  *uint32   `protobuf:"varint,3,opt,name=uint32_value,json=uint32Value" json:"uint32_value,omitempty"`
	Uint64Value  *uint64   `protobuf:"varint,4,opt,name=uint64_value,json=uint64Value" json:"uint64_value,omitempty"`
	FloatValue   *float32  `protobuf:"fixed32,5,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	DoubleValue  *float64  `protobuf:"fixed64,6,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
	BoolValue    *bool     `protobuf:"varint,7,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	EnumValue    *TestEnum `protobuf:"varint,8,opt,name=enum_value,json=enumValue,enum=testdata.TestEnum" json:"enum_value,omitempty"`
	StringValue  *string   `protobuf:"bytes,9,req,name=string_value,json=stringValue" json:"string_value,omitempty"`
	Fixed64Value *uint64   `protobuf:"fixed64,10,opt,name=fixed64_value,json=fixed64Value" json:"fixed64_value,omitempty"`
}

func (x *AllSupportedTypes) Reset() {
	*x = AllSupportedTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllSupportedTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllSupportedTypes) ProtoMessage() {}

func (x *AllSupportedTypes) ProtoReflect() protoreflect.Message {
	mi := &file_testing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllSupportedTypes.ProtoReflect.Descriptor instead.
func (*AllSupportedTypes) Descriptor() ([]byte, []int) {
	return file_testing_proto_rawDescGZIP(), []int{0}
}

func (x *AllSupportedTypes) GetInt32Value() int32 {
	if x != nil && x.Int32Value != nil {
		return *x.Int32Value
	}
	return 0
}

func (x *AllSupportedTypes) GetInt64Value() int64 {
	if x != nil && x.Int64Value != nil {
		return *x.Int64Value
	}
	return 0
}

func (x *AllSupportedTypes) GetUint32Value() uint32 {
	if x != nil && x.Uint32Value != nil {
		return *x.Uint32Value
	}
	return 0
}

func (x *AllSupportedTypes) GetUint64Value() uint64 {
	if x != nil && x.Uint64Value != nil {
		return *x.Uint64Value
	}
	return 0
}

func (x *AllSupportedTypes) GetFloatValue() float32 {
	if x != nil && x.FloatValue != nil {
		return *x.FloatValue
	}
	return 0
}

func (x *AllSupportedTypes) GetDoubleValue() float64 {
	if x != nil && x.DoubleValue != nil {
		return *x.DoubleValue
	}
	return 0
}

func (x *AllSupportedTypes) GetBoolValue() bool {
	if x != nil && x.BoolValue != nil {
		return *x.BoolValue
	}
	return false
}

func (x *AllSupportedTypes) GetEnumValue() TestEnum {
	if x != nil && x.EnumValue != nil {
		return *x.EnumValue
	}
	return TestEnum_TestEnum0
}

func (x *AllSupportedTypes) GetStringValue() string {
	if x != nil && x.StringValue != nil {
		return *x.StringValue
	}
	return ""
}

func (x *AllSupportedTypes) GetFixed64Value() uint64 {
	if x != nil && x.Fixed64Value != nil {
		return *x.Fixed64Value
	}
	return 0
}

type WithWellKnownTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int64Value    *int64                    `protobuf:"varint,1,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	WrappedInt64  *wrapperspb.Int64Value    `protobuf:"bytes,2,opt,name=wrapped_int64,json=wrappedInt64" json:"wrapped_int64,omitempty"`
	StringValue   []string                  `protobuf:"bytes,3,rep,name=string_value,json=stringValue" json:"string_value,omitempty"`
	WrappedString []*wrapperspb.StringValue `protobuf:"bytes,4,rep,name=wrapped_string,json=wrappedString" json:"wrapped_string,omitempty"`
}

func (x *WithWellKnownTypes) Reset() {
	*x = WithWellKnownTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithWellKnownTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithWellKnownTypes) ProtoMessage() {}

func (x *WithWellKnownTypes) ProtoReflect() protoreflect.Message {
	mi := &file_testing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithWellKnownTypes.ProtoReflect.Descriptor instead.
func (*WithWellKnownTypes) Descriptor() ([]byte, []int) {
	return file_testing_proto_rawDescGZIP(), []int{1}
}

func (x *WithWellKnownTypes) GetInt64Value() int64 {
	if x != nil && x.Int64Value != nil {
		return *x.Int64Value
	}
	return 0
}

func (x *WithWellKnownTypes) GetWrappedInt64() *wrapperspb.Int64Value {
	if x != nil {
		return x.WrappedInt64
	}
	return nil
}

func (x *WithWellKnownTypes) GetStringValue() []string {
	if x != nil {
		return x.StringValue
	}
	return nil
}

func (x *WithWellKnownTypes) GetWrappedString() []*wrapperspb.StringValue {
	if x != nil {
		return x.WrappedString
	}
	return nil
}

type InnerType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (x *InnerType) Reset() {
	*x = InnerType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerType) ProtoMessage() {}

func (x *InnerType) ProtoReflect() protoreflect.Message {
	mi := &file_testing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerType.ProtoReflect.Descriptor instead.
func (*InnerType) Descriptor() ([]byte, []int) {
	return file_testing_proto_rawDescGZIP(), []int{2}
}

func (x *InnerType) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type NestedType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InnerType []*InnerType `protobuf:"bytes,1,rep,name=inner_type,json=innerType" json:"inner_type,omitempty"`
}

func (x *NestedType) Reset() {
	*x = NestedType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedType) ProtoMessage() {}

func (x *NestedType) ProtoReflect() protoreflect.Message {
	mi := &file_testing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedType.ProtoReflect.Descriptor instead.
func (*NestedType) Descriptor() ([]byte, []int) {
	return file_testing_proto_rawDescGZIP(), []int{3}
}

func (x *NestedType) GetInnerType() []*InnerType {
	if x != nil {
		return x.InnerType
	}
	return nil
}

type ComplexType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NestedRepeatedType []*NestedType `protobuf:"bytes,1,rep,name=nested_repeated_type,json=nestedRepeatedType" json:"nested_repeated_type,omitempty"`
	InnerType          *InnerType    `protobuf:"bytes,2,opt,name=inner_type,json=innerType" json:"inner_type,omitempty"`
}

func (x *ComplexType) Reset() {
	*x = ComplexType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexType) ProtoMessage() {}

func (x *ComplexType) ProtoReflect() protoreflect.Message {
	mi := &file_testing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexType.ProtoReflect.Descriptor instead.
func (*ComplexType) Descriptor() ([]byte, []int) {
	return file_testing_proto_rawDescGZIP(), []int{4}
}

func (x *ComplexType) GetNestedRepeatedType() []*NestedType {
	if x != nil {
		return x.NestedRepeatedType
	}
	return nil
}

func (x *ComplexType) GetInnerType() *InnerType {
	if x != nil {
		return x.InnerType
	}
	return nil
}

type ContainsRecursive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field *RecursiveType `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
}

func (x *ContainsRecursive) Reset() {
	*x = ContainsRecursive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainsRecursive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainsRecursive) ProtoMessage() {}

func (x *ContainsRecursive) ProtoReflect() protoreflect.Message {
	mi := &file_testing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainsRecursive.ProtoReflect.Descriptor instead.
func (*ContainsRecursive) Descriptor() ([]byte, []int) {
	return file_testing_proto_rawDescGZIP(), []int{5}
}

func (x *ContainsRecursive) GetField() *RecursiveType {
	if x != nil {
		return x.Field
	}
	return nil
}

type RecursiveType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field *ContainsRecursive `protobuf:"bytes,2,opt,name=field" json:"field,omitempty"`
}

func (x *RecursiveType) Reset() {
	*x = RecursiveType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecursiveType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecursiveType) ProtoMessage() {}

func (x *RecursiveType) ProtoReflect() protoreflect.Message {
	mi := &file_testing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecursiveType.ProtoReflect.Descriptor instead.
func (*RecursiveType) Descriptor() ([]byte, []int) {
	return file_testing_proto_rawDescGZIP(), []int{6}
}

func (x *RecursiveType) GetField() *ContainsRecursive {
	if x != nil {
		return x.Field
	}
	return nil
}

type RecursiveTypeTopMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field *RecursiveTypeTopMessage `protobuf:"bytes,2,opt,name=field" json:"field,omitempty"`
}

func (x *RecursiveTypeTopMessage) Reset() {
	*x = RecursiveTypeTopMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecursiveTypeTopMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecursiveTypeTopMessage) ProtoMessage() {}

func (x *RecursiveTypeTopMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecursiveTypeTopMessage.ProtoReflect.Descriptor instead.
func (*RecursiveTypeTopMessage) Descriptor() ([]byte, []int) {
	return file_testing_proto_rawDescGZIP(), []int{7}
}

func (x *RecursiveTypeTopMessage) GetField() *RecursiveTypeTopMessage {
	if x != nil {
		return x.Field
	}
	return nil
}

var File_testing_proto protoreflect.FileDescriptor

var file_testing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x02, 0x0a, 0x11, 0x41, 0x6c,
	0x6c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x65, 0x6e,
	0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e,
	0x75, 0x6d, 0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x12, 0x57, 0x69, 0x74, 0x68, 0x57, 0x65,
	0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x40, 0x0a,
	0x0d, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0c, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x21, 0x0a, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x0a, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x89, 0x01, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x14,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x12, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x42, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x12, 0x2d, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x42, 0x0a, 0x0d,
	0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73,
	0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x22, 0x52, 0x0a, 0x17, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x54, 0x6f, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x54, 0x6f, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x2a, 0x28, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x0d, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x30, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x10, 0x01, 0x42, 0x3d,
	0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
}

var (
	file_testing_proto_rawDescOnce sync.Once
	file_testing_proto_rawDescData = file_testing_proto_rawDesc
)

func file_testing_proto_rawDescGZIP() []byte {
	file_testing_proto_rawDescOnce.Do(func() {
		file_testing_proto_rawDescData = protoimpl.X.CompressGZIP(file_testing_proto_rawDescData)
	})
	return file_testing_proto_rawDescData
}

var file_testing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_testing_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_testing_proto_goTypes = []interface{}{
	(TestEnum)(0),                   // 0: testdata.TestEnum
	(*AllSupportedTypes)(nil),       // 1: testdata.AllSupportedTypes
	(*WithWellKnownTypes)(nil),      // 2: testdata.WithWellKnownTypes
	(*InnerType)(nil),               // 3: testdata.InnerType
	(*NestedType)(nil),              // 4: testdata.NestedType
	(*ComplexType)(nil),             // 5: testdata.ComplexType
	(*ContainsRecursive)(nil),       // 6: testdata.ContainsRecursive
	(*RecursiveType)(nil),           // 7: testdata.RecursiveType
	(*RecursiveTypeTopMessage)(nil), // 8: testdata.RecursiveTypeTopMessage
	(*wrapperspb.Int64Value)(nil),   // 9: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil),  // 10: google.protobuf.StringValue
}
var file_testing_proto_depIdxs = []int32{
	0,  // 0: testdata.AllSupportedTypes.enum_value:type_name -> testdata.TestEnum
	9,  // 1: testdata.WithWellKnownTypes.wrapped_int64:type_name -> google.protobuf.Int64Value
	10, // 2: testdata.WithWellKnownTypes.wrapped_string:type_name -> google.protobuf.StringValue
	3,  // 3: testdata.NestedType.inner_type:type_name -> testdata.InnerType
	4,  // 4: testdata.ComplexType.nested_repeated_type:type_name -> testdata.NestedType
	3,  // 5: testdata.ComplexType.inner_type:type_name -> testdata.InnerType
	7,  // 6: testdata.ContainsRecursive.field:type_name -> testdata.RecursiveType
	6,  // 7: testdata.RecursiveType.field:type_name -> testdata.ContainsRecursive
	8,  // 8: testdata.RecursiveTypeTopMessage.field:type_name -> testdata.RecursiveTypeTopMessage
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_testing_proto_init() }
func file_testing_proto_init() {
	if File_testing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllSupportedTypes); i {
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
		file_testing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithWellKnownTypes); i {
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
		file_testing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerType); i {
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
		file_testing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedType); i {
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
		file_testing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexType); i {
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
		file_testing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainsRecursive); i {
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
		file_testing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecursiveType); i {
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
		file_testing_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecursiveTypeTopMessage); i {
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
			RawDescriptor: file_testing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testing_proto_goTypes,
		DependencyIndexes: file_testing_proto_depIdxs,
		EnumInfos:         file_testing_proto_enumTypes,
		MessageInfos:      file_testing_proto_msgTypes,
	}.Build()
	File_testing_proto = out.File
	file_testing_proto_rawDesc = nil
	file_testing_proto_goTypes = nil
	file_testing_proto_depIdxs = nil
}
