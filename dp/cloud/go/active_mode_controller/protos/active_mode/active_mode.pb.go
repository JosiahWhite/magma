// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dp/protos/active_mode.proto

package active_mode

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CbsdState int32

const (
	CbsdState_Unregistered CbsdState = 0
	CbsdState_Registered   CbsdState = 1
)

var CbsdState_name = map[int32]string{
	0: "Unregistered",
	1: "Registered",
}

var CbsdState_value = map[string]int32{
	"Unregistered": 0,
	"Registered":   1,
}

func (x CbsdState) String() string {
	return proto.EnumName(CbsdState_name, int32(x))
}

func (CbsdState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{0}
}

type GrantState int32

const (
	GrantState_Granted    GrantState = 0
	GrantState_Authorized GrantState = 1
	GrantState_Unsync     GrantState = 2
)

var GrantState_name = map[int32]string{
	0: "Granted",
	1: "Authorized",
	2: "Unsync",
}

var GrantState_value = map[string]int32{
	"Granted":    0,
	"Authorized": 1,
	"Unsync":     2,
}

func (x GrantState) String() string {
	return proto.EnumName(GrantState_name, int32(x))
}

func (GrantState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{1}
}

type GetStateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStateRequest) Reset()         { *m = GetStateRequest{} }
func (m *GetStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetStateRequest) ProtoMessage()    {}
func (*GetStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{0}
}

func (m *GetStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateRequest.Unmarshal(m, b)
}
func (m *GetStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateRequest.Marshal(b, m, deterministic)
}
func (m *GetStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateRequest.Merge(m, src)
}
func (m *GetStateRequest) XXX_Size() int {
	return xxx_messageInfo_GetStateRequest.Size(m)
}
func (m *GetStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateRequest proto.InternalMessageInfo

type State struct {
	Cbsds                []*Cbsd  `protobuf:"bytes,1,rep,name=cbsds,proto3" json:"cbsds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{1}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetCbsds() []*Cbsd {
	if m != nil {
		return m.Cbsds
	}
	return nil
}

type Cbsd struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string                `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FccId                string                `protobuf:"bytes,3,opt,name=fcc_id,json=fccId,proto3" json:"fcc_id,omitempty"`
	SerialNumber         string                `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	State                CbsdState             `protobuf:"varint,5,opt,name=state,proto3,enum=CbsdState" json:"state,omitempty"`
	DesiredState         CbsdState             `protobuf:"varint,6,opt,name=desired_state,json=desiredState,proto3,enum=CbsdState" json:"desired_state,omitempty"`
	Grants               []*Grant              `protobuf:"bytes,7,rep,name=grants,proto3" json:"grants,omitempty"`
	Channels             []*Channel            `protobuf:"bytes,8,rep,name=channels,proto3" json:"channels,omitempty"`
	LastSeenTimestamp    int64                 `protobuf:"varint,9,opt,name=last_seen_timestamp,json=lastSeenTimestamp,proto3" json:"last_seen_timestamp,omitempty"`
	GrantAttempts        int32                 `protobuf:"varint,10,opt,name=grant_attempts,json=grantAttempts,proto3" json:"grant_attempts,omitempty"`
	EirpCapabilities     *EirpCapabilities     `protobuf:"bytes,11,opt,name=eirp_capabilities,json=eirpCapabilities,proto3" json:"eirp_capabilities,omitempty"`
	DbData               *DatabaseCbsd         `protobuf:"bytes,12,opt,name=db_data,json=dbData,proto3" json:"db_data,omitempty"`
	Preferences          *FrequencyPreferences `protobuf:"bytes,13,opt,name=preferences,proto3" json:"preferences,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Cbsd) Reset()         { *m = Cbsd{} }
func (m *Cbsd) String() string { return proto.CompactTextString(m) }
func (*Cbsd) ProtoMessage()    {}
func (*Cbsd) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{2}
}

func (m *Cbsd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cbsd.Unmarshal(m, b)
}
func (m *Cbsd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cbsd.Marshal(b, m, deterministic)
}
func (m *Cbsd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cbsd.Merge(m, src)
}
func (m *Cbsd) XXX_Size() int {
	return xxx_messageInfo_Cbsd.Size(m)
}
func (m *Cbsd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cbsd.DiscardUnknown(m)
}

var xxx_messageInfo_Cbsd proto.InternalMessageInfo

func (m *Cbsd) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cbsd) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Cbsd) GetFccId() string {
	if m != nil {
		return m.FccId
	}
	return ""
}

