// Code generated by protoc-gen-go.
// source: easel_service.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	easel_service.proto

It has these top-level messages:
	NewEaselRequest
	NewEaselResponse
	DeleteEaselRequest
	DeleteEaselResponse
	NewPaletteRequest
	NewPaletteResponse
	DeletePaletteRequest
	DeletePaletteResponse
	PingRequest
	PongResponse
	EaselInfo
	PaletteInfo
	ListupRequest
	ListupResponse
	VertexAttribute
	UniformFloatValue
	UniformIntValue
	UniformVariable
	ArrayBuffer
	PaletteUpdate
	UpdatePaletteRequest
	UpdatePaletteResponse
	RenderRequest
	RenderResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type NewEaselRequest struct {
	EaselId string `protobuf:"bytes,1,opt,name=easel_id,json=easelId" json:"easel_id,omitempty"`
}

func (m *NewEaselRequest) Reset()                    { *m = NewEaselRequest{} }
func (m *NewEaselRequest) String() string            { return proto1.CompactTextString(m) }
func (*NewEaselRequest) ProtoMessage()               {}
func (*NewEaselRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewEaselRequest) GetEaselId() string {
	if m != nil {
		return m.EaselId
	}
	return ""
}

type NewEaselResponse struct {
	EaselId string `protobuf:"bytes,1,opt,name=easel_id,json=easelId" json:"easel_id,omitempty"`
}

func (m *NewEaselResponse) Reset()                    { *m = NewEaselResponse{} }
func (m *NewEaselResponse) String() string            { return proto1.CompactTextString(m) }
func (*NewEaselResponse) ProtoMessage()               {}
func (*NewEaselResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewEaselResponse) GetEaselId() string {
	if m != nil {
		return m.EaselId
	}
	return ""
}

type DeleteEaselRequest struct {
	EaselId string `protobuf:"bytes,1,opt,name=easel_id,json=easelId" json:"easel_id,omitempty"`
}

func (m *DeleteEaselRequest) Reset()                    { *m = DeleteEaselRequest{} }
func (m *DeleteEaselRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteEaselRequest) ProtoMessage()               {}
func (*DeleteEaselRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeleteEaselRequest) GetEaselId() string {
	if m != nil {
		return m.EaselId
	}
	return ""
}

type DeleteEaselResponse struct {
}

