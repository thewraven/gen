package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const generated = `
//TableName returns the table name which uses the storage interface
func ({{.Initial}}  *{{.Type}}) TableName() string { return table{{.Type}} }

//GetID returns the unique ID of the variable
func ({{.Initial}} *{{.Type}}) GetID() string   { return {{.Initial}}.ID }

//SetID sets the unique ID of the value
func ({{.Initial}} *{{.Type}}) SetID(id string) { {{.Initial}}.ID = id }

//Validate checks if the struct is valid to be inserted or updated
func ({{.Initial}} *{{.Type}}) Validate() bool {
	return true
}

//New returns an empty {{.Type}}
func ({{.Initial}} *{{.Type}}) New() storage.Object {
	return &{{.Type}}{}
}

//NewArray returns an empty slice of {{.Type}}
func ({{.Initial}} *{{.Type}}) NewArray() interface{} {
	list := make([]{{.Type}}, 0)
	return &list
}

//FuncMap returns a map of custom functions to be used on the REST generator
func ({{.Initial}} *{{.Type}}) FuncMap() storage.FuncMap {
	funcs := make(storage.FuncMap)
	return funcs
}
`

type Info struct {
	Initial string
	Type    string
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s [file] [type]\n", os.Args[0])
		return
	}
	parser, err := template.New("").Parse(generated)
	if err != nil {
		panic(err)
	}
	path, typeName := os.Args[1], os.Args[2]
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	initial := strings.ToLower(typeName[0:1])
	info := Info{
		Initial: initial,
		Type:    typeName,
	}
	err = parser.Execute(file, info)
	if err != nil {
		panic(err)
	}
}
