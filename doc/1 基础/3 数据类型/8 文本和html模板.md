# 文本模板demo

```golang
var studentTmp = `
name: {{.Name}}
sex: {{.Sex}}
age: {{.Age}}
`

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
	tmp, err := template.New("student template").Parse(studentTmp)
	if err != nil {
		t.Fatal(err)
	}
	student := &Student{
		Name: "darren",
		Sex: Male,
		Age: 26,
	}
	buffer :=&bytes.Buffer{}
	err = tmp.Execute(buffer, student)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Execute result is: %s", buffer.String())
}
```

## if

```
var studentTmp = `
{{if eq .Age 0}}
I'm adult.
{{else if lt .Age 18}}
I'm just adult.
{{else}}
I'm already adult.
{{end}}
`
```

|      | 说明         |
| ---- | ------------ |
| eq   | 等于         |
| ne   | 不等于       |
| gt   | 大于         |
| ge   | 大于或者等于 |
| lt   | 小于         |
| le   | 小于或者等于 |

## for

```golang
var studentTmp = `
Students' name are{{range .}} {{.Name}}{{end}}.`
```

