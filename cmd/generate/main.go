package main

import (
	"bytes"
	"context"
	"go/format"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/google/go-github/v32/github"
)

const repoTemplate = `
// Code generated .* DO NOT EDIT.
package repos

var Repos = map[string]string{
{{- range .}}
	"{{.Name}}": "{{.URL}}",
{{- end}}
}
`

type repo struct {
	Name string
	URL  string
}

func main() {
	if len(os.Args) < 2 {
		panic("not enough arguments")
	}
	org := os.Args[1]
	client := github.NewClient(nil)
	opts := &github.RepositoryListByOrgOptions{
		Type: "public",
	}
	out := []*repo{}
	for {
		repos, resp, err := client.Repositories.ListByOrg(context.Background(), org, opts)
		if err != nil {
			panic(err)
		}
		for _, r := range repos {
			if r.Language != nil && *r.Language == "Go" {
				out = append(out, &repo{
					Name: *r.Name,
					URL:  *r.HTMLURL,
				})
			}
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	var b bytes.Buffer
	tmpl := template.Must(template.New("repos").Parse(repoTemplate))
	tmpl.Execute(&b, out)
	goFile, err := format.Source(b.Bytes())
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("repo.go", goFile, 0644)
	return
}
