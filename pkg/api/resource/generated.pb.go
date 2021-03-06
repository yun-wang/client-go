/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/apimachinery/pkg/api/resource/generated.proto
// DO NOT EDIT!

/*
	Package resource is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/apimachinery/pkg/api/resource/generated.proto

	It has these top-level messages:
		Quantity
*/
package resource

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *Quantity) Reset()                    { *m = Quantity{} }
func (*Quantity) ProtoMessage()               {}
func (*Quantity) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func init() {
	proto.RegisterType((*Quantity)(nil), "k8s.io.client-go.pkg.api.resource.Quantity")
}

var fileDescriptorGenerated = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xed, 0x05, 0x95, 0x8c, 0x15, 0x42, 0xa8, 0x83, 0x53, 0x75, 0x42, 0x48, 0xd8, 0x12,
	0x2c, 0x15, 0x23, 0x6f, 0x00, 0x23, 0x9b, 0x13, 0x0e, 0xf7, 0x14, 0x6a, 0x5b, 0xf6, 0x79, 0xe8,
	0xd6, 0x91, 0xb1, 0x23, 0x63, 0xf3, 0x36, 0x1d, 0x3b, 0x32, 0x30, 0x90, 0xf0, 0x22, 0x48, 0x6e,
	0x23, 0xa4, 0x6e, 0xf7, 0x0d, 0xdf, 0xe9, 0xfb, 0x8b, 0xbb, 0x66, 0x1e, 0x25, 0x3a, 0xd5, 0xa4,
	0x0a, 0x82, 0x05, 0x82, 0xa8, 0x7c, 0x63, 0x94, 0xf6, 0xa8, 0x02, 0x44, 0x97, 0x42, 0x0d, 0xca,
	0x80, 0x85, 0xa0, 0x09, 0x5e, 0xa5, 0x0f, 0x8e, 0xdc, 0x78, 0x76, 0x70, 0xe4, 0xbf, 0x23, 0x7d,
	0x63, 0xa4, 0xf6, 0x28, 0x07, 0x67, 0x72, 0x6b, 0x90, 0x16, 0xa9, 0x92, 0xb5, 0x5b, 0x2a, 0xe3,
	0x8c, 0x53, 0x59, 0xad, 0xd2, 0x5b, 0xa6, 0x0c, 0xf9, 0x3a, 0xbc, 0x9c, 0xdc, 0x1f, 0x33, 0xb4,
	0xc7, 0xa5, 0xae, 0x17, 0x68, 0x21, 0xac, 0x72, 0x48, 0x22, 0x7c, 0x57, 0x68, 0x29, 0x52, 0x38,
	0xed, 0x98, 0xcd, 0x8b, 0xd1, 0x53, 0xd2, 0x96, 0x90, 0x56, 0xe3, 0xcb, 0xe2, 0x2c, 0x52, 0x40,
	0x6b, 0xae, 0xf8, 0x94, 0x5f, 0x9f, 0x3f, 0x1f, 0xe9, 0xe1, 0xe2, 0x73, 0x5b, 0xb2, 0x8f, 0xb6,
	0x64, 0x9b, 0xb6, 0x64, 0xdb, 0xb6, 0x64, 0xeb, 0xef, 0x29, 0x7b, 0xbc, 0xd9, 0x75, 0x82, 0xed,
	0x3b, 0xc1, 0xbe, 0x3a, 0xc1, 0xd6, 0xbd, 0xe0, 0xbb, 0x5e, 0xf0, 0x7d, 0x2f, 0xf8, 0x4f, 0x2f,
	0xf8, 0xe6, 0x57, 0xb0, 0x97, 0xd1, 0xb0, 0xe4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x55, 0x28, 0xd5,
	0x78, 0x22, 0x01, 0x00, 0x00,
}
