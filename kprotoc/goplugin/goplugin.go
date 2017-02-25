package goplugin

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/taodev/kproto"
	"github.com/taodev/kproto/kprotoc/plugin"

	"text/template"
)

type goplugin struct {
}

func (g *goplugin) Name() string {
	return "go"
}

func (g *goplugin) Generate(file *kproto.FileDesc) error {
	return Build(file.FileName, file)
}

func init() {
	plugin.RegisterPlugin(new(goplugin))
}

var funcMap = template.FuncMap{
	"PrintWrite": PrintWrite,
	"PrintRead":  PrintRead,
}

func Build(sourceName string, in *kproto.FileDesc) error {
	in.SetPackage("go")
	tmpl, err := template.New("gotemplate").Funcs(funcMap).Parse(codeTmpl)
	if err != nil {
		return err
	}

	outPath := sourceName + ".go"

	fout, err := os.Create(outPath)
	if err != nil {
		return err
	}

	err = tmpl.Execute(fout, in)
	if err != nil {
		fout.Close()
		return err
	}

	fout.Close()

	cmd := exec.Command("gofmt", "-w", "-l", outPath)
	cmd.Stdout = os.Stdout

	if err = cmd.Run(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

var writeCode = `	if err = binary.Write(w, binary.LittleEndian, msg.%s);err != nil {
		return err
	}`

var writeArrayCode = `	if err = binary.Write(w, binary.LittleEndian, uint16(len(msg.%s)));err != nil {
		return err
	}`

//func PrintWrite(field *kproto.FieldDesc) string {
//	code := ""
//	switch field.Type {
//	case "bool", "byte",
//		"int8", "uint8",
//		"int16", "uint16",
//		"int32", "uint32",
//		"int64", "uint64",
//		"float32", "float64":
//		if field.Length > 0 {
//			code = fmt.Sprintf(writeArrayCode, field.Name)
//		}
//		code += fmt.Sprintf(writeCode, field.Name)
//	case "string":
//		code = fmt.Sprintf(`if err = binary.Write(w, binary.LittleEndian, uint16(len(msg.%s)));err != nil {
//		return
//	}
//	if err = binary.Write(w, binary.LittleEndian, []byte(msg.%s));err != nil {
//		return
//	}`, field.Name, field.Name)
//	default:
//		if field.Length > 0 {
//			code = fmt.Sprintf(`if err = msg.%s.Write(w);err != nil {
//			return
//		}`, field.Name)
//		}
//		code = fmt.Sprintf(`if err = msg.%s.Write(w);err != nil {
//			return
//		}`, field.Name)
//	}

//	return code
//}

func PrintWrite(field *kproto.FieldDesc) string {
	code := ""
	switch field.Type {
	case "bool":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteBoolArray(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteBool(msg.%s)", field.Name)
		}
	case "byte":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteByteArray(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteByte(msg.%s)", field.Name)
		}
	case "int8":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteInt8Array(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteInt8(msg.%s)", field.Name)
		}
	case "uint8":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteUint8Array(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteUint8(msg.%s)", field.Name)
		}
	case "int16":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteInt16Array(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteInt16(msg.%s)", field.Name)
		}
	case "uint16":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteUint16Array(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteUint16(msg.%s)", field.Name)
		}
	case "int32":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteInt32Array(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteInt32(msg.%s)", field.Name)
		}
	case "uint32":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteUint32Array(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteUint32(msg.%s)", field.Name)
		}
	case "int64":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteInt64Array(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteInt64(msg.%s)", field.Name)
		}
	case "uint64":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteUint64Array(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteUint64(msg.%s)", field.Name)
		}
	case "float32":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteFloat32Array(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteFloat32(msg.%s)", field.Name)
		}
	case "float64":
		if field.Length > 0 {
			code = fmt.Sprintf("err = w.WriteFloat64Array(msg.%s)", field.Name)
		} else {
			code = fmt.Sprintf("err = w.WriteFloat64(msg.%s)", field.Name)
		}
	case "string":
		code = fmt.Sprintf("err = w.WriteString(msg.%s)", field.Name)
	default:
		if field.Length > 0 {
			code = fmt.Sprintf(`{
				l := len(msg.%s)
				err = w.WriteLength(l)
				if err != nil {
					return
				}
				for i := 0;i < l;i++ {
					if err = msg.%s[i].Write(w);err != nil {
						return
					}
				}
			}`, field.Name, field.Name)
		} else {
			code = fmt.Sprintf("err = msg.%s.Write(w)", field.Name)
		}
	}

	return code
}

func PrintRead(field *kproto.FieldDesc) string {
	code := ""
	switch field.Type {
	case "bool":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadBoolArray()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadBool()", field.Name)
		}
	case "byte":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadByteArray()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadByte()", field.Name)
		}
	case "int8":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadInt8Array()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadInt8()", field.Name)
		}
	case "uint8":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadUint8Array()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadUint8()", field.Name)
		}
	case "int16":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadInt16Array()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadInt16()", field.Name)
		}
	case "uint16":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadUint16Array()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadUint16()", field.Name)
		}
	case "int32":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadInt32Array()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadInt32()", field.Name)
		}
	case "uint32":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadUint32Array()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadUint32()", field.Name)
		}
	case "int64":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadInt64Array()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadInt64()", field.Name)
		}
	case "uint64":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadUint64Array()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadUint64()", field.Name)
		}
	case "float32":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadFloat32Array()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadFloat32()", field.Name)
		}
	case "float64":
		if field.Length > 0 {
			code = fmt.Sprintf("msg.%s, err = r.ReadFloat64Array()", field.Name)
		} else {
			code = fmt.Sprintf("msg.%s, err = r.ReadFloat64()", field.Name)
		}
	case "string":
		code = fmt.Sprintf("msg.%s, err = r.ReadString()", field.Name)
	default:
		if field.Length > 0 {
			code = fmt.Sprintf(`{
				var l int
				l, err = r.ReadLength()
				if err != nil {
					return
				}
				msg.%s = make([]%s, l)
				for i := 0;i < l;i++ {
					if err = msg.%s[i].Read(r);err != nil {
						return
					}
				}
			}`, field.Name, field.Type, field.Name)
		} else {
			code = fmt.Sprintf("err = msg.%s.Read(r)", field.Name)
		}
	}

	return code
}
