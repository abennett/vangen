package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/abennett/vangen/repos"
	"github.com/go-chi/chi"
)

const (
	repoNameParam   = "repo_name"
	vanityDomainEnv = "VANITY_DOMAIN"
	portEnv         = "PORT"
)

type VanityService struct {
	url      string
	template *template.Template
	repos    map[string]string
}

type repo struct {
	Name         string
	GithubURL    string
	VanityDomain string
}

func NewVanityService(url string) *VanityService {
	tmpl, err := template.New("goPage").Parse(pageTemplate)
	if err != nil {
		panic(err)
	}
	return &VanityService{
		url:      url,
		template: tmpl,
		repos:    repos.Repos,
	}
}

func (vs *VanityService) GetRepo(name string) (*repo, bool) {
	githubURL, ok := vs.repos[name]
	if !ok {
		return nil, false
	}
	return &repo{
		Name:         name,
		GithubURL:    githubURL,
		VanityDomain: vs.url,
	}, true
}

func (vs *VanityService) mux() *chi.Mux {
	r := chi.NewMux()
	r.Get("/{repo_name}", vs.repoHandlerFunc)
	return r
}

func (vs *VanityService) repoHandlerFunc(w http.ResponseWriter, r *http.Request) {
	repoName := chi.URLParam(r, repoNameParam)
	repo, ok := vs.GetRepo(repoName)
	if !ok {
		http.Error(w, repoName+" not found", http.StatusNotFound)
		return
	}
	if err := vs.template.Execute(w, repo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func main() {
	port, ok := os.LookupEnv(portEnv)
	if !ok {
		port = ":8080"
	}
	domain, ok := os.LookupEnv(vanityDomainEnv)
	if !ok {
		log.Fatal("must specify VANITY_DOMAIN")
	}
	vs := NewVanityService(domain)
	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, vs.mux()))
}
