package kproto

import (
	"errors"
	"fmt"
)

type PackageDesc struct {
	Lang string
	Name string
}

type FieldDesc struct {
	Name   string
	Type   string
	Size   int
	Length int
}

type MessageDesc struct {
	Id     uint16
	Name   string
	Fields []*FieldDesc
	owner  *FileDesc
}

func (msg *MessageDesc) AddField(name, typ string, l, sz int) {
	field := &FieldDesc{}
	field.Name = name
	field.Type = typ
	field.Length = l
	field.Size = sz

	msg.Fields = append(msg.Fields, field)
}

func (msg *MessageDesc) MaxSize() (n int, err error) {
	for _, v := range msg.Fields {
		fsize := 0
		switch v.Type {
		case "bool", "byte", "int8", "uint8", "string":
			fsize = 1
		case "int16", "uint16":
			fsize = 2
		case "int32", "uint32", "float32":
			fsize = 4
		case "int64", "uint64", "float64":
			fsize = 8
		default:
			if v.Type == msg.Name {
				err = errors.New("字段类型不能是自身")
				return
			}

			other := msg.owner.GetMessage(v.Type)
			if other == nil {
				err = fmt.Errorf("无效字段类型: %s:%s", msg.Name, v.Type)
				return
			}

			fsize, err = other.MaxSize()
			if err != nil {
				return
			}
		}

		if v.Length > 0 {
			fsize *= int(v.Length)
			fsize += 2
		}

		n += fsize
	}

	return
}

type MethodDesc struct {
	Name    string
	Request string
	Reply   string
}

type RPCDesc struct {
	Name    string
	Methods []*MethodDesc
}

func (rpc *RPCDesc) AddMethod(name, request, reply string) {
	method := &MethodDesc{}
	method.Name = name
	method.Request = request
	method.Reply = reply

	rpc.Methods = append(rpc.Methods, method)
}

type FileDesc struct {
	FileName    string
	IDCounter   uint16
	PackageName string
	Packages    []*PackageDesc
	Messages    []*MessageDesc
	RPCs        []*RPCDesc
}

func (f *FileDesc) AddPackage(lang string, name string) *PackageDesc {
	pkg := &PackageDesc{}
	pkg.Lang = lang
	pkg.Name = name

	f.Packages = append(f.Packages, pkg)

	return pkg
}

func (f *FileDesc) SetPackage(lang string) {
	for _, v := range f.Packages {
		if v.Lang == lang {
			f.PackageName = v.Name
		}
	}
}

func (f *FileDesc) AddMessage(id uint16, name string) *MessageDesc {
	msg := &MessageDesc{}
	msg.Id = id
	msg.Name = name
	msg.Fields = make([]*FieldDesc, 0)
	msg.owner = f

	f.Messages = append(f.Messages, msg)

	return msg
}

func (f *FileDesc) GetMessage(name string) *MessageDesc {
	for _, v := range f.Messages {
		if v.Name == name {
			return v
		}
	}

	return nil
}

func (f *FileDesc) AddRPC(name string) *RPCDesc {
	rpc := &RPCDesc{}
	rpc.Name = name
	rpc.Methods = make([]*MethodDesc, 0)

	f.RPCs = append(f.RPCs, rpc)

	return rpc
}

func NewProtoFile() *FileDesc {
	f := &FileDesc{}
	f.Packages = make([]*PackageDesc, 0)
	f.Messages = make([]*MessageDesc, 0)
	f.RPCs = make([]*RPCDesc, 0)

	return f
}
