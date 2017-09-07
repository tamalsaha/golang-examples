package main

import (
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func example_0() {
	tmpStr := "Hello {{.}}\n"
	tmp := template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, "world"))

	type student struct {
		Name string
		Age  int
	}
	student_1 := student{"Alice", 21}
	tmpStr = "Name: {{.Name}} and Age: {{.Age}}\n"
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, student_1))
}

func example_1() {
	tmpStr := "{{23 -}} < {{- 45}}\n"
	tmp := template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, nil))

	tmpStr = `***
	{{- "Hello" -}}
	***`
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, nil))
}

func example_2() {
	arr_1 := []int{11, 22, 33}
	map_1 := map[int]string{5: "AA", 6: "BB", 7: "CC"}
	map_2 := map[string]string{"a": "A", "b": "B", "c": "C"}

	tmpStr := "Value: {{index . 0}}\n"
	tmp := template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, arr_1))

	tmpStr = "Value: {{index . 5}}\n"
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, map_1))

	tmpStr = "Value: {{.a}}\n"
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, map_2))
}

func example_3() {
	tmpStr := "Iterate: {{range .}} Value = {{.}} {{end}}\n"
	tmp := template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, []int{11, 22, 33}))

	tmpStr = "Iterate: {{range $i, $v := .}} {{$i}}:{{$v}} {{end}}\n"
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, []int{11, 22, 33}))

	tmpStr = "Iterate: {{range .}} Value = {{.}} {{else}}Data is empty{{end}}\n"
	tmp = template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, []int{}))
}

func example_4() {
	type person struct {
		Name   string
		Emails []string
	}
	me := person{
		Name:   "Dipta",
		Emails: []string{"abc@gmail.com", "pqr@gmail.com"},
	}

	tmpStr := `{{$name := .Name}}
{{range .Emails}} Name: {{$name}} Email: {{.}}
{{end}}`
	tmp := template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, me))
}

func example_5() {
	tmpStr := `Grade:
{{- if ge . 80}} A
{{- else if ge . 60}} B
{{- else if ge . 40}} C
{{- else}} Fail
{{- end}}
`
	tmp := template.Must(template.New("test").Parse(tmpStr))
	check(tmp.Execute(os.Stdout, 65))
	check(tmp.Execute(os.Stdout, 80))
	check(tmp.Execute(os.Stdout, 35))
	check(tmp.Execute(os.Stdout, 45))
}

func main() {
	example_0()
}