func (m *Cbsd) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *Cbsd) GetState() CbsdState {
	if m != nil {
		return m.State
	}
	return CbsdState_Unregistered
}

func (m *Cbsd) GetDesiredState() CbsdState {
	if m != nil {
		return m.DesiredState
	}
	return CbsdState_Unregistered
}

func (m *Cbsd) GetGrants() []*Grant {
	if m != nil {
		return m.Grants
	}
	return nil
}

func (m *Cbsd) GetChannels() []*Channel {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *Cbsd) GetLastSeenTimestamp() int64 {
	if m != nil {
		return m.LastSeenTimestamp
	}
	return 0
}

func (m *Cbsd) GetGrantAttempts() int32 {
	if m != nil {
		return m.GrantAttempts
	}
	return 0
}

func (m *Cbsd) GetEirpCapabilities() *EirpCapabilities {
	if m != nil {
		return m.EirpCapabilities
	}
	return nil
}

func (m *Cbsd) GetDbData() *DatabaseCbsd {
	if m != nil {
		return m.DbData
	}
	return nil
}

func (m *Cbsd) GetPreferences() *FrequencyPreferences {
	if m != nil {
		return m.Preferences
	}
	return nil
}

type Grant struct {
	Id                     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State                  GrantState `protobuf:"varint,2,opt,name=state,proto3,enum=GrantState" json:"state,omitempty"`
	HeartbeatIntervalSec   int64      `protobuf:"varint,3,opt,name=heartbeat_interval_sec,json=heartbeatIntervalSec,proto3" json:"heartbeat_interval_sec,omitempty"`
	LastHeartbeatTimestamp int64      `protobuf:"varint,4,opt,name=last_heartbeat_timestamp,json=lastHeartbeatTimestamp,proto3" json:"last_heartbeat_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}   `json:"-"`
	XXX_unrecognized       []byte     `json:"-"`
	XXX_sizecache          int32      `json:"-"`
}

func (m *Grant) Reset()         { *m = Grant{} }
func (m *Grant) String() string { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()    {}
func (*Grant) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{3}
}

func (m *Grant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Grant.Unmarshal(m, b)
}
func (m *Grant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Grant.Marshal(b, m, deterministic)
}
func (m *Grant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grant.Merge(m, src)
}
func (m *Grant) XXX_Size() int {
	return xxx_messageInfo_Grant.Size(m)
}
func (m *Grant) XXX_DiscardUnknown() {
	xxx_messageInfo_Grant.DiscardUnknown(m)
}

var xxx_messageInfo_Grant proto.InternalMessageInfo

func (m *Grant) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Grant) GetState() GrantState {
	if m != nil {
		return m.State
	}
	return GrantState_Granted
}

func (m *Grant) GetHeartbeatIntervalSec() int64 {
	if m != nil {
		return m.HeartbeatIntervalSec
	}
	return 0
}

func (m *Grant) GetLastHeartbeatTimestamp() int64 {
	if m != nil {
		return m.LastHeartbeatTimestamp
	}
	return 0
}

type Channel struct {
	LowFrequencyHz       int64                `protobuf:"varint,1,opt,name=low_frequency_hz,json=lowFrequencyHz,proto3" json:"low_frequency_hz,omitempty"`
	HighFrequencyHz      int64                `protobuf:"varint,2,opt,name=high_frequency_hz,json=highFrequencyHz,proto3" json:"high_frequency_hz,omitempty"`
	MaxEirp              *wrappers.FloatValue `protobuf:"bytes,3,opt,name=max_eirp,json=maxEirp,proto3" json:"max_eirp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{4}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetLowFrequencyHz() int64 {
	if m != nil {
		return m.LowFrequencyHz
	}
	return 0
}

