// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

/*
Package ploioproto is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	ApplicationListFilter
	ApplicationList
	ApplicationID
	Application
	Service
	Port
	Env
*/
package ploioproto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ApplicationListFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ApplicationListFilter) Reset()                    { *m = ApplicationListFilter{} }
func (m *ApplicationListFilter) String() string            { return proto.CompactTextString(m) }
func (*ApplicationListFilter) ProtoMessage()               {}
func (*ApplicationListFilter) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

func (m *ApplicationListFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplicationList struct {
	Applications []*Application `protobuf:"bytes,1,rep,name=Applications" json:"Applications,omitempty"`
}

func (m *ApplicationList) Reset()                    { *m = ApplicationList{} }
func (m *ApplicationList) String() string            { return proto.CompactTextString(m) }
func (*ApplicationList) ProtoMessage()               {}
func (*ApplicationList) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

func (m *ApplicationList) GetApplications() []*Application {
	if m != nil {
		return m.Applications
	}
	return nil
}

type ApplicationID struct {
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *ApplicationID) Reset()                    { *m = ApplicationID{} }
func (m *ApplicationID) String() string            { return proto.CompactTextString(m) }
func (*ApplicationID) ProtoMessage()               {}
func (*ApplicationID) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{2} }

func (m *ApplicationID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Application struct {
	ID       string     `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" gorm:"primary_key"`
	Name     string     `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Owner    string     `protobuf:"bytes,3,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Repo     string     `protobuf:"bytes,4,opt,name=Repo,proto3" json:"Repo,omitempty"`
	Services []*Service `protobuf:"bytes,5,rep,name=Services" json:"Services,omitempty"`
	Env      []*Env     `protobuf:"bytes,6,rep,name=Env" json:"Env,omitempty"`
}

func (m *Application) Reset()                    { *m = Application{} }
func (m *Application) String() string            { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()               {}
func (*Application) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{3} }

func (m *Application) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Application) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *Application) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Application) GetEnv() []*Env {
	if m != nil {
		return m.Env
	}
	return nil
}

type Service struct {
	ID            string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" gorm:"primary_key"`
	ApplicationID string  `protobuf:"bytes,2,opt,name=ApplicationID,proto3" json:"ApplicationID,omitempty"`
	Name          string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Type          string  `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Ports         []*Port `protobuf:"bytes,5,rep,name=Ports" json:"Ports,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{4} }

func (m *Service) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Service) GetApplicationID() string {
	if m != nil {
		return m.ApplicationID
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Service) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type Port struct {
	ID         string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" gorm:"primary_key"`
	ServiceID  string `protobuf:"bytes,2,opt,name=ServiceID,proto3" json:"ServiceID,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Protocol   string `protobuf:"bytes,4,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	Port       string `protobuf:"bytes,5,opt,name=Port,proto3" json:"Port,omitempty"`
	TargetPort string `protobuf:"bytes,6,opt,name=TargetPort,proto3" json:"TargetPort,omitempty"`
	NodePort   string `protobuf:"bytes,7,opt,name=NodePort,proto3" json:"NodePort,omitempty"`
}

func (m *Port) Reset()                    { *m = Port{} }
func (m *Port) String() string            { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()               {}
func (*Port) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{5} }

func (m *Port) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Port) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Port) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Port) GetTargetPort() string {
	if m != nil {
		return m.TargetPort
	}
	return ""
}

func (m *Port) GetNodePort() string {
	if m != nil {
		return m.NodePort
	}
	return ""
}

type Env struct {
	ID            string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" gorm:"primary_key"`
	ApplicationID string `protobuf:"bytes,2,opt,name=ApplicationID,proto3" json:"ApplicationID,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Type          string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Value         string `protobuf:"bytes,5,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *Env) Reset()                    { *m = Env{} }
func (m *Env) String() string            { return proto.CompactTextString(m) }
func (*Env) ProtoMessage()               {}
func (*Env) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{6} }

func (m *Env) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Env) GetApplicationID() string {
	if m != nil {
		return m.ApplicationID
	}
	return ""
}

func (m *Env) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Env) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Env) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ApplicationListFilter)(nil), "ploioproto.ApplicationListFilter")
	proto.RegisterType((*ApplicationList)(nil), "ploioproto.ApplicationList")
	proto.RegisterType((*ApplicationID)(nil), "ploioproto.ApplicationID")
	proto.RegisterType((*Application)(nil), "ploioproto.Application")
	proto.RegisterType((*Service)(nil), "ploioproto.Service")
	proto.RegisterType((*Port)(nil), "ploioproto.Port")
	proto.RegisterType((*Env)(nil), "ploioproto.Env")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PloioAPI service

