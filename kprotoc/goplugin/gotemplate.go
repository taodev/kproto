package goplugin

var codeTmpl = `package {{.PackageName}}
import kproto "github.com/taodev/kproto"
	
const (
{{range .Messages}}	{{.Name}}ID = {{.Id}}
{{end}})
{{range .Messages}}
type {{.Name}} struct { {{range .Fields}}
	{{.Name}} {{if gt .Length 0}}[]{{end}}{{.Type}} {{end}}
}
{{end}}{{range .RPCs}}
type I{{.Name}} interface { {{range .Methods}}
	{{.Name}}(req *{{.Request}}) (reply *{{.Reply}}, err error){{end}}
}{{end}}

{{range .Messages}}
func (msg *{{.Name}}) Write(w *kproto.ByteWriter) (err error) {
	{{range $i, $v := .Fields}}{{PrintWrite $v}}
	if err != nil {
		return
	}
	{{end}}return
}
	
func (msg *{{.Name}}) Read(r *kproto.ByteReader) (err error) {
	{{range $i, $v := .Fields}}{{PrintRead $v}}
	if err != nil {
		return
	}
	{{end}}return
}

func (msg *{{.Name}}) MaxSize() int {
	return {{.MaxSize }}
}
{{end}}

`
