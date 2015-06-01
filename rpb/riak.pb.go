// Code generated by protoc-gen-go.
// source: riak.proto
// DO NOT EDIT!

/*
Package riak is a generated protocol buffer package.

It is generated from these files:
	riak.proto

It has these top-level messages:
	RpbErrorResp
	RpbGetServerInfoResp
	RpbPair
	RpbGetBucketReq
	RpbGetBucketResp
	RpbSetBucketReq
	RpbResetBucketReq
	RpbGetBucketTypeReq
	RpbSetBucketTypeReq
	RpbModFun
	RpbCommitHook
	RpbBucketProps
	RpbAuthReq
*/
package riak

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// Used by riak_repl bucket fixup
type RpbBucketProps_RpbReplMode int32

const (
	RpbBucketProps_FALSE    RpbBucketProps_RpbReplMode = 0
	RpbBucketProps_REALTIME RpbBucketProps_RpbReplMode = 1
	RpbBucketProps_FULLSYNC RpbBucketProps_RpbReplMode = 2
	RpbBucketProps_TRUE     RpbBucketProps_RpbReplMode = 3
)

var RpbBucketProps_RpbReplMode_name = map[int32]string{
	0: "FALSE",
	1: "REALTIME",
	2: "FULLSYNC",
	3: "TRUE",
}
var RpbBucketProps_RpbReplMode_value = map[string]int32{
	"FALSE":    0,
	"REALTIME": 1,
	"FULLSYNC": 2,
	"TRUE":     3,
}

func (x RpbBucketProps_RpbReplMode) Enum() *RpbBucketProps_RpbReplMode {
	p := new(RpbBucketProps_RpbReplMode)
	*p = x
	return p
}
func (x RpbBucketProps_RpbReplMode) String() string {
	return proto.EnumName(RpbBucketProps_RpbReplMode_name, int32(x))
}
func (x *RpbBucketProps_RpbReplMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpbBucketProps_RpbReplMode_value, data, "RpbBucketProps_RpbReplMode")
	if err != nil {
		return err
	}
	*x = RpbBucketProps_RpbReplMode(value)
	return nil
}

// Error response - may be generated for any Req
type RpbErrorResp struct {
	Errmsg           []byte  `protobuf:"bytes,1,req,name=errmsg" json:"errmsg,omitempty"`
	Errcode          *uint32 `protobuf:"varint,2,req,name=errcode" json:"errcode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbErrorResp) Reset()         { *m = RpbErrorResp{} }
func (m *RpbErrorResp) String() string { return proto.CompactTextString(m) }
func (*RpbErrorResp) ProtoMessage()    {}

func (m *RpbErrorResp) GetErrmsg() []byte {
	if m != nil {
		return m.Errmsg
	}
	return nil
}

func (m *RpbErrorResp) GetErrcode() uint32 {
	if m != nil && m.Errcode != nil {
		return *m.Errcode
	}
	return 0
}

// Get server info request - no message defined, just send RpbGetServerInfoReq message code
type RpbGetServerInfoResp struct {
	Node             []byte `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	ServerVersion    []byte `protobuf:"bytes,2,opt,name=server_version" json:"server_version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetServerInfoResp) Reset()         { *m = RpbGetServerInfoResp{} }
func (m *RpbGetServerInfoResp) String() string { return proto.CompactTextString(m) }
func (*RpbGetServerInfoResp) ProtoMessage()    {}

func (m *RpbGetServerInfoResp) GetNode() []byte {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *RpbGetServerInfoResp) GetServerVersion() []byte {
	if m != nil {
		return m.ServerVersion
	}
	return nil
}

// Key/value pair - used for user metadata, indexes, search doc fields
type RpbPair struct {
	Key              []byte `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbPair) Reset()         { *m = RpbPair{} }
func (m *RpbPair) String() string { return proto.CompactTextString(m) }
func (*RpbPair) ProtoMessage()    {}

