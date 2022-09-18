// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/ibc_query/v1/crosschainquery.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// QueryResult
type QueryResult int32

const (
	// UNSPECIFIED
	QueryResult_QUERY_RESULT_UNSPECIFIED QueryResult = 0
	// SUCCESS
	QueryResult_QUERY_RESULT_SUCCESS QueryResult = 1
	// FAILURE
	QueryResult_QUERY_RESULT_FAILURE QueryResult = 2
	// TIMEOUT
	QueryResult_QUERY_RESULT_TIMEOUT QueryResult = 3
)

var QueryResult_name = map[int32]string{
	0: "QUERY_RESULT_UNSPECIFIED",
	1: "QUERY_RESULT_SUCCESS",
	2: "QUERY_RESULT_FAILURE",
	3: "QUERY_RESULT_TIMEOUT",
}

var QueryResult_value = map[string]int32{
	"QUERY_RESULT_UNSPECIFIED": 0,
	"QUERY_RESULT_SUCCESS":     1,
	"QUERY_RESULT_FAILURE":     2,
	"QUERY_RESULT_TIMEOUT":     3,
}

func (x QueryResult) String() string {
	return proto.EnumName(QueryResult_name, int32(x))
}

func (QueryResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bcb4ce27fb7943fa, []int{0}
}

