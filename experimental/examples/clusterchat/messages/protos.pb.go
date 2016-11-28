// Code generated by protoc-gen-gogo.
// source: protos.proto
// DO NOT EDIT!

/*
	Package messages is a generated protocol buffer package.

	It is generated from these files:
		protos.proto

	It has these top-level messages:
		ConnectRequest
		ConnectResponse
		Unit
		SayRequest
		NickRequest
*/
package messages

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import actor "github.com/AsynkronIT/gam/actor"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ConnectRequest struct {
	ClientStreamPID *actor.PID `protobuf:"bytes,1,opt,name=ClientStreamPID,json=clientStreamPID" json:"ClientStreamPID,omitempty"`
}

func (m *ConnectRequest) Reset()                    { *m = ConnectRequest{} }
func (*ConnectRequest) ProtoMessage()               {}
func (*ConnectRequest) Descriptor() ([]byte, []int) { return fileDescriptorProtos, []int{0} }

func (m *ConnectRequest) GetClientStreamPID() *actor.PID {
	if m != nil {
		return m.ClientStreamPID
	}
	return nil
}

type ConnectResponse struct {
	Message string `protobuf:"bytes,1,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
}

func (m *ConnectResponse) Reset()                    { *m = ConnectResponse{} }
func (*ConnectResponse) ProtoMessage()               {}
func (*ConnectResponse) Descriptor() ([]byte, []int) { return fileDescriptorProtos, []int{1} }

type Unit struct {
}

func (m *Unit) Reset()                    { *m = Unit{} }
func (*Unit) ProtoMessage()               {}
func (*Unit) Descriptor() ([]byte, []int) { return fileDescriptorProtos, []int{2} }

type SayRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=UserName,json=userName,proto3" json:"UserName,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
}

func (m *SayRequest) Reset()                    { *m = SayRequest{} }
func (*SayRequest) ProtoMessage()               {}
func (*SayRequest) Descriptor() ([]byte, []int) { return fileDescriptorProtos, []int{3} }

type NickRequest struct {
	OldUserName string `protobuf:"bytes,1,opt,name=OldUserName,json=oldUserName,proto3" json:"OldUserName,omitempty"`
	NewUserName string `protobuf:"bytes,2,opt,name=NewUserName,json=newUserName,proto3" json:"NewUserName,omitempty"`
}

func (m *NickRequest) Reset()                    { *m = NickRequest{} }
func (*NickRequest) ProtoMessage()               {}
func (*NickRequest) Descriptor() ([]byte, []int) { return fileDescriptorProtos, []int{4} }

func init() {
	proto.RegisterType((*ConnectRequest)(nil), "messages.ConnectRequest")
	proto.RegisterType((*ConnectResponse)(nil), "messages.ConnectResponse")
	proto.RegisterType((*Unit)(nil), "messages.Unit")
	proto.RegisterType((*SayRequest)(nil), "messages.SayRequest")
	proto.RegisterType((*NickRequest)(nil), "messages.NickRequest")
}
func (this *ConnectRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ConnectRequest)
	if !ok {
		that2, ok := that.(ConnectRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.ClientStreamPID.Equal(that1.ClientStreamPID) {
		return false
	}
	return true
}
func (this *ConnectResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ConnectResponse)
	if !ok {
		that2, ok := that.(ConnectResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *Unit) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Unit)
	if !ok {
		that2, ok := that.(Unit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *SayRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SayRequest)
	if !ok {
		that2, ok := that.(SayRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.UserName != that1.UserName {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *NickRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NickRequest)
	if !ok {
		that2, ok := that.(NickRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.OldUserName != that1.OldUserName {
		return false
	}
	if this.NewUserName != that1.NewUserName {
		return false
	}
	return true
}
func (this *ConnectRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&messages.ConnectRequest{")
	if this.ClientStreamPID != nil {
		s = append(s, "ClientStreamPID: "+fmt.Sprintf("%#v", this.ClientStreamPID)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ConnectResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&messages.ConnectResponse{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Unit) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&messages.Unit{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SayRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&messages.SayRequest{")
	s = append(s, "UserName: "+fmt.Sprintf("%#v", this.UserName)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NickRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&messages.NickRequest{")
	s = append(s, "OldUserName: "+fmt.Sprintf("%#v", this.OldUserName)+",\n")
	s = append(s, "NewUserName: "+fmt.Sprintf("%#v", this.NewUserName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProtos(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringProtos(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func (m *ConnectRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ClientStreamPID != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(m.ClientStreamPID.Size()))
		n1, err := m.ClientStreamPID.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *ConnectResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *Unit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SayRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.UserName)))
		i += copy(dAtA[i:], m.UserName)
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *NickRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NickRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.OldUserName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.OldUserName)))
		i += copy(dAtA[i:], m.OldUserName)
	}
	if len(m.NewUserName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.NewUserName)))
		i += copy(dAtA[i:], m.NewUserName)
	}
	return i, nil
}

func encodeFixed64Protos(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Protos(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintProtos(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ConnectRequest) Size() (n int) {
	var l int
	_ = l
	if m.ClientStreamPID != nil {
		l = m.ClientStreamPID.Size()
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func (m *ConnectResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func (m *Unit) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *SayRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func (m *NickRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.OldUserName)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	l = len(m.NewUserName)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func sovProtos(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtos(x uint64) (n int) {
	return sovProtos(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ConnectRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConnectRequest{`,
		`ClientStreamPID:` + strings.Replace(fmt.Sprintf("%v", this.ClientStreamPID), "PID", "actor.PID", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConnectResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConnectResponse{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Unit) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Unit{`,
		`}`,
	}, "")
	return s
}
func (this *SayRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SayRequest{`,
		`UserName:` + fmt.Sprintf("%v", this.UserName) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NickRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NickRequest{`,
		`OldUserName:` + fmt.Sprintf("%v", this.OldUserName) + `,`,
		`NewUserName:` + fmt.Sprintf("%v", this.NewUserName) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProtos(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ConnectRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConnectRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientStreamPID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClientStreamPID == nil {
				m.ClientStreamPID = &actor.PID{}
			}
			if err := m.ClientStreamPID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
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
func (m *ConnectResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConnectResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
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
func (m *Unit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Unit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
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
func (m *SayRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SayRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
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
func (m *NickRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NickRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NickRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldUserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OldUserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewUserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewUserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
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
func skipProtos(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtos
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
					return 0, ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtos
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProtos
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtos
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProtos(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProtos = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtos   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protos.proto", fileDescriptorProtos) }

var fileDescriptorProtos = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x3b, 0xf7, 0x12, 0xe0, 0x9e, 0xde, 0x40, 0x32, 0xd1, 0x04, 0xbb, 0x98, 0x90, 0xae,
	0x4c, 0xc4, 0x36, 0x41, 0x1f, 0x40, 0x81, 0x98, 0xb0, 0x10, 0x91, 0xca, 0x03, 0x0c, 0x75, 0x02,
	0x0d, 0x74, 0x06, 0x3b, 0x83, 0x86, 0x9d, 0x8f, 0xe0, 0x63, 0xb8, 0xf2, 0x39, 0x5c, 0xb2, 0x74,
	0x29, 0xe3, 0xc6, 0x25, 0x8f, 0x60, 0x28, 0x94, 0x82, 0xba, 0x6a, 0x7a, 0xce, 0xff, 0x7f, 0x3d,
	0xff, 0x5f, 0xf8, 0x3f, 0x8e, 0x84, 0x12, 0xd2, 0x89, 0x1f, 0x38, 0x1f, 0x32, 0x29, 0x69, 0x9f,
	0x49, 0xab, 0xd2, 0x0f, 0xd4, 0x60, 0xd2, 0x73, 0x7c, 0x11, 0xba, 0xe7, 0x72, 0xca, 0x87, 0x91,
	0xe0, 0xcd, 0x1b, 0xb7, 0x4f, 0x43, 0x97, 0xfa, 0x4a, 0x44, 0xee, 0xb6, 0xcf, 0xbe, 0x80, 0x42,
	0x5d, 0x70, 0xce, 0x7c, 0xd5, 0x61, 0x77, 0x13, 0x26, 0x15, 0x3e, 0x85, 0x62, 0x7d, 0x14, 0x30,
	0xae, 0x3c, 0x15, 0x31, 0x1a, 0xb6, 0x9b, 0x8d, 0x12, 0x2a, 0xa3, 0x43, 0xb3, 0x0a, 0x4e, 0xec,
	0x77, 0xda, 0xcd, 0x46, 0xa7, 0xe8, 0xef, 0x4a, 0xec, 0x23, 0x28, 0x6e, 0x38, 0x72, 0x2c, 0xb8,
	0x64, 0xb8, 0x04, 0xb9, 0xcb, 0xd5, 0x51, 0x31, 0xe0, 0x5f, 0x27, 0xb7, 0xbe, 0xd1, 0xce, 0x42,
	0xa6, 0xcb, 0x03, 0x65, 0xd7, 0x00, 0x3c, 0x3a, 0x4d, 0x3e, 0x6c, 0x41, 0xbe, 0x2b, 0x59, 0xd4,
	0xa2, 0x61, 0x62, 0xc8, 0x4f, 0xd6, 0xef, 0xdb, 0xac, 0x3f, 0xbb, 0xac, 0x6b, 0x30, 0x5b, 0x81,
	0x3f, 0x4c, 0x20, 0x65, 0x30, 0xaf, 0x46, 0xb7, 0xdf, 0x38, 0xa6, 0x48, 0x47, 0x4b, 0x45, 0x8b,
	0x3d, 0x6c, 0x14, 0x2b, 0x9c, 0xc9, 0xd3, 0x51, 0xf5, 0x05, 0x01, 0xd4, 0x07, 0x54, 0x79, 0x2c,
	0xba, 0x67, 0x11, 0x3e, 0x83, 0xdc, 0x3a, 0x1a, 0x2e, 0x39, 0x49, 0xcd, 0xce, 0x6e, 0x6b, 0xd6,
	0xc1, 0x2f, 0x9b, 0x55, 0x0f, 0xb6, 0x81, 0x8f, 0xe1, 0xaf, 0x47, 0xa7, 0x78, 0x2f, 0xd5, 0xa4,
	0xb1, 0xad, 0x42, 0x3a, 0x8d, 0x4b, 0x31, 0xb0, 0x0b, 0x99, 0x65, 0x24, 0xbc, 0x9f, 0x6e, 0xb6,
	0x22, 0xfe, 0x34, 0xd4, 0x2a, 0xb3, 0x39, 0x31, 0xde, 0xe6, 0xc4, 0x58, 0xcc, 0x09, 0x7a, 0xd4,
	0x04, 0x3d, 0x6b, 0x82, 0x5e, 0x35, 0x41, 0x33, 0x4d, 0xd0, 0xbb, 0x26, 0xe8, 0x53, 0x13, 0x63,
	0xa1, 0x09, 0x7a, 0xfa, 0x20, 0x46, 0x2f, 0x1b, 0xff, 0xf9, 0x93, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x10, 0x16, 0x32, 0x93, 0x41, 0x02, 0x00, 0x00,
}
