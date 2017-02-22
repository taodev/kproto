package plugin

import "github.com/taodev/kproto"

type Plugin interface {
	Name() string
	Generate(file *kproto.FileDesc) error
}

var plugins []Plugin

func RegisterPlugin(p Plugin) {
	plugins = append(plugins, p)
}

func GetPlugin(name string) Plugin {
	for _, p := range plugins {
		if p.Name() == name {
			return p
		}
	}

	return nil
}

func EnumPlugins() []string {
	names := make([]string, len(plugins))

	for i, p := range plugins {
		names[i] = p.Name()
	}

	return names
}