// CrossChainQuery
type CrossChainQuery struct {
	Id                    string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path                  string        `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	LocalTimeoutHeight    *types.Height `protobuf:"bytes,3,opt,name=local_timeout_height,json=localTimeoutHeight,proto3" json:"local_timeout_height,omitempty"`
	LocalTimeoutTimestamp uint64        `protobuf:"varint,4,opt,name=local_timeout_timestamp,json=localTimeoutTimestamp,proto3" json:"local_timeout_timestamp,omitempty"`
	QueryHeight           uint64        `protobuf:"varint,5,opt,name=query_height,json=queryHeight,proto3" json:"query_height,omitempty"`
	ClientId              string        `protobuf:"bytes,6,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *CrossChainQuery) Reset()         { *m = CrossChainQuery{} }
func (m *CrossChainQuery) String() string { return proto.CompactTextString(m) }
func (*CrossChainQuery) ProtoMessage()    {}
func (*CrossChainQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcb4ce27fb7943fa, []int{0}
}
func (m *CrossChainQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CrossChainQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CrossChainQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CrossChainQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossChainQuery.Merge(m, src)
}
func (m *CrossChainQuery) XXX_Size() int {
	return m.Size()
}
func (m *CrossChainQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossChainQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CrossChainQuery proto.InternalMessageInfo

func (m *CrossChainQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CrossChainQuery) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CrossChainQuery) GetLocalTimeoutHeight() *types.Height {
	if m != nil {
		return m.LocalTimeoutHeight
	}
	return nil
}

func (m *CrossChainQuery) GetLocalTimeoutTimestamp() uint64 {
	if m != nil {
		return m.LocalTimeoutTimestamp
	}
	return 0
}

func (m *CrossChainQuery) GetQueryHeight() uint64 {
	if m != nil {
		return m.QueryHeight
	}
	return 0
}

func (m *CrossChainQuery) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

// CrossChainQueryResult
type CrossChainQueryResult struct {
	Id     string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Result QueryResult `protobuf:"varint,2,opt,name=result,proto3,enum=ibc.applications.ibc_query.v1.QueryResult" json:"result,omitempty"`
	Data   []byte      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *CrossChainQueryResult) Reset()         { *m = CrossChainQueryResult{} }
func (m *CrossChainQueryResult) String() string { return proto.CompactTextString(m) }
func (*CrossChainQueryResult) ProtoMessage()    {}
func (*CrossChainQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcb4ce27fb7943fa, []int{1}
}
func (m *CrossChainQueryResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CrossChainQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CrossChainQueryResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CrossChainQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossChainQueryResult.Merge(m, src)
}
func (m *CrossChainQueryResult) XXX_Size() int {
	return m.Size()
}
func (m *CrossChainQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossChainQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_CrossChainQueryResult proto.InternalMessageInfo

func (m *CrossChainQueryResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CrossChainQueryResult) GetResult() QueryResult {
	if m != nil {
		return m.Result
	}
	return QueryResult_QUERY_RESULT_UNSPECIFIED
}

func (m *CrossChainQueryResult) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("ibc.applications.ibc_query.v1.QueryResult", QueryResult_name, QueryResult_value)
	proto.RegisterType((*CrossChainQuery)(nil), "ibc.applications.ibc_query.v1.CrossChainQuery")
	proto.RegisterType((*CrossChainQueryResult)(nil), "ibc.applications.ibc_query.v1.CrossChainQueryResult")
}

func init() {
	proto.RegisterFile("ibc/applications/ibc_query/v1/crosschainquery.proto", fileDescriptor_bcb4ce27fb7943fa)
}

var fileDescriptor_bcb4ce27fb7943fa = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xae, 0xbb, 0x51, 0x31, 0x77, 0x1a, 0x95, 0xb5, 0x89, 0x68, 0x40, 0x28, 0x3b, 0x55, 0x93,
	0x66, 0xab, 0x2b, 0xe2, 0xc2, 0x89, 0x85, 0x4c, 0x44, 0x2a, 0x3f, 0xe6, 0x24, 0x07, 0xb8, 0x44,
	0x89, 0x13, 0x35, 0x96, 0x92, 0x3a, 0xc4, 0x4e, 0xa5, 0x9e, 0xf8, 0x17, 0xf8, 0xb3, 0x38, 0xee,
	0xc8, 0x11, 0xb5, 0x7f, 0x06, 0x17, 0x64, 0xa7, 0x43, 0x1d, 0x95, 0x38, 0xf9, 0xe9, 0xfb, 0xfc,
	0x7d, 0xef, 0xbd, 0x4f, 0x0f, 0x4e, 0x78, 0xc2, 0x48, 0x5c, 0x55, 0x05, 0x67, 0xb1, 0xe2, 0x62,
	0x2e, 0x09, 0x4f, 0x58, 0xf4, 0xb5, 0xc9, 0xea, 0x25, 0x59, 0x8c, 0x09, 0xab, 0x85, 0x94, 0x2c,
	0x8f, 0xf9, 0xdc, 0x40, 0xb8, 0xaa, 0x85, 0x12, 0xe8, 0x19, 0x4f, 0x18, 0xde, 0x16, 0xe1, 0xbf,
	0x22, 0xbc, 0x18, 0x9f, 0x3e, 0xd7, 0x9e, 0x4c, 0xd4, 0x19, 0x61, 0x05, 0xcf, 0xe6, 0xca, 0x18,
	0x99, 0xaa, 0xd5, 0x9f, 0xfd, 0x06, 0xf0, 0x91, 0xa3, 0x9d, 0x1d, 0xed, 0x7c, 0xa3, 0x75, 0xe8,
	0x08, 0x76, 0x79, 0x6a, 0x81, 0x21, 0x18, 0x1d, 0xd0, 0x2e, 0x4f, 0x11, 0x82, 0xfb, 0x55, 0xac,
	0x72, 0xab, 0x6b, 0x10, 0x53, 0xa3, 0x29, 0x3c, 0x2e, 0x04, 0x8b, 0x8b, 0x48, 0xf1, 0x32, 0x13,
	0x8d, 0x8a, 0xf2, 0x8c, 0xcf, 0x72, 0x65, 0xed, 0x0d, 0xc1, 0xa8, 0x7f, 0x79, 0xaa, 0xa7, 0xc0,
	0xba, 0x2f, 0xde, 0x74, 0x5b, 0x8c, 0xf1, 0x3b, 0xf3, 0x83, 0x22, 0xa3, 0x0b, 0x5a, 0x59, 0x8b,
	0xa1, 0x57, 0xf0, 0xf1, 0x7d, 0x37, 0xfd, 0x4a, 0x15, 0x97, 0x95, 0xb5, 0x3f, 0x04, 0xa3, 0x7d,
	0x7a, 0xb2, 0x2d, 0x0a, 0xee, 0x48, 0xf4, 0x02, 0x1e, 0x9a, 0x55, 0xef, 0xba, 0x3f, 0x30, 0x9f,
	0xfb, 0x06, 0xdb, 0x58, 0x3f, 0x81, 0x07, 0xed, 0x08, 0x11, 0x4f, 0xad, 0x9e, 0xd9, 0xe0, 0x61,
	0x0b, 0x78, 0xe9, 0xd9, 0x37, 0x78, 0xf2, 0xcf, 0xf2, 0x34, 0x93, 0x4d, 0xa1, 0x76, 0x22, 0xb8,
	0x82, 0xbd, 0xda, 0x30, 0x26, 0x84, 0xa3, 0xcb, 0x73, 0xfc, 0xdf, 0xdc, 0xf1, 0x96, 0x17, 0xdd,
	0x28, 0x75, 0x8c, 0x69, 0xac, 0x62, 0x13, 0xd1, 0x21, 0x35, 0xf5, 0xf9, 0x12, 0xf6, 0xb7, 0xdb,
	0x3e, 0x85, 0xd6, 0x4d, 0xe8, 0xd2, 0xcf, 0x11, 0x75, 0xfd, 0x70, 0x1a, 0x44, 0xe1, 0x07, 0xff,
	0x93, 0xeb, 0x78, 0xd7, 0x9e, 0xfb, 0x76, 0xd0, 0x41, 0x16, 0x3c, 0xbe, 0xc7, 0xfa, 0xa1, 0xe3,
	0xb8, 0xbe, 0x3f, 0x00, 0x3b, 0xcc, 0xf5, 0x1b, 0x6f, 0x1a, 0x52, 0x77, 0xd0, 0xdd, 0x61, 0x02,
	0xef, 0xbd, 0xfb, 0x31, 0x0c, 0x06, 0x7b, 0x57, 0xe1, 0x8f, 0x95, 0x0d, 0x6e, 0x57, 0x36, 0xf8,
	0xb5, 0xb2, 0xc1, 0xf7, 0xb5, 0xdd, 0xb9, 0x5d, 0xdb, 0x9d, 0x9f, 0x6b, 0xbb, 0xf3, 0xe5, 0xf5,
	0x8c, 0xab, 0xbc, 0x49, 0x30, 0x13, 0x25, 0x61, 0x42, 0x96, 0xc2, 0x5c, 0xe2, 0xc5, 0x4c, 0x90,
	0xc5, 0x4b, 0x52, 0x8a, 0xb4, 0x29, 0x32, 0xa9, 0x0f, 0x55, 0x92, 0xc9, 0xf8, 0x42, 0x33, 0xed,
	0x8d, 0xaa, 0x65, 0x95, 0xc9, 0xa4, 0x67, 0xee, 0x6a, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4f,
	0xd5, 0x0b, 0xc8, 0xce, 0x02, 0x00, 0x00,
}

