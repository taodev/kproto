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

func PrintWrite(field *kproto.FieldDesc) string {
	code := ""
	switch field.Type {
	case "bool":
		code = fmt.Sprintf("w.WriteBool(msg.%s)", field.Name)
	case "byte":
		code = fmt.Sprintf("w.WriteByte(msg.%s)", field.Name)
	case "int8":
		code = fmt.Sprintf("w.WriteInt8(msg.%s)", field.Name)
	case "uint8":
		code = fmt.Sprintf("w.WriteUint8(msg.%s)", field.Name)
	case "int16":
		code = fmt.Sprintf("w.WriteInt16(msg.%s)", field.Name)
	case "uint16":
		code = fmt.Sprintf("w.WriteUint16(msg.%s)", field.Name)
	case "int32":
		code = fmt.Sprintf("w.WriteInt32(msg.%s)", field.Name)
	case "uint32":
		code = fmt.Sprintf("w.WriteUint32(msg.%s)", field.Name)
	case "int64":
		code = fmt.Sprintf("w.WriteInt64(msg.%s)", field.Name)
	case "uint64":
		code = fmt.Sprintf("w.WriteUint64(msg.%s)", field.Name)
	case "float32":
		code = fmt.Sprintf("w.WriteFloat32(msg.%s)", field.Name)
	case "float64":
		code = fmt.Sprintf("w.WriteFloat64(msg.%s)", field.Name)
	case "string":
		code = fmt.Sprintf("w.WriteString(msg.%s)", field.Name)
	default:
		code = fmt.Sprintf("w.WriteStruct(&msg.%s)", field.Name)
	}

	return code
}
