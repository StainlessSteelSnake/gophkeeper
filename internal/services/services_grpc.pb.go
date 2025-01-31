// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: services.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GophKeeper_Register_FullMethodName            = "/services.GophKeeper/Register"
	GophKeeper_Login_FullMethodName               = "/services.GophKeeper/Login"
	GophKeeper_Logout_FullMethodName              = "/services.GophKeeper/Logout"
	GophKeeper_GetUserRecords_FullMethodName      = "/services.GophKeeper/GetUserRecords"
	GophKeeper_GetUserRecord_FullMethodName       = "/services.GophKeeper/GetUserRecord"
	GophKeeper_ChangeUserRecord_FullMethodName    = "/services.GophKeeper/ChangeUserRecord"
	GophKeeper_DeleteUserRecord_FullMethodName    = "/services.GophKeeper/DeleteUserRecord"
	GophKeeper_AddLoginPassword_FullMethodName    = "/services.GophKeeper/AddLoginPassword"
	GophKeeper_GetLoginPassword_FullMethodName    = "/services.GophKeeper/GetLoginPassword"
	GophKeeper_ChangeLoginPassword_FullMethodName = "/services.GophKeeper/ChangeLoginPassword"
	GophKeeper_AddBankCard_FullMethodName         = "/services.GophKeeper/AddBankCard"
	GophKeeper_GetBankCard_FullMethodName         = "/services.GophKeeper/GetBankCard"
	GophKeeper_ChangeBankCard_FullMethodName      = "/services.GophKeeper/ChangeBankCard"
	GophKeeper_AddText_FullMethodName             = "/services.GophKeeper/AddText"
	GophKeeper_GetText_FullMethodName             = "/services.GophKeeper/GetText"
	GophKeeper_ChangeText_FullMethodName          = "/services.GophKeeper/ChangeText"
	GophKeeper_AddBytes_FullMethodName            = "/services.GophKeeper/AddBytes"
	GophKeeper_GetBytes_FullMethodName            = "/services.GophKeeper/GetBytes"
	GophKeeper_ChangeBytes_FullMethodName         = "/services.GophKeeper/ChangeBytes"
)

// GophKeeperClient is the client API for GophKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophKeeperClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	GetUserRecords(ctx context.Context, in *GetUserRecordsRequest, opts ...grpc.CallOption) (*GetUserRecordsResponse, error)
	GetUserRecord(ctx context.Context, in *GetUserRecordRequest, opts ...grpc.CallOption) (*GetUserRecordResponse, error)
	ChangeUserRecord(ctx context.Context, in *ChangeUserRecordRequest, opts ...grpc.CallOption) (*ChangeUserRecordResponse, error)
	DeleteUserRecord(ctx context.Context, in *DeleteUserRecordRequest, opts ...grpc.CallOption) (*DeleteUserRecordResponse, error)
	AddLoginPassword(ctx context.Context, in *AddLoginPasswordRequest, opts ...grpc.CallOption) (*AddLoginPasswordResponse, error)
	GetLoginPassword(ctx context.Context, in *GetLoginPasswordRequest, opts ...grpc.CallOption) (*GetLoginPasswordResponse, error)
	ChangeLoginPassword(ctx context.Context, in *ChangeLoginPasswordRequest, opts ...grpc.CallOption) (*ChangeLoginPasswordResponse, error)
	AddBankCard(ctx context.Context, in *AddBankCardRequest, opts ...grpc.CallOption) (*AddBankCardResponse, error)
	GetBankCard(ctx context.Context, in *GetBankCardRequest, opts ...grpc.CallOption) (*GetBankCardResponse, error)
	ChangeBankCard(ctx context.Context, in *ChangeBankCardRequest, opts ...grpc.CallOption) (*ChangeBankCardResponse, error)
	AddText(ctx context.Context, in *AddTextRequest, opts ...grpc.CallOption) (*AddTextResponse, error)
	GetText(ctx context.Context, in *GetTextRequest, opts ...grpc.CallOption) (*GetTextResponse, error)
	ChangeText(ctx context.Context, in *ChangeTextRequest, opts ...grpc.CallOption) (*ChangeTextResponse, error)
	AddBytes(ctx context.Context, in *AddBytesRequest, opts ...grpc.CallOption) (*AddBytesResponse, error)
	GetBytes(ctx context.Context, in *GetBytesRequest, opts ...grpc.CallOption) (*GetBytesResponse, error)
	ChangeBytes(ctx context.Context, in *ChangeBytesRequest, opts ...grpc.CallOption) (*ChangeBytesResponse, error)
}

type gophKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophKeeperClient(cc grpc.ClientConnInterface) GophKeeperClient {
	return &gophKeeperClient{cc}
}

func (c *gophKeeperClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, GophKeeper_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, GophKeeper_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, GophKeeper_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetUserRecords(ctx context.Context, in *GetUserRecordsRequest, opts ...grpc.CallOption) (*GetUserRecordsResponse, error) {
	out := new(GetUserRecordsResponse)
	err := c.cc.Invoke(ctx, GophKeeper_GetUserRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetUserRecord(ctx context.Context, in *GetUserRecordRequest, opts ...grpc.CallOption) (*GetUserRecordResponse, error) {
	out := new(GetUserRecordResponse)
	err := c.cc.Invoke(ctx, GophKeeper_GetUserRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) ChangeUserRecord(ctx context.Context, in *ChangeUserRecordRequest, opts ...grpc.CallOption) (*ChangeUserRecordResponse, error) {
	out := new(ChangeUserRecordResponse)
	err := c.cc.Invoke(ctx, GophKeeper_ChangeUserRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) DeleteUserRecord(ctx context.Context, in *DeleteUserRecordRequest, opts ...grpc.CallOption) (*DeleteUserRecordResponse, error) {
	out := new(DeleteUserRecordResponse)
	err := c.cc.Invoke(ctx, GophKeeper_DeleteUserRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) AddLoginPassword(ctx context.Context, in *AddLoginPasswordRequest, opts ...grpc.CallOption) (*AddLoginPasswordResponse, error) {
	out := new(AddLoginPasswordResponse)
	err := c.cc.Invoke(ctx, GophKeeper_AddLoginPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetLoginPassword(ctx context.Context, in *GetLoginPasswordRequest, opts ...grpc.CallOption) (*GetLoginPasswordResponse, error) {
	out := new(GetLoginPasswordResponse)
	err := c.cc.Invoke(ctx, GophKeeper_GetLoginPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) ChangeLoginPassword(ctx context.Context, in *ChangeLoginPasswordRequest, opts ...grpc.CallOption) (*ChangeLoginPasswordResponse, error) {
	out := new(ChangeLoginPasswordResponse)
	err := c.cc.Invoke(ctx, GophKeeper_ChangeLoginPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) AddBankCard(ctx context.Context, in *AddBankCardRequest, opts ...grpc.CallOption) (*AddBankCardResponse, error) {
	out := new(AddBankCardResponse)
	err := c.cc.Invoke(ctx, GophKeeper_AddBankCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetBankCard(ctx context.Context, in *GetBankCardRequest, opts ...grpc.CallOption) (*GetBankCardResponse, error) {
	out := new(GetBankCardResponse)
	err := c.cc.Invoke(ctx, GophKeeper_GetBankCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) ChangeBankCard(ctx context.Context, in *ChangeBankCardRequest, opts ...grpc.CallOption) (*ChangeBankCardResponse, error) {
	out := new(ChangeBankCardResponse)
	err := c.cc.Invoke(ctx, GophKeeper_ChangeBankCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) AddText(ctx context.Context, in *AddTextRequest, opts ...grpc.CallOption) (*AddTextResponse, error) {
	out := new(AddTextResponse)
	err := c.cc.Invoke(ctx, GophKeeper_AddText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetText(ctx context.Context, in *GetTextRequest, opts ...grpc.CallOption) (*GetTextResponse, error) {
	out := new(GetTextResponse)
	err := c.cc.Invoke(ctx, GophKeeper_GetText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) ChangeText(ctx context.Context, in *ChangeTextRequest, opts ...grpc.CallOption) (*ChangeTextResponse, error) {
	out := new(ChangeTextResponse)
	err := c.cc.Invoke(ctx, GophKeeper_ChangeText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) AddBytes(ctx context.Context, in *AddBytesRequest, opts ...grpc.CallOption) (*AddBytesResponse, error) {
	out := new(AddBytesResponse)
	err := c.cc.Invoke(ctx, GophKeeper_AddBytes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetBytes(ctx context.Context, in *GetBytesRequest, opts ...grpc.CallOption) (*GetBytesResponse, error) {
	out := new(GetBytesResponse)
	err := c.cc.Invoke(ctx, GophKeeper_GetBytes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) ChangeBytes(ctx context.Context, in *ChangeBytesRequest, opts ...grpc.CallOption) (*ChangeBytesResponse, error) {
	out := new(ChangeBytesResponse)
	err := c.cc.Invoke(ctx, GophKeeper_ChangeBytes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophKeeperServer is the server API for GophKeeper service.
// All implementations must embed UnimplementedGophKeeperServer
// for forward compatibility
type GophKeeperServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	GetUserRecords(context.Context, *GetUserRecordsRequest) (*GetUserRecordsResponse, error)
	GetUserRecord(context.Context, *GetUserRecordRequest) (*GetUserRecordResponse, error)
	ChangeUserRecord(context.Context, *ChangeUserRecordRequest) (*ChangeUserRecordResponse, error)
	DeleteUserRecord(context.Context, *DeleteUserRecordRequest) (*DeleteUserRecordResponse, error)
	AddLoginPassword(context.Context, *AddLoginPasswordRequest) (*AddLoginPasswordResponse, error)
	GetLoginPassword(context.Context, *GetLoginPasswordRequest) (*GetLoginPasswordResponse, error)
	ChangeLoginPassword(context.Context, *ChangeLoginPasswordRequest) (*ChangeLoginPasswordResponse, error)
	AddBankCard(context.Context, *AddBankCardRequest) (*AddBankCardResponse, error)
	GetBankCard(context.Context, *GetBankCardRequest) (*GetBankCardResponse, error)
	ChangeBankCard(context.Context, *ChangeBankCardRequest) (*ChangeBankCardResponse, error)
	AddText(context.Context, *AddTextRequest) (*AddTextResponse, error)
	GetText(context.Context, *GetTextRequest) (*GetTextResponse, error)
	ChangeText(context.Context, *ChangeTextRequest) (*ChangeTextResponse, error)
	AddBytes(context.Context, *AddBytesRequest) (*AddBytesResponse, error)
	GetBytes(context.Context, *GetBytesRequest) (*GetBytesResponse, error)
	ChangeBytes(context.Context, *ChangeBytesRequest) (*ChangeBytesResponse, error)
	mustEmbedUnimplementedGophKeeperServer()
}

// UnimplementedGophKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophKeeperServer struct {
}

func (UnimplementedGophKeeperServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGophKeeperServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGophKeeperServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedGophKeeperServer) GetUserRecords(context.Context, *GetUserRecordsRequest) (*GetUserRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRecords not implemented")
}
func (UnimplementedGophKeeperServer) GetUserRecord(context.Context, *GetUserRecordRequest) (*GetUserRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRecord not implemented")
}
func (UnimplementedGophKeeperServer) ChangeUserRecord(context.Context, *ChangeUserRecordRequest) (*ChangeUserRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserRecord not implemented")
}
func (UnimplementedGophKeeperServer) DeleteUserRecord(context.Context, *DeleteUserRecordRequest) (*DeleteUserRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserRecord not implemented")
}
func (UnimplementedGophKeeperServer) AddLoginPassword(context.Context, *AddLoginPasswordRequest) (*AddLoginPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLoginPassword not implemented")
}
func (UnimplementedGophKeeperServer) GetLoginPassword(context.Context, *GetLoginPasswordRequest) (*GetLoginPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginPassword not implemented")
}
func (UnimplementedGophKeeperServer) ChangeLoginPassword(context.Context, *ChangeLoginPasswordRequest) (*ChangeLoginPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeLoginPassword not implemented")
}
func (UnimplementedGophKeeperServer) AddBankCard(context.Context, *AddBankCardRequest) (*AddBankCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBankCard not implemented")
}
func (UnimplementedGophKeeperServer) GetBankCard(context.Context, *GetBankCardRequest) (*GetBankCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankCard not implemented")
}
func (UnimplementedGophKeeperServer) ChangeBankCard(context.Context, *ChangeBankCardRequest) (*ChangeBankCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeBankCard not implemented")
}
func (UnimplementedGophKeeperServer) AddText(context.Context, *AddTextRequest) (*AddTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddText not implemented")
}
func (UnimplementedGophKeeperServer) GetText(context.Context, *GetTextRequest) (*GetTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetText not implemented")
}
func (UnimplementedGophKeeperServer) ChangeText(context.Context, *ChangeTextRequest) (*ChangeTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeText not implemented")
}
func (UnimplementedGophKeeperServer) AddBytes(context.Context, *AddBytesRequest) (*AddBytesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBytes not implemented")
}
func (UnimplementedGophKeeperServer) GetBytes(context.Context, *GetBytesRequest) (*GetBytesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBytes not implemented")
}
func (UnimplementedGophKeeperServer) ChangeBytes(context.Context, *ChangeBytesRequest) (*ChangeBytesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeBytes not implemented")
}
func (UnimplementedGophKeeperServer) mustEmbedUnimplementedGophKeeperServer() {}

// UnsafeGophKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophKeeperServer will
// result in compilation errors.
type UnsafeGophKeeperServer interface {
	mustEmbedUnimplementedGophKeeperServer()
}

func RegisterGophKeeperServer(s grpc.ServiceRegistrar, srv GophKeeperServer) {
	s.RegisterService(&GophKeeper_ServiceDesc, srv)
}

func _GophKeeper_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetUserRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetUserRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetUserRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetUserRecords(ctx, req.(*GetUserRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetUserRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetUserRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetUserRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetUserRecord(ctx, req.(*GetUserRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_ChangeUserRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).ChangeUserRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_ChangeUserRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).ChangeUserRecord(ctx, req.(*ChangeUserRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_DeleteUserRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).DeleteUserRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_DeleteUserRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).DeleteUserRecord(ctx, req.(*DeleteUserRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_AddLoginPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLoginPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).AddLoginPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_AddLoginPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).AddLoginPassword(ctx, req.(*AddLoginPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetLoginPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetLoginPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetLoginPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetLoginPassword(ctx, req.(*GetLoginPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_ChangeLoginPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeLoginPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).ChangeLoginPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_ChangeLoginPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).ChangeLoginPassword(ctx, req.(*ChangeLoginPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_AddBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBankCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).AddBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_AddBankCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).AddBankCard(ctx, req.(*AddBankCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetBankCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetBankCard(ctx, req.(*GetBankCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_ChangeBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeBankCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).ChangeBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_ChangeBankCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).ChangeBankCard(ctx, req.(*ChangeBankCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_AddText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).AddText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_AddText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).AddText(ctx, req.(*AddTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetText(ctx, req.(*GetTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_ChangeText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).ChangeText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_ChangeText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).ChangeText(ctx, req.(*ChangeTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_AddBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBytesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).AddBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_AddBytes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).AddBytes(ctx, req.(*AddBytesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBytesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetBytes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetBytes(ctx, req.(*GetBytesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_ChangeBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeBytesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).ChangeBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_ChangeBytes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).ChangeBytes(ctx, req.(*ChangeBytesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GophKeeper_ServiceDesc is the grpc.ServiceDesc for GophKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GophKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.GophKeeper",
	HandlerType: (*GophKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _GophKeeper_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _GophKeeper_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _GophKeeper_Logout_Handler,
		},
		{
			MethodName: "GetUserRecords",
			Handler:    _GophKeeper_GetUserRecords_Handler,
		},
		{
			MethodName: "GetUserRecord",
			Handler:    _GophKeeper_GetUserRecord_Handler,
		},
		{
			MethodName: "ChangeUserRecord",
			Handler:    _GophKeeper_ChangeUserRecord_Handler,
		},
		{
			MethodName: "DeleteUserRecord",
			Handler:    _GophKeeper_DeleteUserRecord_Handler,
		},
		{
			MethodName: "AddLoginPassword",
			Handler:    _GophKeeper_AddLoginPassword_Handler,
		},
		{
			MethodName: "GetLoginPassword",
			Handler:    _GophKeeper_GetLoginPassword_Handler,
		},
		{
			MethodName: "ChangeLoginPassword",
			Handler:    _GophKeeper_ChangeLoginPassword_Handler,
		},
		{
			MethodName: "AddBankCard",
			Handler:    _GophKeeper_AddBankCard_Handler,
		},
		{
			MethodName: "GetBankCard",
			Handler:    _GophKeeper_GetBankCard_Handler,
		},
		{
			MethodName: "ChangeBankCard",
			Handler:    _GophKeeper_ChangeBankCard_Handler,
		},
		{
			MethodName: "AddText",
			Handler:    _GophKeeper_AddText_Handler,
		},
		{
			MethodName: "GetText",
			Handler:    _GophKeeper_GetText_Handler,
		},
		{
			MethodName: "ChangeText",
			Handler:    _GophKeeper_ChangeText_Handler,
		},
		{
			MethodName: "AddBytes",
			Handler:    _GophKeeper_AddBytes_Handler,
		},
		{
			MethodName: "GetBytes",
			Handler:    _GophKeeper_GetBytes_Handler,
		},
		{
			MethodName: "ChangeBytes",
			Handler:    _GophKeeper_ChangeBytes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