type PloioAPIClient interface {
	ListApplications(ctx context.Context, in *ApplicationListFilter, opts ...grpc.CallOption) (*ApplicationList, error)
	GetApplication(ctx context.Context, in *ApplicationID, opts ...grpc.CallOption) (*Application, error)
	UpsertApplication(ctx context.Context, in *Application, opts ...grpc.CallOption) (*Application, error)
	DeleteApplication(ctx context.Context, in *ApplicationID, opts ...grpc.CallOption) (*ApplicationID, error)
}

type ploioAPIClient struct {
	cc *grpc.ClientConn
}

func NewPloioAPIClient(cc *grpc.ClientConn) PloioAPIClient {
	return &ploioAPIClient{cc}
}

func (c *ploioAPIClient) ListApplications(ctx context.Context, in *ApplicationListFilter, opts ...grpc.CallOption) (*ApplicationList, error) {
	out := new(ApplicationList)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/ListApplications", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ploioAPIClient) GetApplication(ctx context.Context, in *ApplicationID, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/GetApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ploioAPIClient) UpsertApplication(ctx context.Context, in *Application, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/UpsertApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ploioAPIClient) DeleteApplication(ctx context.Context, in *ApplicationID, opts ...grpc.CallOption) (*ApplicationID, error) {
	out := new(ApplicationID)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/DeleteApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PloioAPI service

type PloioAPIServer interface {
	ListApplications(context.Context, *ApplicationListFilter) (*ApplicationList, error)
	GetApplication(context.Context, *ApplicationID) (*Application, error)
	UpsertApplication(context.Context, *Application) (*Application, error)
	DeleteApplication(context.Context, *ApplicationID) (*ApplicationID, error)
}

func RegisterPloioAPIServer(s *grpc.Server, srv PloioAPIServer) {
	s.RegisterService(&_PloioAPI_serviceDesc, srv)
}

func _PloioAPI_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationListFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/ListApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).ListApplications(ctx, req.(*ApplicationListFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _PloioAPI_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/GetApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).GetApplication(ctx, req.(*ApplicationID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PloioAPI_UpsertApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).UpsertApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/UpsertApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).UpsertApplication(ctx, req.(*Application))
	}
	return interceptor(ctx, in, info, handler)
}

func _PloioAPI_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/DeleteApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).DeleteApplication(ctx, req.(*ApplicationID))
	}
	return interceptor(ctx, in, info, handler)
}

