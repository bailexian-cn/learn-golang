package commonpkg_fmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

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
	js := &IssuesSearchResult{}
	if err := json.Unmarshal(ba, js); err != nil {
		panic(err)
	}
	for _, item := range js.Items {
		s := fmt.Sprintf("#%-5d#%-15s#%-150s#",
			item.Number, item.User.Login, item.Title)
		fmt.Println(s)
	}
}
