package main

import (
	"bailexian.cn/learn-golang/1-basic/6-commonpkg/fmt/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func main() {
	//apath := "github_issues_harbor.json"
	_, filename, _, ok := runtime.Caller(0)
	var abPath string
	if ok {
		abPath = path.Dir(filename)
	}
	abPath = abPath + "/github_issues_harbor.json"
	ba, err := ioutil.ReadFile(abPath)
	if err != nil {
		fmt.Errorf("%v\n", err)
	}
	js := &model.IssuesSearchResult{}
	if err := json.Unmarshal(ba, js); err != nil {panic(err)}

	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

	if err := report.Execute(os.Stdout, js); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
