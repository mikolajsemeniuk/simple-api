package handler

import "net/http"

type ArticleHandler struct{}

func (h *ArticleHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list"))
}
func (h *ArticleHandler) Read(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("read"))
}
func (h *ArticleHandler) Add(w http.ResponseWriter, r *http.Request) {
	// we want to validate data
	// we want to get cache if cache exists return it save it
	// if doesnt return get value if value is earlier than today send it to queue
	// save data to elastic search
	// save it to postgres as well
	w.Write([]byte("add"))
}
func (h *ArticleHandler) Update(w http.ResponseWriter, r *http.Request) {
	// we want to validate data
	// we want to get cache if cache exists return it save it
	// if doesnt return get value if value is earlier than today send it to queue
	// save data to postgres
	w.Write([]byte("update"))
}
func (h *ArticleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// remove data from postgres
	// remove data from elastic
	w.Write([]byte("delete"))
}
