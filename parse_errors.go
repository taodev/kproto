package kproto

import (
	"errors"
	"fmt"
)

var (
	proto_error_info []string

	file_name string
	file_line int
)

func combo_errors() error {
	var errstr string
	for _, v := range proto_error_info {
		errstr += v + "\n"
	}

	return errors.New(errstr)
}

func error_print(err string) {
	info := fmt.Sprintf("%s:%d: %s", file_name, file_line, err)
	proto_error_info = append(proto_error_info, info)
}

func has_error() bool {
	return len(proto_error_info) > 0
}
