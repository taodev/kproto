package kproto

type PackageDesc struct {
	Lang string
	Name string
}

type FieldDesc struct {
	Name   string
	Type   string
	Length uint16
}

type MessageDesc struct {
	Id     uint16
	Name   string
	Fields []*FieldDesc
}

func (msg *MessageDesc) AddField(name, typ string, l uint16) {
	field := &FieldDesc{}
	field.Name = name
	field.Type = typ
	field.Length = l

	msg.Fields = append(msg.Fields, field)
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

	f.Messages = append(f.Messages, msg)

	return msg
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
