// Code generated by protoc-gen-go.
// source: configuration.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Permission int32

const (
	Permission_PERM_NONE                    Permission = 0
	Permission_PERM_ALL_PERMISSIONS         Permission = 1
	Permission_PERM_ACCOUNT_LIST_LOGS       Permission = 2
	Permission_PERM_ACCOUNT_LIST_MAPS       Permission = 3
	Permission_PERM_LOG_CREATE              Permission = 4
	Permission_PERM_LOG_DELETE              Permission = 5
	Permission_PERM_MAP_CREATE              Permission = 6
	Permission_PERM_MAP_DELETE              Permission = 7
	Permission_PERM_LOG_RAW_ADD             Permission = 8
	Permission_PERM_LOG_READ_ENTRY          Permission = 9
	Permission_PERM_LOG_READ_HASH           Permission = 10
	Permission_PERM_LOG_PROVE_INCLUSION     Permission = 11
	Permission_PERM_MAP_SET_VALUE           Permission = 12
	Permission_PERM_MAP_GET_VALUE           Permission = 13
	Permission_PERM_MAP_MUTATION_READ_ENTRY Permission = 14
	Permission_PERM_MAP_MUTATION_READ_HASH  Permission = 15
)

var Permission_name = map[int32]string{
	0:  "PERM_NONE",
	1:  "PERM_ALL_PERMISSIONS",
	2:  "PERM_ACCOUNT_LIST_LOGS",
	3:  "PERM_ACCOUNT_LIST_MAPS",
	4:  "PERM_LOG_CREATE",
	5:  "PERM_LOG_DELETE",
	6:  "PERM_MAP_CREATE",
	7:  "PERM_MAP_DELETE",
	8:  "PERM_LOG_RAW_ADD",
	9:  "PERM_LOG_READ_ENTRY",
	10: "PERM_LOG_READ_HASH",
	11: "PERM_LOG_PROVE_INCLUSION",
	12: "PERM_MAP_SET_VALUE",
	13: "PERM_MAP_GET_VALUE",
	14: "PERM_MAP_MUTATION_READ_ENTRY",
	15: "PERM_MAP_MUTATION_READ_HASH",
}
var Permission_value = map[string]int32{
	"PERM_NONE":                    0,
	"PERM_ALL_PERMISSIONS":         1,
	"PERM_ACCOUNT_LIST_LOGS":       2,
	"PERM_ACCOUNT_LIST_MAPS":       3,
	"PERM_LOG_CREATE":              4,
	"PERM_LOG_DELETE":              5,
	"PERM_MAP_CREATE":              6,
	"PERM_MAP_DELETE":              7,
	"PERM_LOG_RAW_ADD":             8,
	"PERM_LOG_READ_ENTRY":          9,
	"PERM_LOG_READ_HASH":           10,
	"PERM_LOG_PROVE_INCLUSION":     11,
	"PERM_MAP_SET_VALUE":           12,
	"PERM_MAP_GET_VALUE":           13,
	"PERM_MAP_MUTATION_READ_ENTRY": 14,
	"PERM_MAP_MUTATION_READ_HASH":  15,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}