var _PloioAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ploioproto.PloioAPI",
	HandlerType: (*PloioAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListApplications",
			Handler:    _PloioAPI_ListApplications_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _PloioAPI_GetApplication_Handler,
		},
		{
			MethodName: "UpsertApplication",
			Handler:    _PloioAPI_UpsertApplication_Handler,
		},
		{
			MethodName: "DeleteApplication",
			Handler:    _PloioAPI_DeleteApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x93, 0x38, 0x4d, 0xa6, 0xd0, 0x9f, 0xa1, 0x80, 0x09, 0x88, 0xb6, 0x2b, 0x54, 0x55,
	0x42, 0x24, 0x52, 0xb9, 0xc1, 0xa9, 0xc8, 0x6d, 0x65, 0x81, 0x42, 0x64, 0x02, 0x57, 0xe4, 0x84,
	0xc5, 0xac, 0x70, 0xbc, 0xab, 0x8d, 0x13, 0x94, 0xc7, 0xe0, 0x29, 0x78, 0x0b, 0xae, 0x5c, 0x78,
	0x06, 0x1e, 0x82, 0x27, 0x40, 0xbb, 0xeb, 0xc4, 0xdb, 0x08, 0x47, 0xb9, 0xf5, 0x36, 0xdf, 0x7c,
	0x9f, 0x67, 0xbf, 0x6f, 0xac, 0x81, 0x56, 0x24, 0x58, 0x47, 0x48, 0x9e, 0x71, 0x04, 0x91, 0x70,
	0xc6, 0x75, 0xdd, 0x7e, 0x16, 0xb3, 0xec, 0xcb, 0x74, 0xd8, 0x19, 0xf1, 0x71, 0x37, 0xe6, 0x31,
	0xef, 0xea, 0xf6, 0x70, 0xfa, 0x59, 0x23, 0x0d, 0x74, 0x65, 0x3e, 0x25, 0x4f, 0xe1, 0xee, 0xb9,
	0x10, 0x09, 0x1b, 0x45, 0x19, 0xe3, 0xe9, 0x1b, 0x36, 0xc9, 0x2e, 0x59, 0x92, 0x51, 0x89, 0x08,
	0xf5, 0x34, 0x1a, 0x53, 0xcf, 0x39, 0x72, 0x4e, 0x5b, 0xa1, 0xae, 0x49, 0x0f, 0x76, 0x57, 0xc4,
	0xf8, 0x12, 0x6e, 0x59, 0xad, 0x89, 0xe7, 0x1c, 0xd5, 0x4e, 0xb7, 0xcf, 0xee, 0x77, 0x0a, 0x47,
	0x1d, 0x8b, 0x0f, 0xaf, 0x89, 0xc9, 0x21, 0xdc, 0xb6, 0x70, 0xe0, 0xe3, 0x0e, 0x54, 0x03, 0x3f,
	0x7f, 0xb2, 0x1a, 0xf8, 0xe4, 0x97, 0x03, 0xdb, 0x96, 0x02, 0x4f, 0x0a, 0xfe, 0xd5, 0xbd, 0xbf,
	0x7f, 0x0e, 0x31, 0xe6, 0x72, 0xfc, 0x82, 0x08, 0xc9, 0xc6, 0x91, 0x9c, 0x7f, 0xfc, 0x4a, 0xe7,
	0x44, 0x7d, 0xa7, 0xcc, 0xf7, 0x94, 0xf9, 0xaa, 0x31, 0xaf, 0x6a, 0x3c, 0x00, 0xf7, 0xed, 0xb7,
	0x94, 0x4a, 0xaf, 0xa6, 0x9b, 0x06, 0x28, 0x65, 0x48, 0x05, 0xf7, 0xea, 0x46, 0xa9, 0x6a, 0xec,
	0x42, 0xf3, 0x1d, 0x95, 0x33, 0x36, 0xa2, 0x13, 0xcf, 0xd5, 0x79, 0xee, 0xd8, 0x79, 0x72, 0x2e,
	0x5c, 0x8a, 0xf0, 0x18, 0x6a, 0x17, 0xe9, 0xcc, 0x6b, 0x68, 0xed, 0xae, 0xad, 0xbd, 0x48, 0x67,
	0xa1, 0xe2, 0xc8, 0x0f, 0x07, 0xb6, 0x72, 0xfd, 0xc6, 0x29, 0x9e, 0xac, 0xac, 0x27, 0x8f, 0xb3,
	0xb2, 0xb3, 0x45, 0xd6, 0x9a, 0x95, 0x15, 0xa1, 0x3e, 0x98, 0x0b, 0xba, 0x48, 0xa5, 0x6a, 0x3c,
	0x01, 0xb7, 0xcf, 0x65, 0xb6, 0x88, 0xb4, 0x67, 0xdb, 0x54, 0x44, 0x68, 0x68, 0xf2, 0xdb, 0x81,
	0xba, 0xaa, 0x36, 0xb6, 0xf9, 0x08, 0x5a, 0x79, 0xb2, 0xa5, 0xc5, 0xa2, 0xf1, 0x5f, 0x7b, 0x6d,
	0x68, 0xf6, 0xd5, 0xbb, 0x23, 0x9e, 0xe4, 0x16, 0x97, 0x58, 0xe9, 0xd5, 0xeb, 0x9e, 0x6b, 0xf4,
	0xda, 0xc9, 0x63, 0x80, 0x41, 0x24, 0x63, 0x9a, 0x69, 0xa6, 0xa1, 0x19, 0xab, 0xa3, 0xe6, 0xf5,
	0xf8, 0x27, 0xaa, 0xd9, 0x2d, 0x33, 0x6f, 0x81, 0xc9, 0x77, 0x47, 0xff, 0x9c, 0x1b, 0x5a, 0xfa,
	0x01, 0xb8, 0x1f, 0xa2, 0x64, 0x4a, 0xf3, 0x38, 0x06, 0x9c, 0xfd, 0xac, 0x42, 0xb3, 0xaf, 0xb6,
	0x7f, 0xde, 0x0f, 0x70, 0x00, 0x7b, 0xea, 0x92, 0xec, 0xc3, 0xc0, 0xe3, 0x92, 0xfb, 0x29, 0xee,
	0xb3, 0xfd, 0x70, 0x8d, 0x84, 0x54, 0xf0, 0x12, 0x76, 0xae, 0xa8, 0x3d, 0x14, 0x1f, 0x94, 0x7c,
	0x10, 0xf8, 0xed, 0xb2, 0x73, 0x25, 0x15, 0xbc, 0x82, 0xfd, 0xf7, 0x62, 0x42, 0xe5, 0xb5, 0x51,
	0x65, 0xfa, 0x75, 0x83, 0x5e, 0xc3, 0xbe, 0x4f, 0x13, 0x9a, 0xd1, 0x0d, 0x3d, 0x95, 0x53, 0xa4,
	0x32, 0x6c, 0xe8, 0xf6, 0xf3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xc5, 0x05, 0x89, 0x04,
	0x05, 0x00, 0x00,
}
