package goplugin

var codeTmpl = `package {{.PackageName}}
import kproto "github.com/taodev/kproto"
	
const (
{{range .Messages}}	{{.Name}}ID = {{.Id}}
{{end}})
{{range .Messages}}
type {{.Name}} struct { {{range .Fields}}
	{{.Name}} {{.Type}} {{end}}
}
{{end}}{{range .RPCs}}
type I{{.Name}} interface { {{range .Methods}}
	{{.Name}}(req *{{.Request}}) (reply *{{.Reply}}, err error){{end}}
}{{end}}

{{range .Messages}}
func (msg *{{.Name}}) Write(w *kproto.Buffer) error {
	var err error
	{{range $i, $v := .Fields}}if err = {{PrintWrite $v}}; err != nil {
		return err
	}
	{{end}}return err
}

func (msg *{{.Name}}) MaxSize() int {
	return {{.MaxSize }}
}
{{end}}

`
