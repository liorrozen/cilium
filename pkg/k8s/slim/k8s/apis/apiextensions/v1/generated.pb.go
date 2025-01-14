// Copyright 2017-2021 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/apiextensions/v1/generated.proto

package v1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *CustomResourceDefinition) Reset()      { *m = CustomResourceDefinition{} }
func (*CustomResourceDefinition) ProtoMessage() {}
func (*CustomResourceDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ae25e910fba1c55, []int{0}
}
func (m *CustomResourceDefinition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustomResourceDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CustomResourceDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomResourceDefinition.Merge(m, src)
}
func (m *CustomResourceDefinition) XXX_Size() int {
	return m.Size()
}
func (m *CustomResourceDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomResourceDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_CustomResourceDefinition proto.InternalMessageInfo

func (m *CustomResourceDefinitionList) Reset()      { *m = CustomResourceDefinitionList{} }
func (*CustomResourceDefinitionList) ProtoMessage() {}
func (*CustomResourceDefinitionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ae25e910fba1c55, []int{1}
}
func (m *CustomResourceDefinitionList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustomResourceDefinitionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CustomResourceDefinitionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomResourceDefinitionList.Merge(m, src)
}
func (m *CustomResourceDefinitionList) XXX_Size() int {
	return m.Size()
}
func (m *CustomResourceDefinitionList) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomResourceDefinitionList.DiscardUnknown(m)
}

var xxx_messageInfo_CustomResourceDefinitionList proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CustomResourceDefinition)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.apis.apiextensions.v1.CustomResourceDefinition")
	proto.RegisterType((*CustomResourceDefinitionList)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.apis.apiextensions.v1.CustomResourceDefinitionList")
}

func init() {
	proto.RegisterFile("github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/apiextensions/v1/generated.proto", fileDescriptor_2ae25e910fba1c55)
}

var fileDescriptor_2ae25e910fba1c55 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0x87, 0x33, 0x5e, 0x04, 0x89, 0x5c, 0xb8, 0x64, 0x25, 0x72, 0x19, 0xc5, 0x95, 0x9b, 0x4e,
	0x88, 0x8b, 0xe2, 0xae, 0x25, 0x2d, 0x85, 0x42, 0x8b, 0xe0, 0xae, 0xdd, 0x8d, 0x71, 0x8c, 0xd3,
	0x38, 0x33, 0x21, 0x33, 0x09, 0xed, 0xa6, 0xf4, 0x09, 0xa4, 0x8f, 0xe5, 0xd2, 0xa5, 0x2b, 0xa9,
	0xe9, 0x8b, 0x94, 0x19, 0xff, 0xb5, 0xb5, 0x52, 0x6c, 0x37, 0x99, 0x33, 0x87, 0x9c, 0xef, 0xf7,
	0x9d, 0x10, 0xbb, 0x13, 0x52, 0x35, 0x4c, 0x7b, 0x28, 0x10, 0xcc, 0x0d, 0xe8, 0x88, 0xa6, 0x9b,
	0x23, 0x8e, 0x42, 0x37, 0x6a, 0x4b, 0x57, 0x8e, 0x28, 0x33, 0x05, 0x8e, 0xa9, 0x79, 0x90, 0x7b,
	0x45, 0xb8, 0xa4, 0x82, 0x4b, 0x37, 0xf3, 0xdc, 0x90, 0x70, 0x92, 0x60, 0x45, 0xfa, 0x28, 0x4e,
	0x84, 0x12, 0xce, 0xc9, 0x16, 0x88, 0x96, 0xa4, 0xf5, 0x11, 0x47, 0x21, 0x8a, 0xda, 0x12, 0x69,
	0xa0, 0x29, 0x34, 0x10, 0x7d, 0x00, 0xa2, 0xcc, 0xab, 0x5e, 0x1c, 0x68, 0xc4, 0x88, 0xc2, 0x5f,
	0x88, 0x54, 0x8f, 0xde, 0x71, 0x42, 0x11, 0x0a, 0xd7, 0xb4, 0x7b, 0xe9, 0xc0, 0xdc, 0xcc, 0xc5,
	0x54, 0xab, 0xd7, 0x35, 0x10, 0x51, 0xa1, 0x99, 0x0c, 0x07, 0x43, 0xca, 0x49, 0xf2, 0x60, 0x12,
	0x93, 0x94, 0x2b, 0xca, 0xc8, 0x0e, 0xff, 0xf8, 0xbb, 0x01, 0x19, 0x0c, 0x09, 0xc3, 0x9f, 0xe7,
	0x1a, 0x63, 0x60, 0x57, 0xce, 0x52, 0xa9, 0x04, 0xeb, 0x12, 0x29, 0xd2, 0x24, 0x20, 0xe7, 0x64,
	0x40, 0x39, 0x55, 0x54, 0x70, 0x27, 0xb1, 0x4b, 0x7a, 0x9f, 0x3e, 0x56, 0xb8, 0x02, 0xea, 0xa0,
	0x59, 0x6e, 0xf9, 0xe8, 0xc0, 0x0f, 0xaa, 0xe7, 0x51, 0xe6, 0xa1, 0x4e, 0xef, 0x8e, 0x04, 0xea,
	0x9a, 0x28, 0xec, 0x3b, 0x93, 0x79, 0xcd, 0xca, 0xe7, 0x35, 0x7b, 0xdb, 0xeb, 0x6e, 0x72, 0x1a,
	0xe3, 0x82, 0xfd, 0x7f, 0x9f, 0xd0, 0x15, 0x95, 0xca, 0xe1, 0x3b, 0x52, 0xa7, 0x3f, 0x95, 0xd2,
	0x3c, 0xa3, 0xf4, 0x6f, 0xa5, 0x54, 0x5a, 0x77, 0xb6, 0x42, 0xce, 0xa3, 0x5d, 0xa4, 0x8a, 0x30,
	0x59, 0x29, 0xd4, 0xff, 0x34, 0xcb, 0xad, 0x1b, 0xf4, 0xcb, 0x5f, 0x0a, 0xed, 0xdb, 0xce, 0xff,
	0xbb, 0xb2, 0x28, 0x5e, 0xea, 0xbc, 0xee, 0x32, 0xd6, 0x6f, 0x4e, 0x16, 0xd0, 0x9a, 0x2e, 0xa0,
	0x35, 0x5b, 0x40, 0xeb, 0x29, 0x87, 0x60, 0x92, 0x43, 0x30, 0xcd, 0x21, 0x98, 0xe5, 0x10, 0xbc,
	0xe4, 0x10, 0x3c, 0xbf, 0x42, 0xeb, 0xb6, 0x90, 0x79, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81,
	0x0e, 0xc6, 0x6f, 0x3e, 0x03, 0x00, 0x00,
}

func (m *CustomResourceDefinition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustomResourceDefinition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustomResourceDefinition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CustomResourceDefinitionList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustomResourceDefinitionList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustomResourceDefinitionList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CustomResourceDefinition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *CustomResourceDefinitionList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CustomResourceDefinition) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CustomResourceDefinition{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CustomResourceDefinitionList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]CustomResourceDefinition{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "CustomResourceDefinition", "CustomResourceDefinition", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&CustomResourceDefinitionList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CustomResourceDefinition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CustomResourceDefinition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustomResourceDefinition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CustomResourceDefinitionList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CustomResourceDefinitionList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustomResourceDefinitionList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, CustomResourceDefinition{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
