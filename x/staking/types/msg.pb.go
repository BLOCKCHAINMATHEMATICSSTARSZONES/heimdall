// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: staking/v1beta1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_maticnetwork_heimdall_types "github.com/maticnetwork/heimdall/types"
	github_com_maticnetwork_heimdall_types_common "github.com/maticnetwork/heimdall/types/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// MsgValidatorJoin defines a message to join a new validator.
type MsgValidatorJoin struct {
	From            string                                                     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ID              github_com_maticnetwork_heimdall_types.ValidatorID         `protobuf:"varint,2,opt,name=id,proto3,casttype=github.com/maticnetwork/heimdall/types.ValidatorID" json:"id,omitempty"`
	ActivationEpoch uint64                                                     `protobuf:"varint,3,opt,name=activation_epoch,json=activationEpoch,proto3" json:"activation_epoch,omitempty"`
	Amount          github_com_cosmos_cosmos_sdk_types.Int                     `protobuf:"bytes,4,opt,name=amount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount,omitempty"`
	SignerPubKey    *github_com_maticnetwork_heimdall_types_common.PubKey      `protobuf:"bytes,5,opt,name=signer_pub_key,json=signerPubKey,proto3,casttype=github.com/maticnetwork/heimdall/types/common.PubKey" json:"signer_pub_key,omitempty"`
	TxHash          github_com_maticnetwork_heimdall_types_common.HeimdallHash `protobuf:"bytes,6,opt,name=tx_hash,json=txHash,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallHash" json:"tx_hash,omitempty"`
	LogIndex        uint64                                                     `protobuf:"varint,7,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	BlockNumber     uint64                                                     `protobuf:"varint,8,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Nonce           uint64                                                     `protobuf:"varint,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *MsgValidatorJoin) Reset()         { *m = MsgValidatorJoin{} }
func (m *MsgValidatorJoin) String() string { return proto.CompactTextString(m) }
func (*MsgValidatorJoin) ProtoMessage()    {}
func (*MsgValidatorJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_2118b48eedb93a64, []int{0}
}
func (m *MsgValidatorJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgValidatorJoin.Unmarshal(m, b)
}
func (m *MsgValidatorJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgValidatorJoin.Marshal(b, m, deterministic)
}
func (m *MsgValidatorJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgValidatorJoin.Merge(m, src)
}
func (m *MsgValidatorJoin) XXX_Size() int {
	return xxx_messageInfo_MsgValidatorJoin.Size(m)
}
func (m *MsgValidatorJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgValidatorJoin.DiscardUnknown(m)
}

var xxx_messageInfo_MsgValidatorJoin proto.InternalMessageInfo

// MsgValidatorJoinResponse defines ValidatorJoin response type.
type MsgValidatorJoinResponse struct {
}

