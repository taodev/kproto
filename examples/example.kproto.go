package examples

import kproto "github.com/taodev/kproto"

const (
	AuthMsgID   = 101
	Login1MsgID = 102
)

type AuthMsg struct {
	Test uint32
}

type Login1Msg struct {
	ID       uint64
	UserName string
	Password string
	Auth     AuthMsg
	LastIP   string
	Auths    []AuthMsg
}

type ILoginService interface {
	Login(req *Login1Msg) (reply *AuthMsg, err error)
}

func (msg *AuthMsg) Write(w *kproto.ByteWriter) (err error) {
	err = w.WriteUint32(msg.Test)
	if err != nil {
		return
	}
	return
}

func (msg *AuthMsg) Read(r *kproto.ByteReader) (err error) {
	msg.Test, err = r.ReadUint32()
	if err != nil {
		return
	}
	return
}

func (msg *AuthMsg) MaxSize() int {
	return 4
}

func (msg *Login1Msg) Write(w *kproto.ByteWriter) (err error) {
	err = w.WriteUint64(msg.ID)
	if err != nil {
		return
	}
	err = w.WriteString(msg.UserName)
	if err != nil {
		return
	}
	err = w.WriteString(msg.Password)
	if err != nil {
		return
	}
	err = msg.Auth.Write(w)
	if err != nil {
		return
	}
	err = w.WriteString(msg.LastIP)
	if err != nil {
		return
	}
	{
		l := len(msg.Auths)
		err = w.WriteLength(l)
		if err != nil {
			return
		}
		for i := 0; i < l; i++ {
			if err = msg.Auths[i].Write(w); err != nil {
				return
			}
		}
	}
	if err != nil {
		return
	}
	return
}

func (msg *Login1Msg) Read(r *kproto.ByteReader) (err error) {
	msg.ID, err = r.ReadUint64()
	if err != nil {
		return
	}
	msg.UserName, err = r.ReadString()
	if err != nil {
		return
	}
	msg.Password, err = r.ReadString()
	if err != nil {
		return
	}
	err = msg.Auth.Read(r)
	if err != nil {
		return
	}
	msg.LastIP, err = r.ReadString()
	if err != nil {
		return
	}
	{
		var l int
		l, err = r.ReadLength()
		if err != nil {
			return
		}
		msg.Auths = make([]AuthMsg, l)
		for i := 0; i < l; i++ {
			if err = msg.Auths[i].Read(r); err != nil {
				return
			}
		}
	}
	if err != nil {
		return
	}
	return
}

func (msg *Login1Msg) MaxSize() int {
	return 57
}