func (m *CrossChainQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrossChainQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CrossChainQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintCrosschainquery(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x32
	}
	if m.QueryHeight != 0 {
		i = encodeVarintCrosschainquery(dAtA, i, uint64(m.QueryHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.LocalTimeoutTimestamp != 0 {
		i = encodeVarintCrosschainquery(dAtA, i, uint64(m.LocalTimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if m.LocalTimeoutHeight != nil {
		{
			size, err := m.LocalTimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCrosschainquery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintCrosschainquery(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCrosschainquery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CrossChainQueryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrossChainQueryResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CrossChainQueryResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintCrosschainquery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Result != 0 {
		i = encodeVarintCrosschainquery(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCrosschainquery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCrosschainquery(dAtA []byte, offset int, v uint64) int {
	offset -= sovCrosschainquery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CrossChainQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCrosschainquery(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovCrosschainquery(uint64(l))
	}
	if m.LocalTimeoutHeight != nil {
		l = m.LocalTimeoutHeight.Size()
		n += 1 + l + sovCrosschainquery(uint64(l))
	}
	if m.LocalTimeoutTimestamp != 0 {
		n += 1 + sovCrosschainquery(uint64(m.LocalTimeoutTimestamp))
	}
	if m.QueryHeight != 0 {
		n += 1 + sovCrosschainquery(uint64(m.QueryHeight))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovCrosschainquery(uint64(l))
	}
	return n
}

func (m *CrossChainQueryResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCrosschainquery(uint64(l))
	}
	if m.Result != 0 {
		n += 1 + sovCrosschainquery(uint64(m.Result))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCrosschainquery(uint64(l))
	}
	return n
}

func sovCrosschainquery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCrosschainquery(x uint64) (n int) {
	return sovCrosschainquery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CrossChainQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrosschainquery
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
			return fmt.Errorf("proto: CrossChainQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrossChainQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalTimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainquery
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
				return ErrInvalidLengthCrosschainquery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LocalTimeoutHeight == nil {
				m.LocalTimeoutHeight = &types.Height{}
			}
			if err := m.LocalTimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalTimeoutTimestamp", wireType)
			}
			m.LocalTimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LocalTimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryHeight", wireType)
			}
			m.QueryHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QueryHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrosschainquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrosschainquery
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
func (m *CrossChainQueryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrosschainquery
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
			return fmt.Errorf("proto: CrossChainQueryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrossChainQueryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= QueryResult(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrosschainquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrosschainquery
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
func skipCrosschainquery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCrosschainquery
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
					return 0, ErrIntOverflowCrosschainquery
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
					return 0, ErrIntOverflowCrosschainquery
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
				return 0, ErrInvalidLengthCrosschainquery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCrosschainquery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCrosschainquery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCrosschainquery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCrosschainquery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCrosschainquery = fmt.Errorf("proto: unexpected end of group")
)
