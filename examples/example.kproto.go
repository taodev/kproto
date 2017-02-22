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
	UserName uint32
	Password string
	Auth     AuthMsg
	LastIP   string
}

type ILoginService interface {
	Login(req *Login1Msg) (reply *AuthMsg, err error)
}

func (msg *AuthMsg) Write(w *kproto.Buffer) error {
	var err error
	if err = w.WriteUint32(msg.Test); err != nil {
		return err
	}
	return err
}

func (msg *Login1Msg) Write(w *kproto.Buffer) error {
	var err error
	if err = w.WriteUint32(msg.UserName); err != nil {
		return err
	}
	if err = w.WriteString(msg.Password); err != nil {
		return err
	}
	if err = w.WriteStruct(&msg.Auth); err != nil {
		return err
	}
	if err = w.WriteString(msg.LastIP); err != nil {
		return err
	}
	return err
}
