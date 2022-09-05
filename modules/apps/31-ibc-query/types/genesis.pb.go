// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/ibc_query/v1/genesis.proto

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
	return fileDescriptor_c259616820459d03, []int{0}
}

// GenesisState defines the ICS31 ibc-query genesis state
type GenesisState struct {
	Queries []*CrossChainQuery       `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
	Results []*CrossChainQueryResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c259616820459d03, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetQueries() []*CrossChainQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *GenesisState) GetResults() []*CrossChainQueryResult {
	if m != nil {
		return m.Results
	}
	return nil
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
	return fileDescriptor_c259616820459d03, []int{1}
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
	return fileDescriptor_c259616820459d03, []int{2}
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
	proto.RegisterType((*GenesisState)(nil), "ibc.applications.ibc_query.v1.GenesisState")
	proto.RegisterType((*CrossChainQuery)(nil), "ibc.applications.ibc_query.v1.CrossChainQuery")
	proto.RegisterType((*CrossChainQueryResult)(nil), "ibc.applications.ibc_query.v1.CrossChainQueryResult")
}

func init() {
	proto.RegisterFile("ibc/applications/ibc_query/v1/genesis.proto", fileDescriptor_c259616820459d03)
}

var fileDescriptor_c259616820459d03 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0x33, 0x4e, 0xfe, 0xe9, 0xbf, 0x93, 0xa8, 0x44, 0xa3, 0x56, 0x58, 0x05, 0x4c, 0xc8,
	0x2a, 0x2a, 0xea, 0x58, 0x49, 0x2b, 0x36, 0xac, 0xa8, 0x71, 0xa9, 0xa5, 0x50, 0xa8, 0x3f, 0x16,
	0xb0, 0xb1, 0xfc, 0x31, 0x8a, 0x47, 0xb2, 0x33, 0xc6, 0x33, 0x8e, 0x94, 0x15, 0xaf, 0xc0, 0x63,
	0xf0, 0x28, 0x2c, 0xbb, 0x64, 0x89, 0x92, 0xc7, 0x60, 0x83, 0x3c, 0x4e, 0x50, 0x4a, 0x10, 0x88,
	0x95, 0xaf, 0xee, 0x3d, 0xe7, 0x5c, 0xcf, 0x4f, 0xba, 0xf0, 0x29, 0x0d, 0x23, 0x3d, 0xc8, 0xf3,
	0x94, 0x46, 0x81, 0xa0, 0x6c, 0xc6, 0x75, 0x1a, 0x46, 0xfe, 0x87, 0x92, 0x14, 0x0b, 0x7d, 0x3e,
	0xd2, 0xa7, 0x64, 0x46, 0x38, 0xe5, 0x38, 0x2f, 0x98, 0x60, 0xe8, 0x11, 0x0d, 0x23, 0xbc, 0x2d,
	0xc6, 0x3f, 0xc5, 0x78, 0x3e, 0x3a, 0x7e, 0x5c, 0x65, 0x45, 0xac, 0x20, 0x7a, 0x94, 0x52, 0x32,
	0x13, 0x55, 0x40, 0x5d, 0xd5, 0xfe, 0xc1, 0x67, 0x00, 0xbb, 0xaf, 0xea, 0x44, 0x47, 0x04, 0x82,
	0xa0, 0x2b, 0xb8, 0x57, 0xb9, 0x29, 0xe1, 0x2a, 0xe8, 0x37, 0x87, 0x9d, 0x31, 0xc6, 0x7f, 0x5c,
	0x81, 0x8d, 0x82, 0x71, 0x6e, 0x24, 0x01, 0x9d, 0xdd, 0x54, 0x2d, 0x7b, 0x63, 0x47, 0xd7, 0x70,
	0xaf, 0x20, 0xbc, 0x4c, 0x05, 0x57, 0x15, 0x99, 0x74, 0xfe, 0x8f, 0x49, 0xd2, 0x6c, 0x6f, 0x42,
	0x06, 0xdf, 0x01, 0xbc, 0xf7, 0x8b, 0x04, 0x1d, 0x40, 0x85, 0xc6, 0x2a, 0xe8, 0x83, 0xe1, 0xbe,
	0xad, 0xd0, 0x18, 0x21, 0xd8, 0xca, 0x03, 0x91, 0xa8, 0x8a, 0xec, 0xc8, 0x1a, 0x4d, 0xe0, 0x61,
	0xca, 0xa2, 0x20, 0xf5, 0x05, 0xcd, 0x08, 0x2b, 0x85, 0x9f, 0x10, 0x3a, 0x4d, 0x84, 0xda, 0xec,
	0x83, 0x61, 0x67, 0x7c, 0x2c, 0x7f, 0xaa, 0x42, 0x84, 0xd7, 0x60, 0xe6, 0x23, 0x7c, 0x25, 0x15,
	0x36, 0x92, 0x3e, 0xb7, 0xb6, 0xd5, 0x3d, 0xf4, 0x0c, 0xde, 0xbf, 0x9b, 0x56, 0x7d, 0xb9, 0x08,
	0xb2, 0x5c, 0x6d, 0xf5, 0xc1, 0xb0, 0x65, 0x1f, 0x6d, 0x9b, 0xdc, 0xcd, 0x10, 0x3d, 0x81, 0x5d,
	0xf9, 0xd0, 0xcd, 0xf6, 0xff, 0xa4, 0xb8, 0x23, 0x7b, 0xeb, 0xe8, 0x07, 0x70, 0xbf, 0xfe, 0x05,
	0x9f, 0xc6, 0x6a, 0x5b, 0xbe, 0xe0, 0xff, 0xba, 0x61, 0xc5, 0x83, 0x8f, 0xf0, 0xe8, 0xb7, 0x7c,
	0x76, 0x10, 0x5c, 0xc0, 0x76, 0x4d, 0x4c, 0x42, 0x38, 0x18, 0x9f, 0xfc, 0x85, 0xfa, 0x36, 0xeb,
	0xb5, 0xb3, 0xc2, 0x18, 0x07, 0x22, 0x90, 0x88, 0xba, 0xb6, 0xac, 0x4f, 0x16, 0xb0, 0xb3, 0xbd,
	0xf6, 0x21, 0x54, 0x6f, 0x3c, 0xd3, 0x7e, 0xe7, 0xdb, 0xa6, 0xe3, 0x4d, 0x5c, 0xdf, 0xbb, 0x76,
	0xde, 0x9a, 0x86, 0x75, 0x69, 0x99, 0x2f, 0x7b, 0x0d, 0xa4, 0xc2, 0xc3, 0x3b, 0x53, 0xc7, 0x33,
	0x0c, 0xd3, 0x71, 0x7a, 0x60, 0x67, 0x72, 0xf9, 0xc2, 0x9a, 0x78, 0xb6, 0xd9, 0x53, 0x76, 0x26,
	0xae, 0xf5, 0xda, 0x7c, 0xe3, 0xb9, 0xbd, 0xe6, 0x85, 0xf7, 0x65, 0xa9, 0x81, 0xdb, 0xa5, 0x06,
	0xbe, 0x2d, 0x35, 0xf0, 0x69, 0xa5, 0x35, 0x6e, 0x57, 0x5a, 0xe3, 0xeb, 0x4a, 0x6b, 0xbc, 0x7f,
	0x3e, 0xa5, 0x22, 0x29, 0x43, 0x1c, 0xb1, 0x4c, 0x8f, 0x18, 0xcf, 0x98, 0x3c, 0x96, 0xd3, 0x29,
	0xd3, 0xe7, 0xe7, 0x7a, 0xc6, 0xe2, 0x32, 0x25, 0xbc, 0xba, 0x25, 0xae, 0x9f, 0x8d, 0x4e, 0xab,
	0x49, 0x7d, 0x46, 0x62, 0x91, 0x13, 0x1e, 0xb6, 0xe5, 0x09, 0x9c, 0xfd, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0x74, 0x4f, 0x38, 0x71, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Results[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Queries) > 0 {
		for iNdEx := len(m.Queries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Queries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
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
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x32
	}
	if m.QueryHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.QueryHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.LocalTimeoutTimestamp != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LocalTimeoutTimestamp))
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
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Id)))
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
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Result != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, e := range m.Queries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *CrossChainQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LocalTimeoutHeight != nil {
		l = m.LocalTimeoutHeight.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LocalTimeoutTimestamp != 0 {
		n += 1 + sovGenesis(uint64(m.LocalTimeoutTimestamp))
	}
	if m.QueryHeight != 0 {
		n += 1 + sovGenesis(uint64(m.QueryHeight))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
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
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Result != 0 {
		n += 1 + sovGenesis(uint64(m.Result))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Queries = append(m.Queries, &CrossChainQuery{})
			if err := m.Queries[len(m.Queries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &CrossChainQueryResult{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *CrossChainQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
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
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
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
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
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
					return ErrIntOverflowGenesis
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
					return ErrIntOverflowGenesis
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
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
				return ErrIntOverflowGenesis
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
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
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
					return ErrIntOverflowGenesis
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
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
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
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