func (Permission) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ServerConfig struct {
	ServerCertPath           string             `protobuf:"bytes,1,opt,name=server_cert_path,json=serverCertPath" json:"server_cert_path,omitempty"`
	ServerKeyPath            string             `protobuf:"bytes,2,opt,name=server_key_path,json=serverKeyPath" json:"server_key_path,omitempty"`
	RestListenBind           string             `protobuf:"bytes,3,opt,name=rest_listen_bind,json=restListenBind" json:"rest_listen_bind,omitempty"`
	InsecureServerForTesting bool               `protobuf:"varint,4,opt,name=insecure_server_for_testing,json=insecureServerForTesting" json:"insecure_server_for_testing,omitempty"`
	Accounts                 []*ResourceAccount `protobuf:"bytes,5,rep,name=accounts" json:"accounts,omitempty"`
	BoltDbPath               string             `protobuf:"bytes,6,opt,name=bolt_db_path,json=boltDbPath" json:"bolt_db_path,omitempty"`
	GrpcListenProtocol       string             `protobuf:"bytes,7,opt,name=grpc_listen_protocol,json=grpcListenProtocol" json:"grpc_listen_protocol,omitempty"`
	GrpcListenBind           string             `protobuf:"bytes,8,opt,name=grpc_listen_bind,json=grpcListenBind" json:"grpc_listen_bind,omitempty"`
	RestServer               bool               `protobuf:"varint,9,opt,name=rest_server,json=restServer" json:"rest_server,omitempty"`
	GrpcServer               bool               `protobuf:"varint,10,opt,name=grpc_server,json=grpcServer" json:"grpc_server,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ServerConfig) GetServerCertPath() string {
	if m != nil {
		return m.ServerCertPath
	}
	return ""
}

func (m *ServerConfig) GetServerKeyPath() string {
	if m != nil {
		return m.ServerKeyPath
	}
	return ""
}

func (m *ServerConfig) GetRestListenBind() string {
	if m != nil {
		return m.RestListenBind
	}
	return ""
}

func (m *ServerConfig) GetInsecureServerForTesting() bool {
	if m != nil {
		return m.InsecureServerForTesting
	}
	return false
}

func (m *ServerConfig) GetAccounts() []*ResourceAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *ServerConfig) GetBoltDbPath() string {
	if m != nil {
		return m.BoltDbPath
	}
	return ""
}

func (m *ServerConfig) GetGrpcListenProtocol() string {
	if m != nil {
		return m.GrpcListenProtocol
	}
	return ""
}

func (m *ServerConfig) GetGrpcListenBind() string {
	if m != nil {
		return m.GrpcListenBind
	}
	return ""
}

func (m *ServerConfig) GetRestServer() bool {
	if m != nil {
		return m.RestServer
	}
	return false
}

func (m *ServerConfig) GetGrpcServer() bool {
	if m != nil {
		return m.GrpcServer
	}
	return false
}

type AccessPolicy struct {
	ApiKey        string       `protobuf:"bytes,1,opt,name=api_key,json=apiKey" json:"api_key,omitempty"`
	NameMatch     string       `protobuf:"bytes,2,opt,name=name_match,json=nameMatch" json:"name_match,omitempty"`
	AllowedFields []string     `protobuf:"bytes,3,rep,name=allowed_fields,json=allowedFields" json:"allowed_fields,omitempty"`
	Permissions   []Permission `protobuf:"varint,4,rep,packed,name=permissions,enum=com.continusec.verifiabledatastructures.configuration.Permission" json:"permissions,omitempty"`
}

func (m *AccessPolicy) Reset()                    { *m = AccessPolicy{} }
func (m *AccessPolicy) String() string            { return proto.CompactTextString(m) }
func (*AccessPolicy) ProtoMessage()               {}
func (*AccessPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *AccessPolicy) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *AccessPolicy) GetNameMatch() string {
	if m != nil {
		return m.NameMatch
	}
	return ""
}

func (m *AccessPolicy) GetAllowedFields() []string {
	if m != nil {
		return m.AllowedFields
	}
	return nil
}

func (m *AccessPolicy) GetPermissions() []Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type ResourceAccount struct {
	Id     string          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Policy []*AccessPolicy `protobuf:"bytes,2,rep,name=policy" json:"policy,omitempty"`
}

func (m *ResourceAccount) Reset()                    { *m = ResourceAccount{} }
func (m *ResourceAccount) String() string            { return proto.CompactTextString(m) }
func (*ResourceAccount) ProtoMessage()               {}
func (*ResourceAccount) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ResourceAccount) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResourceAccount) GetPolicy() []*AccessPolicy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerConfig)(nil), "com.continusec.verifiabledatastructures.configuration.ServerConfig")
	proto.RegisterType((*AccessPolicy)(nil), "com.continusec.verifiabledatastructures.configuration.AccessPolicy")
	proto.RegisterType((*ResourceAccount)(nil), "com.continusec.verifiabledatastructures.configuration.ResourceAccount")
	proto.RegisterEnum("com.continusec.verifiabledatastructures.configuration.Permission", Permission_name, Permission_value)
}

func init() { proto.RegisterFile("configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xea, 0x46,
	0x14, 0x86, 0x0b, 0x26, 0x04, 0x0e, 0x04, 0xac, 0x49, 0x94, 0x58, 0x4d, 0xaa, 0xa0, 0x48, 0xad,
	0x50, 0x17, 0xa8, 0x4a, 0xd5, 0x65, 0x17, 0x0e, 0x38, 0x09, 0x8a, 0x01, 0xcb, 0x36, 0xa9, 0xda,
	0x2e, 0x46, 0xf6, 0x78, 0x48, 0x46, 0x35, 0xb6, 0x35, 0x33, 0xa4, 0x62, 0xd3, 0xf7, 0xe8, 0x6b,
	0xf5, 0x3d, 0xfa, 0x0e, 0x57, 0x1e, 0x1b, 0x07, 0xae, 0xee, 0xdd, 0x64, 0xc7, 0x7c, 0xe7, 0x3f,
	0x87, 0xf3, 0xff, 0x07, 0xe0, 0x94, 0xa4, 0xc9, 0x8a, 0xbd, 0x6c, 0x78, 0x20, 0x59, 0x9a, 0x8c,
	0x32, 0x9e, 0xca, 0x14, 0xfd, 0x42, 0xd2, 0xf5, 0x88, 0xa4, 0x89, 0x64, 0xc9, 0x46, 0x50, 0x32,
	0x7a, 0xa3, 0x9c, 0xad, 0x58, 0x10, 0xc6, 0x34, 0x0a, 0x64, 0x20, 0x24, 0xdf, 0x10, 0xb9, 0xe1,
	0x54, 0x8c, 0x0e, 0x9a, 0x6f, 0xfe, 0xd7, 0xa0, 0xeb, 0x51, 0xfe, 0x46, 0xf9, 0x58, 0x71, 0x34,
	0x04, 0x5d, 0xa8, 0x37, 0x26, 0x94, 0x4b, 0x9c, 0x05, 0xf2, 0xd5, 0xa8, 0x0d, 0x6a, 0xc3, 0xb6,
	0xdb, 0x2b, 0xf8, 0x98, 0x72, 0xe9, 0x04, 0xf2, 0x15, 0xfd, 0x00, 0xfd, 0x52, 0xf9, 0x17, 0xdd,
	0x16, 0xc2, 0xba, 0x12, 0x9e, 0x14, 0xf8, 0x89, 0x6e, 0x95, 0x6e, 0x08, 0x3a, 0xa7, 0x42, 0xe2,
	0x98, 0x09, 0x49, 0x13, 0x1c, 0xb2, 0x24, 0x32, 0xb4, 0x62, 0x62, 0xce, 0x6d, 0x85, 0xef, 0x58,
	0x12, 0xa1, 0x5f, 0xe1, 0x92, 0x25, 0x82, 0x92, 0x0d, 0xa7, 0xb8, 0x1c, 0xbd, 0x4a, 0x39, 0x96,
	0x54, 0x48, 0x96, 0xbc, 0x18, 0x8d, 0x41, 0x6d, 0xd8, 0x72, 0x8d, 0x9d, 0xa4, 0x58, 0xfb, 0x3e,
	0xe5, 0x7e, 0x51, 0x47, 0x21, 0xb4, 0x02, 0x42, 0xd2, 0x4d, 0x22, 0x85, 0x71, 0x34, 0xd0, 0x86,
	0x9d, 0xdb, 0xfb, 0xd1, 0x87, 0x52, 0x19, 0xb9, 0x54, 0xa4, 0x1b, 0x4e, 0xa8, 0x59, 0x8c, 0x73,
	0xab, 0xb9, 0x68, 0x00, 0xdd, 0x30, 0x8d, 0x25, 0x8e, 0xc2, 0xc2, 0x71, 0x53, 0x19, 0x81, 0x9c,
	0x4d, 0x42, 0x65, 0xf7, 0x27, 0x38, 0x7b, 0xe1, 0x19, 0xd9, 0xd9, 0x55, 0xd7, 0x21, 0x69, 0x6c,
	0x1c, 0x2b, 0x25, 0xca, 0x6b, 0x85, 0x65, 0xa7, 0xac, 0xe4, 0x01, 0xed, 0x77, 0xa8, 0x80, 0x5a,
	0x45, 0x40, 0xef, 0x6a, 0x15, 0xd0, 0x35, 0x74, 0x54, 0x94, 0x45, 0x38, 0x46, 0x5b, 0x05, 0x02,
	0x39, 0x2a, 0xc2, 0xc8, 0x05, 0x6a, 0x54, 0x29, 0x80, 0x42, 0x90, 0xa3, 0x42, 0x70, 0xf3, 0x5f,
	0x0d, 0xba, 0x26, 0x21, 0x54, 0x08, 0x27, 0x8d, 0x19, 0xd9, 0xa2, 0x0b, 0x38, 0x0e, 0x32, 0x96,
	0x9f, 0xb0, 0x3c, 0x73, 0x33, 0xc8, 0xd8, 0x13, 0xdd, 0xa2, 0xef, 0x00, 0x92, 0x60, 0x4d, 0xf1,
	0x3a, 0x90, 0x64, 0x77, 0xd9, 0x76, 0x4e, 0x66, 0x39, 0x40, 0xdf, 0x43, 0x2f, 0x88, 0xe3, 0xf4,
	0x6f, 0x1a, 0xe1, 0x15, 0xa3, 0x71, 0x24, 0x0c, 0x6d, 0xa0, 0xe5, 0xc7, 0x2f, 0xe9, 0xbd, 0x82,
	0x88, 0x40, 0x27, 0xa3, 0x7c, 0xcd, 0x84, 0x60, 0x69, 0x22, 0x8c, 0xc6, 0x40, 0x1b, 0xf6, 0x6e,
	0xcd, 0x0f, 0x9e, 0xc5, 0xa9, 0x26, 0xb9, 0xfb, 0x53, 0x6f, 0xfe, 0x81, 0xfe, 0x67, 0x17, 0x43,
	0x3d, 0xa8, 0xb3, 0xa8, 0x74, 0x54, 0x67, 0x11, 0xfa, 0x13, 0x9a, 0x99, 0x32, 0x6c, 0xd4, 0xd5,
	0x2f, 0x63, 0xfc, 0xc1, 0x15, 0xf6, 0xb3, 0x73, 0xcb, 0x91, 0x3f, 0xfe, 0xab, 0x01, 0xbc, 0xef,
	0x86, 0x4e, 0xa0, 0xed, 0x58, 0xee, 0x0c, 0xcf, 0x17, 0x73, 0x4b, 0xff, 0x06, 0x19, 0x70, 0xa6,
	0x9e, 0xa6, 0x6d, 0xe3, 0xfc, 0xc3, 0xd4, 0xf3, 0xa6, 0x8b, 0xb9, 0xa7, 0xd7, 0xd0, 0xb7, 0x70,
	0x5e, 0x54, 0xc6, 0xe3, 0xc5, 0x72, 0xee, 0x63, 0x7b, 0xea, 0xf9, 0xd8, 0x5e, 0x3c, 0x78, 0x7a,
	0xfd, 0xcb, 0xb5, 0x99, 0xe9, 0x78, 0xba, 0x86, 0x4e, 0xa1, 0xaf, 0x6a, 0xf6, 0xe2, 0x01, 0x8f,
	0x5d, 0xcb, 0xf4, 0x2d, 0xbd, 0x71, 0x00, 0x27, 0x96, 0x6d, 0xf9, 0x96, 0x7e, 0x54, 0xc1, 0x99,
	0xe9, 0xec, 0x94, 0xcd, 0x03, 0x58, 0x2a, 0x8f, 0xd1, 0x19, 0xe8, 0x55, 0xbb, 0x6b, 0xfe, 0x86,
	0xcd, 0xc9, 0x44, 0x6f, 0xa1, 0x0b, 0x38, 0x7d, 0xa7, 0x96, 0x39, 0xc1, 0xd6, 0xdc, 0x77, 0x7f,
	0xd7, 0xdb, 0xe8, 0x1c, 0xd0, 0x61, 0xe1, 0xd1, 0xf4, 0x1e, 0x75, 0x40, 0x57, 0x60, 0x54, 0xdc,
	0x71, 0x17, 0xcf, 0x16, 0x9e, 0xce, 0xc7, 0xf6, 0x32, 0x77, 0xac, 0x77, 0xaa, 0xae, 0xfc, 0x9b,
	0x3d, 0xcb, 0xc7, 0xcf, 0xa6, 0xbd, 0xb4, 0xf4, 0xee, 0x01, 0x7f, 0xa8, 0xf8, 0x09, 0x1a, 0xc0,
	0x55, 0xc5, 0x67, 0x4b, 0xdf, 0xf4, 0xa7, 0x8b, 0xf9, 0xfe, 0x1e, 0x3d, 0x74, 0x0d, 0x97, 0x5f,
	0x51, 0xa8, 0x85, 0xfa, 0x77, 0x8d, 0x3f, 0xea, 0x59, 0x18, 0x36, 0xd5, 0xdf, 0xf0, 0xe7, 0x4f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0xe6, 0x5a, 0xda, 0x3b, 0x05, 0x00, 0x00,
}