func (m *Channel) GetHighFrequencyHz() int64 {
	if m != nil {
		return m.HighFrequencyHz
	}
	return 0
}

func (m *Channel) GetMaxEirp() *wrappers.FloatValue {
	if m != nil {
		return m.MaxEirp
	}
	return nil
}

type EirpCapabilities struct {
	MinPower             float32  `protobuf:"fixed32,1,opt,name=min_power,json=minPower,proto3" json:"min_power,omitempty"`
	MaxPower             float32  `protobuf:"fixed32,2,opt,name=max_power,json=maxPower,proto3" json:"max_power,omitempty"`
	AntennaGain          float32  `protobuf:"fixed32,3,opt,name=antenna_gain,json=antennaGain,proto3" json:"antenna_gain,omitempty"`
	NumberOfPorts        int32    `protobuf:"varint,4,opt,name=number_of_ports,json=numberOfPorts,proto3" json:"number_of_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EirpCapabilities) Reset()         { *m = EirpCapabilities{} }
func (m *EirpCapabilities) String() string { return proto.CompactTextString(m) }
func (*EirpCapabilities) ProtoMessage()    {}
func (*EirpCapabilities) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{5}
}

func (m *EirpCapabilities) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EirpCapabilities.Unmarshal(m, b)
}
func (m *EirpCapabilities) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EirpCapabilities.Marshal(b, m, deterministic)
}
func (m *EirpCapabilities) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EirpCapabilities.Merge(m, src)
}
func (m *EirpCapabilities) XXX_Size() int {
	return xxx_messageInfo_EirpCapabilities.Size(m)
}
func (m *EirpCapabilities) XXX_DiscardUnknown() {
	xxx_messageInfo_EirpCapabilities.DiscardUnknown(m)
}

var xxx_messageInfo_EirpCapabilities proto.InternalMessageInfo

func (m *EirpCapabilities) GetMinPower() float32 {
	if m != nil {
		return m.MinPower
	}
	return 0
}

func (m *EirpCapabilities) GetMaxPower() float32 {
	if m != nil {
		return m.MaxPower
	}
	return 0
}

func (m *EirpCapabilities) GetAntennaGain() float32 {
	if m != nil {
		return m.AntennaGain
	}
	return 0
}

func (m *EirpCapabilities) GetNumberOfPorts() int32 {
	if m != nil {
		return m.NumberOfPorts
	}
	return 0
}

type DatabaseCbsd struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ShouldDeregister     bool     `protobuf:"varint,2,opt,name=should_deregister,json=shouldDeregister,proto3" json:"should_deregister,omitempty"`
	IsDeleted            bool     `protobuf:"varint,3,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabaseCbsd) Reset()         { *m = DatabaseCbsd{} }
func (m *DatabaseCbsd) String() string { return proto.CompactTextString(m) }
func (*DatabaseCbsd) ProtoMessage()    {}
func (*DatabaseCbsd) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{6}
}

func (m *DatabaseCbsd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseCbsd.Unmarshal(m, b)
}
func (m *DatabaseCbsd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseCbsd.Marshal(b, m, deterministic)
}
func (m *DatabaseCbsd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseCbsd.Merge(m, src)
}
func (m *DatabaseCbsd) XXX_Size() int {
	return xxx_messageInfo_DatabaseCbsd.Size(m)
}
func (m *DatabaseCbsd) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseCbsd.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseCbsd proto.InternalMessageInfo

func (m *DatabaseCbsd) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DatabaseCbsd) GetShouldDeregister() bool {
	if m != nil {
		return m.ShouldDeregister
	}
	return false
}

func (m *DatabaseCbsd) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

