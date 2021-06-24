package rest

import (
	"b-nova-openhub/stapagen/pkg/gen"
	"b-nova-openhub/stapagen/pkg/repo"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/page", getPage).Methods("GET")
	router.HandleFunc("/pages", getPages).Methods("GET")
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/generate", getStatus).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getPage(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	pages := gen.Generate(repo.RepoContents())
	page := getPageById(v.Get("id"), pages)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&page)
}

func getPages(w http.ResponseWriter, r *http.Request) {
	pages := gen.GeneratedPages
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pages)
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	pages := gen.Generate(repo.RepoContents())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pages)
}

func getGenerate(w http.ResponseWriter, r *http.Request) {
	pages := gen.Generate(repo.RepoContents())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pages)
}

func getPageById(id string, pages []gen.StaticPage) *gen.StaticPage {
	var page *gen.StaticPage
	for _, p := range pages {
		if p.Permalink == id {
			page = &p
		}
	}
	return page
}