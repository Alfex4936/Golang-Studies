package main

import (
	"encoding/json"
	"fmt"
)

type K map[string]interface{}

type Error interface {
	Text(msg string) // return error message
	Json() K
}

type Result struct {
	code int
	msg  string
}

func (r *Result) Text(msg string) {
	r.msg = msg
}

func (r Result) Json() K {
	return K{"version": "1.0", "error": r.msg}
}

func Call(e ...Error) {
	msgs := []string{"load failed", "build failed", "injection failed"}
	for i, err := range e {
		err.Text(msgs[i])
		res, _ := json.Marshal(err.Json())
		fmt.Println(string(res))
	}
}

func main() {
	load_failed := &Result{}
	build_failed := &Result{}
	injection_failed := &Result{}

	Call(load_failed, build_failed, injection_failed)
}
