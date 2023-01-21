package __datatype

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

type Sex string

const (
	Male = Sex("male")
	Female = Sex("female")
)

type Student struct{
	Name string
	Sex Sex
	Age uint8
}

func TestTextTemplate(t *testing.T) {
	studentTmp := `
name: {{.Name}}
sex: {{.Sex}}
age: {{.Age}}
`
	student := &Student{
		Name: "darren",
		Sex: Male,
		Age: 26,
	}
	tmp, err := template.New("student template").Parse(studentTmp)
	if err != nil {
		t.Fatal(err)
	}
	buffer :=&bytes.Buffer{}
	err = tmp.Execute(buffer, student)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Execute result is: %s", buffer.String())
}

func TestTemplateIf(t *testing.T) {
	var studentTmp = `
{{if eq .Age 0}}
I'm adult.
{{else if lt .Age 18}}
I'm yong.
{{else if eq .Age 18}}
I'm just adult.
{{else}}
I'm already adult.
{{end}}
`
	student := &Student{
		Name: "darren",
		Sex: Male,
		Age: 26,
	}
	tmp, err := template.New("student template").Parse(studentTmp)
	if err != nil {
		t.Fatal(err)
	}
	buffer :=&bytes.Buffer{}
	err = tmp.Execute(buffer, student)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Execute result is: %s\n", buffer.String())
}

func TestTemplateFor(t *testing.T) {
	var studentTmp = `
Students' name are{{range .}} {{.Name}}{{end}}.`
	students := []*Student{
		{
			Name: "darren",
			Sex: Male,
			Age: 26,
		},
		{
			Name: "lucy",
			Sex: Male,
			Age: 26,
		},
		{
			Name: "hallen",
			Sex: Male,
			Age: 26,
		},
	}
	tmp, err := template.New("student template").Parse(studentTmp)
	if err != nil {
		t.Fatal(err)
	}
	buffer :=&bytes.Buffer{}
	err = tmp.Execute(buffer, students)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Execute result is: %s\n", buffer.String())
}