func (m *MsgValidatorJoinResponse) Reset()         { *m = MsgValidatorJoinResponse{} }
func (m *MsgValidatorJoinResponse) String() string { return proto.CompactTextString(m) }
func (*MsgValidatorJoinResponse) ProtoMessage()    {}
func (*MsgValidatorJoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2118b48eedb93a64, []int{1}
}
func (m *MsgValidatorJoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgValidatorJoinResponse.Unmarshal(m, b)
}
func (m *MsgValidatorJoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgValidatorJoinResponse.Marshal(b, m, deterministic)
}
func (m *MsgValidatorJoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgValidatorJoinResponse.Merge(m, src)
}
func (m *MsgValidatorJoinResponse) XXX_Size() int {
	return xxx_messageInfo_MsgValidatorJoinResponse.Size(m)
}
func (m *MsgValidatorJoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgValidatorJoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgValidatorJoinResponse proto.InternalMessageInfo

// MsgStakeUpdate defines a message to update stake for a validator.
type MsgStakeUpdate struct {
	From        string                                                     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ID          github_com_maticnetwork_heimdall_types.ValidatorID         `protobuf:"varint,2,opt,name=id,proto3,casttype=github.com/maticnetwork/heimdall/types.ValidatorID" json:"id,omitempty"`
	NewAmount   github_com_cosmos_cosmos_sdk_types.Int                     `protobuf:"bytes,3,opt,name=new_amount,json=newAmount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Int" json:"new_amount,omitempty"`
	TxHash      github_com_maticnetwork_heimdall_types_common.HeimdallHash `protobuf:"bytes,4,opt,name=tx_hash,json=txHash,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallHash" json:"tx_hash,omitempty"`
	LogIndex    uint64                                                     `protobuf:"varint,7,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	BlockNumber uint64                                                     `protobuf:"varint,8,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Nonce       uint64                                                     `protobuf:"varint,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *MsgStakeUpdate) Reset()         { *m = MsgStakeUpdate{} }
func (m *MsgStakeUpdate) String() string { return proto.CompactTextString(m) }
func (*MsgStakeUpdate) ProtoMessage()    {}
func (*MsgStakeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_2118b48eedb93a64, []int{2}
}
func (m *MsgStakeUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgStakeUpdate.Unmarshal(m, b)
}
func (m *MsgStakeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgStakeUpdate.Marshal(b, m, deterministic)
}
func (m *MsgStakeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStakeUpdate.Merge(m, src)
}
func (m *MsgStakeUpdate) XXX_Size() int {
	return xxx_messageInfo_MsgStakeUpdate.Size(m)
}
func (m *MsgStakeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStakeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStakeUpdate proto.InternalMessageInfo

// MsgStakeUpdateResponse defines StakeUpdate response type.
type MsgStakeUpdateResponse struct {
}

func (m *MsgStakeUpdateResponse) Reset()         { *m = MsgStakeUpdateResponse{} }
func (m *MsgStakeUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStakeUpdateResponse) ProtoMessage()    {}
func (*MsgStakeUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2118b48eedb93a64, []int{3}
}
func (m *MsgStakeUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgStakeUpdateResponse.Unmarshal(m, b)
}
func (m *MsgStakeUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgStakeUpdateResponse.Marshal(b, m, deterministic)
}
func (m *MsgStakeUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStakeUpdateResponse.Merge(m, src)
}
func (m *MsgStakeUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_MsgStakeUpdateResponse.Size(m)
}
func (m *MsgStakeUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStakeUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStakeUpdateResponse proto.InternalMessageInfo

// MsgSignerUpdate defines a message to update signer of a validator.
type MsgSignerUpdate struct {
	From            string                                                     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ID              github_com_maticnetwork_heimdall_types.ValidatorID         `protobuf:"varint,2,opt,name=id,proto3,casttype=github.com/maticnetwork/heimdall/types.ValidatorID" json:"id,omitempty"`
	NewSignerPubKey *github_com_maticnetwork_heimdall_types_common.PubKey      `protobuf:"bytes,3,opt,name=new_signer_pub_key,json=newSignerPubKey,proto3,casttype=github.com/maticnetwork/heimdall/types/common.PubKey" json:"new_signer_pub_key,omitempty"`
	TxHash          github_com_maticnetwork_heimdall_types_common.HeimdallHash `protobuf:"bytes,4,opt,name=tx_hash,json=txHash,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallHash" json:"tx_hash,omitempty"`
	LogIndex        uint64                                                     `protobuf:"varint,5,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	BlockNumber     uint64                                                     `protobuf:"varint,6,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Nonce           uint64                                                     `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *MsgSignerUpdate) Reset()         { *m = MsgSignerUpdate{} }
func (m *MsgSignerUpdate) String() string { return proto.CompactTextString(m) }
func (*MsgSignerUpdate) ProtoMessage()    {}
func (*MsgSignerUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_2118b48eedb93a64, []int{4}
}
func (m *MsgSignerUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSignerUpdate.Unmarshal(m, b)
}
func (m *MsgSignerUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSignerUpdate.Marshal(b, m, deterministic)
}
func (m *MsgSignerUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignerUpdate.Merge(m, src)
}
func (m *MsgSignerUpdate) XXX_Size() int {
	return xxx_messageInfo_MsgSignerUpdate.Size(m)
}
func (m *MsgSignerUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignerUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignerUpdate proto.InternalMessageInfo

// MsgSignerUpdateResponse defines SignerUpdate response type.
type MsgSignerUpdateResponse struct {
}

func (m *MsgSignerUpdateResponse) Reset()         { *m = MsgSignerUpdateResponse{} }
func (m *MsgSignerUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSignerUpdateResponse) ProtoMessage()    {}
func (*MsgSignerUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2118b48eedb93a64, []int{5}
}
func (m *MsgSignerUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSignerUpdateResponse.Unmarshal(m, b)
}
func (m *MsgSignerUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSignerUpdateResponse.Marshal(b, m, deterministic)
}
func (m *MsgSignerUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignerUpdateResponse.Merge(m, src)
}
func (m *MsgSignerUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_MsgSignerUpdateResponse.Size(m)
}
func (m *MsgSignerUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignerUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignerUpdateResponse proto.InternalMessageInfo

// MsgValidatorExit defines a message to exit as a validator
type MsgValidatorExit struct {
	From              string                                                     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ID                github_com_maticnetwork_heimdall_types.ValidatorID         `protobuf:"varint,2,opt,name=id,proto3,casttype=github.com/maticnetwork/heimdall/types.ValidatorID" json:"id,omitempty"`
	DeactivationEpoch uint64                                                     `protobuf:"varint,3,opt,name=deactivation_epoch,json=deactivationEpoch,proto3" json:"deactivation_epoch,omitempty"`
	TxHash            github_com_maticnetwork_heimdall_types_common.HeimdallHash `protobuf:"bytes,4,opt,name=tx_hash,json=txHash,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallHash" json:"tx_hash,omitempty"`
	LogIndex          uint64                                                     `protobuf:"varint,5,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	BlockNumber       uint64                                                     `protobuf:"varint,6,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Nonce             uint64                                                     `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *MsgValidatorExit) Reset()         { *m = MsgValidatorExit{} }
func (m *MsgValidatorExit) String() string { return proto.CompactTextString(m) }
func (*MsgValidatorExit) ProtoMessage()    {}
func (*MsgValidatorExit) Descriptor() ([]byte, []int) {
	return fileDescriptor_2118b48eedb93a64, []int{6}
}
func (m *MsgValidatorExit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgValidatorExit.Unmarshal(m, b)
}
func (m *MsgValidatorExit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgValidatorExit.Marshal(b, m, deterministic)
}
func (m *MsgValidatorExit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgValidatorExit.Merge(m, src)
}
func (m *MsgValidatorExit) XXX_Size() int {
	return xxx_messageInfo_MsgValidatorExit.Size(m)
}
func (m *MsgValidatorExit) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgValidatorExit.DiscardUnknown(m)
}

var xxx_messageInfo_MsgValidatorExit proto.InternalMessageInfo

// MsgValidatorExitResponse is response type for ValidatorExit RPC method
type MsgValidatorExitResponse struct {
}

func (m *MsgValidatorExitResponse) Reset()         { *m = MsgValidatorExitResponse{} }
func (m *MsgValidatorExitResponse) String() string { return proto.CompactTextString(m) }
func (*MsgValidatorExitResponse) ProtoMessage()    {}
func (*MsgValidatorExitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2118b48eedb93a64, []int{7}
}
func (m *MsgValidatorExitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgValidatorExitResponse.Unmarshal(m, b)
}
func (m *MsgValidatorExitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgValidatorExitResponse.Marshal(b, m, deterministic)
}
func (m *MsgValidatorExitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgValidatorExitResponse.Merge(m, src)
}
func (m *MsgValidatorExitResponse) XXX_Size() int {
	return xxx_messageInfo_MsgValidatorExitResponse.Size(m)
}
func (m *MsgValidatorExitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgValidatorExitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgValidatorExitResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgValidatorJoin)(nil), "heimdall.staking.v1beta1.MsgValidatorJoin")
	proto.RegisterType((*MsgValidatorJoinResponse)(nil), "heimdall.staking.v1beta1.MsgValidatorJoinResponse")
	proto.RegisterType((*MsgStakeUpdate)(nil), "heimdall.staking.v1beta1.MsgStakeUpdate")
	proto.RegisterType((*MsgStakeUpdateResponse)(nil), "heimdall.staking.v1beta1.MsgStakeUpdateResponse")
	proto.RegisterType((*MsgSignerUpdate)(nil), "heimdall.staking.v1beta1.MsgSignerUpdate")
	proto.RegisterType((*MsgSignerUpdateResponse)(nil), "heimdall.staking.v1beta1.MsgSignerUpdateResponse")
	proto.RegisterType((*MsgValidatorExit)(nil), "heimdall.staking.v1beta1.MsgValidatorExit")
	proto.RegisterType((*MsgValidatorExitResponse)(nil), "heimdall.staking.v1beta1.MsgValidatorExitResponse")
}

func init() { proto.RegisterFile("staking/v1beta1/msg.proto", fileDescriptor_2118b48eedb93a64) }

var fileDescriptor_2118b48eedb93a64 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0x75, 0x7e, 0x08, 0x64, 0xc8, 0x07, 0x7c, 0x23, 0xd4, 0x9a, 0x54, 0x4a, 0x28, 0x8b, 0x2a,
	0x20, 0x61, 0x43, 0xda, 0x45, 0x85, 0xaa, 0x4a, 0x44, 0x20, 0x11, 0x5a, 0xaa, 0xca, 0xa8, 0xad,
	0xd4, 0x8d, 0x35, 0xb6, 0x07, 0x67, 0x14, 0x7b, 0x26, 0xca, 0x4c, 0x48, 0xa2, 0xbe, 0x40, 0x57,
	0x55, 0x1f, 0xa1, 0x9b, 0xbe, 0x47, 0x97, 0x5d, 0xb2, 0x64, 0x15, 0x55, 0xe1, 0x0d, 0xba, 0xcc,
	0xaa, 0xf2, 0xd8, 0x49, 0x9d, 0x54, 0x40, 0x58, 0x80, 0xc4, 0xca, 0x9e, 0x7b, 0x4f, 0xee, 0x9d,
	0x9c, 0x73, 0xe6, 0x8e, 0xc1, 0x0a, 0x17, 0xa8, 0x4e, 0xa8, 0xab, 0x9f, 0x6e, 0x5b, 0x58, 0xa0,
	0x6d, 0xdd, 0xe7, 0xae, 0xd6, 0x68, 0x32, 0xc1, 0xa0, 0x5a, 0xc3, 0xc4, 0x77, 0x90, 0xe7, 0x69,
	0x11, 0x46, 0x8b, 0x30, 0xf9, 0x65, 0x97, 0xb9, 0x4c, 0x82, 0xf4, 0xe0, 0x2d, 0xc4, 0xe7, 0x57,
	0x5c, 0xc6, 0x5c, 0x0f, 0xeb, 0x72, 0x65, 0xb5, 0x4e, 0x74, 0x44, 0xbb, 0x61, 0x6a, 0xed, 0x4b,
	0x1a, 0x2c, 0x1d, 0x71, 0xf7, 0x3d, 0xf2, 0x88, 0x83, 0x04, 0x6b, 0x1e, 0x32, 0x42, 0x21, 0x04,
	0xe9, 0x93, 0x26, 0xf3, 0xd5, 0xc4, 0x6a, 0xa2, 0x94, 0x35, 0xe4, 0x3b, 0x7c, 0x0d, 0x92, 0xc4,
	0x51, 0x93, 0xab, 0x89, 0x52, 0xba, 0xf2, 0xa2, 0xdf, 0x2b, 0x26, 0xab, 0x7b, 0x83, 0x5e, 0xb1,
	0xec, 0x12, 0x51, 0x6b, 0x59, 0x9a, 0xcd, 0x7c, 0xdd, 0x47, 0x82, 0xd8, 0x14, 0x8b, 0x36, 0x6b,
	0xd6, 0xf5, 0xe1, 0x0e, 0x75, 0xd1, 0x6d, 0x60, 0xae, 0x8d, 0xea, 0x57, 0xf7, 0x8c, 0x24, 0x71,
	0xe0, 0x3a, 0x58, 0x42, 0xb6, 0x20, 0xa7, 0x48, 0x10, 0x46, 0x4d, 0xdc, 0x60, 0x76, 0x4d, 0x4d,
	0x05, 0xb5, 0x8d, 0xc5, 0xbf, 0xf1, 0xfd, 0x20, 0x0c, 0x2b, 0x20, 0x83, 0x7c, 0xd6, 0xa2, 0x42,
	0x4d, 0x07, 0xdb, 0xa9, 0x6c, 0x0c, 0x7a, 0xc5, 0x27, 0xb1, 0xb6, 0x36, 0xe3, 0x3e, 0xe3, 0xd1,
	0x63, 0x93, 0x3b, 0xf5, 0xa8, 0x65, 0x95, 0x0a, 0x23, 0xfa, 0x25, 0x14, 0x60, 0x81, 0x13, 0x97,
	0xe2, 0xa6, 0xd9, 0x68, 0x59, 0x66, 0x1d, 0x77, 0xd5, 0x99, 0xd5, 0x44, 0x69, 0xbe, 0xbc, 0xac,
	0x85, 0xcc, 0x68, 0x43, 0x66, 0xb4, 0x5d, 0xda, 0xad, 0x3c, 0x1f, 0xf4, 0x8a, 0xcf, 0xa6, 0xfb,
	0x63, 0xba, 0xcd, 0x7c, 0x9f, 0x51, 0xed, 0x6d, 0xcb, 0x7a, 0x85, 0xbb, 0x46, 0x2e, 0xec, 0x12,
	0xae, 0xe0, 0x07, 0x30, 0x2b, 0x3a, 0x66, 0x0d, 0xf1, 0x9a, 0x9a, 0x91, 0x5b, 0x7f, 0x39, 0xe8,
	0x15, 0x77, 0x6e, 0x56, 0xf8, 0x20, 0x8a, 0x1e, 0x20, 0x5e, 0x33, 0x32, 0xa2, 0x13, 0x3c, 0xe1,
	0x23, 0x90, 0xf5, 0x98, 0x6b, 0x12, 0xea, 0xe0, 0x8e, 0x3a, 0x2b, 0x69, 0x9b, 0xf3, 0x98, 0x5b,
	0x0d, 0xd6, 0xf0, 0x31, 0xc8, 0x59, 0x1e, 0xb3, 0xeb, 0x26, 0x6d, 0xf9, 0x16, 0x6e, 0xaa, 0x73,
	0x32, 0x3f, 0x2f, 0x63, 0x6f, 0x64, 0x08, 0x2e, 0x83, 0x19, 0xca, 0xa8, 0x8d, 0xd5, 0xac, 0xcc,
	0x85, 0x8b, 0x9d, 0xf4, 0xe7, 0x6f, 0x45, 0x65, 0x2d, 0x0f, 0xd4, 0x49, 0x3f, 0x18, 0x98, 0x37,
	0x18, 0xe5, 0x78, 0xed, 0x77, 0x12, 0x2c, 0x1c, 0x71, 0xf7, 0x58, 0xa0, 0x3a, 0x7e, 0xd7, 0x70,
	0x90, 0xc0, 0x77, 0x60, 0x95, 0x2a, 0x00, 0x14, 0xb7, 0xcd, 0xc8, 0x03, 0xa9, 0x1b, 0x7b, 0x20,
	0x4b, 0x71, 0x7b, 0x37, 0xb4, 0x41, 0x4c, 0x90, 0xc0, 0x4b, 0xb9, 0x7b, 0x21, 0x88, 0x0a, 0x1e,
	0x8c, 0x73, 0x3e, 0x92, 0xe3, 0x7b, 0x0a, 0x2c, 0x06, 0x29, 0xe9, 0xb9, 0x3b, 0xd3, 0xe3, 0x13,
	0x80, 0x81, 0x1e, 0x13, 0xe7, 0x29, 0x75, 0x2b, 0xe7, 0x69, 0x91, 0xe2, 0xf6, 0xf1, 0x25, 0x47,
	0xea, 0x16, 0x15, 0x9c, 0xb9, 0x46, 0xc1, 0xcc, 0x15, 0x0a, 0xce, 0xfe, 0xab, 0xe0, 0x0a, 0x78,
	0x38, 0x21, 0xd3, 0x48, 0xc2, 0xf3, 0xe4, 0xf8, 0xf8, 0xdd, 0xef, 0x10, 0x71, 0x07, 0x1a, 0x6e,
	0x02, 0xe8, 0xe0, 0x4b, 0x06, 0xf0, 0xff, 0xf1, 0x4c, 0x38, 0x82, 0xef, 0x25, 0xeb, 0x13, 0x83,
	0x2c, 0x60, 0x76, 0x48, 0x7b, 0xf9, 0x47, 0x0a, 0xa4, 0x8e, 0xb8, 0x0b, 0x19, 0xf8, 0x6f, 0xfc,
	0xe6, 0xdb, 0xd0, 0x2e, 0xbb, 0x5a, 0xb5, 0xc9, 0xa9, 0x98, 0x2f, 0x4f, 0x8f, 0x1d, 0x36, 0x86,
	0x04, 0xcc, 0xc7, 0xa7, 0x67, 0xe9, 0xca, 0x12, 0x31, 0x64, 0x7e, 0x6b, 0x5a, 0xe4, 0xa8, 0x95,
	0x07, 0x72, 0x63, 0x93, 0x61, 0xfd, 0xea, 0x0a, 0x31, 0x68, 0x7e, 0x7b, 0x6a, 0xe8, 0xa8, 0x5b,
	0x9c, 0x49, 0x69, 0xe2, 0x29, 0x99, 0x0c, 0xb0, 0xd3, 0x32, 0x19, 0x97, 0xb0, 0x72, 0xf8, 0xb3,
	0x5f, 0x50, 0xce, 0xfa, 0x05, 0xe5, 0x57, 0xbf, 0xa0, 0x7c, 0xbd, 0x28, 0x28, 0x67, 0x17, 0x05,
	0xe5, 0xfc, 0xa2, 0xa0, 0x7c, 0xdc, 0xba, 0xd6, 0x98, 0x1d, 0x7d, 0xf8, 0x6d, 0x25, 0x2d, 0x6a,
	0x65, 0xe4, 0xb8, 0x7a, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x06, 0x39, 0xe5, 0x73, 0x09,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ValidatorJoin defines a method to join a new validator.
	ValidatorJoin(ctx context.Context, in *MsgValidatorJoin, opts ...grpc.CallOption) (*MsgValidatorJoinResponse, error)
	// StakeUpdate defines a method to update stake for an existing validator.
	StakeUpdate(ctx context.Context, in *MsgStakeUpdate, opts ...grpc.CallOption) (*MsgStakeUpdateResponse, error)
	// SignerUpdate defines a method for update singer details of
	// exisitng validator.
	SignerUpdate(ctx context.Context, in *MsgSignerUpdate, opts ...grpc.CallOption) (*MsgSignerUpdateResponse, error)
	// ValidatorExit defines a method to handle validator exit
	ValidatorExit(ctx context.Context, in *MsgValidatorExit, opts ...grpc.CallOption) (*MsgValidatorExitResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ValidatorJoin(ctx context.Context, in *MsgValidatorJoin, opts ...grpc.CallOption) (*MsgValidatorJoinResponse, error) {
	out := new(MsgValidatorJoinResponse)
	err := c.cc.Invoke(ctx, "/heimdall.staking.v1beta1.Msg/ValidatorJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StakeUpdate(ctx context.Context, in *MsgStakeUpdate, opts ...grpc.CallOption) (*MsgStakeUpdateResponse, error) {
	out := new(MsgStakeUpdateResponse)
	err := c.cc.Invoke(ctx, "/heimdall.staking.v1beta1.Msg/StakeUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SignerUpdate(ctx context.Context, in *MsgSignerUpdate, opts ...grpc.CallOption) (*MsgSignerUpdateResponse, error) {
	out := new(MsgSignerUpdateResponse)
	err := c.cc.Invoke(ctx, "/heimdall.staking.v1beta1.Msg/SignerUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ValidatorExit(ctx context.Context, in *MsgValidatorExit, opts ...grpc.CallOption) (*MsgValidatorExitResponse, error) {
	out := new(MsgValidatorExitResponse)
	err := c.cc.Invoke(ctx, "/heimdall.staking.v1beta1.Msg/ValidatorExit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ValidatorJoin defines a method to join a new validator.
	ValidatorJoin(context.Context, *MsgValidatorJoin) (*MsgValidatorJoinResponse, error)
	// StakeUpdate defines a method to update stake for an existing validator.
	StakeUpdate(context.Context, *MsgStakeUpdate) (*MsgStakeUpdateResponse, error)
	// SignerUpdate defines a method for update singer details of
	// exisitng validator.
	SignerUpdate(context.Context, *MsgSignerUpdate) (*MsgSignerUpdateResponse, error)
	// ValidatorExit defines a method to handle validator exit
	ValidatorExit(context.Context, *MsgValidatorExit) (*MsgValidatorExitResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ValidatorJoin(ctx context.Context, req *MsgValidatorJoin) (*MsgValidatorJoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorJoin not implemented")
}
func (*UnimplementedMsgServer) StakeUpdate(ctx context.Context, req *MsgStakeUpdate) (*MsgStakeUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakeUpdate not implemented")
}
func (*UnimplementedMsgServer) SignerUpdate(ctx context.Context, req *MsgSignerUpdate) (*MsgSignerUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignerUpdate not implemented")
}
func (*UnimplementedMsgServer) ValidatorExit(ctx context.Context, req *MsgValidatorExit) (*MsgValidatorExitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorExit not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ValidatorJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgValidatorJoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ValidatorJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.staking.v1beta1.Msg/ValidatorJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ValidatorJoin(ctx, req.(*MsgValidatorJoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StakeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStakeUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StakeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.staking.v1beta1.Msg/StakeUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StakeUpdate(ctx, req.(*MsgStakeUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SignerUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSignerUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SignerUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.staking.v1beta1.Msg/SignerUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SignerUpdate(ctx, req.(*MsgSignerUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ValidatorExit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgValidatorExit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ValidatorExit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.staking.v1beta1.Msg/ValidatorExit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ValidatorExit(ctx, req.(*MsgValidatorExit))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heimdall.staking.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidatorJoin",
			Handler:    _Msg_ValidatorJoin_Handler,
		},
		{
			MethodName: "StakeUpdate",
			Handler:    _Msg_StakeUpdate_Handler,
		},
		{
			MethodName: "SignerUpdate",
			Handler:    _Msg_SignerUpdate_Handler,
		},
		{
			MethodName: "ValidatorExit",
			Handler:    _Msg_ValidatorExit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staking/v1beta1/msg.proto",
}