type FrequencyPreferences struct {
	BandwidthMhz         int32    `protobuf:"varint,1,opt,name=bandwidth_mhz,json=bandwidthMhz,proto3" json:"bandwidth_mhz,omitempty"`
	FrequenciesMhz       []int32  `protobuf:"varint,2,rep,packed,name=frequencies_mhz,json=frequenciesMhz,proto3" json:"frequencies_mhz,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrequencyPreferences) Reset()         { *m = FrequencyPreferences{} }
func (m *FrequencyPreferences) String() string { return proto.CompactTextString(m) }
func (*FrequencyPreferences) ProtoMessage()    {}
func (*FrequencyPreferences) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{7}
}

func (m *FrequencyPreferences) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyPreferences.Unmarshal(m, b)
}
func (m *FrequencyPreferences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyPreferences.Marshal(b, m, deterministic)
}
func (m *FrequencyPreferences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyPreferences.Merge(m, src)
}
func (m *FrequencyPreferences) XXX_Size() int {
	return xxx_messageInfo_FrequencyPreferences.Size(m)
}
func (m *FrequencyPreferences) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyPreferences.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyPreferences proto.InternalMessageInfo

func (m *FrequencyPreferences) GetBandwidthMhz() int32 {
	if m != nil {
		return m.BandwidthMhz
	}
	return 0
}

func (m *FrequencyPreferences) GetFrequenciesMhz() []int32 {
	if m != nil {
		return m.FrequenciesMhz
	}
	return nil
}

type DeleteCbsdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCbsdRequest) Reset()         { *m = DeleteCbsdRequest{} }
func (m *DeleteCbsdRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCbsdRequest) ProtoMessage()    {}
func (*DeleteCbsdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{8}
}

func (m *DeleteCbsdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCbsdRequest.Unmarshal(m, b)
}
func (m *DeleteCbsdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCbsdRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCbsdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCbsdRequest.Merge(m, src)
}
func (m *DeleteCbsdRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCbsdRequest.Size(m)
}
func (m *DeleteCbsdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCbsdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCbsdRequest proto.InternalMessageInfo

func (m *DeleteCbsdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type AcknowledgeCbsdUpdateRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcknowledgeCbsdUpdateRequest) Reset()         { *m = AcknowledgeCbsdUpdateRequest{} }
func (m *AcknowledgeCbsdUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*AcknowledgeCbsdUpdateRequest) ProtoMessage()    {}
func (*AcknowledgeCbsdUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cca56e429e696f1e, []int{9}
}

func (m *AcknowledgeCbsdUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcknowledgeCbsdUpdateRequest.Unmarshal(m, b)
}
func (m *AcknowledgeCbsdUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcknowledgeCbsdUpdateRequest.Marshal(b, m, deterministic)
}
func (m *AcknowledgeCbsdUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcknowledgeCbsdUpdateRequest.Merge(m, src)
}
func (m *AcknowledgeCbsdUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_AcknowledgeCbsdUpdateRequest.Size(m)
}
func (m *AcknowledgeCbsdUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AcknowledgeCbsdUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AcknowledgeCbsdUpdateRequest proto.InternalMessageInfo

func (m *AcknowledgeCbsdUpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterEnum("CbsdState", CbsdState_name, CbsdState_value)
	proto.RegisterEnum("GrantState", GrantState_name, GrantState_value)
	proto.RegisterType((*GetStateRequest)(nil), "GetStateRequest")
	proto.RegisterType((*State)(nil), "State")
	proto.RegisterType((*Cbsd)(nil), "Cbsd")
	proto.RegisterType((*Grant)(nil), "Grant")
	proto.RegisterType((*Channel)(nil), "Channel")
	proto.RegisterType((*EirpCapabilities)(nil), "EirpCapabilities")
	proto.RegisterType((*DatabaseCbsd)(nil), "DatabaseCbsd")
	proto.RegisterType((*FrequencyPreferences)(nil), "FrequencyPreferences")
	proto.RegisterType((*DeleteCbsdRequest)(nil), "DeleteCbsdRequest")
	proto.RegisterType((*AcknowledgeCbsdUpdateRequest)(nil), "AcknowledgeCbsdUpdateRequest")
}

func init() { proto.RegisterFile("dp/protos/active_mode.proto", fileDescriptor_cca56e429e696f1e) }

var fileDescriptor_cca56e429e696f1e = []byte{
	// 943 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x72, 0xe3, 0x44,
	0x13, 0x8d, 0xec, 0xc8, 0x76, 0xda, 0x3f, 0xb1, 0xe7, 0x4b, 0xf2, 0xa9, 0x12, 0x76, 0xcb, 0xab,
	0x5d, 0x82, 0x2b, 0x14, 0x72, 0x55, 0xf8, 0x2d, 0x28, 0xa8, 0x0a, 0xc9, 0x6e, 0x36, 0x17, 0x81,
	0x94, 0x42, 0xb8, 0xe0, 0x66, 0x6a, 0xa4, 0x69, 0xdb, 0x03, 0xd2, 0x48, 0x68, 0xc6, 0xeb, 0xc4,
	0x2f, 0xc2, 0x3b, 0xf0, 0x2c, 0x3c, 0x12, 0x17, 0x94, 0x46, 0xf2, 0xcf, 0x3a, 0x81, 0xcb, 0x3e,
	0xe7, 0x8c, 0xd4, 0xdd, 0x67, 0xba, 0x07, 0x8e, 0x78, 0x3a, 0x4c, 0xb3, 0x44, 0x27, 0x6a, 0xc8,
	0x42, 0x2d, 0xde, 0x21, 0x8d, 0x13, 0x8e, 0x9e, 0x81, 0x0e, 0x8f, 0xc6, 0x49, 0x32, 0x8e, 0xb0,
	0x10, 0x04, 0xd3, 0xd1, 0x10, 0xe3, 0x54, 0x3f, 0x94, 0xe4, 0xf3, 0x4d, 0x72, 0x96, 0xb1, 0x34,
	0xc5, 0x4c, 0x15, 0xbc, 0xdb, 0x83, 0xdd, 0x4b, 0xd4, 0xb7, 0x9a, 0x69, 0xf4, 0xf1, 0xf7, 0x29,
	0x2a, 0xed, 0xbe, 0x02, 0xdb, 0xc4, 0xe4, 0x08, 0xec, 0x30, 0x50, 0x5c, 0x39, 0x56, 0xbf, 0x3a,
	0x68, 0x9e, 0xda, 0xde, 0x79, 0xa0, 0xb8, 0x5f, 0x60, 0xee, 0xdf, 0x55, 0xd8, 0xce, 0x63, 0xd2,
	0x81, 0x8a, 0xe0, 0x8e, 0xd5, 0xb7, 0x06, 0x3b, 0x7e, 0x45, 0x70, 0xf2, 0x7f, 0xa8, 0x4f, 0x15,
	0x66, 0x54, 0x70, 0xa7, 0x62, 0xc0, 0x5a, 0x1e, 0x5e, 0x71, 0xb2, 0x0f, 0xb5, 0x51, 0x18, 0xe6,
	0x78, 0xd5, 0xe0, 0xf6, 0x28, 0x0c, 0xaf, 0x38, 0x79, 0x09, 0x6d, 0x85, 0x99, 0x60, 0x11, 0x95,
	0xd3, 0x38, 0xc0, 0xcc, 0xd9, 0x36, 0x6c, 0xab, 0x00, 0x7f, 0x30, 0x18, 0xe9, 0x83, 0xad, 0xf2,
	0x9c, 0x1c, 0xbb, 0x6f, 0x0d, 0x3a, 0xa7, 0x60, 0x52, 0x29, 0xb2, 0x2e, 0x08, 0x32, 0x84, 0x36,
	0x47, 0x25, 0x32, 0xe4, 0xb4, 0x50, 0xd6, 0x1e, 0x29, 0x5b, 0xa5, 0xa0, 0xa8, 0xee, 0x39, 0xd4,
	0xc6, 0x19, 0x93, 0x5a, 0x39, 0x75, 0x53, 0x5e, 0xcd, 0xbb, 0xcc, 0x43, 0xbf, 0x44, 0xc9, 0x2b,
	0x68, 0x84, 0x13, 0x26, 0x25, 0x46, 0xca, 0x69, 0x18, 0x45, 0xc3, 0x3b, 0x2f, 0x00, 0x7f, 0xc9,
	0x10, 0x0f, 0xfe, 0x17, 0x31, 0xa5, 0xa9, 0x42, 0x94, 0x54, 0x8b, 0x18, 0x95, 0x66, 0x71, 0xea,
	0xec, 0xf4, 0xad, 0x41, 0xd5, 0xef, 0xe5, 0xd4, 0x2d, 0xa2, 0xfc, 0x69, 0x41, 0x90, 0x0f, 0xa1,
	0x63, 0xbe, 0x4f, 0x99, 0xd6, 0xb9, 0x4f, 0xca, 0x81, 0xbe, 0x35, 0xb0, 0xfd, 0xb6, 0x41, 0xcf,
	0x4a, 0x90, 0x7c, 0x07, 0x3d, 0x14, 0x59, 0x4a, 0x43, 0x96, 0xb2, 0x40, 0x44, 0x42, 0x0b, 0x54,
	0x4e, 0xb3, 0x6f, 0x0d, 0x9a, 0xa7, 0x3d, 0xef, 0xb5, 0xc8, 0xd2, 0xf3, 0x35, 0xc2, 0xef, 0xe2,
	0x06, 0x42, 0x8e, 0xa1, 0xce, 0x03, 0xca, 0x99, 0x66, 0x4e, 0xcb, 0x9c, 0x6a, 0x7b, 0x17, 0x4c,
	0xb3, 0x80, 0x29, 0x34, 0x26, 0xd6, 0x78, 0x90, 0xc7, 0xe4, 0x4b, 0x68, 0xa6, 0x19, 0x8e, 0x30,
	0x43, 0x19, 0xa2, 0x72, 0xda, 0x46, 0xbb, 0xef, 0xbd, 0xc9, 0xf2, 0xbb, 0x20, 0xc3, 0x87, 0x9b,
	0x15, 0xe9, 0xaf, 0x2b, 0xdd, 0x3f, 0x2d, 0xb0, 0x4d, 0xbf, 0x1e, 0xf9, 0xff, 0x62, 0x61, 0x55,
	0xc5, 0x18, 0xd0, 0x2c, 0xda, 0xfa, 0x9e, 0x57, 0x9f, 0xc1, 0xc1, 0x04, 0x59, 0xa6, 0x03, 0x64,
	0x9a, 0x0a, 0xa9, 0x31, 0x7b, 0xc7, 0x22, 0xaa, 0x30, 0x34, 0x37, 0xa3, 0xea, 0xef, 0x2d, 0xd9,
	0xab, 0x92, 0xbc, 0xc5, 0x90, 0x7c, 0x05, 0x8e, 0x69, 0xf5, 0xea, 0xe8, 0xaa, 0xdf, 0xdb, 0xe6,
	0xdc, 0x41, 0xce, 0xbf, 0x5d, 0xd0, 0xcb, 0xa6, 0xbb, 0x7f, 0x58, 0x50, 0x2f, 0xad, 0x23, 0x03,
	0xe8, 0x46, 0xc9, 0x8c, 0x8e, 0x16, 0x15, 0xd2, 0xc9, 0xdc, 0x24, 0x5f, 0xf5, 0x3b, 0x51, 0x32,
	0x5b, 0x16, 0xfe, 0x76, 0x4e, 0x4e, 0xa0, 0x37, 0x11, 0xe3, 0xc9, 0xfb, 0xd2, 0x8a, 0x91, 0xee,
	0xe6, 0xc4, 0xba, 0xf6, 0x0b, 0x68, 0xc4, 0xec, 0x9e, 0xe6, 0x3e, 0x98, 0x1a, 0x9a, 0xa7, 0x47,
	0x5e, 0x31, 0x79, 0xde, 0x62, 0xf2, 0xbc, 0x37, 0x51, 0xc2, 0xf4, 0xcf, 0x2c, 0x9a, 0xa2, 0x5f,
	0x8f, 0xd9, 0x7d, 0xee, 0x62, 0x9e, 0x59, 0x77, 0xd3, 0x4e, 0x72, 0x04, 0x3b, 0xb1, 0x90, 0x34,
	0x4d, 0x66, 0x98, 0x99, 0xdc, 0x2a, 0x7e, 0x23, 0x16, 0xf2, 0x26, 0x8f, 0x0d, 0xc9, 0xee, 0x4b,
	0xb2, 0x52, 0x92, 0xec, 0xbe, 0x20, 0x5f, 0x40, 0x8b, 0x49, 0x8d, 0x52, 0x32, 0x3a, 0x66, 0x42,
	0x9a, 0x54, 0x2a, 0x7e, 0xb3, 0xc4, 0x2e, 0x99, 0x90, 0xe4, 0x18, 0x76, 0x8b, 0x39, 0xa3, 0xc9,
	0x88, 0xa6, 0x49, 0xa6, 0x95, 0x69, 0x9e, 0xed, 0xb7, 0x0b, 0xf8, 0xc7, 0xd1, 0x4d, 0x0e, 0xba,
	0xbf, 0x42, 0x6b, 0xfd, 0xc6, 0xac, 0xd9, 0x5c, 0x35, 0x36, 0x7f, 0x0c, 0x3d, 0x35, 0x49, 0xa6,
	0x11, 0xa7, 0x1c, 0x33, 0x1c, 0x0b, 0xa5, 0xcb, 0x7c, 0x1a, 0x7e, 0xb7, 0x20, 0x2e, 0x96, 0x38,
	0x79, 0x06, 0x20, 0x14, 0xe5, 0x18, 0xa1, 0xc6, 0x62, 0xfc, 0x1b, 0xfe, 0x8e, 0x50, 0x17, 0x05,
	0xe0, 0x72, 0xd8, 0x7b, 0xea, 0xc6, 0xe5, 0xab, 0x21, 0x60, 0x92, 0xcf, 0x04, 0xd7, 0x13, 0x1a,
	0x97, 0x46, 0xd9, 0x7e, 0x6b, 0x09, 0x5e, 0x4f, 0xe6, 0xe4, 0x23, 0xd8, 0x5d, 0x38, 0x24, 0x50,
	0x19, 0x59, 0xa5, 0x5f, 0x1d, 0xd8, 0x7e, 0x67, 0x0d, 0xbe, 0x9e, 0xcc, 0xdd, 0x97, 0xd0, 0x2b,
	0x7e, 0x68, 0x26, 0xa0, 0x58, 0x76, 0x9b, 0x65, 0xb9, 0x1e, 0x7c, 0x70, 0x16, 0xfe, 0x26, 0x93,
	0x59, 0x84, 0x7c, 0x6c, 0x94, 0x77, 0x29, 0x5f, 0x2d, 0xc7, 0x4d, 0xfd, 0xc9, 0x27, 0xb0, 0xb3,
	0x5c, 0x30, 0xa4, 0x0b, 0xad, 0x3b, 0xb9, 0x28, 0x1a, 0x79, 0x77, 0x8b, 0x74, 0x00, 0xfc, 0x55,
	0x6c, 0x9d, 0x7c, 0x0e, 0xb0, 0x1a, 0x07, 0xd2, 0x84, 0xba, 0x89, 0x16, 0xd2, 0xb3, 0xa9, 0x9e,
	0x24, 0x99, 0x98, 0xe7, 0x52, 0x02, 0x50, 0xbb, 0x93, 0xea, 0x41, 0x86, 0xdd, 0xca, 0xe9, 0x5f,
	0x16, 0xec, 0x9d, 0x99, 0xc5, 0x7f, 0x9d, 0x70, 0x3c, 0x4f, 0xa4, 0xce, 0x92, 0x28, 0xc2, 0x8c,
	0x1c, 0x43, 0x63, 0xb1, 0xbe, 0x49, 0xd7, 0xdb, 0xd8, 0xe4, 0x87, 0x35, 0xcf, 0x84, 0xee, 0x16,
	0xf9, 0x1a, 0x60, 0x55, 0x3b, 0x21, 0xde, 0xa3, 0x46, 0x1c, 0x1e, 0x3c, 0xba, 0xaf, 0xaf, 0xf3,
	0x67, 0xc4, 0xdd, 0x22, 0x37, 0xb0, 0xff, 0x64, 0x4b, 0xc8, 0x33, 0xef, 0xbf, 0x5a, 0xf5, 0xef,
	0x5f, 0xfc, 0xfe, 0xdb, 0x5f, 0xbe, 0x89, 0xd9, 0x38, 0x66, 0x43, 0x9e, 0x0e, 0xc3, 0x28, 0x99,
	0xf2, 0xe1, 0x38, 0x59, 0x7f, 0xd8, 0x68, 0xb8, 0xac, 0xf0, 0x89, 0x67, 0x2f, 0xa8, 0x19, 0xec,
	0xd3, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x59, 0xee, 0x9d, 0x16, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ActiveModeControllerClient is the client API for ActiveModeController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActiveModeControllerClient interface {
	GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error)
	DeleteCbsd(ctx context.Context, in *DeleteCbsdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AcknowledgeCbsdUpdate(ctx context.Context, in *AcknowledgeCbsdUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type activeModeControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewActiveModeControllerClient(cc grpc.ClientConnInterface) ActiveModeControllerClient {
	return &activeModeControllerClient{cc}
}

func (c *activeModeControllerClient) GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/ActiveModeController/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeModeControllerClient) DeleteCbsd(ctx context.Context, in *DeleteCbsdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ActiveModeController/DeleteCbsd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activeModeControllerClient) AcknowledgeCbsdUpdate(ctx context.Context, in *AcknowledgeCbsdUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ActiveModeController/AcknowledgeCbsdUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActiveModeControllerServer is the server API for ActiveModeController service.
type ActiveModeControllerServer interface {
	GetState(context.Context, *GetStateRequest) (*State, error)
	DeleteCbsd(context.Context, *DeleteCbsdRequest) (*empty.Empty, error)
	AcknowledgeCbsdUpdate(context.Context, *AcknowledgeCbsdUpdateRequest) (*empty.Empty, error)
}

// UnimplementedActiveModeControllerServer can be embedded to have forward compatible implementations.
type UnimplementedActiveModeControllerServer struct {
}

func (*UnimplementedActiveModeControllerServer) GetState(ctx context.Context, req *GetStateRequest) (*State, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (*UnimplementedActiveModeControllerServer) DeleteCbsd(ctx context.Context, req *DeleteCbsdRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCbsd not implemented")
}
func (*UnimplementedActiveModeControllerServer) AcknowledgeCbsdUpdate(ctx context.Context, req *AcknowledgeCbsdUpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcknowledgeCbsdUpdate not implemented")
}

func RegisterActiveModeControllerServer(s *grpc.Server, srv ActiveModeControllerServer) {
	s.RegisterService(&_ActiveModeController_serviceDesc, srv)
}

func _ActiveModeController_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveModeControllerServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActiveModeController/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveModeControllerServer).GetState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveModeController_DeleteCbsd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCbsdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveModeControllerServer).DeleteCbsd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActiveModeController/DeleteCbsd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveModeControllerServer).DeleteCbsd(ctx, req.(*DeleteCbsdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActiveModeController_AcknowledgeCbsdUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcknowledgeCbsdUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActiveModeControllerServer).AcknowledgeCbsdUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActiveModeController/AcknowledgeCbsdUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActiveModeControllerServer).AcknowledgeCbsdUpdate(ctx, req.(*AcknowledgeCbsdUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActiveModeController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ActiveModeController",
	HandlerType: (*ActiveModeControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _ActiveModeController_GetState_Handler,
		},
		{
			MethodName: "DeleteCbsd",
			Handler:    _ActiveModeController_DeleteCbsd_Handler,
		},
		{
			MethodName: "AcknowledgeCbsdUpdate",
			Handler:    _ActiveModeController_AcknowledgeCbsdUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dp/protos/active_mode.proto",
}