func (m *DeleteEaselResponse) Reset()                    { *m = DeleteEaselResponse{} }
func (m *DeleteEaselResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeleteEaselResponse) ProtoMessage()               {}
func (*DeleteEaselResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type NewPaletteRequest struct {
	EaselId string `protobuf:"bytes,1,opt,name=easel_id,json=easelId" json:"easel_id,omitempty"`
}

func (m *NewPaletteRequest) Reset()                    { *m = NewPaletteRequest{} }
func (m *NewPaletteRequest) String() string            { return proto1.CompactTextString(m) }
func (*NewPaletteRequest) ProtoMessage()               {}
func (*NewPaletteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NewPaletteRequest) GetEaselId() string {
	if m != nil {
		return m.EaselId
	}
	return ""
}

type NewPaletteResponse struct {
	EaselId   string `protobuf:"bytes,1,opt,name=easel_id,json=easelId" json:"easel_id,omitempty"`
	PaletteId string `protobuf:"bytes,2,opt,name=palette_id,json=paletteId" json:"palette_id,omitempty"`
}

func (m *NewPaletteResponse) Reset()                    { *m = NewPaletteResponse{} }
func (m *NewPaletteResponse) String() string            { return proto1.CompactTextString(m) }
func (*NewPaletteResponse) ProtoMessage()               {}
func (*NewPaletteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *NewPaletteResponse) GetEaselId() string {
	if m != nil {
		return m.EaselId
	}
	return ""
}

func (m *NewPaletteResponse) GetPaletteId() string {
	if m != nil {
		return m.PaletteId
	}
	return ""
}

type DeletePaletteRequest struct {
	EaselId   string `protobuf:"bytes,1,opt,name=easel_id,json=easelId" json:"easel_id,omitempty"`
	PaletteId string `protobuf:"bytes,2,opt,name=palette_id,json=paletteId" json:"palette_id,omitempty"`
}

func (m *DeletePaletteRequest) Reset()                    { *m = DeletePaletteRequest{} }
func (m *DeletePaletteRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeletePaletteRequest) ProtoMessage()               {}
func (*DeletePaletteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeletePaletteRequest) GetEaselId() string {
	if m != nil {
		return m.EaselId
	}
	return ""
}

func (m *DeletePaletteRequest) GetPaletteId() string {
	if m != nil {
		return m.PaletteId
	}
	return ""
}

type DeletePaletteResponse struct {
}

func (m *DeletePaletteResponse) Reset()                    { *m = DeletePaletteResponse{} }
func (m *DeletePaletteResponse) String() string            { return proto1.CompactTextString(m) }
func (*DeletePaletteResponse) ProtoMessage()               {}
func (*DeletePaletteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type PingRequest struct {
	EaselId   string `protobuf:"bytes,1,opt,name=easel_id,json=easelId" json:"easel_id,omitempty"`
	PaletteId string `protobuf:"bytes,2,opt,name=palette_id,json=paletteId" json:"palette_id,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto1.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PingRequest) GetEaselId() string {
	if m != nil {
		return m.EaselId
	}
	return ""
}

func (m *PingRequest) GetPaletteId() string {
	if m != nil {
		return m.PaletteId
	}
	return ""
}

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PongResponse struct {
	EaselId   string `protobuf:"bytes,1,opt,name=easel_id,json=easelId" json:"easel_id,omitempty"`
	PaletteId string `protobuf:"bytes,2,opt,name=palette_id,json=paletteId" json:"palette_id,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *PongResponse) Reset()                    { *m = PongResponse{} }
func (m *PongResponse) String() string            { return proto1.CompactTextString(m) }
func (*PongResponse) ProtoMessage()               {}
func (*PongResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PongResponse) GetEaselId() string {
	if m != nil {
		return m.EaselId
	}
	return ""
}

func (m *PongResponse) GetPaletteId() string {
	if m != nil {
		return m.PaletteId
	}
	return ""
}

func (m *PongResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EaselInfo struct {
	Id        string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	UpdatedAt string         `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Palettes  []*PaletteInfo `protobuf:"bytes,3,rep,name=palettes" json:"palettes,omitempty"`
}

func (m *EaselInfo) Reset()                    { *m = EaselInfo{} }
func (m *EaselInfo) String() string            { return proto1.CompactTextString(m) }
func (*EaselInfo) ProtoMessage()               {}
func (*EaselInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *EaselInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EaselInfo) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *EaselInfo) GetPalettes() []*PaletteInfo {
	if m != nil {
		return m.Palettes
	}
	return nil
}

type PaletteInfo struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	UpdatedAt string `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *PaletteInfo) Reset()                    { *m = PaletteInfo{} }
func (m *PaletteInfo) String() string            { return proto1.CompactTextString(m) }
func (*PaletteInfo) ProtoMessage()               {}
func (*PaletteInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PaletteInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PaletteInfo) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListupRequest struct {
}

func (m *ListupRequest) Reset()                    { *m = ListupRequest{} }
func (m *ListupRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListupRequest) ProtoMessage()               {}
func (*ListupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type ListupResponse struct {
	Easels []*EaselInfo `protobuf:"bytes,1,rep,name=easels" json:"easels,omitempty"`
}

func (m *ListupResponse) Reset()                    { *m = ListupResponse{} }
func (m *ListupResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListupResponse) ProtoMessage()               {}
func (*ListupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ListupResponse) GetEasels() []*EaselInfo {
	if m != nil {
		return m.Easels
	}
	return nil
}

type VertexAttribute struct {
	ArgumentName string `protobuf:"bytes,1,opt,name=argument_name,json=argumentName" json:"argument_name,omitempty"`
	BufferName   string `protobuf:"bytes,2,opt,name=buffer_name,json=bufferName" json:"buffer_name,omitempty"`
	ElementSize  int32  `protobuf:"varint,3,opt,name=element_size,json=elementSize" json:"element_size,omitempty"`
	Offset       int32  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	Stride       int32  `protobuf:"varint,5,opt,name=stride" json:"stride,omitempty"`
}

func (m *VertexAttribute) Reset()                    { *m = VertexAttribute{} }
func (m *VertexAttribute) String() string            { return proto1.CompactTextString(m) }
func (*VertexAttribute) ProtoMessage()               {}
func (*VertexAttribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *VertexAttribute) GetArgumentName() string {
	if m != nil {
		return m.ArgumentName
	}
	return ""
}

func (m *VertexAttribute) GetBufferName() string {
	if m != nil {
		return m.BufferName
	}
	return ""
}

func (m *VertexAttribute) GetElementSize() int32 {
	if m != nil {
		return m.ElementSize
	}
	return 0
}

func (m *VertexAttribute) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *VertexAttribute) GetStride() int32 {
	if m != nil {
		return m.Stride
	}
	return 0
}

type UniformFloatValue struct {
	ElementSize int32     `protobuf:"varint,1,opt,name=element_size,json=elementSize" json:"element_size,omitempty"`
	Data        []float32 `protobuf:"fixed32,2,rep,packed,name=data" json:"data,omitempty"`
}

func (m *UniformFloatValue) Reset()                    { *m = UniformFloatValue{} }
func (m *UniformFloatValue) String() string            { return proto1.CompactTextString(m) }
func (*UniformFloatValue) ProtoMessage()               {}
func (*UniformFloatValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *UniformFloatValue) GetElementSize() int32 {
	if m != nil {
		return m.ElementSize
	}
	return 0
}

func (m *UniformFloatValue) GetData() []float32 {
	if m != nil {
		return m.Data
	}
	return nil
}

type UniformIntValue struct {
	ElementSize int32   `protobuf:"varint,1,opt,name=element_size,json=elementSize" json:"element_size,omitempty"`
	Data        []int32 `protobuf:"varint,2,rep,packed,name=data" json:"data,omitempty"`
}

func (m *UniformIntValue) Reset()                    { *m = UniformIntValue{} }
func (m *UniformIntValue) String() string            { return proto1.CompactTextString(m) }
func (*UniformIntValue) ProtoMessage()               {}
func (*UniformIntValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *UniformIntValue) GetElementSize() int32 {
	if m != nil {
		return m.ElementSize
	}
	return 0
}

func (m *UniformIntValue) GetData() []int32 {
	if m != nil {
		return m.Data
	}
	return nil
}

type UniformVariable struct {
	Name       string             `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	FloatValue *UniformFloatValue `protobuf:"bytes,2,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	IntValue   *UniformIntValue   `protobuf:"bytes,3,opt,name=int_value,json=intValue" json:"int_value,omitempty"`
	Texture    []byte             `protobuf:"bytes,4,opt,name=texture,proto3" json:"texture,omitempty"`
}

func (m *UniformVariable) Reset()                    { *m = UniformVariable{} }
func (m *UniformVariable) String() string            { return proto1.CompactTextString(m) }
func (*UniformVariable) ProtoMessage()               {}
func (*UniformVariable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *UniformVariable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UniformVariable) GetFloatValue() *UniformFloatValue {
	if m != nil {
		return m.FloatValue
	}
	return nil
}

func (m *UniformVariable) GetIntValue() *UniformIntValue {
	if m != nil {
		return m.IntValue
	}
	return nil
}

func (m *UniformVariable) GetTexture() []byte {
	if m != nil {
		return m.Texture
	}
	return nil
}

type ArrayBuffer struct {
	Name string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Data []float32 `protobuf:"fixed32,2,rep,packed,name=data" json:"data,omitempty"`
}

func (m *ArrayBuffer) Reset()                    { *m = ArrayBuffer{} }
func (m *ArrayBuffer) String() string            { return proto1.CompactTextString(m) }
func (*ArrayBuffer) ProtoMessage()               {}
func (*ArrayBuffer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *ArrayBuffer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArrayBuffer) GetData() []float32 {
	if m != nil {
		return m.Data
	}
	return nil
}

type PaletteUpdate struct {
	// Shader
	VertexShader   string `protobuf:"bytes,1,opt,name=vertex_shader,json=vertexShader" json:"vertex_shader,omitempty"`
	FragmentShader string `protobuf:"bytes,2,opt,name=fragment_shader,json=fragmentShader" json:"fragment_shader,omitempty"`
	// VertexBuffer
	Buffers []*ArrayBuffer `protobuf:"bytes,3,rep,name=buffers" json:"buffers,omitempty"`
	// VertexElementBuffer
	Indecies []int32 `protobuf:"varint,4,rep,packed,name=indecies" json:"indecies,omitempty"`
	// VertexAttributes
	VertexArrtibutes []*VertexAttribute `protobuf:"bytes,5,rep,name=vertex_arrtibutes,json=vertexArrtibutes" json:"vertex_arrtibutes,omitempty"`
	// UniformVariables
	UniformVariables []*UniformVariable `protobuf:"bytes,6,rep,name=uniform_variables,json=uniformVariables" json:"uniform_variables,omitempty"`
}

func (m *PaletteUpdate) Reset()                    { *m = PaletteUpdate{} }
func (m *PaletteUpdate) String() string            { return proto1.CompactTextString(m) }
func (*PaletteUpdate) ProtoMessage()               {}
func (*PaletteUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *PaletteUpdate) GetVertexShader() string {
	if m != nil {
		return m.VertexShader
	}
	return ""
}

func (m *PaletteUpdate) GetFragmentShader() string {
	if m != nil {
		return m.FragmentShader
	}
	return ""
}

func (m *PaletteUpdate) GetBuffers() []*ArrayBuffer {
	if m != nil {
		return m.Buffers
	}
	return nil
}

func (m *PaletteUpdate) GetIndecies() []int32 {
	if m != nil {
		return m.Indecies
	}
	return nil
}

func (m *PaletteUpdate) GetVertexArrtibutes() []*VertexAttribute {
	if m != nil {
		return m.VertexArrtibutes
	}
	return nil
}

func (m *PaletteUpdate) GetUniformVariables() []*UniformVariable {
	if m != nil {
		return m.UniformVariables
	}
	return nil
}

type UpdatePaletteRequest struct {
	EaselId   string         `protobuf:"bytes,1,opt,name=easel_id,json=easelId" json:"easel_id,omitempty"`
	PaletteId string         `protobuf:"bytes,2,opt,name=palette_id,json=paletteId" json:"palette_id,omitempty"`
	Updates   *PaletteUpdate `protobuf:"bytes,3,opt,name=updates" json:"updates,omitempty"`
}

func (m *UpdatePaletteRequest) Reset()                    { *m = UpdatePaletteRequest{} }
func (m *UpdatePaletteRequest) String() string            { return proto1.CompactTextString(m) }
func (*UpdatePaletteRequest) ProtoMessage()               {}
func (*UpdatePaletteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *UpdatePaletteRequest) GetEaselId() string {
	if m != nil {
		return m.EaselId
	}
	return ""
}

func (m *UpdatePaletteRequest) GetPaletteId() string {
	if m != nil {
		return m.PaletteId
	}
	return ""
}

func (m *UpdatePaletteRequest) GetUpdates() *PaletteUpdate {
	if m != nil {
		return m.Updates
	}
	return nil
}

type UpdatePaletteResponse struct {
}

func (m *UpdatePaletteResponse) Reset()                    { *m = UpdatePaletteResponse{} }
func (m *UpdatePaletteResponse) String() string            { return proto1.CompactTextString(m) }
func (*UpdatePaletteResponse) ProtoMessage()               {}
func (*UpdatePaletteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

// Render
type RenderRequest struct {
	EaselId    string         `protobuf:"bytes,1,opt,name=easel_id,json=easelId" json:"easel_id,omitempty"`
	PaletteId  string         `protobuf:"bytes,2,opt,name=palette_id,json=paletteId" json:"palette_id,omitempty"`
	Updates    *PaletteUpdate `protobuf:"bytes,3,opt,name=updates" json:"updates,omitempty"`
	OutQuality float32        `protobuf:"fixed32,4,opt,name=out_quality,json=outQuality" json:"out_quality,omitempty"`
	OutFormat  string         `protobuf:"bytes,5,opt,name=out_format,json=outFormat" json:"out_format,omitempty"`
	OutWidth   int32          `protobuf:"varint,6,opt,name=out_width,json=outWidth" json:"out_width,omitempty"`
	OutHeight  int32          `protobuf:"varint,7,opt,name=out_height,json=outHeight" json:"out_height,omitempty"`
}

func (m *RenderRequest) Reset()                    { *m = RenderRequest{} }
func (m *RenderRequest) String() string            { return proto1.CompactTextString(m) }
func (*RenderRequest) ProtoMessage()               {}
func (*RenderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *RenderRequest) GetEaselId() string {
	if m != nil {
		return m.EaselId
	}
	return ""
}

func (m *RenderRequest) GetPaletteId() string {
	if m != nil {
		return m.PaletteId
	}
	return ""
}

func (m *RenderRequest) GetUpdates() *PaletteUpdate {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *RenderRequest) GetOutQuality() float32 {
	if m != nil {
		return m.OutQuality
	}
	return 0
}

func (m *RenderRequest) GetOutFormat() string {
	if m != nil {
		return m.OutFormat
	}
	return ""
}

func (m *RenderRequest) GetOutWidth() int32 {
	if m != nil {
		return m.OutWidth
	}
	return 0
}

func (m *RenderRequest) GetOutHeight() int32 {
	if m != nil {
		return m.OutHeight
	}
	return 0
}

type RenderResponse struct {
	Output []byte `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (m *RenderResponse) Reset()                    { *m = RenderResponse{} }
func (m *RenderResponse) String() string            { return proto1.CompactTextString(m) }
func (*RenderResponse) ProtoMessage()               {}
func (*RenderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *RenderResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto1.RegisterType((*NewEaselRequest)(nil), "proto.NewEaselRequest")
	proto1.RegisterType((*NewEaselResponse)(nil), "proto.NewEaselResponse")
	proto1.RegisterType((*DeleteEaselRequest)(nil), "proto.DeleteEaselRequest")
	proto1.RegisterType((*DeleteEaselResponse)(nil), "proto.DeleteEaselResponse")
	proto1.RegisterType((*NewPaletteRequest)(nil), "proto.NewPaletteRequest")
	proto1.RegisterType((*NewPaletteResponse)(nil), "proto.NewPaletteResponse")
	proto1.RegisterType((*DeletePaletteRequest)(nil), "proto.DeletePaletteRequest")
	proto1.RegisterType((*DeletePaletteResponse)(nil), "proto.DeletePaletteResponse")
	proto1.RegisterType((*PingRequest)(nil), "proto.PingRequest")
	proto1.RegisterType((*PongResponse)(nil), "proto.PongResponse")
	proto1.RegisterType((*EaselInfo)(nil), "proto.EaselInfo")
	proto1.RegisterType((*PaletteInfo)(nil), "proto.PaletteInfo")
	proto1.RegisterType((*ListupRequest)(nil), "proto.ListupRequest")
	proto1.RegisterType((*ListupResponse)(nil), "proto.ListupResponse")
	proto1.RegisterType((*VertexAttribute)(nil), "proto.VertexAttribute")
	proto1.RegisterType((*UniformFloatValue)(nil), "proto.UniformFloatValue")
	proto1.RegisterType((*UniformIntValue)(nil), "proto.UniformIntValue")
	proto1.RegisterType((*UniformVariable)(nil), "proto.UniformVariable")
	proto1.RegisterType((*ArrayBuffer)(nil), "proto.ArrayBuffer")
	proto1.RegisterType((*PaletteUpdate)(nil), "proto.PaletteUpdate")
	proto1.RegisterType((*UpdatePaletteRequest)(nil), "proto.UpdatePaletteRequest")
	proto1.RegisterType((*UpdatePaletteResponse)(nil), "proto.UpdatePaletteResponse")
	proto1.RegisterType((*RenderRequest)(nil), "proto.RenderRequest")
	proto1.RegisterType((*RenderResponse)(nil), "proto.RenderResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EaselService service

type EaselServiceClient interface {
	// New/Delete
	NewEasel(ctx context.Context, in *NewEaselRequest, opts ...grpc.CallOption) (*NewEaselResponse, error)
	DeleteEasel(ctx context.Context, in *DeleteEaselRequest, opts ...grpc.CallOption) (*DeleteEaselResponse, error)
	NewPalette(ctx context.Context, in *NewPaletteRequest, opts ...grpc.CallOption) (*NewPaletteResponse, error)
	DeletePalette(ctx context.Context, in *DeletePaletteRequest, opts ...grpc.CallOption) (*DeletePaletteResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
	Listup(ctx context.Context, in *ListupRequest, opts ...grpc.CallOption) (*ListupResponse, error)
	// Setup
	UpdatePalette(ctx context.Context, in *UpdatePaletteRequest, opts ...grpc.CallOption) (*UpdatePaletteResponse, error)
	// Render
	Render(ctx context.Context, in *RenderRequest, opts ...grpc.CallOption) (*RenderResponse, error)
}

type easelServiceClient struct {
	cc *grpc.ClientConn
}

func NewEaselServiceClient(cc *grpc.ClientConn) EaselServiceClient {
	return &easelServiceClient{cc}
}

func (c *easelServiceClient) NewEasel(ctx context.Context, in *NewEaselRequest, opts ...grpc.CallOption) (*NewEaselResponse, error) {
	out := new(NewEaselResponse)
	err := grpc.Invoke(ctx, "/proto.EaselService/NewEasel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *easelServiceClient) DeleteEasel(ctx context.Context, in *DeleteEaselRequest, opts ...grpc.CallOption) (*DeleteEaselResponse, error) {
	out := new(DeleteEaselResponse)
	err := grpc.Invoke(ctx, "/proto.EaselService/DeleteEasel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *easelServiceClient) NewPalette(ctx context.Context, in *NewPaletteRequest, opts ...grpc.CallOption) (*NewPaletteResponse, error) {
	out := new(NewPaletteResponse)
	err := grpc.Invoke(ctx, "/proto.EaselService/NewPalette", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *easelServiceClient) DeletePalette(ctx context.Context, in *DeletePaletteRequest, opts ...grpc.CallOption) (*DeletePaletteResponse, error) {
	out := new(DeletePaletteResponse)
	err := grpc.Invoke(ctx, "/proto.EaselService/DeletePalette", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *easelServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := grpc.Invoke(ctx, "/proto.EaselService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *easelServiceClient) Listup(ctx context.Context, in *ListupRequest, opts ...grpc.CallOption) (*ListupResponse, error) {
	out := new(ListupResponse)
	err := grpc.Invoke(ctx, "/proto.EaselService/Listup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *easelServiceClient) UpdatePalette(ctx context.Context, in *UpdatePaletteRequest, opts ...grpc.CallOption) (*UpdatePaletteResponse, error) {
	out := new(UpdatePaletteResponse)
	err := grpc.Invoke(ctx, "/proto.EaselService/UpdatePalette", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *easelServiceClient) Render(ctx context.Context, in *RenderRequest, opts ...grpc.CallOption) (*RenderResponse, error) {
	out := new(RenderResponse)
	err := grpc.Invoke(ctx, "/proto.EaselService/Render", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EaselService service

type EaselServiceServer interface {
	// New/Delete
	NewEasel(context.Context, *NewEaselRequest) (*NewEaselResponse, error)
	DeleteEasel(context.Context, *DeleteEaselRequest) (*DeleteEaselResponse, error)
	NewPalette(context.Context, *NewPaletteRequest) (*NewPaletteResponse, error)
	DeletePalette(context.Context, *DeletePaletteRequest) (*DeletePaletteResponse, error)
	Ping(context.Context, *PingRequest) (*PongResponse, error)
	Listup(context.Context, *ListupRequest) (*ListupResponse, error)
	// Setup
	UpdatePalette(context.Context, *UpdatePaletteRequest) (*UpdatePaletteResponse, error)
	// Render
	Render(context.Context, *RenderRequest) (*RenderResponse, error)
}

func RegisterEaselServiceServer(s *grpc.Server, srv EaselServiceServer) {
	s.RegisterService(&_EaselService_serviceDesc, srv)
}

func _EaselService_NewEasel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEaselRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EaselServiceServer).NewEasel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EaselService/NewEasel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EaselServiceServer).NewEasel(ctx, req.(*NewEaselRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EaselService_DeleteEasel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEaselRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EaselServiceServer).DeleteEasel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EaselService/DeleteEasel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EaselServiceServer).DeleteEasel(ctx, req.(*DeleteEaselRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EaselService_NewPalette_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPaletteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EaselServiceServer).NewPalette(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EaselService/NewPalette",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EaselServiceServer).NewPalette(ctx, req.(*NewPaletteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EaselService_DeletePalette_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePaletteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EaselServiceServer).DeletePalette(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EaselService/DeletePalette",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EaselServiceServer).DeletePalette(ctx, req.(*DeletePaletteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EaselService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EaselServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EaselService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EaselServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EaselService_Listup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EaselServiceServer).Listup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EaselService/Listup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EaselServiceServer).Listup(ctx, req.(*ListupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EaselService_UpdatePalette_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaletteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EaselServiceServer).UpdatePalette(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EaselService/UpdatePalette",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EaselServiceServer).UpdatePalette(ctx, req.(*UpdatePaletteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EaselService_Render_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EaselServiceServer).Render(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EaselService/Render",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EaselServiceServer).Render(ctx, req.(*RenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EaselService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EaselService",
	HandlerType: (*EaselServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewEasel",
			Handler:    _EaselService_NewEasel_Handler,
		},
		{
			MethodName: "DeleteEasel",
			Handler:    _EaselService_DeleteEasel_Handler,
		},
		{
			MethodName: "NewPalette",
			Handler:    _EaselService_NewPalette_Handler,
		},
		{
			MethodName: "DeletePalette",
			Handler:    _EaselService_DeletePalette_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _EaselService_Ping_Handler,
		},
		{
			MethodName: "Listup",
			Handler:    _EaselService_Listup_Handler,
		},
		{
			MethodName: "UpdatePalette",
			Handler:    _EaselService_UpdatePalette_Handler,
		},
		{
			MethodName: "Render",
			Handler:    _EaselService_Render_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "easel_service.proto",
}

func init() { proto1.RegisterFile("easel_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 942 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0x5d, 0x6f, 0x1b, 0x45,
	0x14, 0x8d, 0xbf, 0xed, 0x6b, 0x3b, 0x4e, 0x26, 0x49, 0xbb, 0x71, 0x41, 0x84, 0xe5, 0x01, 0x3f,
	0x14, 0x23, 0x52, 0x21, 0x04, 0x82, 0x87, 0x50, 0x88, 0x1a, 0x54, 0x45, 0x61, 0xa3, 0x86, 0x47,
	0x6b, 0xdc, 0xbd, 0xeb, 0x0c, 0xb2, 0x77, 0xdd, 0xf9, 0x48, 0xda, 0x3e, 0xf1, 0x3f, 0xf8, 0x07,
	0xbc, 0x20, 0x7e, 0x1c, 0xef, 0x68, 0xbe, 0xd6, 0x5e, 0xdb, 0x6a, 0x03, 0x45, 0x7d, 0xb2, 0xe7,
	0xcc, 0x9d, 0x73, 0xcf, 0xdc, 0x39, 0x7b, 0x67, 0x60, 0x0f, 0xa9, 0xc0, 0xe9, 0x48, 0x20, 0xbf,
	0x61, 0xcf, 0x71, 0x38, 0xe7, 0x99, 0xcc, 0x48, 0xcd, 0xfc, 0x84, 0x0f, 0xa1, 0x77, 0x8e, 0xb7,
	0x3f, 0xea, 0x80, 0x08, 0x5f, 0x28, 0x14, 0x92, 0x1c, 0x42, 0xd3, 0x2e, 0x60, 0x71, 0x50, 0x3a,
	0x2a, 0x0d, 0x5a, 0x51, 0xc3, 0x8c, 0xcf, 0xe2, 0xf0, 0x33, 0xd8, 0x59, 0x44, 0x8b, 0x79, 0x96,
	0x0a, 0x7c, 0x53, 0xf8, 0xe7, 0x40, 0x7e, 0xc0, 0x29, 0x4a, 0xbc, 0x2b, 0xff, 0x01, 0xec, 0x15,
	0x16, 0xd8, 0x14, 0xe1, 0x10, 0x76, 0xcf, 0xf1, 0xf6, 0x82, 0x4e, 0x51, 0x4a, 0xbc, 0x03, 0xcd,
	0x39, 0x90, 0xe5, 0xf8, 0xb7, 0x0a, 0x25, 0x1f, 0x02, 0xcc, 0x6d, 0xb4, 0x9e, 0x2c, 0x9b, 0xc9,
	0x96, 0x43, 0xce, 0xe2, 0xf0, 0x02, 0xf6, 0xad, 0xac, 0x3b, 0x4b, 0x78, 0x1b, 0xe3, 0x7d, 0x38,
	0x58, 0x61, 0x74, 0x5b, 0xa5, 0xd0, 0xbe, 0x60, 0xe9, 0xe4, 0x9d, 0x33, 0x90, 0x00, 0x1a, 0x33,
	0x14, 0x82, 0x4e, 0x30, 0xa8, 0xd8, 0x85, 0x6e, 0x18, 0x8e, 0xa1, 0x73, 0x91, 0xe9, 0x14, 0xef,
	0x5a, 0x97, 0x37, 0xe4, 0xf8, 0x15, 0x5a, 0xe6, 0x08, 0xcf, 0xd2, 0x24, 0x23, 0xdb, 0x50, 0xce,
	0xa9, 0xcb, 0xcc, 0xb0, 0xaa, 0x79, 0x4c, 0x25, 0xc6, 0x23, 0x2a, 0x3d, 0xab, 0x43, 0x4e, 0x24,
	0x19, 0x42, 0xd3, 0xa5, 0x10, 0x41, 0xe5, 0xa8, 0x32, 0x68, 0x1f, 0x13, 0xeb, 0xd9, 0xa1, 0x2b,
	0x96, 0x26, 0x8d, 0xf2, 0x98, 0xf0, 0x5b, 0x68, 0x2f, 0x4d, 0xfc, 0xcb, 0x6c, 0x61, 0x0f, 0xba,
	0x4f, 0x99, 0x90, 0x6a, 0xee, 0x4a, 0x1e, 0x7e, 0x03, 0xdb, 0x1e, 0x70, 0x05, 0x1a, 0x40, 0xdd,
	0x14, 0x44, 0x04, 0x25, 0x23, 0x67, 0xc7, 0xc9, 0xc9, 0x77, 0x18, 0xb9, 0xf9, 0xf0, 0x8f, 0x12,
	0xf4, 0xae, 0x90, 0x4b, 0x7c, 0x79, 0x22, 0x25, 0x67, 0x63, 0x25, 0x91, 0x7c, 0x02, 0x5d, 0xca,
	0x27, 0x6a, 0x86, 0xa9, 0x1c, 0xa5, 0x74, 0x86, 0x4e, 0x5a, 0xc7, 0x83, 0xe7, 0x74, 0x86, 0xe4,
	0x23, 0x68, 0x8f, 0x55, 0x92, 0x20, 0xb7, 0x21, 0x56, 0x25, 0x58, 0xc8, 0x04, 0x7c, 0x0c, 0x1d,
	0x9c, 0xa2, 0x21, 0x11, 0xec, 0xb5, 0xad, 0x77, 0x2d, 0x6a, 0x3b, 0xec, 0x92, 0xbd, 0x46, 0x72,
	0x0f, 0xea, 0x59, 0x92, 0x08, 0x94, 0x41, 0xd5, 0x4c, 0xba, 0x91, 0xc6, 0x85, 0xe4, 0x2c, 0xc6,
	0xa0, 0x66, 0x71, 0x3b, 0x0a, 0x7f, 0x82, 0xdd, 0x67, 0x29, 0x4b, 0x32, 0x3e, 0x3b, 0x9d, 0x66,
	0x54, 0x5e, 0xd1, 0xa9, 0x5a, 0xcf, 0x53, 0x5a, 0xcf, 0x43, 0xa0, 0x1a, 0x53, 0x49, 0x83, 0xf2,
	0x51, 0x65, 0x50, 0x8e, 0xcc, 0xff, 0xf0, 0x09, 0xf4, 0x1c, 0xd7, 0x59, 0xfa, 0xdf, 0x98, 0x6a,
	0x8e, 0xe9, 0xcf, 0x52, 0x4e, 0x75, 0x45, 0x39, 0xa3, 0xe3, 0xa9, 0x89, 0x5b, 0xaa, 0x9c, 0xf9,
	0x4f, 0xbe, 0x86, 0x76, 0xa2, 0x65, 0x8f, 0x6e, 0x74, 0x36, 0x53, 0xb1, 0xf6, 0x71, 0xe0, 0x4e,
	0x66, 0x6d, 0x5f, 0x11, 0x24, 0x8b, 0x3d, 0x3e, 0x82, 0x16, 0x4b, 0xfd, 0xc2, 0x8a, 0x59, 0x78,
	0xaf, 0xb8, 0xd0, 0x6f, 0x22, 0x6a, 0x32, 0xbf, 0x9d, 0x00, 0x1a, 0x12, 0x5f, 0x4a, 0xc5, 0xd1,
	0x94, 0xb7, 0x13, 0xf9, 0x61, 0xf8, 0x25, 0xb4, 0x4f, 0x38, 0xa7, 0xaf, 0xbe, 0x37, 0xa7, 0xb5,
	0x51, 0xec, 0xa6, 0x92, 0xfd, 0x55, 0x86, 0xae, 0xf3, 0xed, 0x33, 0xe3, 0x46, 0xed, 0x94, 0x1b,
	0x63, 0x9e, 0x91, 0xb8, 0xa6, 0x31, 0x72, 0xef, 0x14, 0x0b, 0x5e, 0x1a, 0x8c, 0x7c, 0x0a, 0xbd,
	0x84, 0xd3, 0x89, 0xad, 0xab, 0x0d, 0xb3, 0x6e, 0xd9, 0xf6, 0xb0, 0x0b, 0x7c, 0x08, 0x0d, 0xeb,
	0x9f, 0xd5, 0xaf, 0x68, 0x49, 0x6c, 0xe4, 0x43, 0x48, 0x1f, 0x9a, 0x2c, 0x8d, 0xf1, 0x39, 0x43,
	0x11, 0x54, 0xcd, 0x71, 0xe4, 0x63, 0xf2, 0x18, 0x76, 0x9d, 0x2e, 0xca, 0xb9, 0x34, 0xae, 0x16,
	0x41, 0xcd, 0x70, 0xfa, 0xba, 0xad, 0x98, 0x3e, 0xda, 0xb1, 0x0b, 0x4e, 0xf2, 0x78, 0x4d, 0xa2,
	0x6c, 0x71, 0x47, 0x37, 0xee, 0x5c, 0x45, 0x50, 0x2f, 0x90, 0xac, 0x1c, 0x7b, 0xb4, 0xa3, 0x8a,
	0x80, 0x08, 0x7f, 0x2b, 0xc1, 0xbe, 0x2d, 0xd6, 0xff, 0xd5, 0x89, 0xc9, 0x10, 0x1a, 0xb6, 0x19,
	0x08, 0x67, 0x85, 0xfd, 0x62, 0xb3, 0xb1, 0xe9, 0x22, 0x1f, 0xa4, 0x3b, 0xf7, 0x8a, 0x02, 0xd7,
	0xb9, 0xff, 0x2e, 0x41, 0x37, 0xc2, 0x34, 0x46, 0xfe, 0xde, 0x45, 0xe9, 0xf6, 0x91, 0x29, 0x39,
	0x7a, 0xa1, 0xe8, 0x94, 0xc9, 0x57, 0xc6, 0xa0, 0xe5, 0x08, 0x32, 0x25, 0x7f, 0xb6, 0x88, 0xce,
	0xa7, 0x03, 0x74, 0x35, 0xa9, 0x34, 0x7d, 0xa0, 0x15, 0xb5, 0x32, 0x25, 0x4f, 0x0d, 0x40, 0x1e,
	0x80, 0x1e, 0x8c, 0x6e, 0x59, 0x2c, 0xaf, 0x83, 0xba, 0xf9, 0x50, 0x9b, 0x99, 0x92, 0xbf, 0xe8,
	0xb1, 0x5f, 0x7b, 0x8d, 0x6c, 0x72, 0x2d, 0x83, 0x86, 0x99, 0xd5, 0xe1, 0x4f, 0x0c, 0x10, 0x0e,
	0x60, 0xdb, 0x6f, 0xdb, 0xf5, 0x4b, 0xdd, 0x88, 0x94, 0x9c, 0x2b, 0x69, 0x76, 0xdd, 0x89, 0xdc,
	0xe8, 0xf8, 0xf7, 0x2a, 0x74, 0x4c, 0xcf, 0xbc, 0xb4, 0x2f, 0x11, 0xf2, 0x1d, 0x34, 0xfd, 0x73,
	0x82, 0x78, 0x13, 0xac, 0xbc, 0x46, 0xfa, 0xf7, 0xd7, 0x70, 0x57, 0xef, 0x2d, 0x72, 0x0a, 0xed,
	0xa5, 0xd7, 0x02, 0x39, 0x74, 0x91, 0xeb, 0x4f, 0x8e, 0x7e, 0x7f, 0xd3, 0x54, 0xce, 0xf3, 0x18,
	0x60, 0xf1, 0x5c, 0x20, 0xc1, 0x22, 0x61, 0xd1, 0x64, 0xfd, 0xc3, 0x0d, 0x33, 0x39, 0xc9, 0x53,
	0xe8, 0x16, 0x6e, 0x74, 0xf2, 0xa0, 0x90, 0x73, 0x85, 0xea, 0x83, 0xcd, 0x93, 0x39, 0xdb, 0x17,
	0x50, 0xd5, 0xcf, 0x00, 0x92, 0xdf, 0x7c, 0x8b, 0x37, 0x41, 0x7f, 0xcf, 0x63, 0x4b, 0x97, 0x78,
	0xb8, 0x45, 0xbe, 0x82, 0xba, 0xbd, 0xb7, 0x88, 0x37, 0x4b, 0xe1, 0x5e, 0xeb, 0x1f, 0xac, 0xa0,
	0xcb, 0xca, 0x0b, 0x8e, 0xce, 0x95, 0x6f, 0xfa, 0xd2, 0x72, 0xe5, 0x9b, 0x3f, 0x02, 0x23, 0xc3,
	0xda, 0x21, 0x97, 0x51, 0xf8, 0x28, 0x72, 0x19, 0x45, 0xcf, 0x84, 0x5b, 0xe3, 0xba, 0xc1, 0x1f,
	0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xb5, 0x58, 0x02, 0xae, 0x0a, 0x00, 0x00,
}