func (m *RpbPair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Get bucket properties request
type RpbGetBucketReq struct {
	Bucket           []byte `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Type             []byte `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetBucketReq) Reset()         { *m = RpbGetBucketReq{} }
func (m *RpbGetBucketReq) String() string { return proto.CompactTextString(m) }
func (*RpbGetBucketReq) ProtoMessage()    {}

func (m *RpbGetBucketReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbGetBucketReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Get bucket properties response
type RpbGetBucketResp struct {
	Props            *RpbBucketProps `protobuf:"bytes,1,req,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbGetBucketResp) Reset()         { *m = RpbGetBucketResp{} }
func (m *RpbGetBucketResp) String() string { return proto.CompactTextString(m) }
func (*RpbGetBucketResp) ProtoMessage()    {}

func (m *RpbGetBucketResp) GetProps() *RpbBucketProps {
	if m != nil {
		return m.Props
	}
	return nil
}

// Set bucket properties request
type RpbSetBucketReq struct {
	Bucket           []byte          `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Props            *RpbBucketProps `protobuf:"bytes,2,req,name=props" json:"props,omitempty"`
	Type             []byte          `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbSetBucketReq) Reset()         { *m = RpbSetBucketReq{} }
func (m *RpbSetBucketReq) String() string { return proto.CompactTextString(m) }
func (*RpbSetBucketReq) ProtoMessage()    {}

func (m *RpbSetBucketReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbSetBucketReq) GetProps() *RpbBucketProps {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *RpbSetBucketReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Reset bucket properties request
type RpbResetBucketReq struct {
	Bucket           []byte `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Type             []byte `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbResetBucketReq) Reset()         { *m = RpbResetBucketReq{} }
func (m *RpbResetBucketReq) String() string { return proto.CompactTextString(m) }
func (*RpbResetBucketReq) ProtoMessage()    {}

func (m *RpbResetBucketReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbResetBucketReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Get bucket properties request
type RpbGetBucketTypeReq struct {
	Type             []byte `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetBucketTypeReq) Reset()         { *m = RpbGetBucketTypeReq{} }
func (m *RpbGetBucketTypeReq) String() string { return proto.CompactTextString(m) }
func (*RpbGetBucketTypeReq) ProtoMessage()    {}

func (m *RpbGetBucketTypeReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Set bucket properties request
type RpbSetBucketTypeReq struct {
	Type             []byte          `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Props            *RpbBucketProps `protobuf:"bytes,2,req,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbSetBucketTypeReq) Reset()         { *m = RpbSetBucketTypeReq{} }
func (m *RpbSetBucketTypeReq) String() string { return proto.CompactTextString(m) }
func (*RpbSetBucketTypeReq) ProtoMessage()    {}

func (m *RpbSetBucketTypeReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *RpbSetBucketTypeReq) GetProps() *RpbBucketProps {
	if m != nil {
		return m.Props
	}
	return nil
}

// Module-Function pairs for commit hooks and other bucket properties
// that take functions
type RpbModFun struct {
	Module           []byte `protobuf:"bytes,1,req,name=module" json:"module,omitempty"`
	Function         []byte `protobuf:"bytes,2,req,name=function" json:"function,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbModFun) Reset()         { *m = RpbModFun{} }
func (m *RpbModFun) String() string { return proto.CompactTextString(m) }
func (*RpbModFun) ProtoMessage()    {}

func (m *RpbModFun) GetModule() []byte {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *RpbModFun) GetFunction() []byte {
	if m != nil {
		return m.Function
	}
	return nil
}

// A commit hook, which may either be a modfun or a JavaScript named
// function
type RpbCommitHook struct {
	Modfun           *RpbModFun `protobuf:"bytes,1,opt,name=modfun" json:"modfun,omitempty"`
	Name             []byte     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RpbCommitHook) Reset()         { *m = RpbCommitHook{} }
func (m *RpbCommitHook) String() string { return proto.CompactTextString(m) }
func (*RpbCommitHook) ProtoMessage()    {}

func (m *RpbCommitHook) GetModfun() *RpbModFun {
	if m != nil {
		return m.Modfun
	}
	return nil
}

func (m *RpbCommitHook) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

// Bucket properties
type RpbBucketProps struct {
	// Declared in riak_core_app
	NVal          *uint32          `protobuf:"varint,1,opt,name=n_val" json:"n_val,omitempty"`
	AllowMult     *bool            `protobuf:"varint,2,opt,name=allow_mult" json:"allow_mult,omitempty"`
	LastWriteWins *bool            `protobuf:"varint,3,opt,name=last_write_wins" json:"last_write_wins,omitempty"`
	Precommit     []*RpbCommitHook `protobuf:"bytes,4,rep,name=precommit" json:"precommit,omitempty"`
	HasPrecommit  *bool            `protobuf:"varint,5,opt,name=has_precommit,def=0" json:"has_precommit,omitempty"`
	Postcommit    []*RpbCommitHook `protobuf:"bytes,6,rep,name=postcommit" json:"postcommit,omitempty"`
	HasPostcommit *bool            `protobuf:"varint,7,opt,name=has_postcommit,def=0" json:"has_postcommit,omitempty"`
	ChashKeyfun   *RpbModFun       `protobuf:"bytes,8,opt,name=chash_keyfun" json:"chash_keyfun,omitempty"`
	// Declared in riak_kv_app
	Linkfun     *RpbModFun `protobuf:"bytes,9,opt,name=linkfun" json:"linkfun,omitempty"`
	OldVclock   *uint32    `protobuf:"varint,10,opt,name=old_vclock" json:"old_vclock,omitempty"`
	YoungVclock *uint32    `protobuf:"varint,11,opt,name=young_vclock" json:"young_vclock,omitempty"`
	BigVclock   *uint32    `protobuf:"varint,12,opt,name=big_vclock" json:"big_vclock,omitempty"`
	SmallVclock *uint32    `protobuf:"varint,13,opt,name=small_vclock" json:"small_vclock,omitempty"`
	Pr          *uint32    `protobuf:"varint,14,opt,name=pr" json:"pr,omitempty"`
	R           *uint32    `protobuf:"varint,15,opt,name=r" json:"r,omitempty"`
	W           *uint32    `protobuf:"varint,16,opt,name=w" json:"w,omitempty"`
	Pw          *uint32    `protobuf:"varint,17,opt,name=pw" json:"pw,omitempty"`
	Dw          *uint32    `protobuf:"varint,18,opt,name=dw" json:"dw,omitempty"`
	Rw          *uint32    `protobuf:"varint,19,opt,name=rw" json:"rw,omitempty"`
	BasicQuorum *bool      `protobuf:"varint,20,opt,name=basic_quorum" json:"basic_quorum,omitempty"`
	NotfoundOk  *bool      `protobuf:"varint,21,opt,name=notfound_ok" json:"notfound_ok,omitempty"`
	// Used by riak_kv_multi_backend
	Backend []byte `protobuf:"bytes,22,opt,name=backend" json:"backend,omitempty"`
	// Used by riak_search bucket fixup
	Search *bool                       `protobuf:"varint,23,opt,name=search" json:"search,omitempty"`
	Repl   *RpbBucketProps_RpbReplMode `protobuf:"varint,24,opt,name=repl,enum=RpbBucketProps_RpbReplMode" json:"repl,omitempty"`
	// Search index
	SearchIndex []byte `protobuf:"bytes,25,opt,name=search_index" json:"search_index,omitempty"`
	// KV Datatypes
	Datatype []byte `protobuf:"bytes,26,opt,name=datatype" json:"datatype,omitempty"`
	// KV strong consistency
	Consistent *bool `protobuf:"varint,27,opt,name=consistent" json:"consistent,omitempty"`
	// KV fast path
	WriteOnce        *bool  `protobuf:"varint,28,opt,name=write_once" json:"write_once,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbBucketProps) Reset()         { *m = RpbBucketProps{} }
func (m *RpbBucketProps) String() string { return proto.CompactTextString(m) }
func (*RpbBucketProps) ProtoMessage()    {}

const Default_RpbBucketProps_HasPrecommit bool = false
const Default_RpbBucketProps_HasPostcommit bool = false

func (m *RpbBucketProps) GetNVal() uint32 {
	if m != nil && m.NVal != nil {
		return *m.NVal
	}
	return 0
}

func (m *RpbBucketProps) GetAllowMult() bool {
	if m != nil && m.AllowMult != nil {
		return *m.AllowMult
	}
	return false
}

func (m *RpbBucketProps) GetLastWriteWins() bool {
	if m != nil && m.LastWriteWins != nil {
		return *m.LastWriteWins
	}
	return false
}

func (m *RpbBucketProps) GetPrecommit() []*RpbCommitHook {
	if m != nil {
		return m.Precommit
	}
	return nil
}

func (m *RpbBucketProps) GetHasPrecommit() bool {
	if m != nil && m.HasPrecommit != nil {
		return *m.HasPrecommit
	}
	return Default_RpbBucketProps_HasPrecommit
}

func (m *RpbBucketProps) GetPostcommit() []*RpbCommitHook {
	if m != nil {
		return m.Postcommit
	}
	return nil
}

func (m *RpbBucketProps) GetHasPostcommit() bool {
	if m != nil && m.HasPostcommit != nil {
		return *m.HasPostcommit
	}
	return Default_RpbBucketProps_HasPostcommit
}

func (m *RpbBucketProps) GetChashKeyfun() *RpbModFun {
	if m != nil {
		return m.ChashKeyfun
	}
	return nil
}

func (m *RpbBucketProps) GetLinkfun() *RpbModFun {
	if m != nil {
		return m.Linkfun
	}
	return nil
}

func (m *RpbBucketProps) GetOldVclock() uint32 {
	if m != nil && m.OldVclock != nil {
		return *m.OldVclock
	}
	return 0
}

func (m *RpbBucketProps) GetYoungVclock() uint32 {
	if m != nil && m.YoungVclock != nil {
		return *m.YoungVclock
	}
	return 0
}

func (m *RpbBucketProps) GetBigVclock() uint32 {
	if m != nil && m.BigVclock != nil {
		return *m.BigVclock
	}
	return 0
}

func (m *RpbBucketProps) GetSmallVclock() uint32 {
	if m != nil && m.SmallVclock != nil {
		return *m.SmallVclock
	}
	return 0
}

func (m *RpbBucketProps) GetPr() uint32 {
	if m != nil && m.Pr != nil {
		return *m.Pr
	}
	return 0
}

func (m *RpbBucketProps) GetR() uint32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *RpbBucketProps) GetW() uint32 {
	if m != nil && m.W != nil {
		return *m.W
	}
	return 0
}

func (m *RpbBucketProps) GetPw() uint32 {
	if m != nil && m.Pw != nil {
		return *m.Pw
	}
	return 0
}

func (m *RpbBucketProps) GetDw() uint32 {
	if m != nil && m.Dw != nil {
		return *m.Dw
	}
	return 0
}

func (m *RpbBucketProps) GetRw() uint32 {
	if m != nil && m.Rw != nil {
		return *m.Rw
	}
	return 0
}

func (m *RpbBucketProps) GetBasicQuorum() bool {
	if m != nil && m.BasicQuorum != nil {
		return *m.BasicQuorum
	}
	return false
}

func (m *RpbBucketProps) GetNotfoundOk() bool {
	if m != nil && m.NotfoundOk != nil {
		return *m.NotfoundOk
	}
	return false
}

func (m *RpbBucketProps) GetBackend() []byte {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *RpbBucketProps) GetSearch() bool {
	if m != nil && m.Search != nil {
		return *m.Search
	}
	return false
}

func (m *RpbBucketProps) GetRepl() RpbBucketProps_RpbReplMode {
	if m != nil && m.Repl != nil {
		return *m.Repl
	}
	return RpbBucketProps_FALSE
}

func (m *RpbBucketProps) GetSearchIndex() []byte {
	if m != nil {
		return m.SearchIndex
	}
	return nil
}

func (m *RpbBucketProps) GetDatatype() []byte {
	if m != nil {
		return m.Datatype
	}
	return nil
}

func (m *RpbBucketProps) GetConsistent() bool {
	if m != nil && m.Consistent != nil {
		return *m.Consistent
	}
	return false
}

func (m *RpbBucketProps) GetWriteOnce() bool {
	if m != nil && m.WriteOnce != nil {
		return *m.WriteOnce
	}
	return false
}

// Authentication request
type RpbAuthReq struct {
	User             []byte `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	Password         []byte `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbAuthReq) Reset()         { *m = RpbAuthReq{} }
func (m *RpbAuthReq) String() string { return proto.CompactTextString(m) }
func (*RpbAuthReq) ProtoMessage()    {}

func (m *RpbAuthReq) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RpbAuthReq) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func init() {
	proto.RegisterEnum("RpbBucketProps_RpbReplMode", RpbBucketProps_RpbReplMode_name, RpbBucketProps_RpbReplMode_value)
}
