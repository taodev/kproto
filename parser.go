package kproto

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type KProtoType int

const (
	KPT_Unknown KProtoType = iota
	KPT_Comment
	KPT_Package
	KPT_Message
	KPT_Field
	KPT_RPC
	KPT_Method
	KPT_SpaceLine
)

var (
	ErrParse   = errors.New("代码编译出错")
	ErrLine    = errors.New("错误行")
	ErrComment = errors.New("注释行出错")
	errPackage = errors.New("包定义错误")
)

var KPTString = []string{
	"unknown",
	"comment",
	"package",
	"message",
	"field",
	"rpc",
	"method",
	"spaceline",
}

type ParseFunc func(string, *FileDesc) error

var LineFuncSet = []ParseFunc{
	nil,
	nil,
	parse_package,
	parse_message,
	parse_field,
	parse_rpc,
	parse_method,
	nil,
}

var (
	last_message *MessageDesc
	last_rpc     *RPCDesc
	last_type    KProtoType

	has_package bool
)

func LoadProtoFile1(fname string) (out *FileDesc, err error) {
	file, err := os.Open(fname)
	if err != nil {
		return
	}
	defer file.Close()

	file_name = fname
	file_line = 0

	last_message = nil
	last_rpc = nil
	last_type = KPT_Unknown
	has_package = false

	rbuf := bufio.NewReader(file)

	out = NewProtoFile()
	out.FileName = file_name

	var lineBytes []byte

	quit := false
	for !quit {
		file_line++

		lineBytes, _, err = rbuf.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
				goto exit
			} else {
				return
			}
		}

		line := string(lineBytes)

		lineType := regexp_line_type(line)

		switch lineType {
		case KPT_Package, KPT_Message, KPT_Field, KPT_RPC, KPT_Method:
			{
				fn := LineFuncSet[lineType]
				if fn == nil {
					continue
				}

				if err = fn(line, out); err != nil {
					err = nil
					goto exit
				}
			}
		case KPT_SpaceLine, KPT_Comment:
			continue
		default:
			error_print(ErrLine.Error())
			continue
		}
	}

exit:
	if has_error() {
		err = combo_errors()
	}

	return
}

func parse_package(code string, out *FileDesc) error {
	has_package = true
	last_message = nil

	if last_type != KPT_Unknown {
		error_print("包定义必须在所有消息跟RPC定义之前")
		return errPackage
	}

	last_type = KPT_Unknown

	args, ok := regexp_find(code, `^\s*package\s+(\w*)?\s*(:)?\s*(\w+)\s*(//.*)?$`, 5)
	if !ok {
		error_print("语法错误: " + code)
		return errPackage
	}

	var lang, name string
	lang = args[1]
	if len(args[2]) > 0 {
		if len(lang) <= 0 {
			error_print("语法错误: 冒号前无语言标识[go, ts, ...]")
			return errPackage
		}
	}

	name = args[3]

	out.AddPackage(lang, name)

	return nil
}

func parse_message(code string, out *FileDesc) error {
	last_type = KPT_Message

	if !has_package {
		error_print("文件头必须定义package")
		return nil
	}

	var name string
	var id uint16

	last_message = nil
	last_rpc = nil

	args, ok := regexp_find(code, `^\s*message\s+(\w*)\s*:\s*(\w*)?\s*(//.*)?$`, 4)
	if !ok {
		error_print("语法错误: 消息定义出错")
		return nil
	}

	name = args[1]
	if len(args[2]) > 0 {
		num, err := strconv.ParseUint(args[2], 10, 16)
		if err != nil {
			error_print(fmt.Sprintf("ID错误: %s:%s", name, args[2]))
			return nil
		}

		id = uint16(num)
	}

	out.IDCounter++
	if id > 0 {
		out.IDCounter = id
	}

	last_message = out.AddMessage(out.IDCounter, name)
	return nil
}

func parse_field(code string, out *FileDesc) (err error) {
	if (last_type != KPT_Message) || last_message == nil {
		error_print("字段必须定义在消息下")
		return
	}

	var name, typ string
	var length uint16

	args, ok := regexp_find(code, `^\s+([A-Z][a-z0-9A-Z_]{0,31})\s*([a-z0-9A-Z_]{0,32})\s*(:)?\s*([0-9]*)\s*(//.*)?$`, 6)
	if !ok {
		error_print("字段错误")
		return
	}

	name = args[1]
	typ = args[2]

	if args[3] == ":" {
		if len(args[4]) <= 0 {
			error_print("数组必须定义长度")
			return
		}
	}

	if len(args[4]) > 0 {
		if args[3] != ":" {
			error_print(`类型与长度必须用":"间隔`)
			return
		}

		var num uint64
		num, err = strconv.ParseUint(args[4], 10, 16)
		if err != nil {
			error_print(fmt.Sprintf("ID错误: %s:%s", name, args[4]))
			err = nil
			return
		}

		length = uint16(num)
	}

	if typ == "string" && length <= 0 {
		error_print(fmt.Sprintf("string类型必须定义最大长度"))
		return
	}

	last_message.AddField(name, typ, length)

	return
}

func parse_rpc(code string, out *FileDesc) (err error) {
	last_type = KPT_RPC

	var name string
	last_message = nil
	last_rpc = nil

	args, ok := regexp_find(code, `^\s*rpc\s+([A-Z][a-z0-9A-Z_]{0,31})\s*:\s*(//.*)?$`, 3)
	if !ok {
		error_print("RPC定义出错")
		return
	}

	name = args[1]

	last_rpc = out.AddRPC(name)

	return
}

func parse_method(code string, out *FileDesc) (err error) {
	if (last_type != KPT_RPC) || last_rpc == nil {
		error_print("方法必须定义在函数下")
		return
	}

	var name, request, reply string
	args, ok := regexp_find(code, `^\s+([A-Z][a-z0-9A-Z_]{0,31})\s*\(\s*([a-z0-9A-Z_]{0,32})\s*\)\s*([a-z0-9A-Z_]{0,32})\s*(//.*)?$`, 5)
	if !ok {
		error_print("方法定义错误")
		return
	}

	name = args[1]
	request = args[2]
	reply = args[3]

	last_rpc.AddMethod(name, request, reply)

	return
}

func regexp_line_type(line string) KProtoType {
	matched := false
	matched, _ = regexp.MatchString(`^\s*$`, line)
	if matched {
		return KPT_SpaceLine
	}

	matched, _ = regexp.MatchString(`^\s*//`, line)
	if matched {
		return KPT_Comment
	}

	matched, _ = regexp.MatchString(`^\s*package\s+`, line)
	if matched {
		return KPT_Package
	}

	matched, _ = regexp.MatchString(`^\s*message\s+`, line)
	if matched {
		return KPT_Message
	}

	matched, _ = regexp.MatchString(`^\s*rpc\s+`, line)
	if matched {
		return KPT_RPC
	}

	matched, _ = regexp.MatchString(`^\s*\w+\s*\(`, line)
	if matched {
		return KPT_Method
	}

	matched, _ = regexp.MatchString(`^\s*\w+\s+\w+`, line)
	if matched {
		return KPT_Field
	}

	return KPT_Unknown
}

func regexp_find(code string, pattern string, count int) (args []string, ok bool) {
	ok = false
	r := regexp.MustCompile(pattern)
	if r == nil {
		return
	}

	result := r.FindAllStringSubmatch(code, 1)
	if len(result) != 1 {
		return
	}

	if len(result[0]) != count {
		return
	}

	args = result[0]
	ok = true

	return
}
