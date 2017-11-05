package item

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeHandler(s Service) http.Handler {

	getRoot := &getRootHandler{s}
	getItem := &getItemHandler{s}
	addItem := &addItemHandler{s}

	r := mux.NewRouter()
	r.StrictSlash(true)
	r.Path("/items/{id}").Handler(getItem).Methods("GET")
	r.Path("/items/").Handler(getRoot).Methods("GET")
	r.Path("/items/").Handler(addItem).Methods("POST")
	return r
}

type getRootHandler struct{ service Service }
type getItemHandler struct{ service Service }
type addItemHandler struct{ service Service }

func (h *getRootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i, err := h.service.GetRoot()
	if err != nil {
		body := errorResponse{
			Code:        http.StatusInternalServerError,
			Message:     http.StatusText(http.StatusInternalServerError),
			Description: err.Error(),
		}
		renderErrJSON(w, body)
		return
	}
	renderItemJSON(w, i)
}

func (h *getItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := h.service.GetItem(id)
	if err == ErrItemNotFound {
		body := errorResponse{
			Code:        http.StatusNotFound,
			Message:     http.StatusText(http.StatusNotFound),
			Description: err.Error(),
		}
		renderErrJSON(w, body)
		return
	}
	if err != nil {
		body := errorResponse{
			Code:        http.StatusInternalServerError,
			Message:     http.StatusText(http.StatusInternalServerError),
			Description: err.Error(),
		}
		renderErrJSON(w, body)
		return
	}
	renderItemJSON(w, i)
}

func (h *addItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var resp response
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		body := errorResponse{
			Code:        http.StatusInternalServerError,
			Message:     http.StatusText(http.StatusInternalServerError),
			Description: err.Error(),
		}
		renderErrJSON(w, body)
		return
	}

	id, err := h.service.AddItem(resp.Item.Name, resp.Item.Parent, resp.Item.Kind, resp.Item.Owner)

	renderIDJSON(w, id)
}

func renderErrJSON(w http.ResponseWriter, err errorResponse) {
	body := response{Error: &err}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(body)
}

func renderItemJSON(w http.ResponseWriter, i *Item) {
	item := &itemResponse{
		ID:       i.ID,
		Name:     i.Name,
		Kind:     i.Kind,
		Parent:   i.Parent,
		Path:     i.Path,
		Owner:    i.Owner,
		Children: i.Children,
	}
	renderJSON(w, "", item, nil)
}

func renderIDJSON(w http.ResponseWriter, id string) {
	renderJSON(w, id, nil, nil)
}

func renderJSON(w http.ResponseWriter, id string, item *itemResponse, err *errorResponse) {
	body := response{ID: id, Item: item, Error: err}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

type response struct {
	ID    string         `json:"id,omitempty"`
	Item  *itemResponse  `json:"item,omitempty"`
	Error *errorResponse `json:"error,omitempty"`
}

type itemResponse struct {
	ID       string   `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	Parent   string   `json:"parent"`
	Path     string   `json:"path"`
	Owner    string   `json:"owner,omitempty"`
	Children []string `json:"children"`
}

type errorResponse struct {
	Code        int    `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}
